package request

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"xyb/form"
	"xyb/model"
)

const (
	LoginUrl      = "https://xcx.xybsyw.com/login/login!wx.action"
	PersonInfoUrl = "https://xcx.xybsyw.com/account/LoadAccountInfo.action"
	TraineeIdUrl  = "https://xcx.xybsyw.com/student/clock/GetPlan!getDefault.action"
	PositionUrl   = "https://xcx.xybsyw.com/student/clock/GetPlan!detail.action"
	NewSignUrl    = "https://xcx.xybsyw.com/student/clock/Post!autoClock.action"
	ReSignUrl     = "https://xcx.xybsyw.com/student/clock/Post!updateClock.action"
	SignStatusUrl = "https://xcx.xybsyw.com/student/clock/GetPlan!detail.action"
	PostBlogUrl   = "https://xcx.xybsyw.com/student/blog/Blog!save.action"
)

func basicReq(url, cookie string, form io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, form)
	if err != nil {
		return nil, err
	}

	AddHeadersToReq(req, cookie)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Login(form *form.Forms) (*model.LoginInfo, error) {
	body, err := basicReq(LoginUrl, "",
		form.Reader())
	if err != nil {
		return nil, err
	}

	var r model.LoginInfo
	_ = json.Unmarshal(body, &r)
	return &r, nil
}

func GetPersonInfo(cookie string) (*model.PersonInfo, error) {
	body, err := basicReq(PersonInfoUrl, cookie, nil)
	if err != nil {
		return nil, err
	}

	var p model.PersonInfo
	_ = json.Unmarshal(body, &p)
	return &p, nil
}

func GetTraineeId(cookie string) (*model.TraineeInfo, error) {
	body, err := basicReq(TraineeIdUrl, cookie, nil)
	if err != nil {
		return nil, err
	}

	var t model.TraineeInfo
	_ = json.Unmarshal(body, &t)
	return &t, nil
}

func GetPosition(cookie, traineeId string) (*model.PositionInfo, error) {
	body, err := basicReq(PositionUrl, cookie,
		strings.NewReader("traineeId="+traineeId))
	if err != nil {
		return nil, err
	}

	var p model.PositionInfo
	_ = json.Unmarshal(body, &p)
	return &p, nil
}

func GetSignStatus(cookie, traineeId string) bool {
	body, err := basicReq(SignStatusUrl, cookie,
		strings.NewReader("traineeId="+traineeId))
	if err != nil {
		return false
	}

	var s model.SignStatusInfo
	_ = json.Unmarshal(body, &s)

	if s.Data.ClockInfo.InTime == "" {
		return false
	}
	return true
}

func ReSign(signForm *form.Forms, cookie string) error {

	body, err := basicReq(ReSignUrl, cookie,
		signForm.Reader())

	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}

func NewSign(signForm *form.Forms, cookie string) error {

	body, err := basicReq(NewSignUrl, cookie,
		signForm.Reader())

	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
