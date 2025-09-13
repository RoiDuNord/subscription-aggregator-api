package models

import (
	"github.com/google/uuid"
)

// Subscription описывает подписку
// swagger:model Subscription
type Subscription struct {
	ID          int       `json:"id,omitempty"`
	UserID      uuid.UUID `json:"user_id"`
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	StartDate   string    `json:"start_date"`
}
