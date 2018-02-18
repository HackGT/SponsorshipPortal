package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Prod   bool `default:"false"`
	Server *ServerConfig
	DB     *DatabaseConfig
}

func Load() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	config.Server, err = LoadServerConfig()
	if err != nil {
		return nil, err
	}
	config.DB, err = LoadDatabaseConfig()
	if err != nil {
		return nil, err
	}

	return &config, nil
}

type ServerConfig struct {
	Host         string        `default:""`
	Port         int           `default:"3000"`
	WriteTimeout time.Duration `default:"15s"`
	ReadTimeout  time.Duration `default:"15s"`
	IdleTimeout  time.Duration `default:"60s"`
	ShutdownWait time.Duration `default:"15s"`
}

func LoadServerConfig() (*ServerConfig, error) {
	var config ServerConfig
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (config *ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%v", config.Host, config.Port)
}

type DatabaseConfig struct {
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	var config DatabaseConfig
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
