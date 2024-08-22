package db

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}
type Config struct {
	Database `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	file, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
