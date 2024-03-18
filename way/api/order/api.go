package order

import "github.com/gin-gonic/gin"

func RegisterOrder(c *gin.Engine) {
	c.POST("/order", CreateOrder)
}
