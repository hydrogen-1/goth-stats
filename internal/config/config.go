package config

import (
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbUser         string `envconfig:"DB_USER" default:"postgres"`
	DbPassword     string `envconfig:"DB_PASSWORD" default:"secret"`
	DbPasswordFile string `envconfig:"DB_PASSWORD_FILE" default:"none"`
	DbHost         string `envconfig:"DB_HOST" default:"localhost"`
}

func ParseConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	if config.DbPasswordFile != "none" {
		password, err := os.ReadFile(config.DbPasswordFile)
		if err != nil {
			return nil, err
		}
		config.DbPassword = strings.TrimSpace(string(password))
	}
	return &config, nil
}
