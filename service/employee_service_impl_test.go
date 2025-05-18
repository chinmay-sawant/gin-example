package service

import (
	"errors"
	"testing"

	"github.com/chinmay-sawant/gin-example/models"
	"github.com/chinmay-sawant/gin-example/repo/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type EmployeeServiceTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mocks.MockEmployeeRepository
	svc  EmployeeService
}

func (suite *EmployeeServiceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockEmployeeRepository(suite.ctrl)
	suite.svc = NewEmployeeService(suite.repo)
}

func (suite *EmployeeServiceTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestEmployeeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(EmployeeServiceTestSuite))
}

func (suite *EmployeeServiceTestSuite) TestGetAllEmployees() {
	employees := []models.Employee{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Position: "Dev", Salary: 50000},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Position: "QA", Salary: 40000},
	}
	suite.repo.EXPECT().FindAll().Return(employees, nil)

	result, err := suite.svc.GetAllEmployees()
	suite.NoError(err)
	suite.Equal(employees, result)
}

func (suite *EmployeeServiceTestSuite) TestGetEmployeeByID() {
	employee := models.Employee{ID: 1, Name: "Alice", Email: "alice@example.com", Position: "Dev", Salary: 50000}
	suite.repo.EXPECT().FindByID(uint(1)).Return(employee, nil)

	result, err := suite.svc.GetEmployeeByID(1)
	suite.NoError(err)
	suite.Equal(employee, result)

	suite.repo.EXPECT().FindByID(uint(2)).Return(models.Employee{}, errors.New("not found"))
	_, err = suite.svc.GetEmployeeByID(2)
	suite.Error(err)
}

func (suite *EmployeeServiceTestSuite) TestCreateEmployee() {
	employee := models.Employee{Name: "John", Email: "john@example.com", Position: "Dev", Salary: 60000}
	created := employee
	created.ID = 1
	suite.repo.EXPECT().Create(employee).Return(created, nil)

	result, err := suite.svc.CreateEmployee(employee)
	suite.NoError(err)
	suite.Equal(created, result)
}

func (suite *EmployeeServiceTestSuite) TestUpdateEmployee() {
	updated := models.Employee{ID: 1, Name: "Updated", Email: "updated@example.com", Position: "Lead", Salary: 80000}
	suite.repo.EXPECT().Update(uint(1), updated).Return(updated, nil)

	result, err := suite.svc.UpdateEmployee(1, updated)
	suite.NoError(err)
	suite.Equal(updated, result)
}

func (suite *EmployeeServiceTestSuite) TestDeleteEmployee() {
	suite.repo.EXPECT().Delete(uint(1)).Return(nil)

	err := suite.svc.DeleteEmployee(1)
	suite.NoError(err)

	suite.repo.EXPECT().Delete(uint(2)).Return(errors.New("not found"))
	err = suite.svc.DeleteEmployee(2)
	suite.Error(err)
}
