package e

import "fmt"

const (
	HomeDataKey = "homeData" //首页数据
	OpenId      = "openid"   //openid
)
//redis的组合键名
func KeyStr(prefix, key string) string {
	return fmt.Sprintf("%s:%s", prefix, key)
}
