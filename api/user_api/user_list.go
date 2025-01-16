package user_api

import "github.com/gin-gonic/gin"

func (UserApi) UserListView(c *gin.Context) {
	c.String(200, "用户列表")
}
