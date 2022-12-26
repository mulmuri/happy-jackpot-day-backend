package redis

import (
	"github.com/go-redis/redis"

	"backend/api"
)



type RedisLayer interface {
	CreateAuth(userid uint64, atd api.TokenDetails, rtd api.TokenDetails) error
	FetchAuth(authD *api.AccessDetails) (uint64, error)
	DeleteAuth(givenUuid string) (uint64, error)
	DeleteTokens(authD *api.AccessDetails) error
}


type RedisORM struct {
	*redis.Client
}

var RDB *RedisORM



func ConnectRedis(address, dsn string) {
	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	})
	
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	RDB = &RedisORM{
		Client: client,
	}
}



