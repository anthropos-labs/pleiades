/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

//go:build mage
package main

import (
	"fmt"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

type Test mg.Namespace

// run all tests
func (Test) Cleanroom() error {
	fmt.Println("running all tests in cleanroom environment")
	mg.SerialDeps(Clean.Build, Clean.Cache, Install.Godeps)
	return sh.RunWithV(nil, "go", "test", "-v", "./...")
}

// run all tests
func (Test) All() error {
	fmt.Println("running all tests")
	return sh.RunWithV(nil, "go", "test", "-v", "./...")
}

// run server server tests
func (Test) Server() error {
	fmt.Println("running server tests")
	return sh.RunWithV(nil, "go", "test", "-v", "./pkg/server/...")
}

// run config tests
func (Test) Config() error {
	fmt.Println("running config tests")
	return sh.RunWithV(nil, "go", "test", "-v", "./pkg/configuration/...")
}

// run fsm tests
func (Test) FSM() error {
	fmt.Println("running fsm tests")
	return sh.RunWithV(nil, "go", "test", "-v", "./pkg/fsm/...")
}

// run routing tests
func (Test) Routing() error {
	fmt.Println("running routing tests")
	return sh.RunWithV(nil, "go", "test", "-v", "./pkg/routing/...")
}
