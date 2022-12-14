/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package server

import (
	"context"
	"net/http"

	"a13s.io/api/kvstore/v1/kvstorev1connect"
	"a13s.io/api/raft/v1/raftv1connect"
	"a13s.io/pleiades/pkg/server/eventing"
	"a13s.io/pleiades/pkg/server/runtime"
	"a13s.io/pleiades/pkg/server/serverutils"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/lni/dragonboat/v3"
	dconfig "github.com/lni/dragonboat/v3/config"
	dlog "github.com/lni/dragonboat/v3/logger"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func init() {
	dlog.SetLoggerFactory(serverutils.DragonboatLoggerFactory)
}

// singletons
var (
	httpServer *http.Server
	nodeHost *dragonboat.NodeHost
)

type HttpServeMuxBuilderParams struct {
	fx.In

	Logger   zerolog.Logger
	Handlers []runtime.ServiceHandler `group:"routes"`
}

type HttpServeMuxBuilderResults struct {
	fx.Out

	Mux *http.ServeMux
}

func NewHttpServeMux(params HttpServeMuxBuilderParams) HttpServeMuxBuilderResults {
	mux := http.NewServeMux()

	// add grpc reflection for grpcurl and other tools
	reflector := grpcreflect.NewStaticReflector(
		kvstorev1connect.KvStoreServiceName,
		raftv1connect.HostServiceName)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	for _, route := range params.Handlers {
		params.Logger.Debug().Str("path", route.Path()).Msg("registering handler")
		mux.Handle(route.Path(), route)
	}
	return HttpServeMuxBuilderResults{Mux: mux}
}

type HttpServerBuilderParams struct {
	fx.In

	Logger zerolog.Logger
	Config *viper.Viper
	Mux    *http.ServeMux
}

type HttpServerBuilderResults struct {
	fx.Out

	Server *http.Server
}

func NewHttpServer(lc fx.Lifecycle, params HttpServerBuilderParams) HttpServerBuilderResults {
	httpServer = &http.Server{
		Addr:    params.Config.GetString("server.host.grpcListenAddress"),
		Handler: h2c.NewHandler(params.Mux, &http2.Server{}),
	}

	params.Logger.Info().Str("grpc-addr", params.Config.GetString("server.host.grpcListenAddress")).Msg("http listen address")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go httpServer.ListenAndServe()
			params.Logger.Info().Msg("started http server")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})
	return HttpServerBuilderResults{Server: httpServer}
}

type NodeHostBuilderParams struct {
	fx.In

	Logger         zerolog.Logger
	NodeHostConfig dconfig.NodeHostConfig
	Server         *eventing.Server
}

func NewNodeHost(lc fx.Lifecycle, params NodeHostBuilderParams) (*dragonboat.NodeHost, error) {
	handler, err := params.Server.GetRaftSystemEventListener()
	if err != nil {
		params.Logger.Error().Err(err).Msg("can't build raft system listeners")
		return nil, err
	}

	params.NodeHostConfig.SystemEventListener = handler
	params.NodeHostConfig.RaftEventListener = handler

	nodeHost, err = dragonboat.NewNodeHost(params.NodeHostConfig)
	if err != nil {
		params.Logger.Error().Err(err).Msg("can't build node host")
	}

	// dragonboat starts itself when New() is created, this is purely for the startup sequence
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		return nil
	}})

	return nodeHost, err
}

// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group.
func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(runtime.ServiceHandler)),
		fx.ResultTags(`group:"routes"`),
	)
}
