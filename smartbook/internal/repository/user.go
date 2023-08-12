package repository

import (
	"context"
	"smartbook/internal/domain"
	"smartbook/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO //层层递进
}

//直接调用dao数据库
func (r *UserRepository) NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
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
