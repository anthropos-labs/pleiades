#!/bin/sh

#
# Copyright (c) 2022 Anthropos Labs, Inc.
#
# Licensed under the PolyForm Strict License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License here:
#  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
#

mkdir bin

SHA=$(git rev-parse --short HEAD)
export SHA

go build \
  -ldflags "-X pkg.Sha=${SHA} -X pkg.GoVersion=$(go version) -X pkg.Version='22.6.30-dev.1'" \
  -o bin/pleiades-"${SHA}" \
  main.go