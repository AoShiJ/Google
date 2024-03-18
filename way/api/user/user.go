package user

import (
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"way/lx/utils"
)

func GetUserInfo(c *gin.Context) {

	s, _ := c.Get("to")
	logs.Info(s, "sssssssssss")
	c.JSON(http.StatusOK, "成功")
	return

}
func SendMobile(c *gin.Context) {
	utils.SendMobile(c.PostForm("mobile"))
	c.JSON(http.StatusOK, "发送手机号")
	return
}
