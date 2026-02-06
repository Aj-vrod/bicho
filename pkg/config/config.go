package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPPort int `envconfig:"HTTP_PORT" default:"8080"`
}

func LoadFromEnv() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
