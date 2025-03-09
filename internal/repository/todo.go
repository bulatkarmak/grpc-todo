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

func (r *toDoRepository) ListTasks(ctx context.Context) ([]domain.Task, error) {
	var tasks []domain.Task

	rows, err := r.db.QueryContext(ctx, `SELECT * FROM tasks`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("в таблице нет задач")
		}
		return nil, fmt.Errorf("не получилось получить строки из таблицы: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task domain.Task
		if err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, fmt.Errorf("не получилось отсканировать строку: %w", err)
		}
		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("не получилось итерироваться по строкам: %w", err)
	}

	return tasks, nil
}

func (r *toDoRepository) UpdateTask(ctx context.Context, params *domain.UpdateTaskParams) (*domain.Task, error) {
	query := "UPDATE tasks SET"
	args := []interface{}{}
	i := 1

	if params.Title != nil {
		query += fmt.Sprintf(" title=$%d,", i)
		args = append(args, *params.Title)
		i++
	}

	if params.Description != nil {
		query += fmt.Sprintf(" description=$%d,", i)
		args = append(args, *params.Description)
		i++
	}

	if params.IsCompleted != nil {
		query += fmt.Sprintf(" is_completed=$%d,", i)
		args = append(args, *params.IsCompleted)
		i++
	}

	query += fmt.Sprintf(" updated_at = NOW() WHERE id=$%d RETURNING id, title, description, is_completed, created_at, updated_at", i)
	args = append(args, params.ID)

	row := r.db.QueryRowContext(ctx, query, args...)

	task := &domain.Task{}

	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		fmt.Printf("%v\n%v", query, args)
		return nil, fmt.Errorf("ошибка при обновлении task: %w", err)
	}

	return task, nil
}
