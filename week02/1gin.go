package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon" //网页缩略图
)

func main() {
	//创建一个服务
	ginserver := gin.Default()
	ginserver.Use(favicon.New("favicon.png"))

	//如果连上点数据库

	//访问网站，处理我们的请求get
	ginserver.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "你好~小可爱！"})
	})

	ginserver.POST("/post", func(context *gin.Context) {
		context.String(http.StatusOK, "这里是post路由")
	})

	//查询参数
	ginserver.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "这里是post路由", name)
	})

	//通配符:通配符里面的*是不可以单独出现的。
	ginserver.GET("/views/*.html", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "这是通配符")
	})

	/*
		restful风格
		//get /users/xxx 查询
		//delete  /users/xx 删掉
		//put /users/xx 注册
		//post /users/xx 修改
	*/

	//服务端运行起来
	ginserver.Run(":8082")
}
