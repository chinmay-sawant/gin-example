package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chinmay-sawant/gin-example/models"
	"github.com/chinmay-sawant/gin-example/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type EmployeeControllerTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	svc  *mocks.MockEmployeeService
	r    *gin.Engine
}

func (suite *EmployeeControllerTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.svc = mocks.NewMockEmployeeService(suite.ctrl)
	gin.SetMode(gin.TestMode)
	suite.r = gin.Default()
	controller := &employeeControllerImpl{employeeService: suite.svc}
	v1 := suite.r.Group("/api/v1")
	controller.RegisterRoutes(v1)
}

func (suite *EmployeeControllerTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestEmployeeControllerTestSuite(t *testing.T) {
	suite.Run(t, new(EmployeeControllerTestSuite))
}

func (suite *EmployeeControllerTestSuite) TestGetEmployeesHandler() {
	employees := []models.Employee{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Position: "Dev", Salary: 50000},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Position: "QA", Salary: 40000},
	}
	suite.svc.EXPECT().GetAllEmployees().Return(employees, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/employees/", nil)
	suite.r.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Alice")
	suite.Contains(w.Body.String(), "Bob")
}

func (suite *EmployeeControllerTestSuite) TestGetEmployeeHandler() {
	employee := models.Employee{ID: 1, Name: "Alice", Email: "alice@example.com", Position: "Dev", Salary: 50000}
	suite.svc.EXPECT().GetEmployeeByID(uint(1)).Return(employee, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/employees/1", nil)
	suite.r.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Alice")

	suite.svc.EXPECT().GetEmployeeByID(uint(2)).Return(models.Employee{}, errors.New("not found"))
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/employees/2", nil)
	suite.r.ServeHTTP(w, req)
	suite.Equal(http.StatusNotFound, w.Code)
}

func (suite *EmployeeControllerTestSuite) TestCreateEmployeeHandler() {
	input := `{"name":"John","email":"john@example.com","position":"Dev","salary":60000}`
	created := models.Employee{ID: 1, Name: "John", Email: "john@example.com", Position: "Dev", Salary: 60000}
	suite.svc.EXPECT().CreateEmployee(gomock.Any()).Return(created, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/employees/", strings.NewReader(input))
	req.Header.Set("Content-Type", "application/json")
	suite.r.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)
	suite.Contains(w.Body.String(), "John")
}

func (suite *EmployeeControllerTestSuite) TestUpdateEmployeeHandler() {
	input := `{"name":"Updated","email":"updated@example.com","position":"Lead","salary":80000}`
	updated := models.Employee{ID: 1, Name: "Updated", Email: "updated@example.com", Position: "Lead", Salary: 80000}
	suite.svc.EXPECT().UpdateEmployee(uint(1), gomock.Any()).Return(updated, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/employees/1", strings.NewReader(input))
	req.Header.Set("Content-Type", "application/json")
	suite.r.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Updated")
}

func (suite *EmployeeControllerTestSuite) TestDeleteEmployeeHandler() {
	suite.svc.EXPECT().DeleteEmployee(uint(1)).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/employees/1", nil)
	suite.r.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Employee deleted successfully")
}
