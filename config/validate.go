package config

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

func (dbCfg *DBConfig) Validate() error {
	if dbCfg.Port <= 0 || dbCfg.Port > 65535 {
		return fmt.Errorf("DB_PORT must be in the range 1-65535, got: %d", dbCfg.Port)
	}

	if strings.TrimSpace(dbCfg.Host) == "" {
		return errors.New("DB_HOST is required")
	}

	if strings.TrimSpace(dbCfg.User) == "" {
		return errors.New("DB_USER is required")
	}
	if strings.TrimSpace(dbCfg.Name) == "" {
		return errors.New("DB_NAME is required")
	}

	if strings.TrimSpace(dbCfg.Password) == "" {
		return errors.New("DB_PASSWORD is required (compiled with required tag)")
	}

	validSSLModes := []string{"disable", "enable", "prefer", "require"}
	isValid := slices.Contains(validSSLModes, dbCfg.SSLMode)
	if !isValid {
		return fmt.Errorf("DB_SSLMODE must be one of: %s", strings.Join(validSSLModes, ", "))
	}

	validDBTypes := []string{"postgres", "mysql", "mssql", "sqlite"}
	isAllowed := slices.Contains(validDBTypes, dbCfg.Type)
	if !isAllowed {
		return fmt.Errorf("DB_TYPE must be one of: %s", strings.Join(validDBTypes, ", "))
	}

	return nil
}

func (srvCfg *ServerConfig) Validate() error {
	if srvCfg.Port <= 0 || srvCfg.Port > 65535 {
		return fmt.Errorf("SERVER_PORT must be in the range 1-65535, got: %d", srvCfg.Port)
	}

	if strings.TrimSpace(srvCfg.Host) == "" {
		return errors.New("SERVER_HOST is required")
	}

	return nil
}
