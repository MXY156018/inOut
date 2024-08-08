/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-15 16:19:35
 * @LastEditors: MXY156018 1132022296@qq.com
 * @LastEditTime: 2022-05-23 21:41:09
 */
/**
 * @Author: Dong
 * @Description:获得当前月，当前周，当前季度的初始和结束日期
 * @File:  tools
 * @Date: 2020/08/06 16:24
 */
package utils

import (
	"fmt"
	"time"
)

const (
	DayInt                       = 24 * 3600
	DayDur                       = 24 * time.Hour
	DateFmt                      = "2006-01-02"
	DatetimeFmt                  = "2006-01-02 15:04:05"
	StartTimeSuffix              = " 00:00:00"
	EndTimeSuffix                = " 23:59:59"
	GetDateBetweenForDate  int32 = 1 // 按日处理
	GetDateBetweenForMonth int32 = 2 // 按月处理
	GetDateBetweenForYear  int32 = 3 // 按年处理
)

func GetDate(date time.Time, types string) (int64, int64) {
	var start, end int64
	if types == "today" {
		start, end = GetTodayUnix(date)
	} else if types == "yesterday" {
		start, end = GetPreDayUnix(date)
	} else if types == "week" {
		start, end = GetWeekDayUnix(date)
	} else if types == "last week" {
		start, end = GetPreWeekUnix(date)
	} else if types == "month" {
		start, end = GetMonthDayUnix(date)
	} else {
		return start, end
	}
	return start, end
}

//获取本日时间戳
func GetTodayUnix(date time.Time) (int64, int64) {
	currentYear, currentMonth, currentDay := date.Date()
	currentLocation := date.Location()
	today := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)
	defaultFormat := DatetimeFmt
	return Datetime2TimeObj(today.Format("2006-01-02")+" 00:00:00", defaultFormat).Unix(), Datetime2TimeObj(today.Format("2006-01-02")+" 23:59:59", defaultFormat).Unix()
}

// Datetime2TimeObj
func Datetime2TimeObj(target string, format ...string) time.Time {
	defaultFormat := DatetimeFmt
	if len(format) > 0 {
		defaultFormat = format[0]
	}
	timeObj, err := time.ParseInLocation(defaultFormat, target, time.Local)
	if err != nil {
		return time.Time{}
	}
	return timeObj
}

/**
 * @Author Dong
 * @Description 获得当前月的初始和结束日期
 * @Date 16:29 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetMonthDayUnix(date time.Time) (int64, int64) {

	currentYear, currentMonth, _ := date.Date()
	currentLocation := date.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	defaultFormat := DatetimeFmt
	return Datetime2TimeObj(time.Unix(f, 0).Format("2006-01-02")+" 00:00:00", defaultFormat).Unix(), Datetime2TimeObj(time.Unix(l, 0).Format("2006-01-02")+" 23:59:59", defaultFormat).Unix()
}

func GetLastDayOfMonth2(date time.Time) time.Time {
	currentYear, currentMonth, _ := date.Date()
	currentLocation := date.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	l := lastOfMonth.Unix()
	return time.Unix(l, 0)
}

//获取本月最后一天
func GetLastDayOfMonth(date time.Time) string {
	l := GetLastDayOfMonth2(date)
	return l.Format("2006-01-02") + " 23:59:59"
}

func GetSection(date time.Time) (string, string) {
	currentYear, currentMonth, _ := date.Date()
	currentLocation := date.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	day := date.Local().Day()
	month := date.Local().Month()
	year := date.Local().Year()
	var startday, endday int
	if day >= 21 {
		startday = 21
		endday = lastOfMonth.Day()
	} else if day >= 11 {
		startday = 11
		endday = 21
	} else if day >= 1 {
		startday = 1
		endday = 11
	}
	return fmt.Sprintf("%d-%d-%d", year, month, startday) + " 00:00:00", fmt.Sprintf("%d-%d-%d", year, month, endday) + " 23:59:59"
}

/**
 * @Author Dong
 * @Description 获得当前周的初始和结束日期
 * @Date 16:32 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetWeekDayUnix(date time.Time) (int64, int64) {
	offset := int(time.Monday - date.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - date.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	defaultFormat := DateFmt
	return Datetime2TimeObj(time.Unix(f, 0).Format("2006-01-02")+" 00:00:00", defaultFormat).Unix(), Datetime2TimeObj(time.Unix(l, 0).Format("2006-01-02")+" 23:59:59", defaultFormat).Unix()
}

/**
 * @Author Dong
 * @Description //获得当前季度的初始和结束日期
 * @Date 16:33 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetQuarterDay() (string, string) {
	year := time.Now().Format("2006")
	month := int(time.Now().Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
	}
	return firstOfQuarter, lastOfQuarter
}

// 获取本周开始及结算unix时间戳
//
//now 时间
func GetPreWeekUnix(now time.Time) (int64, int64) {
	ad, _ := time.ParseDuration("-168h")
	now = now.Add(ad)

	start, end := GetWeekDayUnix(now)
	return start, end
}

// 获取前日的unix时间戳
//
//now 时间
func GetPreDayUnix(now time.Time) (int64, int64) {
	ad, _ := time.ParseDuration("-24h")
	now = now.Add(ad)

	start, end := GetTodayUnix(now)
	return start, end
}

// 获取日期
//
//date 时间
//
//fmt 格式
//
//s 偏移量
func GetDateStr(date time.Time, fmtstr string, s string) string {
	ad, _ := time.ParseDuration(s)
	now := date.Add(ad)
	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()
	today := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)

	return today.Format(fmtstr)
}

// 获取当天0点的Unix 时间戳
//
//date 时间
//
//s 时间偏移
func GetDateUnix(date time.Time, s string) int64 {
	ad, _ := time.ParseDuration(s)
	now := date.Add(ad)
	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()
	today := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)

	return today.Unix()
}

// 获取 月-日 格式时间
//
// date 时间
//
//s 便宜量， 如 1d
func GetMonthDayStr(date time.Time, s string) string {
	ad, _ := time.ParseDuration(s)
	now := date.Add(ad)
	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()
	today := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)
	month := today.Local().Month()
	day := today.Local().Day()
	return fmt.Sprintf("%d-%d", month, day)
}

/**
 * @description: 获取上月一号的日期
 * @param {time.Time} now
 * @return {*}
 */
func GetPreMonth(now time.Time) string {
	year := now.Year()
	var month int
	if now.Month() == 1 {
		month = 12
		year -= 1
	} else {
		month = int(now.Month()) - 1
	}
	return fmt.Sprintf("%d-%02d-%02d", year, int(month), 1)
}

// 获取本月日期字符串
func GetMonth(now time.Time) string {
	year := now.Year()
	month := int(now.Month())

	return fmt.Sprintf("%d-%02d-%02d", year, int(month), 1)
}

/**
 * @description: 获取最近三个月的起止日期
 * @param {time.Time} now
 * @return {*}
 */
func GetLastThreeMonth(date time.Time) (string, string) {
	month := date.Month()
	year := date.Year()

	var nowmonth, nowyear int
	if month < 4 {
		nowmonth = int(month) + 9
		nowyear = year - 1
	} else {
		nowmonth = int(month) - 3
		nowyear = year
	}
	return fmt.Sprintf("%d-%d-%d", nowyear, int(nowmonth), date.Day()), fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day()) + " 23:59:59"
}
func UnixToString(date int64) string {
	return time.Unix(date, 0).Format(DatetimeFmt)
}

func StringToTime(date string) (time.Time, error) {
	return time.ParseInLocation(DateFmt, date, time.Local)
}
