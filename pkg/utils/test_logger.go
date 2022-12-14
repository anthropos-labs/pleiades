/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package utils

import (
	"testing"

	"github.com/rs/zerolog"
)

func NewTestLogger(t *testing.T) zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	testWriter := zerolog.NewTestWriter(t)
	return zerolog.New(testWriter)
}

func NewFuzzLogger(t *testing.F) zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	testWriter := zerolog.NewTestWriter(t)
	return zerolog.New(testWriter)
}
