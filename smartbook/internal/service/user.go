/*
service是逻辑层
*/
package service

import (
	"context"
	"errors"
	"fmt"
	"smartbook/internal/domain"
	"smartbook/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("邮箱/帐号或密码不对") //我们不能准确地告诉用户是帐号不对/密码不对，因为准确之后，攻击者会知道会存在某一个用户的。

type UserService interface {
	Login(ctx context.Context, email, password string) (domain.User, error)
	SignUp(ctx context.Context, u domain.User) error
	FindOrCreate(ctx context.Context, phone string)
	Profile(ctx context.Context, id int64)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *userService) Login(ctx context.Context, email string, password string) (domain.User, error) {
	//先找用户
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound { //如果用户找不到
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		fmt.Println(err)
		return domain.User{}, err
	}
	//比较密码了
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		//debug，打印到日志
		return domain.User{}, ErrInvalidUserOrPassword //这里比较密码，如果密码对不上，那么就是返回邮箱/帐号/密码不对
	}
	return u, nil
}

func (svc *userService) SignUp(ctx context.Context, u domain.User) error {
	//你要考虑加密放在哪里的问题了
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return err
	}
	u.Password = string(hash)
	//然后就是，村起来
	return svc.repo.Create(ctx, u)
}

func (svc *userService) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	u, err := svc.repo.FindByPhone(ctx, phone)
	if err != repository.ErrUserNotFound {
		//绝大多数请求都会到达这里，nil也会到达这里，不为ErrUserNotFound也会到达这里
		return u, err
	}
	// 在系统资源不足，触发降级之后，不执行慢路径了
	//if ctx.Value("降级") == "true" {
	//	return domain.User{}, errors.New("系统降级了")
	//}
	// 这个叫做慢路径
	// 你明确知道，没有这个用户
	u = domain.User{
		Phone: phone,
	}
	err = svc.repo.Create(ctx, u)
	if err != nil && err != repository.ErrUserDuplicate {
		return u, err
	}
	// 因为这里会遇到主从延迟的问题
	return svc.repo.FindByPhone(ctx, phone)
}

func (svc *userService) Profile(ctx context.Context, id int64) (domain.User, error) {
	u, err := svc.repo.FindById(ctx, id)
	return u, err
}

func PahtDownGrade(ctx context.Context, quick, slow func()) {
	quick()
	if ctx.Value("降级") == true {
		return
	}
	slow()
}
