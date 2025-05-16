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

// RegisterRoutes registers the employee routes with the given router group.
// RegisterRoutes sets up the employee API routes on the specified router group

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
// GetEmployees godoc
// @Summary Get all employees
// @Description Retrieves all employees from the database
// @Tags employees
// @Accept json
// @Produce json
// @Success 200 {array} models.Employee
// @Failure 500 {object} map[string]interface{} "Error response"
// @Router /employees [get]
func (ec *EmployeeController) GetEmployees(c *gin.Context) {
	employees, err := ec.employeeService.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

// GetEmployee handles GET request to fetch a specific employee by ID
// @Summary Get employee by ID
// @Description Retrieves a specific employee by their ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} models.Employee
// @Failure 400 {object} map[string]interface{} "Invalid employee ID"
// @Failure 404 {object} map[string]interface{} "Employee not found"
// @Router /employees/{id} [get]
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
// @Summary Create employee
// @Description Creates a new employee record
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body models.Employee true "Employee object"
// @Success 201 {object} models.Employee
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 500 {object} map[string]interface{} "Error response"
// @Router /employees [post]
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
// @Summary Update employee
// @Description Updates an existing employee record
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body models.Employee true "Updated employee object"
// @Success 200 {object} models.Employee
// @Failure 400 {object} map[string]interface{} "Invalid employee ID or request data"
// @Failure 500 {object} map[string]interface{} "Error response"
// @Router /employees/{id} [put]
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
// @Summary Delete employee
// @Description Removes an employee from the database
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]interface{} "Success message"
// @Failure 400 {object} map[string]interface{} "Invalid employee ID"
// @Failure 500 {object} map[string]interface{} "Error response"
// @Router /employees/{id} [delete]
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
