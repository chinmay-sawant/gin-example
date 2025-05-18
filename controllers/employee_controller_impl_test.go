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
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setupRouterWithController(svc *mocks.MockEmployeeService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	controller := &employeeControllerImpl{employeeService: svc}
	v1 := r.Group("/api/v1")
	controller.RegisterRoutes(v1)
	return r
}

func TestGetEmployeesHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := mocks.NewMockEmployeeService(ctrl)
	employees := []models.Employee{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Position: "Dev", Salary: 50000},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Position: "QA", Salary: 40000},
	}
	svc.EXPECT().GetAllEmployees().Return(employees, nil)

	r := setupRouterWithController(svc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/employees/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Alice")
	assert.Contains(t, w.Body.String(), "Bob")
}

func TestGetEmployeeHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := mocks.NewMockEmployeeService(ctrl)
	employee := models.Employee{ID: 1, Name: "Alice", Email: "alice@example.com", Position: "Dev", Salary: 50000}
	svc.EXPECT().GetEmployeeByID(uint(1)).Return(employee, nil)

	r := setupRouterWithController(svc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/employees/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Alice")

	svc.EXPECT().GetEmployeeByID(uint(2)).Return(models.Employee{}, errors.New("not found"))
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/employees/2", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateEmployeeHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := mocks.NewMockEmployeeService(ctrl)
	input := `{"name":"John","email":"john@example.com","position":"Dev","salary":60000}`
	created := models.Employee{ID: 1, Name: "John", Email: "john@example.com", Position: "Dev", Salary: 60000}
	svc.EXPECT().CreateEmployee(gomock.Any()).Return(created, nil)

	r := setupRouterWithController(svc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/employees/", strings.NewReader(input))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John")
}

func TestUpdateEmployeeHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := mocks.NewMockEmployeeService(ctrl)
	input := `{"name":"Updated","email":"updated@example.com","position":"Lead","salary":80000}`
	updated := models.Employee{ID: 1, Name: "Updated", Email: "updated@example.com", Position: "Lead", Salary: 80000}
	svc.EXPECT().UpdateEmployee(uint(1), gomock.Any()).Return(updated, nil)

	r := setupRouterWithController(svc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/employees/1", strings.NewReader(input))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated")
}

func TestDeleteEmployeeHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := mocks.NewMockEmployeeService(ctrl)
	svc.EXPECT().DeleteEmployee(uint(1)).Return(nil)

	r := setupRouterWithController(svc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/employees/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Employee deleted successfully")
}
