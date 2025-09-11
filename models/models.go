package models

const (
	StatusCreated   = "created"
	StatusRunning   = "running"
	StatusCompleted = "completed"
	StatusUpdated   = "updated"
	ResultCompleted = "subscription completed successfully"
)

type Subscription struct {
	UserID      string `json:"user_id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	StartDate   string `json:"start_date"`
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
