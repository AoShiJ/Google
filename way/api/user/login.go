package user

import (
	"demo/lx/framework/http"
	demo "demo/lx/message"
	"way/lx/server"

	"github.com/gin-gonic/gin"
	"way/lx/consts"
)

func Login(c *gin.Context) {
	n := new(demo.LoginRequest)
	if err := c.BindJSON(&n); err != nil {
		http.Res(c, err.Error(), consts.PPM_ERROR, nil)
		return
	}
	login, err := server.UserLogin(c, n.Username, n.Password)
	if err != nil {
		http.Res(c, err.Error(), consts.PPM_ERROR, nil)
		return
	}
	http.Res(c, "登录成功", consts.SUCCESS, login)
	return
}
