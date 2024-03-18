package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Msg  string
	Code int64
	Data interface{}
}

func Res(c *gin.Context, msg string, code int64, data interface{}) {
	httpCode := http.StatusOK
	if code > 20000 {
		httpCode = http.StatusBadGateway
	}
	c.JSON(httpCode, Response{
		Msg:  msg,
		Code: code,
		Data: data,
	})
	return
}
