package user_api

import (
	"fast_gin/global"
	"fast_gin/middleware"
	"fast_gin/models"
	"fast_gin/utils/captcha"
	"fast_gin/utils/jwts"
	"fast_gin/utils/pwd"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Username    string `json:"username" binding:"required" label:"用户名"`
	Password    string `json:"password" binding:"required" label:"密码"`
	CaptchaID   string `json:"captchaID"`
	CaptchaCode string `json:"captchaCode"`
}

func (UserApi) LoginView(c *gin.Context) {
	cr := middleware.GetBind[LoginRequest](c)

	if global.Config.Site.Login.Captcha {
		if cr.CaptchaID == "" || cr.CaptchaCode == "" {
			res.FailWithMsg("请输入图片验证码", c)
			return
		}
		if !captcha.CaptchaStore.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
			res.FailWithMsg("图片验证码验证失败", c)
			return
		}
	}

	var user models.UserModel
	err := global.DB.Take(&user, "username = ?", cr.Username).Error
	if err != nil {
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	token, err := jwts.SetToken(jwts.Claims{
		UserID: user.ID,
		RoleID: user.RoleID,
	})
	if err != nil {
		logrus.Errorf("生成token失败 %s", err)
		res.FailWithMsg("登录失败", c)
		return
	}

	res.OkWithData(token, c)
	return
}
