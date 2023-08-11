package dao

import (
	"context"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}
func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	return

}

// User直接对应数据库表结构
// 有些人叫做entity，有些人叫做model，有些人叫做PO(persistent object)
type User struct {
	Id       int64
	Email    string
	Password string

	//创建时间：毫秒数
	Ctime int64

	//更新时间：毫秒数
	Utime int64
}
