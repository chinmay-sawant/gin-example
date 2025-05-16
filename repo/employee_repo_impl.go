package repo

import (
	"errors"

	"github.com/chinmay-sawant/gin-example/db"
	"github.com/chinmay-sawant/gin-example/models"
	"gorm.io/gorm"
)

type employeeRepositoryImpl struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepositoryImpl{}
}

func (r *employeeRepositoryImpl) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	result := db.DB.Find(&employees)
	return employees, result.Error
}

func (r *employeeRepositoryImpl) FindByID(id uint) (models.Employee, error) {
	var employee models.Employee
	result := db.DB.First(&employee, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return employee, errors.New("employee not found")
		}
		return employee, result.Error
	}
	return employee, nil
}

func (r *employeeRepositoryImpl) Create(employee models.Employee) (models.Employee, error) {
	result := db.DB.Create(&employee)
	return employee, result.Error
}

func (r *employeeRepositoryImpl) Update(id uint, employee models.Employee) (models.Employee, error) {
	var existingEmployee models.Employee
	result := db.DB.First(&existingEmployee, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return existingEmployee, errors.New("employee not found")
		}
		return existingEmployee, result.Error
	}

	existingEmployee.Name = employee.Name
	existingEmployee.Email = employee.Email
	existingEmployee.Position = employee.Position
	existingEmployee.Salary = employee.Salary
	existingEmployee.JoinDate = employee.JoinDate

	db.DB.Save(&existingEmployee)
	return existingEmployee, nil
}

func (r *employeeRepositoryImpl) Delete(id uint) error {
	var employee models.Employee
	result := db.DB.First(&employee, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("employee not found")
		}
		return result.Error
	}
	db.DB.Delete(&employee)
	return nil
}
