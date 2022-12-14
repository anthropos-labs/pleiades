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
	"time"

	kvstorev1 "a13s.io/api/kvstore/v1"
	"a13s.io/api/kvstore/v1/kvstorev1connect"
	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

// kvGetCmd represents the kvGet command
var kvPutCmd = &cobra.Command{
	Use:   "put",
	Short: "put a key",
	Run:   putKey,
}

func init() {
	kvCmd.AddCommand(kvPutCmd)

	kvPutCmd.PersistentFlags().BytesBase64VarP(&payload, "value", "v", []byte{}, "a base64 encoded value")
	kvPutCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "key to use")
	kvPutCmd.PersistentFlags().Uint32VarP(&keyVersion, "version", "n", 0, "key version")
}

var (
	keyVersion uint32
)

func putKey(cmd *cobra.Command, args []string) {
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

	now := time.Now().UnixMilli()
	descriptor, err := client.PutKey(context.Background(), connect.NewRequest(&kvstorev1.PutKeyRequest{
		AccountId:  accountId,
		BucketName: bucketName,
		KeyValuePair: &kvstorev1.KeyValue{
			Key:            []byte(key),
			CreateRevision: now,
			ModRevision:    now,
			Version:        keyVersion,
			Value:          payload,
			Lease:          0,
		}}))
	if err != nil {
		logger.Fatal().Err(err).Msg("can't put key")
	}

	print(protojson.Format(descriptor.Msg))
}
