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
	Host string `env:"SERVER_HOST" envDefault:"localhost"`
	Port int    `env:"SERVER_PORT" envDefault:"8080"`
}

type DBConfig struct {
	Type     string `env:"DB_TYPE" envDefault:"postgres"`
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER" envDefault:"user"`
	Password string `env:"DB_PASSWORD,required"`
	Name     string `env:"DB_NAME" envDefault:"db"`
	SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

type AppConfig struct {
	SrvCfg ServerConfig
	DBCfg  DBConfig
}

// Однократное выполнение загрузки конфигурации (уточни у Макса)
var (
	once    sync.Once
	cfg     *AppConfig
	loadErr error
)

func MustLoad() (*AppConfig, error) {
	once.Do(func() {
		srvCfg := &ServerConfig{}
		dbCfg := &DBConfig{}

		configs := []Config{srvCfg, dbCfg}
		for _, c := range configs {
			if err := c.Load(); err != nil {
				slog.Error("Ошибка при загрузке конфигурации", "error", err)
				loadErr = fmt.Errorf("не удалось загрузить конфигурацию: %w", err)
				return
			}
			if err := c.Validate(); err != nil {
				slog.Error("Ошибка валидации конфигурации", "error", err)
				loadErr = fmt.Errorf("не удалось валидировать конфигурацию: %w", err)
				return
			}
		}

		cfg = &AppConfig{
			SrvCfg: *srvCfg,
			DBCfg:  *dbCfg,
		}

		slog.Info("Все конфигурации успешно загружены")
	})

	return cfg, loadErr
}
