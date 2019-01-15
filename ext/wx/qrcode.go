package wx

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"hatgo/ext"
	"hatgo/pkg/e"
	"hatgo/pkg/link"
	"hatgo/pkg/logs"
	"hatgo/pkg/util"
)

type ReqQrCode struct {
	Path      string `json:"path"`       //默认路径
	Scene     string `json:"scene"`      //默认路径
	IsHyaline bool   `json:"is_hyaline"` //是否需要透明底色
}
type ResQrCodeData struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//小程序码
func XCXQRCode() (string, error) {
	v, err := link.Rd.Get(e.QRCode).Result()
	if err == redis.Nil {
		ak, err := AccessToken()
		if err != nil {
			return "", logs.SysErr(err)
		}

		req := new(ReqQrCode)
		req.Path = "/pages/index/main" //路径
		req.Scene = "fansAuth"
		req.IsHyaline = true
		bt, err := json.Marshal(req)
		if err != nil {
			return "", logs.SysErr(err)
		}
		url := fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", ak)
		body, _ := util.PostCurl(url, bt, util.JSONHeader)
		res := new(ResQrCodeData)
		err = json.Unmarshal(body, res)
		if err != nil {
			path, err := ext.QiniuByteUpload(body, "qrcode")
			if err != nil {
				return "", logs.SysErr(err)
			}
			err = link.Rd.Set(e.QRCode, path, 0).Err()
			if err != nil {
				return "", logs.SysErr(err)
			}
			return path, nil
		}
		return "", fmt.Errorf(res.Errmsg)
	} else if err != nil {
		return "", logs.SysErr(err)
	}
	return v, nil
}
