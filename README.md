# URL Shortener
This project aims to showcase best practices in Go (Golang) by implementing the Hexagonal Architecture (Ports and Adapters) pattern. It emphasizes a clean separation of concerns between business logic and infrastructure, enabling high testability, scalability, and maintainability. The design follows principles such as dependency inversion, interface-driven development, and modularity, making the codebase easier to adapt to changing requirements or external systems (e.g., databases, APIs). Through this approach, the project serves as a practical example of building robust, flexible, and production-ready Go applications.

## 🏗️ Architecture

This project follows the Hexagonal Architecture (also known as Ports and Adapters) pattern, which provides a clear separation of concerns and makes the application more maintainable and testable. The architecture is organized into the following layers:

- **Handler Layer**: Implements use cases and orchestrates the flow of data. It is the entry point of the application.
- **Module Layer**: Contains the core business logic and entities.
- **Persistence Layer**: Handles concerns like databases.

### Project Structure

```
.
├── cmd/                      # Application entry points
|   └── main.go               # Main entry point of the application
├── config/                   # Configuration management
|   └── sqlc.yaml             # Defines how SQL queries are compiled into Go code using sqlc
|   └── config.yaml           # Centralized configuration file for application settings
├── internal/                 # Private application code
│   ├── handler/              # Handles HTTP requests and orchestrates use cases
|   |   └── middleware        # Contains reusable functions that intercept and process HTTP requests and responses
|   |   └── rest              # Contains HTTP handlers
|   |       └── gin           # Contains Gin-based HTTP handlers
|   |       └── handler.go    # Contains handlers interface
│   ├── module/               # Core business logic and entities
|   |   └── handler.go        # Contains modules interface
│   └── storage/              # Persistence layer for database interactions
|       └── storage.go        # Contains storages interface  
├── platform/                 # Platform-specific code
|   ├── logger/               # Logging(e.g., Zap integration)
|   └── utils/                # Utility functions (e.g., helpers functions)
└── initiator/                # Application initialization
    └── config.go             # Intialization of configuration
    └── db.go                 # Intializtion of database
    └── handler.go            # Intialization of handler layer
    └── initiato.go           # Sets up and wires dependencies before the application starts running
    └── logger.go             # Initialization of logger(eg. zap)
    └── module.go             # Initializtion of module layer
    └── persistence.go        # Initializtion of persistence layer
    └── routes.go             # Initialization of routes
```

## 🚀 Features

- URL shortening with custom aliases
- URL redirection
- PostgreSQL database integration
- RESTful API endpoints
- Input validation
- Structured logging
- Configuration management
- Clean architecture implementation

## 🛠️ Technology Stack

- **Go 1.24**: Core programming language
- **Gin**: Web framework
- **CockroachDB**: Database
- **pgx**: PostgreSQL driver
- **Zap**: Structured logging
- **Viper**: Configuration management
- **Ozzo Validation**: Input validation

## 📋 Prerequisites

- Go 1.24 or higher
- CockroachDB
- Make (optional, for using Makefile commands)

## 🔧 Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd url_shortener
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up your environment variables (create a `config.yaml` file based on the example in `config/`)

4. Run the application:
   ```bash
   go run cmd/main.go
   ```

## 🏃‍♂️ Running with Make

The project includes a Makefile for common operations:

```bash
make run        # Run the application
make sqlc       # To Generate GO Code from sql file
make test       # Run tests
```

