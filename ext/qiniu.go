package ext

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"hatgo/pkg/logs"
	"hatgo/pkg/util"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"time"
)

const (
	host       = "http://p2otxz81j.bkt.clouddn.com"
	hostBase   = "http://fans.dcwen.top"
	accessKey  = "YwBMfAjdDqGQMWrwWgQrkHoES8h_sfQ4oJT7esdG"
	secretKey  = "b-laMNJSLbOyGj-W7qfyFOGWEtvinnaeOLZtAs2-"
	bucket     = "fans"
	isUseHttps = false
	zoneKey    = "huaNan"
)
const tokenExpire = time.Hour //1个小时token
const maxSize = 1 << 20 * 10  //最大10M
var (
	cfg       *storage.Config
	putPolicy *storage.PutPolicy
	putExtra  *storage.PutExtra
	mac       *qbox.Mac
	ret       *storage.PutRet
	zone      = map[string]*storage.Zone{
		"huaDong": &storage.ZoneHuadong,
		"huaBei":  &storage.ZoneHuabei,
		"huaNan":  &storage.ZoneHuanan,
		"beiMei":  &storage.ZoneBeimei,
	}
)

func init() {
	cfg = new(storage.Config)
	cfg.Zone = zone[zoneKey]
	cfg.UseHTTPS = isUseHttps
	cfg.UseCdnDomains = isUseHttps

	ret = new(storage.PutRet)
	putExtra = new(storage.PutExtra)
}

//官方token有效期为1小时
func token() (string, error) {
	putPolicy = &storage.PutPolicy{
		Scope: bucket,
	}
	mac = qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken, nil

}

//数据流上传
func QiniuUpload(file *multipart.FileHeader, pathName string) (path string, err error) {
	f, err := file.Open()
	defer f.Close()
	if file.Size > maxSize {
		return "", logs.SysErr(fmt.Errorf("上传的文件太大了，客官"))
	}
	if err != nil {
		return "", logs.SysErr(err)
	}
	bf, err := ioutil.ReadAll(f)
	if err != nil {
		return "", logs.SysErr(err)
	}
	upToken, err := token()
	if err != nil {
		return "", logs.SysErr(err)
	}
	//存储后的新地址
	ext := filepath.Ext(file.Filename) //图片格式
	key := fmt.Sprintf("%s/%v%s", pathName, util.Md5(fmt.Sprintf("%d", time.Now().UnixNano())), ext)
	formUploader := storage.NewFormUploader(cfg)

	err = formUploader.Put(context.Background(), ret, upToken, key, bytes.NewReader(bf), int64(len(bf)), putExtra)
	if err != nil {
		return "", logs.SysErr(err)
	}
	imgUrl := ""
	if hostBase == "" {
		imgUrl = fmt.Sprintf("%s/%s", host, key)
	} else {
		imgUrl = fmt.Sprintf("%s/%s", hostBase, key)
	}
	return imgUrl, nil
}

//字节数组的上传
func QiniuByteUpload(body []byte, pathName string) (path string, err error) {
	ext := ".png" //图片格式
	key := fmt.Sprintf("%s/%v%s", pathName, util.Md5(fmt.Sprintf("%d", time.Now().UnixNano())), ext)
	upToken, err := token()
	if err != nil {
		return "", logs.SysErr(err)
	}
	formUploader := storage.NewFormUploader(cfg)
	err = formUploader.Put(context.Background(), ret, upToken, key, bytes.NewReader(body), int64(len(body)), putExtra)
	if err != nil {
		return "", logs.SysErr(err)
	}
	imgUrl := ""
	if hostBase == "" {
		imgUrl = fmt.Sprintf("%s/%s", host, key)
	} else {
		imgUrl = fmt.Sprintf("%s/%s", hostBase, key)
	}
	return imgUrl, nil
}

//删除文件
func QiniuDelUrl(urlPath string) error {
	u, err := url.Parse(urlPath)
	if err != nil {
		return logs.SysErr(err)
	}
	key := string([]byte(u.Path)[1:])
	mac := qbox.NewMac(accessKey, secretKey)
	bucketManager := storage.NewBucketManager(mac, cfg)
	return bucketManager.Delete(bucket, key)
}

func fileInfo(key string) {
	mac := qbox.NewMac(accessKey, secretKey)
	bucketManager := storage.NewBucketManager(mac, cfg)

	fileInfo, sErr := bucketManager.Stat(bucket, key)
	if sErr != nil {
		fmt.Println(sErr)
		return
	}
	fmt.Println(fileInfo.String())
	//可以解析文件的PutTime
	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
}
