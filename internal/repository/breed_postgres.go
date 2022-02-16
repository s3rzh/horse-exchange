package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/s3rzh/horse-exchange/entity"
)

type Breed interface {
	GetBreeds() ([]entity.Breed, error)
}

type BreedPostgres struct {
	db *sqlx.DB
}

func NewBreedPostgres(db *sqlx.DB) *BreedPostgres {
	return &BreedPostgres{db: db}
}

func (b *BreedPostgres) GetBreeds() ([]entity.Breed, error) {
	var breeds []entity.Breed

	query := fmt.Sprintf("SELECT b.* FROM %s b", breedsTable)
	err := b.db.Select(&breeds, query)

	return breeds, err
}
