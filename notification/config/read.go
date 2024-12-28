package config

import (
	"encoding/json"
	"os"
)

func ReadConfig(path string) (NatsConfig, error) {
	var config NatsConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return NatsConfig{}, err
	}
	return config, json.Unmarshal(data, &config)
}

func MustReadConfig(path string) NatsConfig {
	config, err := ReadConfig(path)
	if err != nil {
		panic(err)
	}
	return config
}
