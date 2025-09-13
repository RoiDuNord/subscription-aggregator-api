package config

import (
	"fmt"
	"log/slog"
	"sync"
)

type Config interface {
	Load() error
	Validate() error
}

type ServerConfig struct {
	Host string `env:"SERVER_HOST"`
	Port int    `env:"SERVER_PORT"`
}

type DBConfig struct {
	Type     string `env:"DB_TYPE"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	SSLMode  string `env:"DB_SSLMODE"`
}

type AppConfig struct {
	once    sync.Once
	SrvCfg  ServerConfig
	DBCfg   DBConfig
	loadErr error
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		SrvCfg: ServerConfig{},
		DBCfg:  DBConfig{},
	}
}

func (appcfg *AppConfig) MustLoad() error {
	appcfg.once.Do(func() {
		configs := []Config{&appcfg.SrvCfg, &appcfg.DBCfg}
		for _, c := range configs {
			if err := c.Load(); err != nil {
				slog.Error("Error loading configuration", "error", err)
				appcfg.loadErr = fmt.Errorf("failed to load configuration: %w", err)
				return
			}
		}

		slog.Info("Configuration data loaded successfully")
	})

	return appcfg.loadErr
}
