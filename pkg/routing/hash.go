
/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package routing

import (
	"bytes"

	"github.com/dgryski/go-farm"
)

type farmHash struct {
	buf bytes.Buffer
}

func (f *farmHash) Write(p []byte) (n int, err error) {
	return f.buf.Write(p)
}

func (f *farmHash) Reset() {
	f.buf.Reset()
}

func (f *farmHash) Sum64() uint64 {
	return farm.Hash64(f.buf.Bytes())
}
