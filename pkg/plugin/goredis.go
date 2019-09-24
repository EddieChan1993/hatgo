package plugin

import (
	"fmt"
	"github.com/go-redis/redis"
	"hatgo/pkg/s"
	"log"
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
		db         int
	)
	sec, err := s.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis':%v", err)
	}
	host = sec.Key("HOST").MustString("127.0.0.1:6379")
	pass = sec.Key("PASSWORD").MustString("")
	db = sec.Key("DB").MustInt(0)
	Rd = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
		DB:       db,
	})


	_, err = Rd.Ping().Result()
	fmt.Printf("%s %s\n", redisLogIH, "ping redis")
	fmt.Printf("%s host:%s pass:%s db:%d\n", redisLogIH, host, pass, db)
	if err != nil {
		log.Printf("%s %s\n", redisLogWH, err)
	} else {
		log.Printf("%s %s\n", redisLogIH, "redis's connecting is ok")
	}
}
