package web

import (
	"fmt"
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

/*
functions:UserHandler 我准备在它上面定义跟用户有关的路由
arguments:nil
return:经过编译的邮箱和密码的正则表达式
tips:
*/
type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

/*
functions:注册路由,含有所有的注册路由,这是作为用户处理userhandler的方法
arguments:
return:
tips:这是使用了路由组的
*/
func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	//注册了多个路由
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}

/*
functions:注册所有路由
arguments:
return:
tips:非路由组的情况
*/
func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	//注册了多个路由
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile) //u是邮件和密码的正则表达式结构体的方法,方法的
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	var req SignUpReq
	// Bind 方法会根据 Content-Type 来解析你的数据到 req 里面
	// 解析错了，就会直接写回一个 400 的错误
	if err := ctx.Bind(&req); err != nil {
		return
	}

	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		// 记录日志
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	fmt.Printf("%v", req)
	// 这边就是数据库操作
}

func (u *UserHandler) Login(ctx *gin.Context) {
}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}
