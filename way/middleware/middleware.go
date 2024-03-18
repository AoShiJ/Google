package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"way/lx/utils"
)

func AuthToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	to, err := utils.GetJwtToken(time.Now().Unix(), int64(time.Second*5), token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "无效 token")
		c.Abort()
	}
	c.Set("to", to)
}
