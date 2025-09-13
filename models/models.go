package models

import (
	"github.com/google/uuid"
)

// Subscription описывает подписку
//
// swagger:model Subscription
type Subscription struct {
	// ID подписки
	// example: 123
	ID int `json:"id,omitempty"`

	// Идентификатор пользователя, владеющего подпиской
	// example: 550e8400-e29b-41d4-a716-446655440000
	UserID uuid.UUID `json:"user_id"`

	// Название сервиса подписки
	// example: Netflix
	ServiceName string `json:"service_name"`

	// Цена подписки в целых единицах
	// example: 1499
	Price int `json:"price"`

	// Дата начала подписки в формате ISO8601
	// example: 2023-07-18
	StartDate string `json:"start_date"`
}
