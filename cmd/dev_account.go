/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */
package cmd

import (
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "operations on accounts!",
}

var (
	accountId uint64
	accountOwner string
)

func init() {
	devCmd.AddCommand(accountCmd)
	accountCmd.PersistentFlags().Uint64Var(&accountId, "account-id", 0, "the account to operate on")
}
