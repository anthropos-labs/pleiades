
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
	"math/rand"
	"time"
)

func RandomInt(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := r.Perm(max - min + 1)
	return p[min]
}
