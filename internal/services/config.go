package services

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SecretKey string `yaml:"secret_key"`
}

func loadConfig() ([]byte, error) {
	file, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, err
	}
	var config Config

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return []byte(config.SecretKey), nil
}
