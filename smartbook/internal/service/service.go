package service

import (
	"context"
	"smartbook/internal/domain"
)

type UserService struct {
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	return
}
