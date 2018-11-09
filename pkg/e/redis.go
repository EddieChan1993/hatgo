package e

import "fmt"

const (
	HOME_DATA_KEY = "homeData" //首页数据
	OPEN_ID       = "openid"   //openid
)
//redis的组合键名
func KeyStr(prefix, key string) string {
	return fmt.Sprintf("%s:%s", prefix, key)
}
