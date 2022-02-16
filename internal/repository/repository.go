package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Horse
	Breed
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Horse: NewHorsePostgres(db),
		Breed: NewBreedPostgres(db),
		Task:  NewTaskPostgres(db),
	}
}
