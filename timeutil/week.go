package timeutil

import (
	"fmt"
	"time"
)

//计算某天是星期几的
func Test(t time.Time) int {
	year := t.Year()
	month := t.Month()
	day := t.Day()
	fmt.Println(int(t.Weekday()))
	k := year % 100
	fmt.Println(k)
	j := year / 100
	return (day + int(26*(month+1)/10) + k + k/4 + j/4 + 5) % 7
}
