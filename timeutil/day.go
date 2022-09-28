package timeutil

import "time"

const (
	TimeFormat     = "2006-01-02 15:04:05"
	TimeDateFormat = "2006-01-02"
)

//获取该时间所在的月份的第一天
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return d
}

//获取该时间所在的月份的最后一天
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//
func WeekFirstDay() time.Time {
	n := time.Now()
	day := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local) // 当天0点
	day = day.AddDate(0, 0, 0-(int(day.Weekday())+7)%7)                    // 周日一0点
	return day
}

type LiveTime struct {
	Start int64
	End   int64
}

//判断两个时间段是否重合
func checkoutTimeOverlap(t1, t2 LiveTime) bool {
	return t1.Start < t2.End && t1.End > t2.Start
}
