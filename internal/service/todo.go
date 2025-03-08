package service

import (
	"context"
	"github.com/bulatkarmak/grpc-todo/internal/domain"
)

type ToDoRepository interface {
	CreateTask(ctx context.Context, params *domain.CreateTaskParams) (*domain.Task, error)
	GetTask(ctx context.Context, taskID int64) (*domain.Task, error)
}

type ToDoService interface {
	CreateTask(ctx context.Context, params *domain.CreateTaskParams) (*domain.Task, error)
	GetTask(ctx context.Context, taskID int64) (*domain.Task, error)
	ListTasks(ctx context.Context) ([]domain.Task, error)
	UpdateTask(ctx context.Context, task *domain.Task) (*domain.Task, error)
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
	task, err := s.repo.CreateTask(ctx, params)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *toDoService) GetTask(ctx context.Context, taskID int64) (*domain.Task, error) {
	task, err := s.repo.GetTask(ctx, taskID)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *toDoService) ListTasks(ctx context.Context) ([]domain.Task, error) {
	return nil, nil
}

func (s *toDoService) UpdateTask(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	return &domain.Task{}, nil
}

func (s *toDoService) DeleteTask(ctx context.Context, taskID int64) error {
	return nil
}
