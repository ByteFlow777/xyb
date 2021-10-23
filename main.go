package main

import (
	"github.com/robfig/cron/v3"
	"math/rand"
	"os"
	"time"
	"xyb/form"
	"xyb/notice"
	"xyb/request"
)

func Sign() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Second * 60 * time.Duration(r.Intn(30)))
	InitServiceInfo(os.Args[1])
	println(os.Args[1])
	var cookie = SystemInfo.Cookie
	var info = &SystemInfo

	status := request.GetSignStatus(info.Cookie, info.TraineeId)

	var err error
	signForm := form.NewSignForm(&SystemInfo)
	if status {
		err = request.ReSign(signForm, cookie)
	} else {
		err = request.NewSign(signForm, cookie)
	}
	if err != nil {
		notice.Error(err)
		return
	}
	notice.Success("签到成功")
}

func main() {
	ch := make(chan struct{})
	c := cron.New()
	c.AddFunc("30 8 * * *", func() {
		Sign()
	})
	c.AddFunc("00 18 * * *", func() {
		Sign()
	})
	c.Start()
	<-ch
}
