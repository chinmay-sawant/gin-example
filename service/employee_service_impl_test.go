package service

import (
	"errors"
	"testing"

	"github.com/chinmay-sawant/gin-example/models"
	"github.com/chinmay-sawant/gin-example/repo/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllEmployees(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockEmployeeRepository(ctrl)
	employees := []models.Employee{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Position: "Dev", Salary: 50000},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Position: "QA", Salary: 40000},
	}
	repoMock.EXPECT().FindAll().Return(employees, nil)

	svc := NewEmployeeService(repoMock)
	result, err := svc.GetAllEmployees()
	assert.NoError(t, err)
	assert.Equal(t, employees, result)
}

func TestGetEmployeeByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockEmployeeRepository(ctrl)
	employee := models.Employee{ID: 1, Name: "Alice", Email: "alice@example.com", Position: "Dev", Salary: 50000}
	repoMock.EXPECT().FindByID(uint(1)).Return(employee, nil)

	svc := NewEmployeeService(repoMock)
	result, err := svc.GetEmployeeByID(1)
	assert.NoError(t, err)
	assert.Equal(t, employee, result)

	repoMock.EXPECT().FindByID(uint(2)).Return(models.Employee{}, errors.New("not found"))
	_, err = svc.GetEmployeeByID(2)
	assert.Error(t, err)
}

func TestCreateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockEmployeeRepository(ctrl)
	employee := models.Employee{Name: "John", Email: "john@example.com", Position: "Dev", Salary: 60000}
	created := employee
	created.ID = 1
	repoMock.EXPECT().Create(employee).Return(created, nil)

	svc := NewEmployeeService(repoMock)
	result, err := svc.CreateEmployee(employee)
	assert.NoError(t, err)
	assert.Equal(t, created, result)
}

func TestUpdateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockEmployeeRepository(ctrl)
	updated := models.Employee{ID: 1, Name: "Updated", Email: "updated@example.com", Position: "Lead", Salary: 80000}
	repoMock.EXPECT().Update(uint(1), updated).Return(updated, nil)

	svc := NewEmployeeService(repoMock)
	result, err := svc.UpdateEmployee(1, updated)
	assert.NoError(t, err)
	assert.Equal(t, updated, result)
}

func TestDeleteEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockEmployeeRepository(ctrl)
	repoMock.EXPECT().Delete(uint(1)).Return(nil)

	svc := NewEmployeeService(repoMock)
	err := svc.DeleteEmployee(1)
	assert.NoError(t, err)

	repoMock.EXPECT().Delete(uint(2)).Return(errors.New("not found"))
	err = svc.DeleteEmployee(2)
	assert.Error(t, err)
}
