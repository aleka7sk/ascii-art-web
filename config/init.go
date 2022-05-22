package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port string
}

func InitConfig() (Config, error) {
	bytes, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
