/*
handler包：装路由
*/
package web

import (
	"fmt"
	"net/http"
	"regexp"
	"smartbook/internal/domain"
	"smartbook/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
这就是层层传导的做法：可以只调用上一层的错误，而不需要直接调用到dao层面的err变量
这样就不会跨层打交道了。
*/
var ErrUserDuplicateEmail = service.ErrUserDuplicateEmail //这个错误提示是层层传递的。
var ErrUserDuplicateEmailV1 = fmt.Errorf("%w邮箱冲突", service.ErrUserDuplicateEmail)

const biz = "login"

// 确保 UserHandler 上实现了 handler 接口
var _ handler = &UserHandler{}

// 这个更优雅
var _ handler = (*UserHandler)(nil)

// 路由的处理
type UserHandler struct {
	svc         service.UserService
	codeSvc     service.CodeService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

/*
functions:生成一个信的路由邮箱和密码的正则表达式结构体
arguments:nil
return:指针
tips:就是一个简单的golang自带的正则表达式编译
*/
func NewUserHandler(svc service.UserService, codeSvc service.CodeService) *UserHandler {
	const ( //如果正则表达式比较复杂，可以多写几个逻辑要求来对邮箱和密码进行处理，弥补自己正则表达式的书写缺陷
		//定义email的正则表达式:随意搭配+@+随意搭配+.com，自己写的正则表达式
		emailRegexPttern = `[a-zA-Z0-9]+@\w+(\.com)$`
		//定义密码的正则表达式：必须含有英文和中文和特殊符号,且不小于8位数,自己写的要求可以填入非空字符，要求8个以上
		passwordRegexPttern = `\S{8,50}`
	)
	emailExp := regexp.MustCompile(emailRegexPttern)
	passwordExp := regexp.MustCompile(passwordRegexPttern)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
		codeSvc:     codeSvc,
	}
}

// create method of UserHandler
func (u *UserHandler) RegitsterRouter(server *gin.Engine) {
	ug := server.Group("/users") //next写道这里了
	server.GET("/profile", u.ProfileJWT)
	server.POST("/login", u.LoginJWT)
	//server.POST("/login", u.Login)
	server.POST("/signup", u.SignUp)
	server.POST("/edit", u.Edit)
	// PUT "/login/sms/code" 发验证码
	// POST "/login/sms/code" 校验验证码
	// POST /sms/login/code
	// POST /code/sms
	//下面的代码等我做好作业之后再处理
	//ug.POST("/login_sms/code/send", u.SendLoginSMSCode)
	//ug.POST("/login_sms", u.LoginSMS)
}

/*
这是分组管理的方法
避免路由种类太多导致管理麻烦，或者输入前缀容易出错
*/
func (u *UserHandler) RegitsterRouterV1(ug *gin.RouterGroup) {
	//ug := server.Group("/users") //注意：users后面不再加斜杠
	ug.GET("/profile", u.Profile)
	ug.POST("/login", u.Login)
	ug.POST("/signup", u.SignUp)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	//定义注册的string结构体
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var req SignUpReq
	//对上下文进行数据绑定解释
	/*
		理解：
		可以理解Bind方法会根据Content-Type来解释我的数据到达req里去
		解释错了，会直接返回一个400错误，其他啥都不用管
	*/
	fmt.Println("开始对接收到的文档进行处理")
	if err := ctx.Bind(&req); err != nil {
		return
	}
	emailIsExist := u.emailExp.MatchString(req.Email)
	if emailIsExist == false {
		//检测到存在,也就是通过了测试，邮箱没有问题
		ctx.String(http.StatusBadRequest, "你的邮箱格式不对,请重新输入")
		return
	}
	passwordIsExist := u.passwordExp.MatchString(req.Password)
	if passwordIsExist == false {
		//不通过
		ctx.String(http.StatusBadRequest, "你的密码格式不对,请重新输入")
		return
	} else {
		//通过了测试，邮箱没有问题
		if req.Password != req.ConfirmPassword {
			ctx.String(http.StatusBadRequest, "两次输入的密码不一致")
			return
		}
	}

	//调用svc的方法
	err := u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == ErrUserDuplicateEmail { //通过层层调用的方式去调用各种
		ctx.String(http.StatusOK, "邮箱冲突")
	}
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusOK, "系统异常")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	fmt.Println("req是：", req)
}

/*
functions:Web页面的登录函数
arguments:
return:
tips:需要层层调用，然后与数据库内的数据进行比对，如果重复的就报错
*/
func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil { //Bind函数会将req里面的标签作为标志物区查找ctx里面的emial，然后提取数值出来给自己的email变量
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "用户名/密码不对")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	//登录成功了
	//设置session
	sess := sessions.Default(ctx)
	//我就可以随便设置session的数值了，也就是放在session里面的数值
	sess.Set("userId", user.Id)
	sess.Save()

	ctx.String(http.StatusOK, "登录成功")
	return
}

func (u *UserHandler) Edit(ctx *gin.Context) {

}
func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "您已经访问到了profile")
}

var a string = `
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
	`
