package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	fmt.Println("auth请求")
	c.Next()
	fmt.Println("auth响应")
}
