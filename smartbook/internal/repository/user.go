package repository

import (
	"context"
	"fmt"
	"smartbook/internal/domain"
	"smartbook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO //层层递进,实际就是*gorm.DB数据库
}

// 直接调用dao数据库
func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		fmt.Println(err)
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
	//有缓存的话在这里操作缓存
}

func (r *UserRepository) FindById(int64) {
	//先从cache里面找
	//再从dao里面找
	//找到了再回写到cache里
}
