#!/bin/sh

#
# Copyright (c) 2022 Anthropos Labs, Inc.
#
# Licensed under the PolyForm Strict License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License here:
#  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
#ne/LICENSE
#

set -e

echo "generating config protocols"
capnp compile \
  -I$GOPATH/src/capnproto.org/go/capnp/std \
  -ogo:pkg \
  protocols/v1/config/*.capnp

echo "generating database protocols"
capnp compile \
  -I$GOPATH/src/capnproto.org/go/capnp/std \
  -ogo:pkg \
  protocols/v1/database/*.capnp
