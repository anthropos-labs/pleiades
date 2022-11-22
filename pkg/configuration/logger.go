/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package configuration

import (
	"os"

	"a13s.io/pleiades/pkg"
	zlog "github.com/rs/zerolog"
)

var (
	rootLogger zlog.Logger
)

func init() {
	rootLogger = zlog.New(os.Stdout).
		With().
		Str("sha", pkg.Sha).
		Timestamp().
		Logger().
		Level(zlog.InfoLevel)
}

func NewRootLogger() zlog.Logger {
	return rootLogger
}
