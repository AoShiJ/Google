package order

import (
	"demo/lx/framework/http"
	"github.com/gin-gonic/gin"
	"way/lx/server"
)

func CreateOrder(c *gin.Context) {

	var message struct {
		Name string `json:"name"`
		Num  int    `json:"num"`
	}
	if err := c.BindJSON(&message); err != nil {
		http.Res(c, "绑定参数错误", 502, "")
		return
	}
	create, err := server.OrderCreate(c, message.Name, message.Num)
	if err != nil {
		http.Res(c, "数据错误", 502, "")
		return
	}
	http.Res(c, "创建成功", 200, create)
	return
}
