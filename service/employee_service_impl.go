package service

import (
	"github.com/chinmay-sawant/gin-example/models"
	"github.com/chinmay-sawant/gin-example/repo"
)

// EmployeeServiceImpl implements the EmployeeService interface
type EmployeeServiceImpl struct {
	employeeRepo repo.EmployeeRepository
}

// NewEmployeeService creates a new instance of EmployeeService
func NewEmployeeService(employeeRepo repo.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{employeeRepo: employeeRepo}
}

// GetAllEmployees returns all employees
func (s *EmployeeServiceImpl) GetAllEmployees() ([]models.Employee, error) {
	return s.employeeRepo.FindAll()
}

// GetEmployeeByID returns an employee by ID
func (s *EmployeeServiceImpl) GetEmployeeByID(id uint) (models.Employee, error) {
	return s.employeeRepo.FindByID(id)
}

// CreateEmployee creates a new employee
func (s *EmployeeServiceImpl) CreateEmployee(employee models.Employee) (models.Employee, error) {
	return s.employeeRepo.Create(employee)
}

// UpdateEmployee updates an existing employee
func (s *EmployeeServiceImpl) UpdateEmployee(id uint, employee models.Employee) (models.Employee, error) {
	return s.employeeRepo.Update(id, employee)
}

// DeleteEmployee deletes an employee by ID
func (s *EmployeeServiceImpl) DeleteEmployee(id uint) error {
	return s.employeeRepo.Delete(id)
}
