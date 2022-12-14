
/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package servicegraph

import (
	"gonum.org/v1/gonum/graph"
)

type LifecycleServiceType int

const (
	// TransientServiceType is used for things like RPC requests
	TransientServiceType LifecycleServiceType = 0
	// ScopedServiceType is used for things like individual processes
	ScopedServiceType LifecycleServiceType = 1
	// SingletonServiceType is used for globally unique things
	SingletonServiceType LifecycleServiceType = 2
)

type Service interface {
	graph.Node
	SetNodeID(nid int64)
	GetServiceName() string
	GetServiceType() LifecycleServiceType
	GetDependencies() []Service
	PrepareToRun() error
	ReadyToRun() bool
	IsRunning() bool
	Start(retry bool) error
	Stop(retry, force bool) error
}
