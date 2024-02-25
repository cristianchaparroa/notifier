package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App `yaml:"app"`
	}

	App struct {
		Port string `yaml:"port" env:"APP_PORT"`
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
