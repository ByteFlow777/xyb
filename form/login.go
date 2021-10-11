package form

import (
	"xyb/config"
)

func NewLoginForm() *Forms {
	return &Forms{
		Form{"openId", config.Conf.WechatInfo.OpenId},
		Form{"unionId", config.Conf.WechatInfo.UnionId},
	}
}
