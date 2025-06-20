package initialize

import (
	"Common/global"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func RedisInit() {
	data := Nacos.Redis
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     data.Addr,
		Password: data.Password, // no password set
		DB:       data.Db,       // use default DB
	})

	pong, err := global.Rdb.Ping(global.Rdb.Context()).Result()
	fmt.Println(pong, err)
	if err != nil {
		panic("redis failed to connect")
	}
	fmt.Println("redis connect success")
	// Output: PONG <nil>
}
