//go:build wireinject
// +build wireinject

package di

import (
	"github.com/RohithER12/api_gateway/pkg/api"
	"github.com/RohithER12/api_gateway/pkg/api/handler"
	"github.com/RohithER12/api_gateway/pkg/client"
	"github.com/RohithER12/api_gateway/pkg/config"
	"github.com/google/wire"
)

func InitializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(client.InitClient, client.NewVideoClient, handler.NewVideoHandler, api.NewServeHTTP)
	return &api.Server{}, nil
}
