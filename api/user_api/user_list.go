package user_api

import (
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {
	res.OkWithData("用户列表", c)
}
