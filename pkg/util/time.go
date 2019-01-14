/**
	时间的转换
 */
package util

import (
	"hatgo/pkg/logs"
	"time"
)

const (
	//timeVal
	YMD_HIS = "2006-01-02 15:04:05"
	YMD     = "2006-01-02"
	YM      = "2006-01"
)

//时间戳转为标准格式
func FormatByStamp(timeStamp int64, timeFormat string) string {
	return time.Unix(timeStamp, 0).Format(timeFormat)
}

//标准格式转为时间戳
func StampByFormat(format string, timeFormat string) (int64, error) {
	loc, _ := time.LoadLocation("Local") //获取当地时区
	tm2, err := time.ParseInLocation(timeFormat, format, loc)
	if err != nil {
		return 0, logs.SysErr(err)
	}
	return tm2.Unix(), nil
}

//标准格式转为time
func TimeByFormat(format string, timeFormat string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local") //获取当地时区
	return time.ParseInLocation(timeFormat, format, loc)
}

//当前日期
func NowFormat(timeFormat string) string {
	stamp := time.Now().Unix()
	return FormatByStamp(stamp, timeFormat)
}

/**
	获取当天到添加时间的timeDuration区间时间类型
	days 加上多少天时间
	返回
 */
func ExpireDays(days int64) (time.Duration, error) {
	torrowStam := time.Now().Unix() + 24*60*60
	tomD := FormatByStamp(torrowStam, YMD)
	tomS, err := TimeByFormat(tomD, YMD)
	if err != nil {
		return 0, logs.SysErr(err)
	}
	subD := tomS.Sub(time.Now())
	return subD, nil
}

/**
	获取当天到添加时间的timeDuration区间时间类型
	days 加上多少天时间
	返回
 */
func ExpireSec(sec int64) (time.Duration, error) {
	torrowStam := time.Now().Unix() + sec
	tomD := FormatByStamp(torrowStam, YMD_HIS)
	tomS, err := TimeByFormat(tomD, YMD_HIS)
	if err != nil {
		return 0, logs.SysErr(err)
	}
	subD := tomS.Sub(time.Now())
	return subD, nil
}
