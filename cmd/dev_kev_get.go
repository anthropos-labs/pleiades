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
	"context"
	"net/http"

	kvstorev1 "a13s.io/api/kvstore/v1"
	"a13s.io/api/kvstore/v1/kvstorev1connect"
	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

// kvGetCmd represents the kvGet command
var kvGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a key",
	Run:   getKey,
}

func init() {
	kvCmd.AddCommand(kvGetCmd)

	kvGetCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "key to look for")
}

func getKey(cmd *cobra.Command, args []string) {
	err := cmd.Flags().Parse(args)
	if err != nil {
		log.Fatal().Err(err).Msg("can't parse flags")
	}

	logger := setupLogger(cmd, args)
	logger = logger.With().Uint64("account-id", accountId).Str("bucket", bucketName).Logger()

	if accountId == 0 {
		logger.Fatal().Msg("account id cannot be zero")
	}

	client := kvstorev1connect.NewKvStoreServiceClient(http.DefaultClient, "http://localhost:8080")

	descriptor, err := client.GetKey(context.Background(), connect.NewRequest(&kvstorev1.GetKeyRequest{
		AccountId:  accountId,
		BucketName: bucketName,
		Key:        []byte(key),
	}))
	if err != nil {
		logger.Fatal().Err(err).Msg("can't delete bucket")
	}

	print(protojson.Format(descriptor.Msg))
}
