package user_api

import (
	"fast_gin/middleware"
	"fast_gin/utils/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {
	fmt.Println(middleware.GetAuth(c))
	res.OkWithData("用户列表", c)
}
