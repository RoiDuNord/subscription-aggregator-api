package storage

import (
	"database/sql"
	"fmt"
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

func (storage *SQLStorage) Create(subscription models.Subscription) error {
	query := `
		INSERT INTO subscriptions (service_name, user_id, price, start_date)
		VALUES ($1, $2, $3, $4)
	`

	_, err := storage.db.Exec(query,
		subscription.ServiceName,
		subscription.UserID,
		subscription.Price,
		subscription.StartDate,
	)

	return err
}

func (storage *SQLStorage) GetList() ([]models.Subscription, error) {
	query := `SELECT user_id, service_name, price, start_date, end_date FROM subscriptions`

	rows, err := storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []models.Subscription

	for rows.Next() {
		var subscription models.Subscription
		err := rows.Scan(&subscription.UserID, &subscription.ServiceName, &subscription.Price, &subscription.StartDate)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (storage *SQLStorage) GetByID(id string) (models.Subscription, error) {
	var subscription models.Subscription

	query := `
		SELECT * FROM subscriptions
		WHERE id = $1
	`

	result := storage.db.QueryRow(query, id)

	err := result.Scan(&subscription.UserID, &subscription.ServiceName, &subscription.Price, &subscription.StartDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Subscription{}, fmt.Errorf("subscription with id %s not found", id)
		}
		return models.Subscription{}, err

	}

	return subscription, nil
}

func (storage *SQLStorage) Update(id string, updatedSubscription models.Subscription) error {
	query := `
		UPDATE subscriptions
		SET user_id = $2, service_name = $3, price = $4, start_date = $5
		WHERE id = $1
	`

	result, err := storage.db.Exec(query, id,
		updatedSubscription.UserID,
		updatedSubscription.ServiceName,
		updatedSubscription.Price,
		updatedSubscription.StartDate,
		updatedSubscription.EndDate,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("subscription with id %s not found or no changes made", id)
	}

	return nil
}

func (storage *SQLStorage) Delete(id string) error {
	query := `
		DELETE FROM subscriptions
		WHERE id = $1
	`

	result, err := storage.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("subscription with id %s not found or no changes made", id)
	}

	return nil
}
