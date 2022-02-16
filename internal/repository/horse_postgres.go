package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/s3rzh/horse-exchange/entity"
	"github.com/spf13/viper"
)

type Horse interface {
	GetAllHorses() ([]entity.Horse, error)
	BuyHorseByID(horseID int) (int64, error)
	GetHorsesByBreedID(breedID int) ([]entity.Horse, error)
	GetHorsesByTaskID(taskID int) ([]entity.Horse, error)
	RentHorseByIDWithPeriod(rent entity.Rent) (int, error)
	PayHorseByRentID(rentID int) error
}

type HorsePostgres struct {
	db *sqlx.DB
}

func NewHorsePostgres(db *sqlx.DB) *HorsePostgres {
	return &HorsePostgres{db: db}
}

func (hp *HorsePostgres) GetAllHorses() ([]entity.Horse, error) {
	var horses []entity.Horse

	query := fmt.Sprintf("SELECT h.id, h.name, h.gender, h.age, h.weight, b.name AS breed, h.birthplace, h.habit, h.rentalprice, h.purchaseprice FROM %s h LEFT JOIN %s b ON h.breed_id = b.id WHERE h.sold = FALSE", horsesTable, breedsTable)
	err := hp.db.Select(&horses, query)

	return horses, err
}

func (hp *HorsePostgres) BuyHorseByID(horseID int) (int64, error) {
	query := fmt.Sprintf("UPDATE %s h SET sold = TRUE FROM %s r WHERE h.id = $1 AND h.sold = FALSE AND NOT EXISTS (SELECT FROM rentals AS r WHERE r.horse_id = h.id AND r.paid = true AND r.till > NOW())", horsesTable, rentalsTable)
	result, err := hp.db.Exec(query, horseID)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (hp *HorsePostgres) GetHorsesByBreedID(breedID int) ([]entity.Horse, error) {
	var horses []entity.Horse

	query := fmt.Sprintf("SELECT h.id, h.name, h.gender, h.age, h.weight, b.name AS breed, h.birthplace, h.habit, h.rentalprice, h.purchaseprice FROM %s h LEFT JOIN %s b ON h.breed_id = b.id WHERE h.breed_id = $1 AND h.sold = FALSE", horsesTable, breedsTable)
	err := hp.db.Select(&horses, query, breedID)

	return horses, err
}

func (hp *HorsePostgres) GetHorsesByTaskID(taskID int) ([]entity.Horse, error) {
	var horses []entity.Horse

	query := fmt.Sprintf("SELECT h.id, h.name, h.gender, h.age, h.weight, b.name AS breed, h.birthplace, h.habit, h.rentalprice, h.purchaseprice FROM %s ht LEFT JOIN %s h ON h.id = ht.horse_id LEFT JOIN %s b ON h.breed_id = b.id WHERE ht.task_id = $1 AND h.sold = FALSE", horsesTasksTable, horsesTable, breedsTable)
	err := hp.db.Select(&horses, query, taskID)

	return horses, err
}

func (hp *HorsePostgres) RentHorseByIDWithPeriod(rent entity.Rent) (int, error) {
	var countRents int

	checkRentQuery := fmt.Sprintf("SELECT  COUNT(*) FROM %s r WHERE r.horse_id = $1 AND r.paid = TRUE AND ((r.since <= $2 AND $2 <= r.till) OR (r.since <= $3 AND $3 <= r.till))", rentalsTable)
	_ = hp.db.QueryRow(checkRentQuery, rent.HorseID, rent.Since, rent.Till).Scan(&countRents)

	if countRents > 0 {
		return 0, nil
	}

	var id int
	createRentQuery := fmt.Sprintf("INSERT INTO %s (horse_id, since, till) VALUES ($1, $2, $3) RETURNING id", rentalsTable)

	row := hp.db.QueryRow(createRentQuery, rent.HorseID, rent.Since, rent.Till)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (hp *HorsePostgres) PayHorseByRentID(rentID int) error {
	var rent entity.Rent

	checkRentQuery := fmt.Sprintf("SELECT r.horse_id, r.since, r.till FROM %s r WHERE r.id = $1 AND r.paid = FALSE", rentalsTable)

	err := hp.db.Get(&rent, checkRentQuery, rentID)

	if err != nil {
		return errors.New(viper.GetString("messages.errors.rent_not_reserved_or_paid"))
	}

	var countRents int
	checkAnotherRentQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s r WHERE r.horse_id = $1 AND r.paid = TRUE AND r.id <> $4 AND ((r.since <= $2 AND $2 <= r.till) OR (r.since <= $3 AND $3 <= r.till))", rentalsTable)
	_ = hp.db.QueryRow(checkAnotherRentQuery, rent.HorseID, rent.Since, rent.Till, rentID).Scan(&countRents)

	if countRents > 0 {
		return errors.New(viper.GetString("messages.errors.busy_period"))
	}

	query := fmt.Sprintf("UPDATE %s r SET paid = TRUE WHERE r.id = $1", rentalsTable)
	_, err = hp.db.Exec(query, rentID)

	if err != nil {
		return errors.New(viper.GetString("messages.errors.retry_later"))
	}

	return nil
}
