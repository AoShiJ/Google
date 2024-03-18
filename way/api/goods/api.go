package goods

import "github.com/gin-gonic/gin"

func RegisterGoods(c *gin.Engine) {
	c.POST("/create", CreateGoods)
	c.POST("/query", GetGoodsInfo)
}
