package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cihub/seelog"
)

const (
	FIVE_MINUTE_STEP = 5 * 60       //5分钟
	ONE_HOUR_STEP    = 60 * 60      // 一小时
	ONE_DAY_STEP     = 24 * 60 * 60 //一天
)

func GetUnix13NowTime() int64 {
	return time.Now().UTC().UnixNano() / 1000000
}

func GetUnix13Time(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

//获取传入日期零点时间
func GetZeroTime(t time.Time) time.Time {
	zt, _ := time.ParseInLocation("2006-01-02", t.Format("2006-01-02"), time.Local)
	return zt
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取当天日期零点时间
func GetNowZeroTime() time.Time {
	zt, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	return zt
}

//获取第n个时间间隔的时间
func GetTimeByIndex(zeroTime time.Time, timeStep, index int64) time.Time {

	return zeroTime.Add(time.Duration(timeStep*index) * time.Second)
}

//获取距传入时间最近时间间隔点是一天中的第几个
func GetIndexOfTime(t time.Time, timeStep int64) int64 {
	zeroTime := GetZeroTime(t)
	index := (t.Unix() - zeroTime.Unix()) / timeStep
	return index
}

func GetYearWeek(t int64) int {
	ti := time.Unix(t, 0)
	day := 0
	if ti.Weekday() == 0 {
		day = ti.YearDay() - 6
	} else {
		day = ti.YearDay() - (int(ti.Weekday()) - 1)
	}
	begin, _ := time.Parse("2006-01-02", fmt.Sprintf("%d-01-01", ti.Year()))
	beginweek := int(begin.Weekday())
	yearweek := 0
	switch beginweek {
	case 1:
		yearweek = day / 7
	case 0:
		yearweek = (day - 1) / 7
	default:
		yearweek = (day - 8 + beginweek) / 7
	}
	if yearweek == 0 {
		return GetYearWeek(ti.AddDate(0, 0, -ti.YearDay()).Unix())
	}
	return yearweek
}

//时区转换
func ParseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		return time.Time{}, err
	} else {
		//转成带时区的时间
		lt, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, l)
		//直接转成相对时间
		// fmt.Println(time.Now().In(l).Format("2006-01-02 15:04:05"))
		return lt, nil
	}
}

var weekDay = map[string]int{
	"Sunday":    0,
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
}

func FindWeekTimeByTime(t time.Time) (int64, int64) {
	end := (time.Saturday - t.Weekday()).String()
	plus := int64(86400 * weekDay[end])
	owe := int64(86400 * (6 - 1 - weekDay[end]))
	var endDay time.Time
	var startDay time.Time
	endDay = GetZeroTime(time.Unix(t.Unix()+plus+86399, 0))
	startDay = GetZeroTime(time.Unix(t.Unix()-owe, 0))
	return endDay.Unix(), startDay.Unix()
}

//获取当天和明天的零点时间戳
func GetTimestamp() (beginTimeNum float64) {
	//, endTimeNum
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	beginTimeNum = float64(t.Unix())
	//endTimeNum = beginTimeNum + 86400
	return beginTimeNum //, endTimeNum
}
func GetUnixNowTime(day int, hour int) int64 {

	currentday := time.Now().Day() + day
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	i := time.Date(currentYear, currentMonth, currentday, 0+hour, 0, 0, 0, time.Local)
	// fmt.Printf("i: %v\n", i)
	return i.Unix()
}
func MonthZero() string {
	ts := time.Now().AddDate(0, 0, -1)
	timeYesterday := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location()).Unix()
	timeStr := time.Unix(timeYesterday, 0).Format("200601")
	// fmt.Println(timeStr) // 2022-04-14
	return timeStr
}
func MonthSubOne() string {
	month := MonthZero()
	monthNum, err := strconv.Atoi(month)
	if err != nil {
		seelog.Error("MonthSubOne err")
		return ""
	}
	monthNum--
	return strconv.Itoa(monthNum)
}
func Month() string {
	month := MonthZero()
	monthNum, err := strconv.Atoi(month)
	if err != nil {
		seelog.Error("MonthSubOne err")
		return ""
	}
	return strconv.Itoa(monthNum)
}

func GetTimeByUnix(i int) time.Time {
	return time.Unix(int64(i), 0)
	// dateStr := t.Format("2006/01/02 15:04:05")
}
func GetLastMonthDay() {

}
