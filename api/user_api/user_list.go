package user_api

import (
	"fast_gin/global"
	"fast_gin/middleware"
	"fast_gin/models"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {
	var cr = middleware.GetBind[models.PageInfo](c)
	var list = make([]models.UserModel, 0)
	query := global.DB.Where("")

	if cr.Key != "" {
		query.Where("username like ?", "%"+cr.Key+"%")
	}

	offset := (cr.Page - 1) * cr.Limit

	global.DB.Where(query).Limit(cr.Limit).Offset(offset).Order(cr.Order).Find(&list)

	var count int64
	global.DB.Where(query).Count(&count)

	res.OkWithList(list, count, c)
}
