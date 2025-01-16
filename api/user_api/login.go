package user_api

import (
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
)

func (UserApi) LoginView(c *gin.Context) {
	res.OkWithData("用户登录", c)
	return
}
