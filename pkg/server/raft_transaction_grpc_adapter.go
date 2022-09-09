/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package server

import (
	"context"

	"a13s.io/pleiades/api/v1/database"
	"github.com/rs/zerolog"
)

var _ TransactionsServer = (*raftTransactionGrpcAdapter)(nil)

type raftTransactionGrpcAdapter struct {
	logger zerolog.Logger
	transactionManager ITransactionManager
}

func (r *raftTransactionGrpcAdapter) NewTransaction(ctx context.Context, request *database.NewTransactionRequest) (*database.NewTransactionReply, error) {
	if request.GetClusterId() <= systemShardStop {
		return &database.NewTransactionReply{}, ErrSystemShardRange
	}
	//r.transactionManager.GetTransaction()

	//TODO implement me
	panic("implement me")
}

func (r *raftTransactionGrpcAdapter) CloseSession(ctx context.Context, request *database.CloseTransactionRequest) (*database.CloseTransactionReply, error) {
	//TODO implement me
	panic("implement me")
}

func (r *raftTransactionGrpcAdapter) mustEmbedUnimplementedTransactionsServer() {
	//TODO implement me
	panic("implement me")
}

