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
	"time"

	"a13s.io/pleiades/api/v1/database"
	"a13s.io/pleiades/pkg/configuration"
	"a13s.io/pleiades/pkg/server"
	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		log.Logger.Fatal().Err(err).Msg("can't parse flags")
	}

	var logger zerolog.Logger
	if debug {
		logger = configuration.NewRootLogger().Level(zerolog.DebugLevel)
	} else {
		logger = configuration.NewRootLogger()
	}
	logger = logger.With().Uint64("account-id", accountId).Str("bucket", bucketName).Logger()

	if accountId == 0 {
		logger.Fatal().Msg("account id cannot be zero")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Millisecond)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logger.Fatal().Err(err).Msg("error dialing server")
	}

	client := server.NewKVStoreServiceClient(conn)

	logger.Info().Str("key", key).Msg("getting key")
	descriptor, err := client.GetKey(context.Background(), &database.GetKeyRequest{
		AccountId:  accountId,
		BucketName: bucketName,
		Key:        key,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("can't delete bucket")
	}

	print(proto.MarshalTextString(descriptor))
}
