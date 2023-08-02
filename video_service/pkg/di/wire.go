package di

import (
	"github.com/RohithER12/video_service/pkg/api"
	"github.com/RohithER12/video_service/pkg/api/service"
	"github.com/RohithER12/video_service/pkg/config"
	"github.com/RohithER12/video_service/pkg/db"
	"github.com/RohithER12/video_service/pkg/repository"
	"github.com/google/wire"
)

func InitializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(db.Initdb, repository.NewVideoRepo, service.NewVideoServer, api.NewgrpcServe)
	return &api.Server{}, nil
}
