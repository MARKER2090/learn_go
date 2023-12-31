package main

import (
	//"gitee.com/geekbang/basic-go/webook/internal/web"

	web "tempweb/tempweb"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	//v1 := server.Group("/v1")
	//users := server.Group("/users/v1")
	u := web.NewUserHandler() //创建一个userhandler
	//u.RegisterRoutesV1(server.Group("/users"))
	u.RegisterRoutes(server)
	server.Run(":8080")
}
