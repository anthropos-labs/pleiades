
/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package config

import (
	"testing"

	"a13s.io/pleiades/pkg/protocols/v1/config"
	"a13s.io/pleiades/pkg/utils"
	"capnproto.org/go/capnp/v3/server"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

func TestRegistry(t *testing.T) {
	suite.Run(t, new(RegistryTests))
}

type RegistryTests struct {
	suite.Suite
	logger   zerolog.Logger
	registry *Registry
}

// implement BeforeTest interface
func (s *RegistryTests) SetupTest() {
	s.logger = utils.NewTestLogger(s.T())

	var err error
	s.registry, err = NewRegistry(s.logger)
	s.Require().NoError(err, "there must not be an error creating the Registry")
}

func (s *RegistryTests) TestGet() {
	srv := &server.Server{}
	err := s.registry.PutServer(config.ServiceType_Type_test, srv)
	s.Require().NoError(err, "there must not be an error putting the value")

	value, _ := s.registry.GetServer(config.ServiceType_Type_test)
	s.Assert().Equal(srv, value)
}

func (s *RegistryTests) TestPut() {
	srv := &server.Server{}
	err := s.registry.PutServer(config.ServiceType_Type_test, srv)
	s.Require().NoError(err, "there must not be an error putting the value")

	value, _ := s.registry.GetServer(config.ServiceType_Type_test)
	s.Assert().Equal(srv, value)
}
