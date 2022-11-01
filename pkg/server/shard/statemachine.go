/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package shard

import (
	"a13s.io/pleiades/pkg/server/runtime"
	"github.com/cockroachdb/errors"
)

const (
	testStateMachineType  runtime.StateMachineType = 0
	BBoltStateMachineType runtime.StateMachineType = 1
)

var (
	ErrUnsupportedStateMachine = errors.New("state machine type is unsupported")
)
