package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func GenMD5(str string) string {
	w := md5.New()
	_, err := io.WriteString(w, str)
	if err != nil {
		return ""
	}
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func Now() string {
	return time.Now().Format("2006.01.02")
}

func WeekDay() (string, string) {
	var WeekDayMap = map[string]string{
		"Monday":    "周一",
		"Tuesday":   "周二",
		"Wednesday": "周三",
		"Thursday":  "周四",
		"Friday":    "周五",
		"Saturday":  "周六",
		"Sunday":    "周日",
	}
	now := time.Now()
	return now.Format("2006.01.02"), WeekDayMap[now.Weekday().String()]
}

func GenRandom(end int) int {
	r := rand.New(rand.NewSource(time.Now().Unix() + rand.Int63()))
	return r.Intn(end)
}

func ReadFromPath(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(body)
}
