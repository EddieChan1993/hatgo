package link

import (
	"github.com/go-redis/redis"
	"log"
	"hatgo/pkg/setting"
	"fmt"
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
	host = sec.Key("HOST").String()
	pass = sec.Key("PASSWORD").String()
	Rd := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
	})

	_, err = Rd.Ping().Result()
	fmt.Println("[go-redis][info] ping redis")
	if err != nil {
		fmt.Println(fmt.Sprintf("[go-redis][warning] %s",err))
	}else{
		fmt.Println("[go-redis][info] redis's connecting is ok")
	}
	fmt.Println("----------------------------------")
}
