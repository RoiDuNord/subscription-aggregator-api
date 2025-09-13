package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
)

func (dbCfg *DBConfig) Load() error {
	if err := env.Parse(dbCfg); err != nil {
		return fmt.Errorf("error loading DBConfig from env: %w", err)
	}
	if err := dbCfg.Validate(); err != nil {
		return fmt.Errorf("error validating DBConfig: %w", err)
	}
	return nil
}

func (srvCfg *ServerConfig) Load() error {
	if err := env.Parse(srvCfg); err != nil {
		return fmt.Errorf("error loading ServerConfig from env: %w", err)
	}
	if _, ok := os.LookupEnv("SERVER_PORT"); !ok {
		return fmt.Errorf("environment variable SERVER_PORT is required")
	}
	if _, ok := os.LookupEnv("SERVER_HOST"); !ok {
		return fmt.Errorf("environment variable SERVER_HOST is required")
	}
	if err := srvCfg.Validate(); err != nil {
		return fmt.Errorf("error validating ServerConfig: %w", err)
	}
	return nil
}
