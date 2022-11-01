/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package transactions

import (
	"context"

	kvstorev1 "a13s.io/api/kvstore/v1"
	"a13s.io/pleiades/pkg/server/runtime"
	"github.com/cockroachdb/errors"
	"github.com/lni/dragonboat/v3"
	dclient "github.com/lni/dragonboat/v3/client"
	"github.com/rs/zerolog"
)

var (
	_                        runtime.ITransactionManager = (*TransactionManager)(nil)
	ErrNilTransaction                                    = errors.New("cannot close an empty transaction")
	ErrUnupportedTransaction                             = errors.New("unsupported transaction type")
)

func NewTransactionManager(nh *dragonboat.NodeHost, logger zerolog.Logger) *TransactionManager {
	l := logger.With().Str("component", "session-manager").Logger()
	return &TransactionManager{l, nh, make(map[uint64]*dclient.Session)}
}

type TransactionManager struct {
	logger zerolog.Logger
	nh     *dragonboat.NodeHost

	// todo (sienna): there has to be a better/faster version of this
	sessionCache map[uint64]*dclient.Session
}

func (t *TransactionManager) CloseTransaction(ctx context.Context, transaction *kvstorev1.Transaction) error {
	t.logger.Debug().Uint64("shard", transaction.ShardId).Msg("closing transaction")

	cs, ok := t.sessionCache[transaction.GetClientId()]
	if !ok {
		return errors.New("transaction not found")
	}

	err := t.nh.SyncCloseSession(ctx, cs)
	if err != nil {
		t.logger.Error().Err(err).Msg("can't close transaction")
	}
	delete(t.sessionCache, cs.ClientID)

	return err
}

func (t *TransactionManager) Commit(ctx context.Context, transaction *kvstorev1.Transaction) *kvstorev1.Transaction {
	// nb (sienna): I know, I know. stop judging me.
	// is this hacky? yes.
	// does it work? yes.
	// is it the right thing to do now? no.
	// will it help later? yes.

	t.logger.Debug().Uint64("shard", transaction.ShardId).Msg("closing transaction")

	cs, ok := t.sessionCache[transaction.GetClientId()]
	if !ok {
		return &kvstorev1.Transaction{}
	}

	cs.ProposalCompleted()

	ta := csToTransaction(*cs)
	return ta
}

func (t *TransactionManager) GetNoOpTransaction(shardId uint64) *kvstorev1.Transaction {
	t.logger.Debug().Uint64("shard", shardId).Msg("getting noop transaction")
	cs := t.nh.GetNoOPSession(shardId)
	t.sessionCache[cs.ClientID] = cs
	return csToTransaction(*cs)
}

func (t *TransactionManager) GetTransaction(ctx context.Context, shardId uint64) (*kvstorev1.Transaction, error) {
	t.logger.Debug().Uint64("shard", shardId).Msg("getting transaction")
	cs, err := t.nh.SyncGetSession(ctx, shardId)
	if err != nil {
		t.logger.Error().Err(err).Uint64("shard", shardId).Msg("can't get transaction")
		return nil, err
	}

	t.sessionCache[cs.ClientID] = cs

	return csToTransaction(*cs), nil
}

func (t *TransactionManager) SessionFromClientId(clientId uint64) (*dclient.Session, bool) {
	sess, ok := t.sessionCache[clientId]
	return sess, ok
}

func csToTransaction(cs dclient.Session) *kvstorev1.Transaction {
	return &kvstorev1.Transaction{
		ShardId:       cs.ClusterID,
		ClientId:      cs.ClientID,
		TransactionId: cs.SeriesID,
		RespondedTo:   cs.RespondedTo,
	}
}