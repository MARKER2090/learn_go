/*
学习gin框架2
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //新建一个engin

	router.GET("/hello", func(c *gin.Context) {
		//c.String(http.StatusOK, "你好！这是gin web框架")
		c.String(300, "你好！这是gin web框架")
		fmt.Println(c.Get("hello"))
	})

	router.Run(":8082") //运行框架
}
