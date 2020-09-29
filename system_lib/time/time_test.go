package time

import (
	"testing"
	"time"
)

var loc = local()

func local() *time.Location {
	loc, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		panic(err)
	}
	return loc
}


func Test_tst(t *testing.T) {
	tarTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-04-07 21:03:32", loc)
	if err != nil {
		t.Error(err)
	}
	now := time.Now()
	t.Log("target",tarTime.Format("2006-01-02 15:04:05"), tarTime.Unix())
	t.Log("now",now.Format("2006-01-02 15:04:05"), now.Unix())
	t.Log(tarTime.Sub(now))
	if tarTime.Before(now) {
		t.Error("time dur too small")
		return
	}
	t.Log("success")
}

func Test_Time(t *testing.T) {
	page := 1
	number := 5

	tm := time.Now()

	sub := time.Now().Weekday() - time.Monday
	add := (time.Sunday + 7) - time.Now().Weekday()
	t.Log("before",tm.Format("2006-01-02 15:04"))
	start := tm.AddDate(0,0, -int(sub))
	t.Log("一周开始", start.Format("2006-01-02 15:04"))
	end := tm.AddDate(0,0, int(add))
	t.Log("一周结束", end.Format("2006-01-02 15:04"))
	t.Log(time.Now().Weekday())

	//page 是起始的周
	//page 1 当前周 number 5 取5周
	now := time.Now()

	zz :=  page * number * 7

	//所在最早周
	now.AddDate(0,0, -zz)
}