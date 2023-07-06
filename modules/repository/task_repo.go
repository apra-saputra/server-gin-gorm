package repository

import (
	"gorm.io/gorm"
)

type TaskRepository struct {
	Db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return (&TaskRepository{
		Db: db,
	})
}
