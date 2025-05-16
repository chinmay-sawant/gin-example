package db

import (
	"log"

	"github.com/chinmay-sawant/gin-example/models"
	"github.com/glebarez/sqlite" // Pure Go SQLite driver, doesn't require CGO
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection and performs migrations
func ConnectDatabase() {
	// Using pure Go SQLite implementation - no CGO dependency
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to set up in-memory database:", err)
	}

	// Auto migrate the models
	err = database.AutoMigrate(&models.Employee{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database

	// Insert default employees
	employees := []models.Employee{
		{Name: "Alice Smith", Email: "alice@example.com", Position: "Developer", Salary: 70000},
		{Name: "Bob Johnson", Email: "bob@example.com", Position: "Designer", Salary: 65000},
		{Name: "Charlie Lee", Email: "charlie@example.com", Position: "Manager", Salary: 90000},
		{Name: "Diana King", Email: "diana@example.com", Position: "QA Engineer", Salary: 60000},
		{Name: "Ethan Brown", Email: "ethan@example.com", Position: "DevOps", Salary: 75000},
	}
	DB.CreateInBatches(&employees, 5)

	log.Println("Database connected and migrated successfully!")
}
