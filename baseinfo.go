package main

import (
	"strconv"
	"xyb/config"
	"xyb/form"
	"xyb/model"
	"xyb/notice"
	"xyb/request"
)

var SystemInfo model.SystemInfo

var conf = config.Conf

func InitServiceInfo() {
	loginResponse, err := request.Login(form.NewLoginForm())
	if err != nil {
		notice.Error(err)
		return
	}
	personInfo, err := request.GetPersonInfo(loginResponse.Data.SessionId)
	if err != nil {
		notice.Error(err)
		return
	}
	trainee, err := request.GetTraineeId(loginResponse.Data.SessionId)
	if err != nil {
		notice.Error(err)
		return
	}
	position, err := request.GetPosition(loginResponse.Data.SessionId,
		strconv.Itoa(trainee.Data.ClockVo.TraineeId))
	if err != nil {
		notice.Error(err)
		return
	}

	SystemInfo.WechatInfo = conf.WechatInfo
	SystemInfo.Location = conf.Location
	SystemInfo.Version = conf.Version
	SystemInfo.ServerChan = conf.ServerChan

	SystemInfo.Position = model.Position{
		Lat: strconv.FormatFloat(position.Data.PostInfo.Lat, 'f', 5, 64),
		Lng: strconv.FormatFloat(position.Data.PostInfo.Lng, 'f', 5, 64),
	}

	SystemInfo.ServiceInfo = model.ServiceInfo{
		UserName:  personInfo.Data.Loginer,
		UserId:    strconv.Itoa(loginResponse.Data.LoginerId),
		Cookie:    loginResponse.Data.SessionId,
		TraineeId: strconv.Itoa(trainee.Data.ClockVo.TraineeId),
	}
}
