package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	//ID int64 `gorm:"primaryKey,autoIncrement"`
	Code  string `gorm:"unique"`
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//db,err:=gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/your_db"))
	if err != nil {
		panic("failed to connect database")
	}

	db = db.Debug()

	//迁移schema
	//建立表格
	db.AutoMigrate(&Product{})

	//Create
	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&Product{Code: "G42", Price: 150})

	//Read
	var product Product
	db.First(&product, 1)                 //根据整型主键查找
	db.First(&product, "code = ?", "D42") //查找code字段值为D42的记录
	fmt.Println("这里读取成功了！")

	//Update:更新数值
	db.Model(&product).Update("Price", 200) //这是全部修改的
	/*
		Updates更新多个字段
		也就是一句话会更新price和code两个字段
		set`price`=200,`code`=`F42`
	*/
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})                    //仅更新非零字段,因为gorm.Model还有4个字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"}) //甚至可以通过map来更新

	//Delete 删除product
	db.Delete(&product, 1)               //用法应该和read方法一样的
	db.Delete(&product, "code=?", "F42") //用法应该和read方法一样的
	time.Sleep(time.Second)
}
