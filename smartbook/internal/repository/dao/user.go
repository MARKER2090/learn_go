/*
对数据库进行操作层
由于创建时间和更新时间是偏向于数据库方面的操作的，所以在dao内进行时间方面的操作是较为合理的。
*/
package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

/*
functions：使用insert符合数据库的命名习惯
arguments:
return:
tips：
1.传入都传入结构体，记住！将context作为第一个参数，然后error作为返回值就可以了。这里好习惯!该函数是强耦合的，因为和mysql进行链接的，如果不用mysql，该函数是会报错的。
2.不需要使用分布式锁，因为如果通过并发去查询然后没查到后加入到数据库，会导致重复加入导致报错，且邮箱一直的情况会比较少，所以没必要用大的力气放在极小部分的服务上,迫不得已才使用分布式锁
*/
func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	//存毫秒数：喜好问题，高并发一秒有好千个好几百个用户，按照实际业务去设定即可。最少精确到毫秒数即可
	now := time.Now().UnixMilli()
	//select * from users where email = 123@qq.com for update，在写这个语句进行查询的时候就会进行锁住
	u.Utime = now
	u.Ctime = now
	err := dao.db.WithContext(ctx).Create(&u).Error //自己暂时还不知道这里的意思
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1602 //数值是1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			//邮箱冲突(可以通过gorm.io/go-sql-driver/mysql去查看具体的错误代码)
			return ErrUserDuplicateEmail
		}
	}
	//return dao.db.WithContext(ctx).Create(&u).Error //调用context可以让链路一直保持下去:这里其实我听不懂链路是啥
	return err
}

// User直接对应数据库表结构
// 有些人叫做entity，有些人叫做model，有些人叫做PO(persistent object)
type User struct { //这里的user是直接对标数据库的
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string

	//如果需要添加额外的字段，可以通过这里添加
	//UserDetail//实现用户的不同内容的统一管理

	//创建时间：毫秒数
	Ctime int64 //不使用time.Time，因为含有市区的问题，在后续处理起来会容易出各种问题
	/*
		个人建议就是：所有的与时间的数据全部使用UTC+0，就可以避免很多市区的问题;
		根据用户使用的市区去转化即可。
	*/

	//更新时间：毫秒数
	Utime int64
}

type UserDetail struct {
	name string //姓名
}

type UserCrefend struct { //用户装用户的帐号和密码的，因为帐号和密码的调用是比较高频的，所以可以单独作为一个结构体去被总的表格去嵌入

}
