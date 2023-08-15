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

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
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

func (svc *UserService) Login(ctx context.Context, email string, password string) error {
	//先找用户
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound { //如果用户找不到
		return ErrInvalidUserOrPassword
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	//比较密码了
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		//debug，打印到日志
		return ErrInvalidUserOrPassword //这里比较密码，如果密码对不上，那么就是返回邮箱/帐号/密码不对
	}
	return nil
}
