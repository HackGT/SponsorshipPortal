package config

import (
	"fmt"
	nurl "net/url"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Prod     bool `default:"false"`
	Server   *ServerConfig
	Database *DatabaseConfig
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
	config.Database, err = LoadDatabaseConfig()
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
	Host             string `default="localhost:5432"`
	DbName           string `default="portal"`
	User             string
	Password         string
	ConnectionString string
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	var config DatabaseConfig
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	if config.ConnectionString == "" {

		var userinfo *nurl.Userinfo
		if config.User != "" {
			if config.Password == "" {
				userinfo = nurl.User(config.User)
			} else {
				userinfo = nurl.UserPassword(config.User, config.Password)
			}
		}
		dbUrl := &nurl.URL{
			Scheme: "postgres",
			Host:   config.Host,
			User:   userinfo,
			Path:   config.DbName,
		}
		config.ConnectionString = dbUrl.String()
	} else {
		dbUrl, err := nurl.Parse(config.ConnectionString)
		if err != nil {
			// TODO parse the "<key>=<value>[ <key>=<value>...]" format
		} else {
			config.Host = dbUrl.Host
			if dbUrl.User != nil {
				config.User = dbUrl.User.Username()
				config.Password, _ = dbUrl.User.Password()
			}
		}
	}

	return &config, nil
}
