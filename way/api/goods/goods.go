package goods

import (
	"demo/lx/framework/http"
	"demo/lx/message/goods"
	"github.com/gin-gonic/gin"
	"way/lx/server"
)

func CreateGoods(c *gin.Context) {
	goodsinfo := goods.GoodsInfo{}
	if err := c.BindJSON(&goodsinfo); err != nil {
		http.Res(c, err.Error(), 502, "")
		return
	}
	if goodsinfo.Name == "" {
		http.Res(c, "mi名称不能为空", 502, "")
		return
	}
	createGoods, err := server.CreateGoods(c, &goodsinfo)
	if err != nil {
		return
	}
	http.Res(c, "添加商品成功", 200, createGoods)
	return

}
func GetGoodsInfo(c *gin.Context) {
	name := c.PostForm("name")
	createGoods, err := server.QueryGoods(c, name)
	if err != nil {
		http.Res(c, err.Error(), 502, "")
		return
	}
	http.Res(c, "查询成功", 200, createGoods)
	return
}
