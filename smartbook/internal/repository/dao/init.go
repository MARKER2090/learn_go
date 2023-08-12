package dao

import "gorm.io/gorm"

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{}) //以后要增加新的表格，就需要在这行的}后面自行添加
}
