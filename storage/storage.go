package storage

import (
	"database/sql"
	"subscription-aggregator-api/models"
)

type SQLStorage struct {
	db *sql.DB
}

func NewSQL(db *sql.DB) *SQLStorage {
	return &SQLStorage{
		db: db,
	}
}

func (p *SQLStorage) Create(subscription models.Subscription) error {
	// TODO: INSERT-запрос
	return nil
}

func (p *SQLStorage) GetList() ([]models.Subscription, error) {
	// TODO: SELECT-запрос
	return nil, nil
}

func (p *SQLStorage) GetByID(id string) (models.Subscription, error) {
	// TODO: SELECT-запрос по ID
	var sub models.Subscription
	return sub, nil
}

func (p *SQLStorage) Update(updated models.Subscription) (models.Subscription, error) {
	// TODO: UPDATE-запрос
	return updated, nil
}

func (p *SQLStorage) Delete(id string) error {
	// TODO: DELETE-запрос
	return nil
}
