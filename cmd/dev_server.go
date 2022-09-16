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
	"net"
	"os"
	"sync"
	"time"

	"a13s.io/pleiades/pkg/configuration"
	"a13s.io/pleiades/pkg/server"
	"a13s.io/pleiades/pkg/utils"
	dconfig "github.com/lni/dragonboat/v3/config"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run a development server",
	Long: `runs a development server.

it will boot with 256 predefined shards, configured in 
insecure mode, and will generally be buggy. it will run
the latest and greatest, which means it may or may not 
be usable for consuming applications. there may be unversioned
changes in this command which are not available as part of
the cloud offering. this command is unsupported beyond 
filing bugs against it the team may or may not get to

DO NOT USE THIS IN PRODUCTION`,
	Run: startServer,
}

var reset = true

func init() {
	devCmd.AddCommand(serverCmd)

	serverCmd.LocalFlags().StringVar(&serverConfig.Datastore.BasePath, "datastore-base-path", serverConfig.Datastore.BasePath, "set the default base directory")
	serverCmd.LocalFlags().StringVar(&serverConfig.Datastore.LogDir, "datastore-logs-dir", serverConfig.Datastore.LogDir, "logs for state machines")
	serverCmd.LocalFlags().StringVar(&serverConfig.Datastore.DataDir, "datastore-data-dir", serverConfig.Datastore.DataDir, "data files for state machines")

	serverCmd.LocalFlags().Uint64Var(&serverConfig.Host.DeploymentId, "deployment-id", serverConfig.Host.DeploymentId, "deployment id for this host")
	serverCmd.LocalFlags().StringVar(&serverConfig.Host.GrpcListenAddress, "grpc-listen-addr", serverConfig.Host.GrpcListenAddress, "grpc listen address")
	serverCmd.LocalFlags().StringVar(&serverConfig.Host.ListenAddress, "listen-addr", serverConfig.Host.ListenAddress, "listen address")
	serverCmd.LocalFlags().BoolVar(&serverConfig.Host.NotifyCommit, "notify-commit", serverConfig.Host.NotifyCommit, "enable notification on commit")
	serverCmd.LocalFlags().Uint64Var(&serverConfig.Host.Rtt, "rtt", serverConfig.Host.DeploymentId, "average round-trip-time in milliseconds")
}

func startServer(cmd *cobra.Command, args []string) {
	var logger zerolog.Logger
	if debug {
		logger = configuration.NewRootLogger().Level(zerolog.DebugLevel)
	} else {
		logger = configuration.NewRootLogger().Level(zerolog.DebugLevel)
	}

	logger.Info().Msg("hello from boulder")

	err := cmd.Flags().Parse(args)
	if err != nil {
		logger.Fatal().Err(err).Msg("can't parse flags")
	}

	if debug {

	}

	nhc := dconfig.NodeHostConfig{
		DeploymentID:   config.Server.Host.DeploymentId,
		WALDir:         config.Server.Datastore.LogDir,
		NodeHostDir:    config.Server.Datastore.DataDir,
		RTTMillisecond: config.Server.Host.Rtt,
		RaftAddress:    config.Server.Host.ListenAddress,
		EnableMetrics:  true,
		NotifyCommit:   config.Server.Host.NotifyCommit,
	}

	if config.Server.Host.MutualTLS {
		nhc.MutualTLS = config.Server.Host.MutualTLS
		nhc.CAFile = config.Server.Host.CaFile
		nhc.CertFile = config.Server.Host.CertFile
		nhc.KeyFile = config.Server.Host.KeyFile
	}

	if reset {
		err := os.RemoveAll(config.Server.Datastore.LogDir)
		if err != nil {
			logger.Fatal().Err(err).Str("dir", config.Server.Datastore.LogDir).Msg("can't remove directory")
		}
		err = os.RemoveAll(config.Server.Datastore.DataDir)
		if err != nil {
			logger.Fatal().Err(err).Str("dir", config.Server.Datastore.DataDir).Msg("can't remove directory")
		}
	}

	var opts []grpc.ServerOption

	gServer := grpc.NewServer(opts...)

	s, err := server.New(nhc, gServer, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("can't create pleiades server")
	}

	var wg sync.WaitGroup
	// shardLimit+1
	for i := uint64(1); i < 257; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			err = s.GetRaftShardManager().NewShard(i, i*257, server.BBoltStateMachineType, 300*time.Millisecond)
		}()
		utils.Wait(100 * time.Millisecond)
	}
	wg.Wait()

	logger.Debug().Msg("state machines finished, starting server")

	listen, err := net.Listen("tcp", config.Server.Host.GrpcListenAddress)
	if err != nil {
		logger.Fatal().Err(err).Str("listen-addr", config.Server.Host.GrpcListenAddress).Msg("can't start listener")
	}

	err = gServer.Serve(listen)
	if err != nil {
		logger.Fatal().Err(err).Msg("can't run grpc server")
	}

	s.Stop()
}