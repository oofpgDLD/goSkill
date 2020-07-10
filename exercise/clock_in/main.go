package main

import (
	"github.com/limetext/backend/log"
	"gitlab.33.cn/chat/work/library/work33"
	work33_model "gitlab.33.cn/chat/work/library/work33/model"
	"gitlab.33.cn/chat/work/service/attendance"
	"gitlab.33.cn/chat/work/service/record/model"
	"os"
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

func main() {
	work33.Init(&work33_model.Config{
		Host: "http://47.105.41.71:8082",
		Token: "Y2hhdDMzX3Rva2Vu",
	})

	tarTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-08 20:31:40", loc)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	log.Info(tarTime.Format("2006-01-02 15:04:05"))
	log.Info(time.Now().Format("2006-01-02 15:04:05"))
	now := time.Now()
	dur := tarTime.Sub(now)
	if dur <= 0 {
		panic("time dur too small")
		os.Exit(1)
	}
	t:= time.NewTimer(dur)
	<- t.C
	log.Info("start clock in")
	createTime := time.Now().Unix()
	loc := &model.Location{
		Address: "浙江省杭州市西湖区马塍路6号楼靠近东部软件园(马塍路)",
		Longitude: 120.146157,
		Latitude: 30.279045,
	}
	result, err := attendance.ClockIn("戴笠东", createTime, loc, "")
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("success", "result", result)
}