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
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSchema(t *testing.T) {
	suite.Run(t, new(schemaTestSuite))
}

type schemaTestSuite struct {
	suite.Suite
}
