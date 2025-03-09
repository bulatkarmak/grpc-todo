package domain

import (
	"time"
)

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTaskParams struct {
	Title       string
	Description string
}

type UpdateTaskParams struct {
	ID          int64
	Title       *string
	Description *string
	IsCompleted *bool
}
