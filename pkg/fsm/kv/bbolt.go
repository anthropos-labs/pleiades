/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package kv

import (
	"encoding/binary"
	"os"
	"time"

	"a13s.io/pleiades/api/v1/database"
	"github.com/cockroachdb/errors"
	"github.com/planetscale/vtprotobuf/codec/grpc"
	"github.com/rs/zerolog"
	"go.etcd.io/bbolt"
	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	encoding.RegisterCodec(grpc.Codec{})
}

const (
	monotonicLogBucket = "monotonic"
	monotonicLogKey    = "index"
	descriptorKey      = "_descriptor"
)

var (
	ErrInvalidAccount       = errors.New("invalid account id")
	ErrMissingAccountBucket = errors.New("account bucket not found")
	ErrInvalidBucketName    = errors.New("invalid bucket name")
	ErrEmptyBucketName      = errors.New("empty bucket name")
	ErrInvalidOwner         = errors.New("invalid owner")
	ErrKeyNotFound          = errors.New("key not found")
)

func newBboltStore(shardId, replicaId uint64, dbPath string, logger zerolog.Logger) (*bboltStore, error) {
	l := logger.With().Uint64("shard", shardId).Uint64("replica", replicaId).Logger()
	db, err := bbolt.Open(dbPath, os.FileMode(484), nil)
	if err != nil {
		l.Error().Err(err).Msg("can't open bbolt")
		return nil, err
	}

	b := &bboltStore{
		logger: l,
		db:     db,
	}
	return b, nil
}

type bboltStore struct {
	logger zerolog.Logger
	db     *bbolt.DB
}

func (b *bboltStore) close() error {
	b.logger.Debug().Msg("closing bbolt")
	return b.db.Close()
}

func (b *bboltStore) CreateAccountBucket(request *database.CreateAccountRequest) (*database.CreateAccountReply, error) {
	account := request.GetAccountId()
	if account == 0 {
		b.logger.Trace().Msg("empty account value")
		return &database.CreateAccountReply{}, ErrInvalidAccount
	}

	owner := request.GetOwner()
	if owner == "" {
		b.logger.Trace().Msg("empty owner value")
		return &database.CreateAccountReply{}, ErrInvalidOwner
	}

	now := timestamppb.Now()
	acctDescriptor := &database.AccountDescriptor{
		AccountId:   account,
		Owner:       owner,
		Created:     now,
		LastUpdated: now,
		BucketCount: 0,
		Buckets:     nil,
	}

	err := b.db.Update(func(tx *bbolt.Tx) error {
		accountBuf := make([]byte, 8)
		binary.LittleEndian.PutUint64(accountBuf, account)

		// open the account bucket
		accountBucket, err := tx.CreateBucket(accountBuf)
		if err != nil {
			b.logger.Error().Err(err).Msg("can't open account bucket")
			return errors.Wrap(err, "can't open account bucket")
		}

		// now store the descriptor, which will have updated values if necessary.
		_acctDescriptorPayload, err := acctDescriptor.MarshalVT()
		if err != nil {
			b.logger.Error().Err(err).Msg("can't marshal account descriptor")
			return err
		}

		err = accountBucket.Put([]byte(descriptorKey), _acctDescriptorPayload)

		return err
	})
	if err != nil {
		b.logger.Error().Err(err).Msg("can't create bucket")
	}

	return &database.CreateAccountReply{
		AccountDescriptor: acctDescriptor,
	}, err
}

func (b *bboltStore) GetAccountInfo(request *database.GetAccountDescriptorRequest) (*database.GetAccountDescriptorReply, error) {
	//TODO implement me
	panic("implement me")
}

func (b *bboltStore) DeleteAccountBucket(request *database.DeleteAccountRequest) (*database.DeleteAccountReply, error) {
	account := request.GetAccountId()
	if account == 0 {
		b.logger.Trace().Msg("empty account value")
		return &database.DeleteAccountReply{}, ErrInvalidAccount
	}

	owner := request.GetOwner()
	if owner == "" {
		b.logger.Trace().Msg("empty owner value")
		return &database.DeleteAccountReply{}, ErrInvalidOwner
	}

	err := b.db.Update(func(tx *bbolt.Tx) error {
		accountBuf := make([]byte, 8)
		binary.LittleEndian.PutUint64(accountBuf, account)

		// open the account bucket
		accountBucket := tx.Bucket(accountBuf)
		if accountBucket == nil {
			b.logger.Error().Msg("account bucket not found")
			return errors.Wrap(bbolt.ErrBucketNotFound, ErrMissingAccountBucket.Error())
		}
		// clear the reference
		accountBucket = nil

		err := tx.DeleteBucket(accountBuf)
		if err != nil {
			b.logger.Error().Err(err).Msg("can't delete account bucket")
			return errors.Wrap(err, "can't delete account bucket")
		}

		return nil
	})
	resp := &database.DeleteAccountReply{
		Ok: true,
	}

	if err != nil {
		b.logger.Error().Err(err).Msg("can't delete account bucket")
		resp.Ok = false
	}
	return resp, err
}

func (b *bboltStore) CreateBucket(request *database.CreateBucketRequest) (*database.CreateBucketReply, error) {
	account := request.GetAccountId()
	if account == 0 {
		b.logger.Trace().Msg("empty account value")
		return &database.CreateBucketReply{}, ErrInvalidAccount
	}

	newBucketName := request.GetName()
	if newBucketName == "" {
		b.logger.Trace().Msg("empty bucket name")
		return &database.CreateBucketReply{}, ErrEmptyBucketName
	}

	owner := request.GetOwner()
	if owner == "" {
		b.logger.Trace().Msg("empty owner value")
		return &database.CreateBucketReply{}, ErrInvalidOwner
	}

	now := timestamppb.Now()
	descriptor := &database.BucketDescriptor{
		Owner:       owner,
		Size:        0,
		KeyCount:    0,
		Created:     now,
		LastUpdated: now,
	}

	err := b.db.Update(func(tx *bbolt.Tx) error {
		accountBuf := make([]byte, 8)
		binary.LittleEndian.PutUint64(accountBuf, account)

		// open the account bucket
		accountBucket := tx.Bucket(accountBuf)
		if accountBucket == nil {
			b.logger.Error().Msg("account bucket doesn't exist")
			return errors.Wrap(bbolt.ErrBucketNotFound, "account bucket not found")
		}

		// get the account descriptor
		acctDescriptor := &database.AccountDescriptor{}
		_acctDescriptor := accountBucket.Get([]byte(descriptorKey))

		if err := proto.Unmarshal(_acctDescriptor, acctDescriptor); err != nil {
			b.logger.Error().Err(err).Msg("can't unmarshal account descriptor")
		}

		newBucket, err := accountBucket.CreateBucket([]byte(newBucketName))
		if err != nil {
			b.logger.Error().Err(err).Msg("can't create bucket")
			return errors.Wrap(err, "can't create bucket")
		}

		descriptorPayload, err := descriptor.MarshalVT()
		if err != nil {
			b.logger.Error().Err(err).Msg("can't marshal bucket descriptor")
			return errors.Wrap(err, "can't marshal bucket descriptor")
		}

		err = newBucket.Put([]byte(descriptorKey), descriptorPayload)

		acctDescriptor.Buckets = append(acctDescriptor.Buckets, newBucketName)
		acctDescriptor.BucketCount += 1

		_acctDescriptorPayload, err := acctDescriptor.MarshalVT()
		if err != nil {
			b.logger.Error().Err(err).Msg("can't marshal account descriptor")
			return errors.Wrap(err, "can't marshal account descriptor")
		}

		err = accountBucket.Put([]byte(descriptorKey), _acctDescriptorPayload)

		return errors.Wrap(err, "error updating account descriptor")
	})
	if err != nil {
		b.logger.Error().Err(err).Msg("can't create bucket")
	}

	return &database.CreateBucketReply{
		BucketDescriptor: descriptor,
	}, err
}

func (b *bboltStore) DeleteBucket(request *database.DeleteBucketRequest) (*database.DeleteBucketReply, error) {
	account := request.GetAccountId()
	if account == 0 {
		b.logger.Trace().Msg("empty account value")
		return &database.DeleteBucketReply{}, ErrInvalidAccount
	}

	targetBucketName := request.GetName()
	if targetBucketName == "" {
		b.logger.Trace().Msg("empty bucket targetBucketName")
		return &database.DeleteBucketReply{}, ErrInvalidBucketName
	}

	now := timestamppb.Now()

	err := b.db.Update(func(tx *bbolt.Tx) error {
		accountBuf := make([]byte, 8)
		binary.LittleEndian.PutUint64(accountBuf, account)

		// open the account bucket
		accountBucket := tx.Bucket(accountBuf)
		if accountBucket == nil {
			b.logger.Error().Msg("account bucket doesn't exist")
			return errors.Wrap(bbolt.ErrBucketNotFound, "account bucket not found")
		}

		// get the account descriptor
		acctDescriptor := &database.AccountDescriptor{}
		_acctDescriptor := accountBucket.Get([]byte(descriptorKey))

		if err := proto.Unmarshal(_acctDescriptor, acctDescriptor); err != nil {
			b.logger.Error().Err(err).Msg("can't unmarshal account descriptor")
		}

		// grab the size so we can emit the metric
		bucketSize := uint64(0)
		bucket := tx.Bucket([]byte(targetBucketName))
		if bucket != nil {
			bucketDesc := &database.BucketDescriptor{}
			bucketDescBytes := bucket.Get([]byte(descriptorKey))
			if err := proto.Unmarshal(bucketDescBytes, bucketDesc); err != nil {
				b.logger.Error().Err(err).Msg("can't unmarshal bucket descriptor")
				return errors.Wrap(err, "can't unmarshal bucket descriptor")
			}
			bucketSize = bucketDesc.GetSize()
		}

		// todo (sienna): emit metric here
		b.logger.Debug().Uint64("bucket-size", bucketSize).Msg("deleting bucket with size")

		err := accountBucket.DeleteBucket([]byte(targetBucketName))
		if err != nil {
			b.logger.Error().Err(err).Msg("can't delete bucket")
			return errors.Wrap(err, "can't delete bucket")
		}

		// todo (sienna): at least it's still O(1), but if we can optimize some of the branch prediction, that'd be great lol
		sz := len(acctDescriptor.Buckets)
		for idx := range acctDescriptor.Buckets {
			if acctDescriptor.Buckets[idx] == targetBucketName {
				if sz == 1 {
					acctDescriptor.Buckets = []string{}
					break
				}

				if idx == 0 {
					acctDescriptor.Buckets = acctDescriptor.Buckets[1:]
					break
				}

				// lazy girl's pop and shift
				// grab 0, remove 1, and then append 0 back on the end
				if idx == 1 {
					t := acctDescriptor.Buckets[0]
					acctDescriptor.Buckets = append(acctDescriptor.Buckets, acctDescriptor.Buckets[:2]...)
					acctDescriptor.Buckets = append(acctDescriptor.Buckets, t)
					break
				}

				// it's the last item
				if idx == sz-1 {
					acctDescriptor.Buckets = acctDescriptor.Buckets[:sz-1]
					break
				}

				// at least 3 buckets exist and it's not the first, second, or last one
				acctDescriptor.Buckets = acctDescriptor.Buckets[idx-1 : idx+1]
				break
			}
		}
		acctDescriptor.BucketCount--
		acctDescriptor.LastUpdated = now

		_acctDescriptorPayload, err := acctDescriptor.MarshalVT()
		if err != nil {
			b.logger.Error().Err(err).Msg("can't marshal account descriptor")
			return errors.Wrap(err, "can't marshal account descriptor")
		}

		err = accountBucket.Put([]byte(descriptorKey), _acctDescriptorPayload)

		return errors.Wrap(err, "error updating account descriptor")
	})

	rep := &database.DeleteBucketReply{
		Ok: true,
	}
	if err != nil {
		rep.Ok = false
		b.logger.Error().Err(err).Msg("can't create bucket")
	}

	return rep, err
}

func (b *bboltStore) GetKey(request *database.GetKeyRequest) (*database.GetKeyReply, error) {
	account := request.GetAccountId()
	if account == 0 {
		b.logger.Trace().Msg("empty account value")
		return &database.GetKeyReply{}, ErrInvalidAccount
	}

	bucketName := request.GetBucketName()
	if bucketName == "" {
		b.logger.Trace().Msg("empty bucket name")
		return &database.GetKeyReply{}, ErrInvalidBucketName
	}

	keyName := request.GetKey()
	if keyName == "" {
		b.logger.Trace().Msg("empty key identifier")
		return &database.GetKeyReply{}, errors.New("empty key identifier")
	}

	kvp := &database.KeyValue{}
	err := b.db.View(func(tx *bbolt.Tx) error {
		accountBuf := make([]byte, 8)
		binary.LittleEndian.PutUint64(accountBuf, account)

		accountBucket := tx.Bucket(accountBuf)
		if accountBucket == nil {
			b.logger.Error().Msg("account bucket doesn't exist")
			return errors.Wrap(bbolt.ErrBucketNotFound, "account bucket not found")
		}

		bucket := accountBucket.Bucket([]byte(bucketName))
		if bucket == nil {
			b.logger.Error().Msg("bucket not found")
			return errors.Wrap(bbolt.ErrBucketNotFound, "bucket not found")
		}

		payload := bucket.Get([]byte(keyName))
		if payload == nil {
			b.logger.Trace().Uint64("account-id", account).Str("bucket", keyName).Msg("key not found")
			return ErrKeyNotFound
		}

		return kvp.UnmarshalVT(payload)
	})

	if err != nil {
		b.logger.Error().Err(err).Uint64("account-id", account).Str("bucket", bucketName).Msg("error fetching key")
	}

	return &database.GetKeyReply{
		KeyValuePair: kvp,
	}, errors.Wrap(err, "error fetching key")
}

func (b *bboltStore) PutKey(request *database.PutKeyRequest) (*database.PutKeyReply, error) {
	account := request.GetAccountId()
	if account == 0 {
		b.logger.Trace().Msg("empty account value")
		return &database.PutKeyReply{}, ErrInvalidAccount
	}

	bucketName := request.GetBucketName()
	if bucketName == "" {
		b.logger.Error().Msg("empty bucket name")
		return &database.PutKeyReply{}, ErrInvalidBucketName
	}

	keyValuePair := request.GetKeyValuePair()
	if keyValuePair.GetKey() == "" {
		b.logger.Error().Msg("empty key identifier")
		return &database.PutKeyReply{}, errors.New("empty key identifier")
	}

	now := time.Now()

	err := b.db.Update(func(tx *bbolt.Tx) error {
		accountBuf := make([]byte, 8)
		binary.LittleEndian.PutUint64(accountBuf, account)

		accountBucket := tx.Bucket(accountBuf)
		if accountBucket == nil {
			b.logger.Error().Msg("account bucket doesn't exist")
			return errors.Wrap(bbolt.ErrBucketNotFound, "account bucket not found")
		}

		bucket := accountBucket.Bucket([]byte(bucketName))
		if bucket == nil {
			b.logger.Error().Msg("bucket not found")
			return errors.Wrap(bbolt.ErrBucketNotFound, "bucket not found")
		}

		// compare-and-swap and update some fields
		payload := bucket.Get([]byte(keyValuePair.GetKey()))
		if payload != nil {
			tmp := &database.KeyValue{}
			if err := tmp.UnmarshalVT(payload); err != nil {
				b.logger.Error().Err(err).Msg("key can't be unmarshalled, overwriting")
				goto overwrite
			}

			if tmp.Version > keyValuePair.GetVersion() {
				return errors.New("cannot overwrite existing key with older version")
			}

			keyValuePair.ModRevision = now.UnixMilli()
			keyValuePair.CreateRevision = tmp.CreateRevision
			keyValuePair.Version += tmp.GetVersion()
		}

	overwrite:

		payload, err := keyValuePair.MarshalVT()
		if err != nil {
			b.logger.Error().Err(err).Msg("can't marshal kvp")
			return errors.Wrap(err, "can't marshal kvp")
		}

		err = bucket.Put([]byte(keyValuePair.Key), payload)
		if err != nil {
			b.logger.Error().Err(err).Msg("can't put key")
		}

		// todo (sienna): emit metrics here

		return errors.Wrap(err, "can't store key")
	})

	if err != nil {
		b.logger.Error().Err(errors.Wrap(err, "error storing key")).Uint64("account-id", account).Str("bucket", bucketName).Msg("error storing key")
		// reset so we don't send data twice
		keyValuePair = &database.KeyValue{}
	}

	return &database.PutKeyReply{}, errors.Wrap(err, "error storing key")
}

func (b *bboltStore) DeleteKey(request *database.DeleteKeyRequest) (*database.DeleteKeyReply, error) {
	account := request.GetAccountId()
	if account == 0 {
		b.logger.Trace().Msg("empty account value")
		return &database.DeleteKeyReply{}, ErrInvalidAccount
	}

	bucketName := request.GetBucketName()
	if bucketName == "" {
		b.logger.Trace().Msg("empty bucket name")
		return &database.DeleteKeyReply{}, ErrInvalidBucketName
	}

	keyName := request.GetKey()
	if keyName == "" {
		b.logger.Trace().Msg("empty key identifier")
		return &database.DeleteKeyReply{}, errors.New("empty key identifier")
	}

	err := b.db.Update(func(tx *bbolt.Tx) error {
		accountBuf := make([]byte, 8)
		binary.LittleEndian.PutUint64(accountBuf, account)

		accountBucket := tx.Bucket(accountBuf)
		if accountBucket == nil {
			b.logger.Error().Msg("account bucket doesn't exist")
			return errors.Wrap(bbolt.ErrBucketNotFound, "account bucket not found")
		}

		bucket := accountBucket.Bucket([]byte(bucketName))
		if bucket == nil {
			b.logger.Error().Msg("bucket not found")
			return errors.Wrap(bbolt.ErrBucketNotFound, "bucket not found")
		}

		// compare-and-swap and update some fields
		err := bucket.Delete([]byte(keyName))

		// todo (sienna): emit metrics here

		return errors.Wrap(err, "can't delete key")
	})

	resp := &database.DeleteKeyReply{Ok: true}

	if err != nil {
		b.logger.Error().Err(err).Msg("can't delete key")
	}

	return resp, err
}

func (b *bboltStore) UpdateMonotonicLog(idx uint64) error {
	b.logger.Debug().Uint64("index", idx).Msg("updating monotonic log")
	return b.db.Update(func(tx *bbolt.Tx) error {
		internalBucket, err := tx.CreateBucketIfNotExists([]byte(monotonicLogBucket))
		if err != nil {
			return err
		}

		indexBuf := make([]byte, 8)
		binary.LittleEndian.PutUint64(indexBuf, idx)

		return internalBucket.Put([]byte(monotonicLogKey), indexBuf)
	})
}

func (b *bboltStore) GetMonotonicLogIndex() (uint64, error) {
	idx := uint64(0)

	err := b.db.Update(func(tx *bbolt.Tx) error {
		internalBucket, err := tx.CreateBucketIfNotExists([]byte(monotonicLogBucket))
		if err != nil {
			return err
		}

		val := internalBucket.Get([]byte(monotonicLogKey))
		if val == nil {
			return nil
		}

		idx = binary.LittleEndian.Uint64(val)
		return nil
	})

	return idx, err
}