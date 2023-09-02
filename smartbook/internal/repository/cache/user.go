package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"smartbook/internal/domain"
	"time"

	"github.com/redis/go-redis/v9"
)

var ErrKeyNoExist = redis.Nil

type UserCache interface {
	Get(ctx context.Context, id int64) (domain.User, error)
	Set(ctx context.Context, u domain.User) error
}

type RedisUserCache struct {
	//传单机Redis可以
	//传cluster的Redis也可以
	client     redis.Cmdable
	expiration time.Duration
}

func NewUserCacheV1(addr string) UserCache {
	client := redis.NewClient(&redis.Options{})
	return &RedisUserCache{
		client:     client,
		expiration: time.Minute * 15,
	}
}

//NewUserCache
//A用到了B，B一定要接口=>这个是保证面向对象
//A用到了B，B一定是A的字段=>规避包变量、包方法，非常缺乏扩展性
//A用到了B，A绝对不出是话B，而是外面注入=>保持依赖注入（DI，Dependency Injection）和依赖翻转(ioc)
//expiration is,1m         ????
func NewUserCache(client redis.cmdable) UserCache {
	return &RedisUserCache{
		client:     client,
		expiration: time.Minute * 15,
	}
}

//Get如果没有输入，返回一个特定的error
func (cache *RedisUserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := cache.key(id)
	//如果数据不存在，err = redis.Nil
	val, err := cache.client.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	err = json.Unmarshal(val, &u)
	//if err!=nil{
	//  return domain.User{},err
	//}
	//return u,nil
	return u, err
}

func (cache *RedisUserCache) Set(ctx context.Context, u domain.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := cache.key(u.Id)
	return cache.client.Set(ctx, key, val, cache.expiration).Err()
}

func (cache *RedisUserCache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}

// main 函数里面初始化好
//var RedisClient *redis.Client

//func GetUser(ctx context.Context, id int64) {
//	RedisClient.Get()
//}

//type UnifyCache interface {
//	Get(ctx context.Context, key string)
//	Set(ctx context.Context, key string, val any, expiration time.Duration)
//}
//
//
//type NewRedisCache() UnifyCache {
//
//}
