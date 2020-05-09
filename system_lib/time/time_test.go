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