package service

import (
	"github.com/s3rzh/horse-exchange/entity"
	"github.com/s3rzh/horse-exchange/internal/repository"
)

type Task interface {
	GetTasks() ([]entity.Task, error)
}

type TaskService struct {
	repos repository.Task
}

func NewTaskService(r repository.Task) *TaskService {
	return &TaskService{repos: r}
}

func (b *TaskService) GetTasks() ([]entity.Task, error) {
	return b.repos.GetTasks()
}
