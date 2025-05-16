package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/chinmay-sawant/gin-example/models"
	"github.com/chinmay-sawant/gin-example/repo"
	"github.com/chinmay-sawant/gin-example/service"
	"github.com/gin-gonic/gin"
)

// EmployeeController handles employee-related HTTP requests
type EmployeeController struct {
	employeeService service.EmployeeService
}

// NewEmployeeController creates a new instance of EmployeeController
func NewEmployeeController(repo repo.EmployeeRepository) *EmployeeController {
	return &EmployeeController{
		employeeService: service.NewEmployeeService(repo),
	}
}

// RegisterRoutes registers employee routes to the provided router group
func (ec *EmployeeController) RegisterRoutes(router *gin.RouterGroup) {
	employees := router.Group("/employees")
	{
		employees.GET("/", ec.GetEmployees)
		employees.GET("/:id", ec.GetEmployee)
		employees.POST("/", ec.CreateEmployee)
		employees.PUT("/:id", ec.UpdateEmployee)
		employees.DELETE("/:id", ec.DeleteEmployee)
	}
}

// GetEmployees handles GET request to fetch all employees
func (ec *EmployeeController) GetEmployees(c *gin.Context) {
	employees, err := ec.employeeService.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

// GetEmployee handles GET request to fetch a specific employee by ID
func (ec *EmployeeController) GetEmployee(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	employee, err := ec.employeeService.GetEmployeeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// CreateEmployee handles POST request to create a new employee
func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the join date to current time if not provided
	if employee.JoinDate.IsZero() {
		employee.JoinDate = time.Now()
	}

	createdEmployee, err := ec.employeeService.CreateEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdEmployee)
}

// UpdateEmployee handles PUT request to update an existing employee
func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEmployee, err := ec.employeeService.UpdateEmployee(uint(id), employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedEmployee)
}

// DeleteEmployee handles DELETE request to remove an employee
func (ec *EmployeeController) DeleteEmployee(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	err = ec.employeeService.DeleteEmployee(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
