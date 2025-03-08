package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bulatkarmak/grpc-todo/internal/domain"
	"github.com/bulatkarmak/grpc-todo/internal/service"
)

type toDoRepository struct {
	db *sql.DB
}

func NewToDoRepository(db *sql.DB) service.ToDoRepository {
	return &toDoRepository{db: db}
}

func (r *toDoRepository) CreateTask(ctx context.Context, params *domain.CreateTaskParams) (*domain.Task, error) {
	task := &domain.Task{}

	err := r.db.QueryRowContext(ctx,
		`INSERT INTO tasks(title, description) 
		VALUES($1, $2) 
		RETURNING id, title, description, is_completed, created_at, updated_at`,
		params.Title, params.Description).
		Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("не получилось создать task: %w", err)
	}

	return task, nil
}
