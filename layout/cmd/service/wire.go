// +build wireinject

package main

import (
	"github.com/dch1228/gobestpractices/layout/internal/biz"
	"github.com/dch1228/gobestpractices/layout/internal/conf"
	"github.com/dch1228/gobestpractices/layout/internal/data"
	"github.com/dch1228/gobestpractices/layout/internal/server"
	"github.com/dch1228/gobestpractices/layout/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
)

func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, trace.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
