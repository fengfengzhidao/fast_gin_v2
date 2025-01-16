package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LimitMiddleware(limit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("limit请求")
		c.Next()
		fmt.Println("limit响应")
	}
}
