package main

import (
	"github.com/chinmay-sawant/gin-example/controllers"
	"github.com/chinmay-sawant/gin-example/db"
	"github.com/chinmay-sawant/gin-example/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db.ConnectDatabase()

	// Create a new Gin router
	router := gin.Default()

	employeeRepo := repo.NewEmployeeRepository()
	// Create controllers
	employeeController := controllers.NewEmployeeController(employeeRepo)

	// Routes
	v1 := router.Group("/api/v1")
	employeeController.RegisterRoutes(v1)

	// Start the server
	router.Run(":8080")
}
