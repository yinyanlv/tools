package timer

import (
	"fmt"
	"time"
)

const (
	NanoSecond  time.Duration = 1
	MicroSecond               = NanoSecond * 1000
	MillsSecond               = MicroSecond * 1000
	Second                    = MillsSecond * 1000
	Minute                    = Second * 60
	Hour                      = Minute * 60
	Day                       = Hour * 24
)

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, nil
	}
	return currentTime.Add(duration), nil
}

// time.Parse默认把字符串当作UTC时间
// time.In会做时间转换，time.Format只会格式化时间不会自动做时间转换
// t.Unix只会返回零时区的UTC时间戳，和时区无关
// time.Unix返回当前时间戳对应的本地时间
func TestFormatTime() {
	var layout = "2006-01-02 15:04:05"

	fmt.Println("-1-")

	r1, _ := time.LoadLocation("")
	r2, _ := time.LoadLocation("Asia/Shanghai")
	// 本地当前时间
	fmt.Println(time.Now())
	// UTC当前时间
	fmt.Println(time.Now().In(r1))
	// 上海时区当前时间
	fmt.Println(time.Now().In(r2))

	fmt.Println("-2-")

	// 默认把字符串当作UTC时间
	t1, _ := time.Parse(layout, "2020-11-07 15:10:10")
	fmt.Println(t1.Unix()) // 1604761810
	fmt.Println(t1)
	fmt.Println(t1.Format(layout))
	fmt.Println(t1.In(r1).Format(layout))
	fmt.Println(t1.In(r2).Format(layout))               // 不一致，2020-11-07 23:10:10
	fmt.Println(time.Unix(t1.Unix(), 0).Format(layout)) // 不一致，2020-11-07 23:10:10
	fmt.Println(time.Unix(t1.Unix(), 0).In(r1).Format(layout))
	fmt.Println(time.Unix(t1.Unix(), 0).In(r2).Format(layout)) // 不一致，2020-11-07 23:10:10

	fmt.Println("-3-")

	t2, _ := time.ParseInLocation(layout, "2020-11-07 15:10:10", r2)
	fmt.Println(t2.Unix()) // 1604733010
	fmt.Println(t2)
	fmt.Println(t2.Format(layout))
	fmt.Println(t2.In(r1).Format(layout)) // 不一致，2020-11-07 07:10:10
	fmt.Println(t2.In(r2).Format(layout))
	fmt.Println(time.Unix(t2.Unix(), 0).Format(layout))
	fmt.Println(time.Unix(t2.Unix(), 0).In(r1).Format(layout)) // 不一致，2020-11-07 07:10:10
	fmt.Println(time.Unix(t2.Unix(), 0).In(r2).Format(layout))
}
