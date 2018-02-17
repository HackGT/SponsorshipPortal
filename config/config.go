package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host         string        `default:""`
	Port         int           `default:"3000"`
	WriteTimeout time.Duration `default:"15s"`
	ReadTimeout  time.Duration `default:"15s"`
	IdleTimeout  time.Duration `default:"60s"`
	ShutdownWait time.Duration `default:"15s"`
	Prod         bool          `default:"false"`
}

func Load() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (config *Config) Addr() string {
	return fmt.Sprintf("%s:%v", config.Host, config.Port)
}
