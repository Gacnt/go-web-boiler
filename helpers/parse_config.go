package hlp

import (
	"encoding/json"
	"fmt"
	"os"
)

var AppConfig *Config

type Config struct {
	ViewPath string `json:"view_path"`
}

func ParseConfig() *Config {
	if AppConfig != nil {
		return AppConfig
	}
	file, err := os.Open("./config.json")
	fileInfo, err := file.Stat()
	configFile := make([]byte, fileInfo.Size())
	_, err = file.Read(configFile)
	if err != nil {
		panic(fmt.Errorf("%s: %v", "Could not open config file", err))
	}
	defer file.Close()

	config := &Config{}
	err = json.Unmarshal(configFile, config)
	if err != nil {
		panic(err)
	}

	return config
}
