package service

import (
	"context"
	"errors"
	"github.com/bulatkarmak/grpc-todo/internal/domain"
)

var (
	EmptyTitleErr = errors.New("title не может быть пустым")
	EmptyDescErr  = errors.New("description не может быть пустым")
	LessOneIDErr  = errors.New("id должен быть равен 1 или больше")
)

type ToDoRepository interface {
	CreateTask(ctx context.Context, params *domain.CreateTaskParams) (*domain.Task, error)
	GetTask(ctx context.Context, taskID int64) (*domain.Task, error)
	ListTasks(ctx context.Context) ([]domain.Task, error)
	UpdateTask(ctx context.Context, params *domain.UpdateTaskParams) (*domain.Task, error)
	DeleteTask(ctx context.Context, taskID int64) error
}

type ToDoService interface {
	CreateTask(ctx context.Context, params *domain.CreateTaskParams) (*domain.Task, error)
	GetTask(ctx context.Context, taskID int64) (*domain.Task, error)
	ListTasks(ctx context.Context) ([]domain.Task, error)
	UpdateTask(ctx context.Context, params *domain.UpdateTaskParams) (*domain.Task, error)
	DeleteTask(ctx context.Context, taskID int64) error
}

type toDoService struct {
	repo ToDoRepository
}

func NewToDoService(repo ToDoRepository) ToDoService {
	return &toDoService{
		repo: repo,
	}
}

func (s *toDoService) CreateTask(ctx context.Context, params *domain.CreateTaskParams) (*domain.Task, error) {
	if params.Title == "" {
		return nil, EmptyTitleErr
	}

	if params.Description == "" {
		return nil, EmptyDescErr
	}

	task, err := s.repo.CreateTask(ctx, params)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *toDoService) GetTask(ctx context.Context, taskID int64) (*domain.Task, error) {
	if taskID < 1 {
		return nil, LessOneIDErr
	}

	task, err := s.repo.GetTask(ctx, taskID)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *toDoService) ListTasks(ctx context.Context) ([]domain.Task, error) {
	tasks, err := s.repo.ListTasks(ctx)

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *toDoService) UpdateTask(ctx context.Context, params *domain.UpdateTaskParams) (*domain.Task, error) {
	if params.Title == nil && params.Description == nil && params.IsCompleted == nil {
		return s.GetTask(ctx, params.ID)
	}

	if params.Title != nil && *params.Title == "" {
		return nil, EmptyTitleErr
	}

	if params.Description != nil && *params.Description == "" {
		return nil, EmptyDescErr
	}

	task, err := s.repo.UpdateTask(ctx, params)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *toDoService) DeleteTask(ctx context.Context, taskID int64) error {
	if taskID < 1 {
		return LessOneIDErr
	}

	err := s.repo.DeleteTask(ctx, taskID)

	if err != nil {
		return err
	}

	return nil
}
