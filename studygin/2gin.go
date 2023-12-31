/*
学习gin框架2
学会restful风格去设计路由,路由都放入到web文件夹内，方便管理
具体操作：就是gin1里面的路由全部分开来
*/
package main

import (
	"strings"
	"studygin/web"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //新建一个*gin.Engin

	//解决跨域问题
	router.Use(cors.New(cors.Config{
		/*
			alloworigins和alloworiginfunc可以只用其一
			如果域名不多，可以只使用alloworigins;
			如果不想使用alloworigins，可以使用alloworiginfunc来通过函数来定义可以通过的函数
			next：2023-08-09 00:54:10 继续根据视频进行学习,理解和完善中间件的代码理解
		*/
		//AllowOrigins: []string{"http://localhost:3000"}, //如果这里的大括号内不写东西，就是允许所有的链接，不推荐这么做，因为公司里面的域名就那么几个，就算全写了也不多，没必要这么做，会增加黑客攻击的风险
		//AllowOrigins: []string{"*"},        //不建议这么做
		AllowMethods: []string{"POST", "GET"},                   //如果这里大括号是不写的，就是允许所有的方式
		AllowHeaders: []string{"Content-Type", "authorization"}, //大小写无所谓
		//允许你带有cookie
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.Contains(origin, "localhost") { //请求的网址含有localhost字符串，就通过给他
				return true
			}
			if strings.HasPrefix(origin, "localhost") { //请求的网址带有localhost前缀的时候，就可以这么做
				return true
			}
			//return strings.Contains(origin, "your company.com")
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	c := web.NewUserHandler()
	//对所有的路由都进行注册
	//c.RegitsterRouter(router)
	c.RegitsterRouterV1(router.Group("/users")) //设置分组
	router.Run(":8082")                         //运行框架
}
