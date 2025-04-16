# 🏦 Simple Bank

A modern, production-ready banking backend application built with Go, featuring REST and gRPC APIs, secure authentication, database migrations, and complete CI/CD pipeline.

## ✨ Features

- 🔐 Secure user authentication with JWT and PASETO tokens
- 💰 Account management with balance tracking
- 💸 Money transfers between accounts with transaction history
- 📊 Database schema migration with golang-migrate
- 🔄 Concurrent transaction processing with proper locking
- 🛣️ RESTful API with Gin framework
- 🌐 High-performance gRPC API with Protocol Buffers
- 📚 API documentation with Swagger/OpenAPI
- 🧪 Comprehensive test coverage with testify
- 🐳 Containerization with Docker and Docker Compose
- 🔄 CI/CD with GitHub Actions
- 📧 Email verification system with async processing
- 📝 Structured logging with zerolog
- ⚙️ Configuration management with Viper

## 🛠️ Technology Stack

- **Language**: [Go 1.24](https://golang.org/)
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: [PostgreSQL 12](https://www.postgresql.org/)
- **SQL Generator**: [SQLC](https://sqlc.dev/)
- **API Layer**: 
  - RESTful API with Gin
  - [gRPC](https://grpc.io/) with Protocol Buffers
  - [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- **Authentication**: 
  - [JWT](https://github.com/golang-jwt/jwt)
  - [PASETO](https://github.com/o1egl/paseto)
- **Task Queue**: 
  - [Redis](https://redis.io/)
  - [Asynq](https://github.com/hibiken/asynq)
- **Validation**: [validator](https://github.com/go-playground/validator)
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Logging**: [zerolog](https://github.com/rs/zerolog)
- **Email**: [jordan-wright/email](https://github.com/jordan-wright/email)
- **Testing**: 
  - [Testify](https://github.com/stretchr/testify)
  - [gomock](https://github.com/uber-go/mock)
- **CI/CD**: GitHub Actions
- **Containerization**: Docker, Docker Compose

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

- [Go 1.24+](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master#installation)
- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)
- [protoc](https://grpc.io/docs/protoc-installation/) with Go plugins (for gRPC development)

## 🚀 Quick Start

### Clone the repository

```bash
git clone https://github.com/pawaspy/simple_bank.git
cd simple_bank
```

### Using Docker Compose (Recommended)

The fastest way to get started is using Docker Compose:

```bash
# Start all services (PostgreSQL, Redis, API server)
docker compose up

# Run in detached mode
docker compose up -d
```

### Manual Setup

#### 1. Set up environment variables

Copy the example environment file and configure as needed:

```bash
cp app.env.example app.env
# Edit app.env with your preferred editor
```

#### 2. Start PostgreSQL

```bash
# Create a Docker network for the services
docker network create simplebank-net

# Start PostgreSQL container
make postgres

# Create the database
make createdb
```

#### 3. Database Migration

```bash
# Run all migrations
make migrateup

# Run only one migration forward
make migrateup1

# Roll back all migrations
make migratedown

# Roll back one migration
make migratedown1

# Create a new migration file
make new_migration name=add_users_table
```

#### 4. Generate SQL code

```bash
# Generate Go code from SQL queries
make sqlc
```

#### 5. Start Redis for task queue

```bash
make redis
```

#### 6. Start the server

```bash
make server
```

## 🧪 Testing

```bash
# Run all tests
make test

# Generate mock for testing
make mock
```

## 🖥️ API Access

- **REST API**: http://localhost:8080
- **gRPC API**: localhost:9090
- **Swagger UI**: http://localhost:8080/swagger/index.html

### Testing gRPC API with Evans

```bash
# Install Evans (gRPC client)
go install github.com/ktr0731/evans@latest

# Connect to the gRPC server
make evans
```

## 📚 Detailed Setup Instructions

### Installing Required Tools

#### Install golang-migrate

```bash
# For macOS
brew install golang-migrate

# For Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/bin/migrate
```

#### Install sqlc

```bash
# For macOS
brew install sqlc

# For Linux or other platforms
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

#### Install Protocol Buffer Compiler

```bash
# For macOS
brew install protobuf

# For Linux
apt install -y protobuf-compiler
```

#### Install Go gRPC plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

### Setting up the database schema

```bash
# Generate database schema documentation
make db_docs

# Generate SQL from DBML
make db_schema
```

### Generating Protocol Buffers

```bash
# Generate Go code from Protocol Buffer definitions
make proto
```

## 🔧 Available Make Commands

| Command | Description |
|---------|-------------|
| `make postgres` | Start PostgreSQL container |
| `make stop` | Stop PostgreSQL container |
| `make createdb` | Create database inside container |
| `make dropdb` | Drop the database |
| `make migrateup` | Apply all migrations |
| `make migratedown` | Revert all migrations |
| `make migrateup1` | Apply single migration |
| `make migratedown1` | Revert single migration |
| `make new_migration name=xyz` | Create new migration files |
| `make sqlc` | Generate Go code from SQL |
| `make db_docs` | Generate database documentation |
| `make db_schema` | Generate SQL from DBML |
| `make test` | Run all tests |
| `make server` | Start the API server |
| `make mock` | Generate mocks for testing |
| `make proto` | Generate Protocol Buffer code |
| `make evans` | Start Evans gRPC client |
| `make redis` | Start Redis container |

## 🔒 Authentication

The API supports two authentication methods:

1. **JWT (JSON Web Token)** - Industry standard, stateless tokens
2. **PASETO (Platform-Agnostic Security Tokens)** - More secure alternative to JWT

Both methods support:
- Access tokens for short-term authentication
- Refresh tokens for obtaining new access tokens
- Token revocation

## 📝 API Documentation

- **REST API**: Available via Swagger UI at `http://localhost:8080/swagger/index.html`
- **gRPC API**: Protocol Buffer definitions in the `proto/` directory

## 🚀 Deployment

### Docker

Build and run using Docker:

```bash
# Build Docker image
docker build -t simple-bank:latest .

# Run container
docker run -p 8080:8080 -p 9090:9090 simple-bank:latest
```

### Docker Compose

Deploy the entire stack:

```bash
docker compose up -d
```

## 🧪 CI/CD Pipeline

The project includes GitHub Actions workflows for continuous integration:

- Automated testing on push and pull requests
- Database migrations are run before tests
- Tests run against PostgreSQL in a Docker container

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 👏 Acknowledgements

- [Go](https://golang.org/)
- [SQLC](https://sqlc.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [gRPC](https://grpc.io/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Viper](https://github.com/spf13/viper)
- [zerolog](https://github.com/rs/zerolog)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)
- [Asynq](https://github.com/hibiken/asynq)
