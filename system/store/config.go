package store

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const path = "conf/dbconfig.yml"

func ReadConfig() (conf *Config, err error) {
	var file []byte
	file, err = ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error reading config file")
		return
	}
	conf = &Config{}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		log.Fatal("Error: wrong format of config file")
		return
	}
	return
}

