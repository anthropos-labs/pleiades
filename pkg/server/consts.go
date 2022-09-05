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
	"fmt"
	"math/rand"
)

const (
	testServerPortStart int = 8000
	testServerPortStop  int = 9000
	headerSize int = 10
)

type StreamState uint32

const (
	Valid StreamState = 0
	Invalid StreamState = 1
	StreamStart StreamState = 2
	StreamContinue StreamState = 3
	StreamNoLongerValid StreamState = 4
	StreamEnd StreamState = 5
	InvalidMessageChecksum StreamState = 6
	SystemError StreamState = 7
)

func testServerAddr() string {
	testPort := rand.Intn(testServerPortStop-testServerPortStart) + testServerPortStart
	return fmt.Sprintf("localhost:%d", testPort)
}
