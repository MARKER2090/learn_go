/*
service是逻辑层
*/
package service

import (
	"context"
	"smartbook/internal/domain"
	"smartbook/internal/repository"
)

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
	//然后就是，村起来
	return svc.repo.Create(ctx, u)
}
