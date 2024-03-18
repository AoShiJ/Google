package api

import (
	"github.com/gin-gonic/gin"
	"way/lx/api/goods"
	"way/lx/api/order"
	"way/lx/api/user"
)

func RegisterApi(g *gin.Engine) {
	user.Register(g)
	goods.RegisterGoods(g)
	order.RegisterOrder(g)
}
