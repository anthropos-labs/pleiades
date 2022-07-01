#!/bin/bash
#
# Copyright (c) 2022 Anthropos Labs, Inc.
#
# Licensed under the PolyForm Strict License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License here:
#  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
#

[ -e vars.yaml ] && rm -- vars.yaml

for key in $(vault kv list -format=json concourse/main/pleiades | jq -r '.[]'); do
  echo "rendering ${key}"
  vault kv get -format=yaml concourse/main/pleiades/"${key}" | \
    KEY=$key yq '{env(KEY): .data.data}' >> vars.yaml
done
