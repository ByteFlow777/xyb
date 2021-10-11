package model

type WechatInfo struct {
	OpenId  string `json:"openId"`
	UnionId string `json:"unionId"`
}

type ServiceInfo struct {
	UserName  string `json:"userName"`
	UserId    string `json:"userId"`
	Cookie    string `json:"cookie"`
	TraineeId string `json:"traineeId"`
}

type Position struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Location struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Adcode   string `json:"adcode"`
	Address  string `json:"address"`
}

type Config struct {
	WechatInfo WechatInfo `json:"wechatInfo"`
	Location   Location   `json:"location"`
	ServerChan string     `json:"ServerChan"`
	Version    string     `json:"version"`
	ClientIp   string     `json:"clientIp"`
	PageId     string     `json:"pageId"`
}

type SystemInfo struct {
	ServerChan string `json:"serverChan"`
	Version    string `json:"version"`
	WechatInfo
	Location

	Position
	ServiceInfo
}
