package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	IsComplete  bool `gorm:"default:false"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	Date        time.Time
}
