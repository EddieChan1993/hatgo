package link

import (
	"github.com/go-redis/redis"
	"log"
	"hatgo/pkg/setting"
	"fmt"
)

const (
	redisLogIH = "[go-redis] [info]"
	redisLogWH = "[go-redis] [warning]"
)

var Rd *redis.Client

func init() {
	rd()
}

func rd() {
	var (
		err        error
		pass, host string
	)
	sec, err := setting.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis':%v", err)
	}
	host = sec.Key("HOST").MustString("127.0.0.1:6379")
	pass = sec.Key("PASSWORD").MustString("")
	Rd := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
	})

	_, err = Rd.Ping().Result()
	fmt.Println(fmt.Sprintf("%s %s", redisLogIH, "ping redis"))
	if err != nil {
		fmt.Println(fmt.Sprintf("%s %s", redisLogWH, err))
	} else {
		fmt.Println(fmt.Sprintf("%s %s", redisLogIH, "redis's connecting is ok"))
	}
	fmt.Println("--------------------------------------------------------------")
}
