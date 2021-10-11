package form

import (
	"fmt"
	"strings"
)

type Form struct {
	key string
	val string
}

type Forms []Form

func (f *Forms) Reader() *strings.Reader {
	return strings.NewReader(f.String())
}

func (f *Forms) String() string {
	s := ""
	for _, form := range *f {
		s += fmt.Sprintf("%s=%s&", form.key, form.val)
	}
	// 去除掉末尾多余的一个 & 号
	return s[:len(s)-1]
}
