package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

func (dbCfg *DBConfig) Load() error {
	if err := env.Parse(dbCfg); err != nil {
		return fmt.Errorf("ошибка загрузки DBConfig из env: %w", err)
	}
	if err := dbCfg.Validate(); err != nil {
		return fmt.Errorf("ошибка валидации DBConfig: %w", err)
	}
	return nil
}

func (srvCfg *ServerConfig) Load() error {
	if err := env.Parse(srvCfg); err != nil {
		return fmt.Errorf("ошибка загрузки ServerConfig из env: %w", err)
	}
	if err := srvCfg.Validate(); err != nil {
		return fmt.Errorf("ошибка валидации ServerConfig: %w", err)
	}
	return nil
}
