/*
学习gin框架2
学会restful风格去设计路由,路由都放入到web文件夹内，方便管理
具体操作：就是gin1里面的路由全部分开来
*/
package main

import (
	"studygin/web"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //新建一个*gin.Engin

	////设置分组
	//users := router.Group("/users/ddd")

	////这是属于/user/ddd组的，需要访问的网址是http://localhost:8082/users/ddd/hello
	//users.GET("/hello", func(c *gin.Context) {
	//	c.String(http.StatusOK, "可以访问！200")
	//})
	c := web.NewUserHandler()
	//对所有的路由都进行注册
	c.RegitsterRouter(router)

	router.Run(":8082") //运行框架
}
