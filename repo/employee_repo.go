package repo

import (
	"github.com/chinmay-sawant/gin-example/models"
)

type EmployeeRepository interface {
	FindAll() ([]models.Employee, error)
	FindByID(id uint) (models.Employee, error)
	Create(employee models.Employee) (models.Employee, error)
	Update(id uint, employee models.Employee) (models.Employee, error)
	Delete(id uint) error
}
