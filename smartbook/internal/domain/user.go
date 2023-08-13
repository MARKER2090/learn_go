/*
domain领域层
domain里面的用户user是业务意义上的用户,而dao里面的用户才是数据库里面的用户
*/
package domain

import "time"

// 领域对象User，是DDD的聚合根,是DDD的entitiy
// 有些人叫做BO（business object)
type User struct {
	Email    string
	Password string
	Phone    string
	Ctime    time.Time
}

type Address struct {
}
