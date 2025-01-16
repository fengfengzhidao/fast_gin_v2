package middleware

import (
	"fast_gin/utils/jwts"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	_, err := jwts.CheckToken(token)
	if err != nil {
		c.JSON(200, gin.H{"code": 7, "msg": "认证失败", "data": gin.H{}})
		c.Abort()
		return
	}
	c.Next()
}

func AdminMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	claims, err := jwts.CheckToken(token)
	if err != nil {
		c.JSON(200, gin.H{"code": 7, "msg": "认证失败", "data": gin.H{}})
		c.Abort()
		return
	}
	if claims.RoleID != 1 {
		c.JSON(200, gin.H{"code": 7, "msg": "角色认证失败", "data": gin.H{}})
		c.Abort()
		return
	}
	c.Next()
}
