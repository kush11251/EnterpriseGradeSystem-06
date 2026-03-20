package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

var Config struct {
	Port string `json:"port"`
	Env  string `json:"env"`
}

func Init() {
	flag.Parse()
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(file, &Config)
}