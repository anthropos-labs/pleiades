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
	vip "github.com/spf13/viper"
)

var (
	viper *vip.Viper
)


func init() {
	viper = vip.New()
}

func Get() *vip.Viper {
	if viper == nil {
		print("here")
		viper = vip.New()
	}
	return viper
}
