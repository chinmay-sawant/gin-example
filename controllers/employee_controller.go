package controllers

import (
	"github.com/gin-gonic/gin"
)

// EmployeeController defines the interface for employee controller
type EmployeeController interface {
	RegisterRoutes(router *gin.RouterGroup)
	GetEmployees(c *gin.Context)
	GetEmployee(c *gin.Context)
	CreateEmployee(c *gin.Context)
	UpdateEmployee(c *gin.Context)
	DeleteEmployee(c *gin.Context)
}
