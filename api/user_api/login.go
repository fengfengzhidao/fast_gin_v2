package user_api

import (
	"fast_gin/utils/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

func (UserApi) LoginView(c *gin.Context) {
	var cr LoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	fmt.Println(cr)
	res.OkWithData("用户登录", c)
	return
}
