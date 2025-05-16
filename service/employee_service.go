package service

import (
	"github.com/chinmay-sawant/gin-example/models"
)

// EmployeeService defines the interface for employee operations
type EmployeeService interface {
	GetAllEmployees() ([]models.Employee, error)
	GetEmployeeByID(id uint) (models.Employee, error)
	CreateEmployee(employee models.Employee) (models.Employee, error)
	UpdateEmployee(id uint, employee models.Employee) (models.Employee, error)
	DeleteEmployee(id uint) error
}
