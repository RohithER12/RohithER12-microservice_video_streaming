package routes

import (
	"github.com/RohithER12/api_gateway/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterVideoRoutes(api *gin.RouterGroup, videoHandler handler.VideoHandler) {
	api.POST("/upload", videoHandler.UploadVideo)
	api.GET("/stream/:video_id/:playlist", videoHandler.StreamVideo)
	api.GET("/video/all", videoHandler.FindAllVideo)
}
