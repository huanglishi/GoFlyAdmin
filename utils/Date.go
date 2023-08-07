package utils

import (
	"time"
)

// 时间转时间戳
// timetype时间格式类型 date=日期 datetime=日期时间 datesecond=日期时间秒
func StringTimestamp(timeLayout string, timetype string) int64 {
	timetpl := "2006-01-02 15:04:05"
	if timetype == "date" {
		timetpl = "2006-01-02"
	} else if timetype == "datetime" {
		timetpl = "2006-01-02 15:04"
	}
	times, _ := time.ParseInLocation(timetpl, timeLayout, time.Local)
	timeUnix := times.Unix()
	return timeUnix
}

// 时间戳格式化为日期字符串
// timetype时间格式类型 date=日期 datetime=日期时间 datesecond=日期时间秒
func TimestampString(timedata interface{}, timetype string) string {
	timetpl := "2006-01-02 15:04:05"
	if timetype == "date" {
		timetpl = "2006-01-02"
	} else if timetype == "datetime" {
		timetpl = "2006-01-02 15:04"
	}
	return time.Unix(timedata.(int64), 0).Format(timetpl)
}
