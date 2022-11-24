/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package runtime

// Logger logs the flow logs
type Logger interface {
	// Configure configure a logger with flowname and requestID
	Configure(flowName string, requestId string)
	// Init initialize a logger
	Init() error
	// Log logs a flow log
	Log(str string)
}
