/*
学习gin框架2
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //新建一个engin

	//静态路由
	router.GET("/hello", func(c *gin.Context) {
		//c.String(http.StatusOK, "你好！这是gin web框架")
		c.String(300, "你好！这是gin web框架")
		fmt.Println(c.Get("hello"))
	})

	//参数路由
	router.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "这是你传过来的名字%s", name)
	})

	//通配符路由
	router.GET("/views/*.html", func(c *gin.Context) {
		path := c.Param(".html")
		c.String(http.StatusOK, "这是你传过来的名字%s", path)
	})

	//查询参数
	router.GET("/order", func(c *gin.Context) {
		id := c.Query("id")
		c.String(http.StatusOK, "这是你传过来的名字%s", id)
	})

	router.Run(":8082") //运行框架
}
