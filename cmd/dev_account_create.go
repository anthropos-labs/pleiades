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

	kvstorev1 "a13s.io/api/kvstore/v1"
	"a13s.io/pleiades/pkg/configuration"
	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// accountCreateCmd represents the accountCreate command
var accountCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create an account",
	Run:   createAccount,
}

func init() {
	accountCmd.AddCommand(accountCreateCmd)
	accountCreateCmd.PersistentFlags().StringVar(&accountOwner, "owner", "", "the email owning the account")
}

func createAccount(cmd *cobra.Command, args []string) {
	err := cmd.Flags().Parse(args)
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("can't parse flags")
	}

	var logger zerolog.Logger
	if debug {
		logger = configuration.NewRootLogger().Level(zerolog.DebugLevel)
	} else {
		logger = configuration.NewRootLogger().Level(zerolog.DebugLevel)
	}

	if accountId == 0 {
		logger.Fatal().Msg("account id cannot be zero")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Millisecond)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logger.Fatal().Err(err).Msg("error dialing server")
	}

	client := kvstorev1.NewKvStoreServiceClient(conn)

	logger.Info().Uint64("account-id", accountId).Msg("creating account")
	descriptor, err := client.CreateAccount(context.Background(), &kvstorev1.CreateAccountRequest{AccountId: accountId, Owner: accountOwner})
	if err != nil {
		logger.Fatal().Err(err).Uint64("account-id", accountId).Msg("can't create account")
	}

	print(proto.MarshalTextString(descriptor))
}
