package e

import "fmt"

const (
	AK     = "accessToken"
	AKFlag = "accessTokenFlag" //公众号accessToken
	Ticket = "ticket"          //ticket用于公众号jssdk
	QRCode = "qrcode"          //二维码
)
const (
	HomeDataKey = "homeData" //首页数据
	OpenId      = "openid"   //openid
)

//redis的组合键名
func KeyStr(prefix, key string) string {
	return fmt.Sprintf("%s:%s", prefix, key)
}
