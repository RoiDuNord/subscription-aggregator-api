package models

import (
	"time"
)

const (
	StatusCreated   = "created"
	StatusRunning   = "running"
	StatusCompleted = "completed"
	StatusUpdated   = "updated"
	ResultCompleted = "subscription completed successfully"
)

type Subscription struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"user_id"`
	ServiceName string    `json:"service_name"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// type Subscription struct {
// 	ID           string    `json:"id"`
// 	Name         string    `json:"name"`
// 	Status       string    `json:"status"`
// 	CreationTime time.Time `json:"creation_time"`
// 	StartTime    time.Time `json:"start_time,omitzero"`
// 	FinishTime   time.Time `json:"finish_time,omitzero"`
// 	Duration     float32   `json:"duration_sec,omitempty"`
// 	Result       string    `json:"result,omitempty"`

// 	CancelFunc context.CancelFunc `json:"-"`
// }
