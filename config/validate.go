package config

import (
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

func (dbCfg *DBConfig) Validate() error {
	if dbCfg.Port <= 0 || dbCfg.Port > 65535 {
		return fmt.Errorf("DB_PORT должен быть в диапазоне 1-65535, получено: %d", dbCfg.Port)
	}

	if strings.TrimSpace(dbCfg.Host) == "localhost" && dbCfg.Port == 5432 {
		slog.Warn("Используются дефолтные DB_HOST и DB_PORT")
	}

	if strings.TrimSpace(dbCfg.Host) == "" {
		return errors.New("DB_HOST обязателен")
	}

	if strings.TrimSpace(dbCfg.User) == "" {
		return errors.New("DB_USER обязателен")
	}
	if strings.TrimSpace(dbCfg.Name) == "" {
		return errors.New("DB_NAME обязателен")
	}

	if strings.TrimSpace(dbCfg.Password) == "" {
		return errors.New("DB_PASSWORD обязателен (скомпилировано через required)")
	}

	validSSLModes := []string{"disable", "enable", "prefer", "require"}
	isValid := slices.Contains(validSSLModes, dbCfg.SSLMode)
	if !isValid {
		return fmt.Errorf("DB_SSLMODE должен быть одним из: %s", strings.Join(validSSLModes, ", "))
	}

	validDBTypes := []string{"postgres", "mysql", "mssql", "sqlite"}
	isAllowed := slices.Contains(validDBTypes, dbCfg.Type)
	if !isAllowed {
		return fmt.Errorf("DB_TYPE должен быть одним из: %s", strings.Join(validDBTypes, ", "))
	}

	return nil
}

func (srvCfg *ServerConfig) Validate() error {
	if srvCfg.Port <= 0 || srvCfg.Port > 65535 {
		return fmt.Errorf("SERVER_PORT должен быть в диапазоне 1-65535, получено: %d", srvCfg.Port)
	}

	if srvCfg.Host == "localhost" && srvCfg.Port == 8080 {
		slog.Warn("Используются дефолтные SERVER_HOST и SERVER_PORT")
	}

	if strings.TrimSpace(srvCfg.Host) == "" {
		return errors.New("SERVER_HOST обязателен")
	}

	return nil
}
