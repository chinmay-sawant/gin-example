package main

import (
	"github.com/chinmay-sawant/gin-example/controllers"
	"github.com/chinmay-sawant/gin-example/db"
	"github.com/chinmay-sawant/gin-example/docs"
	"github.com/chinmay-sawant/gin-example/repo"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Inside main()
	docs.SwaggerInfo.Title = "Employee Management API"
	docs.SwaggerInfo.Description = "API for managing employees"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	// Initialize database
	db.ConnectDatabase()

	// Create a new Gin router
	router := gin.Default()
	// Use gin-swagger middleware to expose Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	employeeRepo := repo.NewEmployeeRepository()
	// Create controllers
	employeeController := controllers.NewEmployeeController(employeeRepo)

	// Routes
	v1 := router.Group("/api/v1")
	employeeController.RegisterRoutes(v1)

	// Start the server
	router.Run(":8080")
}
