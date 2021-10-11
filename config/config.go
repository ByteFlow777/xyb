package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"xyb/model"
)

var Conf *model.Config

func init() {
	Conf = ReadConfig()
}

func ReadFromPath(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ReadConfig() *model.Config {
	bytes, err := ReadFromPath(Configuration)
	if err != nil {
		panic(err)
	}

	var config model.Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
