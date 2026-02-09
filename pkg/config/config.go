package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPPort int    `envconfig:"HTTP_PORT" default:"8080"`
	DBDNS    string `envconfig:"DB_DNS" default:"postgres://bicho:bicho-pwd@localhost:5432/bicho?sslmode=disable"`
}

func LoadFromEnv() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
