package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	TaskID   uint      `gorm:"primary key;autoIncrement" json:"id"`
	TaskName *string   `gorm:"unique" json:"name"`
	Note     *string   `json:"note"`
	Deadline time.Time `json:"date"`
}

func MigrateTask(db *gorm.DB) error {
	err := db.AutoMigrate(&Task{})
	return err
}
