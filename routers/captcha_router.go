package routers

import (
	"fast_gin/api"
	"github.com/gin-gonic/gin"
)

func CaptchaRouter(g *gin.RouterGroup) {
	app := api.App.CaptchaApi
	g.GET("captcha/generate", app.GenerateView)
}
