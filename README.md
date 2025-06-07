# Jobs Collection

A collection of Go applications for various job/service tasks.

## 📁 Project Structure

```
jobs/
├── hello/          # Greeting service with environment variable support
│   ├── go.mod
│   ├── main.go
│   ├── Dockerfile
│   └── .dockerignore
├── info/           # Cronjob logging service
│   ├── go.mod
│   ├── main.go
│   ├── Dockerfile
│   └── .dockerignore
├── docker-compose.yml
└── README.md       # This file
```

## 🚀 Applications

### Hello Service (`hello/`)

A simple greeting application that reads from environment variables.

**Features:**

- Reads `CUSTOM_USER` environment variable
- Provides personalized greeting or default message

**Usage:**

```bash
cd hello
go run main.go                    # Output: Hi USER_NOT_SPECIFIED
CUSTOM_USER=Akash go run main.go  # Output: Hi, Akash
```

**Build:**

```bash
cd hello
go build -o hello main.go
./hello
```

**Docker:**

```bash
cd hello
# Build Docker image
docker build -t jobs-hello .

# Run once with default behavior
docker run --rm jobs-hello

# Run once with custom user
docker run --rm -e CUSTOM_USER=Akash jobs-hello
```

### Info Service (`info/`)

A cronjob logging service for scheduled tasks.

**Features:**

- Simple logging output for cronjob monitoring
- Lightweight and fast execution

**Usage:**

```bash
cd info
go run main.go  # Output: This is a cronjob log.
```

**Build:**

```bash
cd info
go build -o info main.go
./info
```

**Docker:**

```bash
cd info
# Build Docker image
docker build -t jobs-info .

# Run once
docker run --rm jobs-info
```

## 🛠️ Development

### Prerequisites

- Go 1.24.3 or later
- Docker (optional, for containerized deployment)
- Git

### Building All Applications

```bash
# Build hello service
cd hello && go build -o ../bin/hello main.go && cd ..

# Build info service
cd info && go build -o ../bin/info main.go && cd ..
```

### Building Docker Images

```bash
# Build all Docker images
docker build -t jobs-hello ./hello
docker build -t jobs-info ./info

# Or build individually
cd hello && docker build -t jobs-hello . && cd ..
cd info && docker build -t jobs-info . && cd ..

# Using Docker Compose (recommended)
docker-compose build
```

### Running Tests

```bash
# Test hello service
cd hello && go test ./... && cd ..

# Test info service
cd info && go test ./... && cd ..
```

## 📦 Module Organization

This project uses **multiple Go modules** (one per service), which is the recommended approach for:

- ✅ Independent applications with separate concerns
- ✅ Different dependency requirements
- ✅ Independent versioning and releases
- ✅ Clear separation of responsibilities

Each service maintains its own `go.mod` file for optimal dependency management and deployment flexibility.

## 🔧 Environment Variables

### Hello Service

| Variable      | Description                        | Default              |
| ------------- | ---------------------------------- | -------------------- |
| `CUSTOM_USER` | Username for personalized greeting | `USER_NOT_SPECIFIED` |

## 📋 Common Tasks

### Adding a New Service

1. Create a new directory: `mkdir newservice`
2. Initialize Go module: `cd newservice && go mod init github.com/akash1047/jobs/newservice`
3. Create your `main.go` file
4. Update this README.md with the new service information

### Deployment

Each service can be deployed independently:

**Native Binary Deployment:**

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o service-linux main.go

# Build for different platforms
GOOS=windows GOARCH=amd64 go build -o service-windows.exe main.go
GOOS=darwin GOARCH=amd64 go build -o service-macos main.go
```

**Docker Deployment:**

```bash
# Build images only
docker-compose build

# Run individual tools once
docker run --rm jobs-hello
docker run --rm jobs-info

# Run with Docker Compose (using profiles)
docker-compose --profile tools run --rm hello
docker-compose --profile tools run --rm info

# For cronjob usage with Docker
# Add to crontab: 0 */6 * * * docker run --rm jobs-info
```

## 🐳 Docker

Each service includes a multi-stage Dockerfile optimized for:

- ✅ **Small image size** - Uses Alpine Linux base
- ✅ **Security** - Runs as non-root user
- ✅ **Efficiency** - Multi-stage build separates build and runtime
- ✅ **Production ready** - Includes CA certificates for HTTPS

**Docker Image Features:**

- Base: `alpine:latest` (~5MB)
- User: Non-root (`appuser:appgroup`)
- Go version: 1.24.3
- Static binary (CGO disabled)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request
6. Create a feature branch
7. Make your changes
8. Test thoroughly
9. Submit a pull request

## 📄 License

This project is part of a personal jobs collection. Please refer to individual service directories for specific licensing information.

---

**Last Updated:** June 7, 2025
