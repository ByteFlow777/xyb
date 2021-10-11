package notice

import (
	"fmt"
	"net/http"
	"xyb/config"
	"xyb/util"
)

type Notice struct {
	title, desc string
	t           string
}

func NewNotice(title, desc string) *Notice {
	return &Notice{
		title: "",
		desc:  "",
		t:     util.Now(),
	}
}

func (n *Notice) Send() {
	url := fmt.Sprintf("https://sc.ftqq.com/%s.send?text=%s&desp=%s",
		config.Conf.ServerChan,
		fmt.Sprintf("%s %s", n.t, n.title), n.desc)

	_, _ = http.Get(url)
}

func Error(err error) {
	NewNotice("发生错误", err.Error()).Send()
}

func Success(desc string) {
	NewNotice("成功了", desc).Send()
}
