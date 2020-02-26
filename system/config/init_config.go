package config

import (
	"encoding/json"
	"io/ioutil"
)

// Path ...
const path = "conf/config.json"

// ReadConfig ...
func ReadConfig() (conf *AppConfig, err error) {
	var file []byte
	file, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	conf = &AppConfig{}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		return nil, err
	}
	return
}
