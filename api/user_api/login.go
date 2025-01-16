package user_api

import "github.com/gin-gonic/gin"

func (UserApi) LoginView(c *gin.Context) {
	c.String(200, "登录")
	return
}
