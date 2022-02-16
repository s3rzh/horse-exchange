package service

import (
	"github.com/s3rzh/horse-exchange/internal/repository"
)

type Service struct {
	Horse
	Breed
	Task
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Horse: NewHorseService(r.Horse),
		Breed: NewBreedService(r.Breed),
		Task:  NewTaskService(r.Task),
	}
}
