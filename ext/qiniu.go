package ext

import (
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/auth/qbox"
	"fmt"
	"context"
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"time"
	"path/filepath"
	"hatgo/pkg/conf"
	"hatgo/pkg/logs"
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
		Scope: conf.QiNiuer.Bucket,
	}
	mac = qbox.NewMac(conf.QiNiuer.AccessKey, conf.QiNiuer.SecretKey)
	upToken = putPolicy.UploadToken(mac)
	cfg = new(storage.Config)
	cfg.Zone = zone[conf.QiNiuer.ZoneKey]
	cfg.UseHTTPS = conf.QiNiuer.IsUseHttps
	cfg.UseCdnDomains = conf.QiNiuer.IsUseHttps

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
	key := fmt.Sprintf("%s/%s/%v%s", conf.QiNiuer.Folder, pathName, time.Now().UnixNano(), filepath.Ext(file.Filename))
	formUploader := storage.NewFormUploader(cfg)
	err = formUploader.Put(context.Background(), ret, upToken, key, bytes.NewReader(bf), int64(len(bf)), putExtra)
	if err != nil {
		return "", logs.SysErr(err)
	}
	return fmt.Sprintf("http://%s/%s", conf.QiNiuer.HostBase, key), nil
}

func fileInfo(key string) {
	mac := qbox.NewMac(conf.QiNiuer.AccessKey, conf.QiNiuer.SecretKey)
	bucketManager := storage.NewBucketManager(mac, cfg)

	fileInfo, sErr := bucketManager.Stat(conf.QiNiuer.Bucket, key)
	if sErr != nil {
		fmt.Println(sErr)
		return
	}
	fmt.Println(fileInfo.String())
	//可以解析文件的PutTime
	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
}
