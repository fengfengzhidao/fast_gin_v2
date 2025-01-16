package routers

import (
	"fast_gin/api"
	"fast_gin/api/user_api"
	"fast_gin/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup) {
	app := api.App.UserApi
	g.POST("users/login", middleware.LimitMiddleware(2), middleware.BindJsonMiddleware[user_api.LoginRequest], app.LoginView)
	g.GET("users", middleware.LimitMiddleware(10), middleware.AdminMiddleware, app.UserListView)
}
