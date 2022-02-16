package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/s3rzh/horse-exchange/entity"
)

type Task interface {
	GetTasks() ([]entity.Task, error)
}

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (b *TaskPostgres) GetTasks() ([]entity.Task, error) {
	var tasks []entity.Task

	query := fmt.Sprintf("SELECT t.* FROM %s t", tasksTable)
	err := b.db.Select(&tasks, query)

	return tasks, err
}
