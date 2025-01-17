package middleware

import (
	"fast_gin/utils/jwts"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	_, err := jwts.CheckToken(token)
	if err != nil {
		res.FailWithMsg("认证失败", c)
		c.Abort()
		return
	}
	c.Next()
}

func AdminMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	claims, err := jwts.CheckToken(token)
	if err != nil {
		res.FailWithMsg("认证失败", c)
		c.Abort()
		return
	}
	if claims.RoleID != 1 {
		res.FailWithMsg("角色认证失败", c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
	c.Next()
}

func GetAuth(c *gin.Context) (cl *jwts.MyClaims) {
	cl = new(jwts.MyClaims)
	_claims, ok := c.Get("claims")
	if !ok {
		return
	}
	cl, ok = _claims.(*jwts.MyClaims)
	return
}
