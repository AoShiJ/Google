package main

import (
	"github.com/gin-gonic/gin"
	"way/lx/api"
)

func main() {
	g := gin.Default()
	api.RegisterApi(g)
	g.Run(":8081")

}
