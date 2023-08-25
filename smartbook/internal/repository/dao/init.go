package dao

import "gorm.io/gorm"

/*
functions:mysql增加一个表格，如果日后想要继续添加表格，需要在后面自行添加
arguments:
return:
tips:
*/
func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{}) //以后要增加新的表格，就需要在这行的}后面自行添加
}
