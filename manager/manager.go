package manager

import (
	"errors"
	"subscription-aggregator-api/models"
)

type SubscriptionStorage interface {
	Create(subscription models.Subscription) error
	GetByID(id string) (models.Subscription, error)
	GetList() ([]models.Subscription, error)
	Update(id string, updated models.Subscription) error
	Delete(id string) error
}

type SubscriptionManager struct {
	storage SubscriptionStorage
}

func New(storage SubscriptionStorage) *SubscriptionManager {
	return &SubscriptionManager{storage: storage}
}

func (sm *SubscriptionManager) CreateSubscription(subscription models.Subscription) error {
	return sm.storage.Create(subscription)
}

func (sm *SubscriptionManager) GetSubscriptionList() ([]models.Subscription, error) {
	return sm.storage.GetList()
}

func (sm *SubscriptionManager) GetSubscription(id string) (models.Subscription, error) {
	return sm.storage.GetByID(id)
}

func (sm *SubscriptionManager) UpdateSubscription(id string, updated models.Subscription) error {
	if id != updated.ID {
		return errors.New("mismatched subscription ID")
	}
	return sm.storage.Update(id, updated)
}

func (sm *SubscriptionManager) DeleteSubscription(id string) error {
	return sm.storage.Delete(id)
}
