package config

import (
	"encoding/json"
	"io/ioutil"
)

// TokenInit ...
func TokenInit() (conf *Token, err error) {
	var file []byte
	file, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	conf = &Token{}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		return nil, err
	}
	return
}
