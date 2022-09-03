
/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package fsm

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"testing"

	db "a13s.io/pleiades/api/v1/database"
	"github.com/lni/dragonboat/v3/statemachine"
	"github.com/stretchr/testify/suite"
	"go.etcd.io/bbolt"
)

type TestBBoltFsm struct {
	suite.Suite
}

func TestBBoltStateMachine(t *testing.T) {
	suite.Run(t, new(TestBBoltFsm))
}

func (bfsm *TestBBoltFsm) TestNewBBoltStateMachine() {
	opts := &bbolt.Options{
		Timeout:         0,
		NoGrowSync:      false,
		NoFreelistSync:  false,
		FreelistType:    bbolt.FreelistMapType,
		ReadOnly:        false,
		InitialMmapSize: 0,
		PageSize:        0,
		NoSync:          false,
		OpenFile:        nil,
		Mlock:           false,
	}

	bfsm.Require().NotPanics(func() {
		NewBBoltStateMachine(1, 1, bfsm.T().TempDir(), opts)
	}, "creating a new bbolt fsm must not throw an error")
}

func (bfsm *TestBBoltFsm) TestBBoltStateMachineOpen() {
	testOpts := &bbolt.Options{
		Timeout:         0,
		NoGrowSync:      false,
		NoFreelistSync:  false,
		FreelistType:    bbolt.FreelistMapType,
		ReadOnly:        false,
		InitialMmapSize: 0,
		PageSize:        0,
		NoSync:          false,
		OpenFile:        nil,
		Mlock:           false,
	}

	fsm := NewBBoltStateMachine(1, 1, bfsm.T().TempDir(), testOpts)

	// verify the core path constructs for humans are available
	dbPath := fsm.dbPath(true)
	bfsm.Assert().Contains(dbPath, "cluster-")
	bfsm.Assert().Contains(dbPath, "node-")
	bfsm.Assert().Contains(dbPath, ".db")

	// verify there's no ".db"
	dbRootPath := fsm.dbPath(false)
	bfsm.Assert().Contains(dbRootPath, "cluster-")
	bfsm.Assert().Contains(dbRootPath, "node-")
	bfsm.Assert().NotContains(dbRootPath, ".db")

	// open a blank database
	var index uint64
	var err error
	bfsm.Require().NotPanics(func() {
		index, err = fsm.Open(make(<-chan struct{}, 1))
	})
	bfsm.Require().Equal(uint64(0), index, "opening a brand new fsm requires a 0 index")
	bfsm.Require().NoError(err, "there must not be an error opening a brand new fsm")
	bfsm.Require().NoError(fsm.db.Close(), "there must not be an error when closing the the brand new database")

	fi, err := os.Lstat(bfsm.T().TempDir())
	val := fi.Mode().Perm()
	bfsm.Assert().NotEmpty(val)
	err = os.RemoveAll(bfsm.T().TempDir())
	bfsm.Require().NoError(err, "there must not be an error when deleting the test directory")

	err = os.MkdirAll(fsm.dbPath(false), os.FileMode(dbDirModeVal))
	bfsm.Require().NoError(err, "there must not be an error when creating the database path")

	// create a database with an existing index
	db, err := bbolt.Open(fsm.dbPath(true), os.FileMode(dbFileModeVal), testOpts)
	bfsm.Require().NoError(err, "there must not be an error when opening the test database")
	bfsm.Require().NotNil(db, "the test database must be opened")

	// set the monotonic key, monotonicLogKey, to be an arbitrary value
	testIndexVal := uint64(55)
	err = db.Update(func(tx *bbolt.Tx) error {
		internalBucket, err := tx.CreateBucketIfNotExists([]byte(monotonicLogBucket))
		if err != nil {
			return err
		}

		val := make([]byte, 8)
		binary.LittleEndian.PutUint64(val, testIndexVal)
		return internalBucket.Put([]byte(monotonicLogKey), val)
	})
	bfsm.Require().NoError(err, "there must not be an error when prepping the monotonic log")

	err = db.Close()
	bfsm.Require().NoError(err, "there must not be an error when closing the test database")

	bfsm.Require().NotPanics(func() {
		index, err = fsm.Open(make(<-chan struct{}, 1))
	})
	bfsm.Require().Equal(testIndexVal, index, fmt.Sprintf("the existing index must be %d", testIndexVal))
	bfsm.Require().NoError(err, "there must not be an error opening an existing fsm")
	bfsm.Require().NoError(fsm.db.Close(), "there must not be an error when closing the database manually")
}

func (bfsm *TestBBoltFsm) TestBBoltStateMachineClose() {
	testOpts := &bbolt.Options{
		Timeout:         0,
		NoGrowSync:      false,
		NoFreelistSync:  false,
		FreelistType:    bbolt.FreelistMapType,
		ReadOnly:        false,
		InitialMmapSize: 0,
		PageSize:        0,
		NoSync:          false,
		OpenFile:        nil,
		Mlock:           false,
	}

	fsm := NewBBoltStateMachine(1, 1, bfsm.T().TempDir(), testOpts)

	index, err := fsm.Open(make(<-chan struct{}))
	bfsm.Require().NoError(err, "there must not be an error when opening the database")
	bfsm.Require().Equal(uint64(0), index, "the index must equal as there are no records")

	err = fsm.Close()
	bfsm.Require().NoError(err, "there must not be an error while closing the database")
	bfsm.Assert().Nil(fsm.db, "the database should be dereferenced")
	bfsm.Require().Panics(func() {
		_ = fsm.db.View(func(tx *bbolt.Tx) error {
			return nil
		})
	}, "there must be a nil reference error when trying to access the database after it's been closed")
}

func (bfsm *TestBBoltFsm) TestBBoltStateMachineUpdate() {
	testOpts := &bbolt.Options{
		Timeout:         0,
		NoGrowSync:      false,
		NoFreelistSync:  false,
		FreelistType:    bbolt.FreelistMapType,
		ReadOnly:        false,
		InitialMmapSize: 0,
		PageSize:        0,
		NoSync:          false,
		OpenFile:        nil,
		Mlock:           false,
	}

	fsm := NewBBoltStateMachine(1, 1, bfsm.T().TempDir(), testOpts)

	index, err := fsm.Open(make(<-chan struct{}))
	bfsm.Require().NoError(err, "there must not be an error when opening the database")
	bfsm.Require().Equal(uint64(0), index, "the index must equal as there are no records")

	rootPrn := &PleiadesResourceName{
		Partition:    GlobalPartition,
		Service:      Pleiades,
		Region:       GlobalRegion,
		AccountId:    testAccountKey,
		ResourceType: Bucket,
		ResourceId:   "test-bucket",
	}

	var testKvps []db.KeyValue

	for i := 0; i < 3; i++ {
		kvp := db.KeyValue{
			Key: []byte(fmt.Sprintf("%s/test-key-%d", rootPrn.ToFsmRootPath("test-bucket"), i)),
			Value: []byte(fmt.Sprintf("test-value-%d", i)),
			CreateRevision: 0,
			ModRevision: 0,
			Version: 1,
			Lease: 0,
		}

		testKvps = append(testKvps, kvp)
	}

	testUpdates := make([]statemachine.Entry, 0)
	for idx := range testKvps {
		marshalled, err := testKvps[idx].MarshalVT()
		bfsm.Require().NoError(err, "there must not be an error marshalling the message")

		testUpdates = append(testUpdates, statemachine.Entry{
			Index:  uint64(idx),
			Cmd:    marshalled,
			Result: statemachine.Result{},
		})
	}

	var endingIndex []statemachine.Entry
	bfsm.Require().NotPanics(func() {
		endingIndex, err = fsm.Update(testUpdates)
	})
	bfsm.Require().NoError(err, "there must not be an error delivering updates")
	bfsm.Require().Equal(
		testUpdates[len(testUpdates)-1].Index,
		endingIndex[len(endingIndex)-1].Index,
		fmt.Sprintf("the ending index must be %d", testUpdates[len(testUpdates)-1].Index))
}

func (bfsm *TestBBoltFsm) TestPrepareSnapshot() {
	testOpts := &bbolt.Options{
		Timeout:         0,
		NoGrowSync:      false,
		NoFreelistSync:  false,
		FreelistType:    bbolt.FreelistMapType,
		ReadOnly:        false,
		InitialMmapSize: 0,
		PageSize:        0,
		NoSync:          false,
		OpenFile:        nil,
		Mlock:           false,
	}

	fsm := NewBBoltStateMachine(1, 1, bfsm.T().TempDir(), testOpts)
	empty, err := fsm.PrepareSnapshot()
	bfsm.Require().Empty(empty, "this must be a noop")
	bfsm.Require().NoError(err, "there must be no error")
}

func (bfsm *TestBBoltFsm) TestSnapshotLifecycle() {
	testOpts := &bbolt.Options{
		Timeout:         0,
		NoGrowSync:      false,
		NoFreelistSync:  false,
		FreelistType:    bbolt.FreelistMapType,
		ReadOnly:        false,
		InitialMmapSize: 0,
		PageSize:        0,
		NoSync:          false,
		OpenFile:        nil,
		Mlock:           false,
	}

	fsm := NewBBoltStateMachine(1, 1, bfsm.T().TempDir(), testOpts)
	index, err := fsm.Open(make(<-chan struct{}))
	bfsm.Require().NoError(err, "there must not be an error when opening the database")
	bfsm.Require().Equal(uint64(0), index, "the index must equal as there are no records")

	rootPrn := &PleiadesResourceName{
		Partition:    GlobalPartition,
		Service:      Pleiades,
		Region:       GlobalRegion,
		AccountId:    testAccountKey,
		ResourceType: Bucket,
		ResourceId:   "test-bucket",
	}

	var testKvps []db.KeyValue

	for i := 0; i < 3; i++ {
		kvp := db.KeyValue{
			Key: []byte(fmt.Sprintf("%s/test-key-%d", rootPrn.ToFsmRootPath("test-bucket"), i)),
			Value: []byte(fmt.Sprintf("test-value-%d", i)),
			CreateRevision: 0,
			ModRevision: 0,
			Version: 1,
			Lease: 0,
		}

		testKvps = append(testKvps, kvp)
	}

	testUpdates := make([]statemachine.Entry, 0)
	for idx := range testKvps {
		payload, err := testKvps[idx].MarshalVT()
		bfsm.Require().NoError(err, "there must not be an error encoding the message")

		testUpdates = append(testUpdates, statemachine.Entry{
			Index:  uint64(idx),
			Cmd:    payload,
			Result: statemachine.Result{},
		})
	}

	var endingIndex []statemachine.Entry
	bfsm.Require().NotPanics(func() {
		endingIndex, err = fsm.Update(testUpdates)
	})
	bfsm.Require().NoError(err, "there must not be an error delivering updates")
	bfsm.Require().Equal(
		testUpdates[len(testUpdates)-1].Index,
		endingIndex[len(endingIndex)-1].Index,
		fmt.Sprintf("the ending index must be %d", testUpdates[len(testUpdates)-1].Index))

	// todo (sienna): replace local buffer with net.Pipe() for better test
	var buf bytes.Buffer
	bfsm.Require().NotPanics(func() {
		err = fsm.SaveSnapshot(context.Background(), &buf, make(<-chan struct{}))
	}, "saving the snapshot to the buffer must not panic")
	bfsm.Require().NoError(err, "there must not be an error when saving the snapshot")

	bfsm.Require().NotPanics(func() {
		err = fsm.RecoverFromSnapshot(&buf, make(<-chan struct{}))
	}, "recovering the snapshot must not panic")
	bfsm.Require().NoError(err, "there must not be an error when recovering from a snapshot")

	dbPath := fsm.dbPath(true)
	bdb, err := bbolt.Open(dbPath, os.FileMode(dbFileModeVal), nil)
	bfsm.Require().NoError(err)

	// todo (sienna): fix this to ensure the finalKvp is stored, not just the value
	var target []byte
	bfsm.Require().NoError(bdb.Update(func(tx *bbolt.Tx) error {
		// rootPrn.ToFsmRootPath("test-bucket") + "/" + "test-key-2"
		bucketHierarchy := strings.Split(rootPrn.ToFsmRootPath("test-bucket")+"/"+"test-key-2", "/")[1:]
		parentBucketName := bucketHierarchy[0]
		childBucketNames := bucketHierarchy[1:]
		parentBucket := tx.Bucket([]byte(parentBucketName))
		resChan := make(chan []byte, 1)
		if err := keyOp(parentBucket, childBucketNames, make([]byte, 0), get, &resChan); err != nil {
			return err
		}
		target = <-resChan
		close(resChan)
		return nil
	}))

	finalKvp := db.KeyValue{}
	err = finalKvp.UnmarshalVT(target)
	bfsm.Require().NoError(err, "there must not be an error unmarshalling the key")

	testKey := testKvps[len(testKvps)-1].Key
	returnedKey := finalKvp.Key

	bfsm.Require().Equal(testKey, returnedKey, "the serialized result must match the initial value")
	bfsm.Require().Equal(testKvps[len(testKvps)-1].Version, finalKvp.Version, "the serialized result must match the initial value")
	bfsm.Require().Equal(testKvps[len(testKvps)-1].ModRevision, finalKvp.ModRevision, "the serialized result must match the initial value")
	bfsm.Require().Equal(testKvps[len(testKvps)-1].Lease, finalKvp.Lease, "the serialized result must match the initial value")
	bfsm.Require().Equal(testKvps[len(testKvps)-1].CreateRevision, finalKvp.CreateRevision, "the serialized result must match the initial value")

	testValue := testKvps[len(testKvps)-1].Value
	returnedValue := finalKvp.Value

	bfsm.Require().Equal(testValue, returnedValue, "the serialized result must match the initial value")
}

func (bfsm *TestBBoltFsm) TestLookup() {
	testOpts := &bbolt.Options{
		Timeout:         0,
		NoGrowSync:      false,
		NoFreelistSync:  false,
		FreelistType:    bbolt.FreelistMapType,
		ReadOnly:        false,
		InitialMmapSize: 0,
		PageSize:        0,
		NoSync:          false,
		OpenFile:        nil,
		Mlock:           false,
	}

	fsm := NewBBoltStateMachine(1, 1, bfsm.T().TempDir(), testOpts)
	index, err := fsm.Open(make(<-chan struct{}))
	bfsm.Require().NoError(err, "there must not be an error when opening the database")
	bfsm.Require().Equal(uint64(0), index, "the index must equal as there are no records")

	rootPrn := &PleiadesResourceName{
		Partition:    GlobalPartition,
		Service:      Pleiades,
		Region:       GlobalRegion,
		AccountId:    testAccountKey,
		ResourceType: Bucket,
		ResourceId:   "test-bucket",
	}

	var testKvps []db.KeyValue

	for i := 0; i < 3; i++ {
		kvp := db.KeyValue{
			Key: []byte(fmt.Sprintf("%s/test-key-%d", rootPrn.ToFsmRootPath("test-bucket"), i)),
			Value: []byte(fmt.Sprintf("test-value-%d", i)),
			CreateRevision: 0,
			ModRevision: 0,
			Version: 1,
			Lease: 0,
		}

		testKvps = append(testKvps, kvp)
	}

	testUpdates := make([]statemachine.Entry, 0)
	for idx := range testKvps {
		marshalled, err := testKvps[idx].MarshalVT()
		bfsm.Require().NoError(err, "there must not be an error marshalling the message")

		testUpdates = append(testUpdates, statemachine.Entry{
			Index:  uint64(idx),
			Cmd:    marshalled,
			Result: statemachine.Result{},
		})
	}

	var endingIndex []statemachine.Entry
	bfsm.Require().NotPanics(func() {
		endingIndex, err = fsm.Update(testUpdates)
	})
	bfsm.Require().NoError(err, "there must not be an error delivering updates")
	bfsm.Require().Equal(
		testUpdates[len(testUpdates)-1].Index,
		endingIndex[len(endingIndex)-1].Index,
		fmt.Sprintf("the ending index must be %d", testUpdates[len(testUpdates)-1].Index))

	val, err := fsm.Lookup(testUpdates[len(testUpdates)-1].Cmd)
	bfsm.Require().NoError(err, "there must not be an error when calling lookup")

	var casted db.KeyValue
	bfsm.Require().NotPanics(func() {
		casted = val.(db.KeyValue)
	}, "casting the lookup value must not panic")

	expectedValue := testKvps[len(testUpdates)-1].Value
	castedValue := casted.Value
	bfsm.Require().Equal(expectedValue, castedValue, "the found value must be identical")

	expectedKey := testKvps[len(testUpdates)-1].Key
	castedKey := casted.Key
	bfsm.Require().Equal(expectedKey, castedKey, "the found value must be identical")

	bfsm.Require().Equal(testKvps[len(testUpdates)-1].Lease, casted.Lease, "the found value must be identical")
	bfsm.Require().Equal(testKvps[len(testUpdates)-1].CreateRevision, casted.CreateRevision, "the found value must be identical")
	bfsm.Require().Equal(testKvps[len(testUpdates)-1].Version, casted.Version, "the found value must be identical")
	bfsm.Require().Equal(testKvps[len(testUpdates)-1].ModRevision, casted.ModRevision, "the found value must be identical")
}

func (bfsm *TestBBoltFsm) TestSync() {
	testOpts := &bbolt.Options{
		Timeout:         0,
		NoGrowSync:      false,
		NoFreelistSync:  false,
		FreelistType:    bbolt.FreelistMapType,
		ReadOnly:        false,
		InitialMmapSize: 0,
		PageSize:        0,
		NoSync:          false,
		OpenFile:        nil,
		Mlock:           false,
	}

	fsm := NewBBoltStateMachine(1, 1, bfsm.T().TempDir(), testOpts)
	index, err := fsm.Open(make(<-chan struct{}))
	bfsm.Require().NoError(err, "there must not be an error when opening the database")
	bfsm.Require().Equal(uint64(0), index, "the index must equal as there are no records")

	rootPrn := &PleiadesResourceName{
		Partition:    GlobalPartition,
		Service:      Pleiades,
		Region:       GlobalRegion,
		AccountId:    testAccountKey,
		ResourceType: Bucket,
		ResourceId:   "test-bucket",
	}

	var testKvps []db.KeyValue

	for i := 0; i < 3; i++ {
		kvp := db.KeyValue{
			Key: []byte(fmt.Sprintf("%s/test-key-%d", rootPrn.ToFsmRootPath("test-bucket"), i)),
			Value: []byte(fmt.Sprintf("test-value-%d", i)),
			CreateRevision: 0,
			ModRevision: 0,
			Version: 1,
			Lease: 0,
		}

		testKvps = append(testKvps, kvp)
	}

	testUpdates := make([]statemachine.Entry, 0)
	for idx := range testKvps {
		marshalled, err := testKvps[idx].MarshalVT()
		bfsm.Require().NoError(err, "there must not be an error marshalling the message")

		testUpdates = append(testUpdates, statemachine.Entry{
			Index:  uint64(idx),
			Cmd:    marshalled,
			Result: statemachine.Result{},
		})
	}

	var endingIndex []statemachine.Entry
	bfsm.Require().NotPanics(func() {
		endingIndex, err = fsm.Update(testUpdates)
	})
	bfsm.Require().NoError(err, "there must not be an error delivering updates")
	bfsm.Require().Equal(
		testUpdates[len(testUpdates)-1].Index,
		endingIndex[len(endingIndex)-1].Index,
		fmt.Sprintf("the ending index must be %d", testUpdates[len(testUpdates)-1].Index))

	bfsm.Require().NotPanics(func() {
		err = fsm.Sync()
	})
	bfsm.Require().NoError(err, "there must not be an error when syncing bbolt to disk")
}
