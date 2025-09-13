package storage

import (
	"database/sql"
	"errors"
	"subscription-aggregator-api/models"
)

var (
	ErrSubscriptionNotFound = errors.New("subscription not found")
	ErrNoSubscriptions      = errors.New("no subscriptions found")
)

type SQLStorage struct {
	db *sql.DB
}

func NewSQL(db *sql.DB) *SQLStorage {
	return &SQLStorage{
		db: db,
	}
}

func (s *SQLStorage) Create(subscription models.Subscription) error {
	query := `
		INSERT INTO subscriptions (service_name, user_id, price, start_date)
		VALUES ($1, $2, $3, $4);
	`

	result, err := s.db.Exec(query,
		subscription.ServiceName,
		subscription.UserID,
		subscription.Price,
		subscription.StartDate,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrSubscriptionNotFound
	}

	return nil
}

func (s *SQLStorage) GetByID(id int) (models.Subscription, error) {
	var subscription models.Subscription

	query := `
		SELECT id, user_id, service_name, price, start_date
		FROM subscriptions
		WHERE id = $1;
	`

	result := s.db.QueryRow(query, id)

	err := result.Scan(&subscription.ID, &subscription.UserID, &subscription.ServiceName, &subscription.Price, &subscription.StartDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Subscription{}, ErrSubscriptionNotFound
		}
		return models.Subscription{}, err
	}

	return subscription, nil
}

func (s *SQLStorage) GetList() ([]models.Subscription, error) {
	query := `
		SELECT id, user_id, service_name, price, start_date
		FROM subscriptions;
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []models.Subscription

	for rows.Next() {
		var subscription models.Subscription
		err := rows.Scan(&subscription.ID, &subscription.UserID, &subscription.ServiceName, &subscription.Price, &subscription.StartDate)
		if err != nil {
			return nil, err
		}
		subscription.StartDate = subscription.StartDate[:10]
		subscriptions = append(subscriptions, subscription)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(subscriptions) == 0 {
		return nil, ErrNoSubscriptions
	}

	return subscriptions, nil
}

func (s *SQLStorage) Update(id int, updatedSubscription models.Subscription) error {
	query := `
		UPDATE subscriptions
		SET user_id = $2, service_name = $3, price = $4, start_date = $5
		WHERE id = $1;
	`

	result, err := s.db.Exec(query, id,
		updatedSubscription.UserID,
		updatedSubscription.ServiceName,
		updatedSubscription.Price,
		updatedSubscription.StartDate,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrSubscriptionNotFound
	}

	return nil
}

func (s *SQLStorage) Delete(id int) error {
	query := `
		DELETE FROM subscriptions
		WHERE id = $1;
	`

	result, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrSubscriptionNotFound
	}

	return nil
}

func (s *SQLStorage) GetTotalSum() (int, error) {
	query := `
		SELECT COALESCE(SUM(price), 0) as total_price
		FROM subscriptions;
	`

	var totalSum int

	result := s.db.QueryRow(query)

	if err := result.Scan(&totalSum); err != nil {
		return 0, err
	}

	return totalSum, nil
}
