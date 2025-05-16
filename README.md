# Employee Management System with Gin

A simple employee management system built with Go and Gin framework, showcasing best practices for organizing a Gin project with controllers, models, and services using interfaces.

## Project Structure

```
.
├── controllers/         # HTTP request handlers (interface-based)
│   └── employee_controller.go
├── db/                  # Database configuration
│   └── database.go
├── models/              # Data models
│   └── employee.go
├── repo/                # Data access layer (repository pattern)
│   ├── employee_repo.go         # Interface
│   └── employee_repo_impl.go    # Implementation
├── service/             # Business logic (interface-based)
│   ├── employee_service.go       # Interface
│   └── employee_service_impl.go  # Implementation
├── go.mod               # Go module file
├── main.go              # Entry point
└── README.md            # Documentation
```

## Features

- Clean architecture with separation of concerns
- Interface-based service and controller layers for better testability and flexibility
- Repository pattern for all database interactions
- RESTful API endpoints for CRUD operations
- SQLite database with GORM ORM

## API Endpoints

- `GET /api/v1/employees` - Get all employees
- `GET /api/v1/employees/:id` - Get a specific employee
- `POST /api/v1/employees` - Create a new employee
- `PUT /api/v1/employees/:id` - Update an existing employee
- `DELETE /api/v1/employees/:id` - Delete an employee

## How to Run

```bash
# Run the application
go run main.go
```

The server will start on http://localhost:8080

## Example API Usage

### Create a new employee:
```bash
curl -X POST http://localhost:8080/api/v1/employees \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","position":"Software Engineer","salary":75000}'
```

### Get all employees:
```bash
curl http://localhost:8080/api/v1/employees
```

## Project Design

This project follows a clean architecture pattern with the following layers:

1. **Controllers**: Handle HTTP requests and responses
2. **Services**: Implement business logic (with interfaces)
3. **Models**: Define data structures
4. **Repository**: Handles all database interactions (no DB logic in service)
5. **Database**: Handles database connections and migrations

## Notes

- All database logic is in the `repo` package, following the repository pattern.
- Both service and controller layers use interfaces and dependency injection.
- The project is easily extensible and testable due to this separation.
