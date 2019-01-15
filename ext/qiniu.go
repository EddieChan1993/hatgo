package ext

import (
	"fans/pkg/link"
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/auth/qbox"
	"fmt"
	"context"
	"bytes"
	"hatgo/pkg/util"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"time"
	"path/filepath"
	"hatgo/pkg/logs"
)

const (
	host       = "http://p2otxz81j.bkt.clouddn.com"
	hostBase   = "http://fans.dcwen.top"
	accessKey  = "YwBMfAjdDqGQMDfWrwWgQrkHoESDD8h_sfQ4oJT7esdG"
	secretKey  = "b-laMNJSLbOyGj-W7343werqfyFOGWEtvinfnaeOLZtAs2-"
	bucket     = "fans"
	folder     = "test"
	isUseHttps = false
	zoneKey    = "huaNan"
)

var (
	cfg       *storage.Config
	putPolicy *storage.PutPolicy
	putExtra  *storage.PutExtra
	mac       *qbox.Mac
	ret       *storage.PutRet
	upToken   string
	zone      = map[string]*storage.Zone{
		"huaDong": &storage.ZoneHuadong,
		"huaBei":  &storage.ZoneHuabei,
		"huaNan":  &storage.ZoneHuanan,
		"beiMei":  &storage.ZoneBeimei,
	}
)

func init() {
	putPolicy = &storage.PutPolicy{
		Scope: bucket,
	}
	mac = qbox.NewMac(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
	cfg = new(storage.Config)
	cfg.Zone = zone[zoneKey]
	cfg.UseHTTPS = isUseHttps
	cfg.UseCdnDomains = isUseHttps

	ret = new(storage.PutRet)
	putExtra = new(storage.PutExtra)
}

//数据流上传
func QiniuUpload(file *multipart.FileHeader, pathName string) (path string, err error) {
	f, err := file.Open()
	defer f.Close()
	if err != nil {
		return "", logs.SysErr(err)
	}
	bf, err := ioutil.ReadAll(f)
	if err != nil {
		return "", logs.SysErr(err)
	}
	//存储后的新地址
	key := fmt.Sprintf("%s/%s/%v%s", folder, pathName, time.Now().UnixNano(), filepath.Ext(file.Filename))
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
	key := fmt.Sprintf("%s/%s/%v%s", link.DbName, pathName, util.Md5(fmt.Sprintf("%d", time.Now().UnixNano())), ext)
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
