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

//今晚的时间戳
func DayNightStamp(nowStamp int64) (int64, error) {
	tomorrow := nowStamp + int64(time.Hour.Seconds()*24)
	tomorrowDate := FormatByStamp(tomorrow, YMD)
	dayNightStamp, err := StampByFormat(tomorrowDate, YMD)
	if err != nil {
		return 0, logs.SysErr(err)
	}
	return dayNightStamp, nil
}

//当前时间到未来指定天数晚上的时间间隔
func ExpireDayNight(days int64) (time.Duration, error) {
	torrowStam := time.Now().Unix() + int64(time.Hour.Seconds()*24)*days
	tomD := FormatByStamp(torrowStam, YMD)
	tomS, err := TimeByFormat(tomD, YMD)
	if err != nil {
		return 0, logs.SysErr(err)
	}
	subD := tomS.Sub(time.Now())
	return subD, nil
}

/**
	指定时间戳到未来指定天数晚上的时间间隔
	stamp 时间戳
	days 未来多少天
 */
func ExpireDaysNight(stamp int64, days int64) (time.Duration, error) {
	Stam := stamp + int64(time.Hour.Seconds()*24)*days
	tomD := FormatByStamp(Stam, YMD)
	tomS, err := TimeByFormat(tomD, YMD)
	if err != nil {
		return 0, logs.SysErr(err)
	}
	nowD2 := FormatByStamp(stamp, YMD_HIS)
	nowS2, err := TimeByFormat(nowD2, YMD_HIS)
	if err != nil {
		return 0, logs.SysErr(err)
	}
	subD := tomS.Sub(nowS2)
	return subD, nil
}
