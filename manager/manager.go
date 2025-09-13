package manager

import (
	"errors"
	"strconv"
	"subscription-aggregator-api/models"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserIDEmpty         = errors.New("user ID cannot be empty")
	ErrServiceNameEmpty    = errors.New("service name cannot be empty")
	ErrInvalidStartDate    = errors.New("start date must be in format YYYY-MM-DD")
	ErrPriceMustBePositive = errors.New("price must be greater than 0")
	ErrIDEmpty             = errors.New("ID must be greater than 0")
)

type BadRequestError struct {
	msg string
}

func (e *BadRequestError) Error() string {
	return e.msg
}

type SubscriptionStorage interface {
	Create(subscription models.Subscription) error
	GetByID(id int) (models.Subscription, error)
	GetList() ([]models.Subscription, error)
	Update(id int, updated models.Subscription) error
	Delete(id int) error
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
		return &BadRequestError{msg: err.Error()}
	}
	return m.storage.Create(subscription)
}

func (m *Manager) GetSubscription(id string) (models.Subscription, error) {
	parsedID, err := validateID(id)
	if err != nil {
		return models.Subscription{}, err
	}
	return m.storage.GetByID(parsedID)
}

func (m *Manager) GetAllSubscriptions() ([]models.Subscription, error) {
	return m.storage.GetList()
}

func (m *Manager) UpdateSubscription(id string, updatedSubscription models.Subscription) error {
	parsedID, err := validateID(id)
	if err != nil {
		return err
	}
	if err := validateSubscription(updatedSubscription); err != nil {
		return err
	}
	return m.storage.Update(parsedID, updatedSubscription)
}

func (m *Manager) DeleteSubscription(id string) error {
	parsedID, err := validateID(id)
	if err != nil {
		return err
	}
	return m.storage.Delete(parsedID)
}

func (m *Manager) GetAllSubscriptionsSum() (int, error) {
	return m.storage.GetTotalSum()
}

func validateSubscription(subscription models.Subscription) error {
	if subscription.UserID == uuid.Nil {
		return ErrUserIDEmpty
	}
	if subscription.ServiceName == "" {
		return ErrServiceNameEmpty
	}
	if err := validateDate(subscription.StartDate); err != nil {
		return err
	}
	if subscription.Price <= 0 {
		return ErrPriceMustBePositive
	}
	return nil
}

func validateDate(date string) error {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ErrInvalidStartDate
	}
	return nil
}

func validateID(id string) (int, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return 0, &BadRequestError{msg: "invalid subscription ID"}
	}
	if parsedID <= 0 {
		return 0, &BadRequestError{msg: ErrIDEmpty.Error()}
	}
	return parsedID, nil
}
