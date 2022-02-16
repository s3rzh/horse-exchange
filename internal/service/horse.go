package service

import (
	"github.com/s3rzh/horse-exchange/entity"
	"github.com/s3rzh/horse-exchange/internal/repository"
)

type Horse interface {
	GetAllHorses() ([]entity.Horse, error)
	BuyHorseByID(horseID int) (int64, error)
	GetHorsesByBreedID(breedID int) ([]entity.Horse, error)
	GetHorsesByTaskID(taskID int) ([]entity.Horse, error)
	RentHorseByIDWithPeriod(rent entity.Rent) (int, error)
	PayHorseByRentID(rentID int) error
}

type HorseService struct {
	repos repository.Horse
}

func NewHorseService(r repository.Horse) *HorseService {
	return &HorseService{repos: r}
}

func (h *HorseService) GetAllHorses() ([]entity.Horse, error) {
	return h.repos.GetAllHorses()
}

func (h *HorseService) BuyHorseByID(horseID int) (int64, error) {
	return h.repos.BuyHorseByID(horseID)
}

func (h *HorseService) GetHorsesByBreedID(breedID int) ([]entity.Horse, error) {
	return h.repos.GetHorsesByBreedID(breedID)
}
func (h *HorseService) GetHorsesByTaskID(taskID int) ([]entity.Horse, error) {
	return h.repos.GetHorsesByTaskID(taskID)
}

func (h *HorseService) RentHorseByIDWithPeriod(rent entity.Rent) (int, error) {
	return h.repos.RentHorseByIDWithPeriod(rent)
}

func (h *HorseService) PayHorseByRentID(rentID int) error {
	return h.repos.PayHorseByRentID(rentID)
}
