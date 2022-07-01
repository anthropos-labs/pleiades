/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package blaze

import (
	"fmt"
	"math/rand"
)

const (
	testServerPortStart int = 8000
	testServerPortStop  int = 9000
)

func testServerAddr() string {
	testPort := rand.Intn(testServerPortStop-testServerPortStart) + testServerPortStart
	return fmt.Sprintf("localhost:%d", testPort)
}
