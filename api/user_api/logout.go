package user_api

import (
	"fast_gin/global"
	"fast_gin/service/redis_ser"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
)

func (UserApi) LogoutView(c *gin.Context) {
	token := c.GetHeader("token")
	if global.Redis == nil {
		res.OkWithMsg("注销成功", c)
		return
	}
	redis_ser.Logout(token)
	res.OkWithMsg("注销成功", c)
}
