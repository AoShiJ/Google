package user

import (
	"github.com/gin-gonic/gin"
	"way/lx/middleware"
)

func Register(c *gin.Engine) {
	c.POST("/login", Login)
	u := c.Group("/user")
	u.Use(middleware.AuthToken)
	{
		u.POST("del", GetUserInfo)
		u.POST("send", SendMobile)
	}
}
