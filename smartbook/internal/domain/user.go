package domain

// 领域对象User，是DDD的聚合根,是DDD的entitiy
// 有些人叫做BO（business object)
type User struct {
	Email    string
	Password string
}

type Address struct {
}
