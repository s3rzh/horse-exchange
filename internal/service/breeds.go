package service

import (
	"github.com/s3rzh/horse-exchange/entity"
	"github.com/s3rzh/horse-exchange/internal/repository"
)

type Breed interface {
	GetBreeds() ([]entity.Breed, error)
}

type BreedService struct {
	repos repository.Breed
}

func NewBreedService(r repository.Breed) *BreedService {
	return &BreedService{repos: r}
}

func (b *BreedService) GetBreeds() ([]entity.Breed, error) {
	return b.repos.GetBreeds()
}
