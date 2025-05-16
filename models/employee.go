package models

import (
	"time"
)

// Employee represents the employee entity
type Employee struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Position  string    `json:"position" binding:"required"`
	Salary    float64   `json:"salary" binding:"required"`
	JoinDate  time.Time `json:"join_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
