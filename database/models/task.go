package models

type Task struct {
	TaskID   uint   `gorm:"primaryKey" json:"task_id"`
	TaskName string `gorm:"unique" json:"name" validate:"required,min=3,max=40"`
	Note     string `json:"note" validate:"required,min=3,max=100"`
	Deadline string `json:"date" validate:"required"`
}
