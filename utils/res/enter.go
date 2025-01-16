package res

import (
	"fast_gin/utils/validate"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Ok(data any, msg string, c *gin.Context) {
	c.JSON(200, Response{
		Code: 0,
		Data: data,
		Msg:  msg,
	})
}

func OkWithMsg(msg string, c *gin.Context) {
	Ok(gin.H{}, msg, c)
}
func OkWithData(data any, c *gin.Context) {
	Ok(data, "成功", c)
}

func Fail(code int, msg string, c *gin.Context) {
	c.JSON(200, Response{
		Code: code,
		Data: gin.H{},
		Msg:  msg,
	})
}

func FailWithMsg(msg string, c *gin.Context) {
	Fail(7, msg, c)
}

func FailWithError(err error, c *gin.Context) {
	msg := validate.ValidateError(err)
	Fail(7, msg, c)
}
