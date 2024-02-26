package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvProduction = "production"
)

type (
	Config struct {
		App  `yaml:"app"`
		SMTP `yaml:"smtp"`
	}

	App struct {
		Env  string `yaml:"env" env:"APP_ENV"`
		Port string `yaml:"port" env:"APP_PORT"`
	}

	SMTP struct {
		Server   string `yaml:"server" env:"SMTP_SERVER"`
		Port     int    `yaml:"port" env:"SMTP_PORT"`
		User     string `yaml:"user" env:"SMTP_USER"`
		Password string `yaml:"password" env:"SMTP_PASSWORD"`
	}
)

// NewConfiguration will create a config struct from environment variables
func NewConfiguration() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return cfg, nil
}
