// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package api

import (
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/rskvp/necode/proto/gen/necode/operatorservice"
	"github.com/rskvp/necode/proto/gen/necode/workflowservice"
	"github.com/rskvp/necode/services/chore/auth"
	"github.com/rskvp/necode/services/chore/config"
	"github.com/rskvp/necode/services/chore/rpc"
	"github.com/rskvp/necode/services/chore/version"
)

type Auth struct {
	Enabled bool
	Options []string
}

type CodecResponse struct {
	Endpoint                   string
	PassAccessToken            bool
	IncludeCredentials         bool
	DecodeEventHistoryDownload bool
}

type SettingsResponse struct {
	Auth                        *Auth
	DefaultNamespace            string
	ShowTemporalSystemNamespace bool
	FeedbackURL                 string
	NotifyOnNewVersion          bool
	Codec                       *CodecResponse
	Version                     string
	DisableWriteActions         bool
	WorkflowTerminateDisabled   bool
	WorkflowCancelDisabled      bool
	WorkflowSignalDisabled      bool
	WorkflowResetDisabled       bool
	BatchActionsDisabled        bool
	HideWorkflowQueryErrors     bool
}

func TemporalAPIHandler(cfgProvider *config.ConfigProviderWithRefresh, apiMiddleware []Middleware) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := auth.ValidateAuthHeaderExists(c, cfgProvider)
		if err != nil {
			return err
		}

		cfg, err := cfgProvider.GetConfig()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		tls, err := rpc.CreateTLSConfig(cfg.TemporalGRPCAddress, &cfg.TLS)
		if err != nil {
			fmt.Printf("unable to read TLS configs: %s", err)
		}
		conn := rpc.CreateGRPCConnection(cfg.TemporalGRPCAddress, tls)
		defer conn.Close()

		mux, err := getTemporalClientMux(c, conn, apiMiddleware)
		if err != nil {
			return err
		}

		mux.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func GetSettings(cfgProvider *config.ConfigProviderWithRefresh) func(echo.Context) error {
	return func(c echo.Context) error {
		cfg, err := cfgProvider.GetConfig()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		var options []string
		if len(cfg.Auth.Providers) != 0 {
			authProviderCfg := cfg.Auth.Providers[0].Options
			for k := range authProviderCfg {
				options = append(options, k)
			}
		}

		settings := &SettingsResponse{
			Auth: &Auth{
				Enabled: cfg.Auth.Enabled,
				Options: options,
			},
			DefaultNamespace:            cfg.DefaultNamespace,
			ShowTemporalSystemNamespace: cfg.ShowTemporalSystemNamespace,
			FeedbackURL:                 cfg.FeedbackURL,
			NotifyOnNewVersion:          cfg.NotifyOnNewVersion,
			Codec: &CodecResponse{
				Endpoint:                   cfg.Codec.Endpoint,
				PassAccessToken:            cfg.Codec.PassAccessToken,
				IncludeCredentials:         cfg.Codec.IncludeCredentials,
				DecodeEventHistoryDownload: cfg.Codec.DecodeEventHistoryDownload,
			},
			Version:                   version.UIVersion,
			DisableWriteActions:       cfg.DisableWriteActions,
			WorkflowTerminateDisabled: cfg.WorkflowTerminateDisabled,
			WorkflowCancelDisabled:    cfg.WorkflowCancelDisabled,
			WorkflowSignalDisabled:    cfg.WorkflowSignalDisabled,
			WorkflowResetDisabled:     cfg.WorkflowResetDisabled,
			BatchActionsDisabled:      cfg.BatchActionsDisabled,
			HideWorkflowQueryErrors:   cfg.HideWorkflowQueryErrors,
		}

		return c.JSON(http.StatusOK, settings)
	}
}

func getTemporalClientMux(c echo.Context, temporalConn *grpc.ClientConn, apiMiddleware []Middleware) (*runtime.ServeMux, error) {
	var muxOpts []runtime.ServeMuxOption

	tMux := runtime.NewServeMux(
		append(muxOpts)..., // This is necessary to get error details properly

	)

	if wfErr := workflowservice.RegisterWorkflowServiceHandler(context.Background(), tMux, temporalConn); wfErr != nil {
		return nil, wfErr
	}

	if opErr := operatorservice.RegisterOperatorServiceHandler(context.Background(), tMux, temporalConn); opErr != nil {
		return nil, opErr
	}
	return tMux, nil
}
