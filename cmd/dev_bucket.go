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

// bucketCmd represents the bucket command
var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "operations on buckets!",
}

var (
	bucketName string
)

func init() {
	devCmd.AddCommand(bucketCmd)
	bucketCmd.PersistentFlags().Uint64Var(&accountId, "account-id", 0, "the account to operate on")
}
