package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"subscription-aggregator-api/config"
	"sync"
)

type DBManager struct {
	once sync.Once
	DB   *sql.DB
	err  error
}

func NewDBManager() *DBManager {
	return &DBManager{}
}

func (dm *DBManager) InitDB(dbCfg config.DBConfig) error {
	dm.once.Do(func() {
		conn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.Name, dbCfg.SSLMode)

		dm.DB, dm.err = sql.Open(dbCfg.Type, conn)
		if dm.err != nil {
			slog.Error("Ошибка открытия БД", "err", dm.err)
			dm.err = fmt.Errorf("не удалось открыть подключение к базе данных: %w", dm.err)
			return
		}

		if pingErr := dm.DB.Ping(); pingErr != nil {
			dm.DB.Close()
			dm.err = fmt.Errorf("не удалось выполнить ping к базе данных: %w", pingErr)
			slog.Error("Ошибка пинга БД", "err", pingErr)
			return
		}

		if err := dm.createTables(); err != nil {
			dm.err = err
			return
		}

		slog.Info("Подключено к БД", "type", dbCfg.Type, "name", dbCfg.Name, "host", dbCfg.Host, "port", dbCfg.Port)

	})
	return dm.err
}

func (dbManager *DBManager) createTables() error {
	query :=
		`CREATE TABLE IF NOT EXISTS subscriptions (
			id SERIAL PRIMARY KEY,
			service_name VARCHAR(50) NOT NULL,
			user_id VARCHAR(50) NOT NULL,
			price INT NOT NULL,
			start_date VARCHAR(7) NOT NULL
		);`

	if _, err := dbManager.DB.Exec(query); err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	slog.Info("Table 'subscriptions' created")

	return nil
}

func (dm *DBManager) Close() error {
	if dm.DB != nil {
		return dm.DB.Close()
	}
	return nil
}
