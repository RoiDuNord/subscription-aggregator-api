package manager

import (
	"errors"
	"subscription-aggregator-api/models"
	"time"
)

type SubscriptionStorage interface {
	Create(subscription models.Subscription) error
	GetByID(id string) (models.Subscription, error)
	GetList() ([]models.Subscription, error)
	Update(id string, updated models.Subscription) error
	Delete(id string) error
	GetTotalSum() (totalSum int, err error)
}

type Manager struct {
	storage SubscriptionStorage
}

func New(storage SubscriptionStorage) *Manager {
	return &Manager{storage: storage}
}

func (m *Manager) CreateSubscription(subscription models.Subscription) error {
	if err := validateSubscription(subscription); err != nil {
		return err
	}
	return m.storage.Create(subscription)
}

func validateSubscription(subscription models.Subscription) error {
	if subscription.UserID == "" {
		return errors.New("UserID cannot be empty")
	}
	if subscription.ServiceName == "" {
		return errors.New("ServiceName cannot be empty")
	}
	if err := validateDate(subscription.StartDate); err != nil {
		return err
	}
	if subscription.Price <= 0 {
		return errors.New("Price must be greater than 0")
	}
	return nil
}

func validateDate(date string) error {
	if _, err := time.Parse("2006-01", date); err != nil {
		return errors.New("invalid start_date format: must be 'YYYY-MM'")
	}
	return nil
}

func (m *Manager) GetSubscription(id string) (models.Subscription, error) {
	return m.storage.GetByID(id)
}

func (m *Manager) GetAllSubscriptions() ([]models.Subscription, error) {
	return m.storage.GetList()
}

func (m *Manager) Get(id string) (models.Subscription, error) {
	if id == "" {
		return models.Subscription{}, errors.New("invalid ID: cannot be empty")
	}
	return m.storage.GetByID(id)
}

func (m *Manager) UpdateSubscription(id string, updated models.Subscription) error {
	if id == "" {
		return errors.New("invalid ID: cannot be empty")
	}
	return m.storage.Update(id, updated)
}

func (m *Manager) DeleteSubscription(id string) error {
	if id == "" {
		return errors.New("invalid ID: cannot be empty")
	}
	return m.storage.Delete(id)
}

func (m *Manager) GetAllSubscriptionsSum() (int, error) {
	return m.storage.GetTotalSum()
}
