package routers

import (
	"fast_gin/api"
	"fast_gin/middleware"
	"github.com/gin-gonic/gin"
)

func ImageRouter(g *gin.RouterGroup) {
	app := api.App.ImageApi
	g.POST("images/upload", middleware.AuthMiddleware, app.UploadView)
}
