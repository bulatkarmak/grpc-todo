package repository

import (
	"context"
	"database/sql"
	"errors"
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

func (r *toDoRepository) GetTask(ctx context.Context, taskID int64) (*domain.Task, error) {
	task := &domain.Task{}

	row := r.db.QueryRowContext(ctx,
		`SELECT id, title, description, is_completed, created_at, updated_at FROM tasks
		WHERE id = $1`, taskID)

	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("task не найдена")
		}
		return nil, fmt.Errorf("не получилось отсканировать row: %w", err)
	}

	return task, nil
}
