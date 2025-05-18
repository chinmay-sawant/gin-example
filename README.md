# Employee Management System with Gin

A simple employee management system built with Go and Gin framework, showcasing best practices for organizing a Gin project with controllers, models, and services using interfaces.

## Project Structure

```
.
├── controllers/         # HTTP request handlers (interface-based)
│   ├── employee_controller.go         # Interface
│   ├── employee_controller_impl.go    # Implementation
│   └── mocks/                        # Generated controller mocks
│       └── mock_employee_controller.go
├── db/                  # Database configuration
│   └── database.go
├── models/              # Data models
│   └── employee.go
├── repo/                # Data access layer (repository pattern)
│   ├── employee_repo.go         # Interface
│   ├── employee_repo_impl.go    # Implementation
│   └── mocks/                  # Generated repo mocks
│       └── mock_employee_repo.go
├── service/             # Business logic (interface-based)
│   ├── employee_service.go       # Interface
│   ├── employee_service_impl.go  # Implementation
│   └── mocks/                   # Generated service mocks
│       └── mock_employee_service.go
## Mocking & Testing

- All interfaces (in `controllers/`, `service/`, and `repo/`) can be mocked for unit testing.
- Mocks are generated using [mockgen](https://go.uber.org/mock/gomock).
- Generated mocks are placed in the `mocks/` subdirectory of each package.
- Example command to generate a mock for the `EmployeeService` interface:
  ```cmd
  mockgen -source=service/employee_service.go -destination=service/mocks/mock_employee_service.go -package=mocks
  ```
- Use these mocks in your tests with [testify](https://github.com/stretchr/testify) for assertions.
├── go.mod               # Go module file
├── main.go              # Entry point
└── README.md            # Documentation
```

## Features

- Clean architecture with separation of concerns
- Interface-based service and controller layers for better testability and flexibility
- Repository pattern for all database interactions (no DB logic in service layer)
- RESTful API endpoints for CRUD operations
- SQLite database with GORM ORM (using pure Go driver, no CGO required)


## API Endpoints

- `GET /api/v1/employees` - Get all employees
- `GET /api/v1/employees/{id}` - Get a specific employee
- `POST /api/v1/employees` - Create a new employee
- `PUT /api/v1/employees/{id}` - Update an existing employee
- `DELETE /api/v1/employees/{id}` - Delete an employee


## How to Run

```bash
# Run the application
go run main.go
```

The server will start on http://localhost:8080

## Swagger/OpenAPI Documentation

Swagger UI is available at: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

If you change your API, regenerate docs with:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

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

1. **Controllers**: Handle HTTP requests and responses (use interfaces)
2. **Services**: Implement business logic (use interfaces, no DB logic)
3. **Models**: Define data structures
4. **Repository**: Handles all database interactions (all DB logic here)
5. **Database**: Handles database connections and migrations

## Notes

- All database logic is in the `repo` package, following the repository pattern.
- Both service and controller layers use interfaces and dependency injection.
- The project is easily extensible and testable due to this separation.
- Uses pure Go SQLite driver for maximum portability (no CGO required).
