package user_api

import (
	"fast_gin/middleware"
	"fast_gin/models"
	"fast_gin/service/common"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {
	var cr = middleware.GetBind[models.PageInfo](c)

	list, count, _ := common.QueryList(models.UserModel{}, common.QueryOption{
		PageInfo: cr,
		Likes:    []string{"username", "nickname"},
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}
