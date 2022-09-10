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
	"github.com/cockroachdb/errors"
)

const (
	systemShardStart = 1
	systemShardStop = 100
)

var (
	errNilTransaction = errors.New("cannot close an empty transaction")
	errUnupportedTransaction = errors.New("unsupported transaction type")
	ErrSystemShardRange = errors.New("shardId is within system shard range")
)