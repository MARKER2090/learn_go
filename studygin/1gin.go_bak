package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon" //网页缩略图
)

func main() {
	//创建一个服务
	router := gin.Default()
	router.Use(favicon.New("favicon.png"))

	//设置分组
	users := router.Group("/users/ddd")

	//这是属于/user/ddd组的，需要访问的网址是http://localhost:8082/users/ddd/hello
	users.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "可以访问！200")
	})

	//静态路由:就是只要访问到了这个地址，就可以返回对应的函数
	router.GET("/hello", func(c *gin.Context) {
		//c.String(http.StatusOK, "你好！这是gin web框架")
		c.String(300, "你好！这是gin web框架")
	})

	//参数路由：冒号一定要有，否则无法识别到name这个参数
	/*
		理解：在/users/:name里冒号后面的name就是参数，c.param("name")就是查询网站的name参数
		可以是任意字符，网站名称http://localhost:8082/users/452452?id=1354，识别到的name是452452
		不识别问号及后面的内容(其他符号是识别的),可以应用于网站识别层次
	*/
	router.GET("/users/:bcd", func(c *gin.Context) {
		name := c.Param("bcd")
		c.String(http.StatusOK, "参数路由接收到的参数是：%s", name)
	})

	//通配符路由
	/*
		理解：
		发现只要是views后面的所有一切都可以读取，不限定是.html格式,那岂不是和参数路由一致了？
	*/
	router.GET("/views/*abc", func(c *gin.Context) {
		path := c.Param("abc")
		c.String(http.StatusOK, "服务器接收到的通配符是：%s", path)
	})

	//查询参数
	/*
		理解：
		网址：http://localhost:8082/order?id=1234,识别到的参数内容是：1234
		意思就是读取=号后面的内容;
		如果网址是http://localhost:8082/order?id=1234&&ip=127.0.9.1,识别到的参数内容还是：1234
	*/
	router.GET("/order", func(c *gin.Context) {
		id := c.Query("id")
		c.String(http.StatusOK, "查询参数路由id:%s", id)
	})

	//服务端运行起来
	router.Run(":8082")
}
