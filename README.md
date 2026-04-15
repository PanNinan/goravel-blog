<div align="center">

<h1>Goravel Blog</h1>

<p>A full-featured blog system built with <a href="https://github.com/goravel/goravel">Goravel</a> — a Laravel-style Go web framework.</p>

[![Go Version](https://img.shields.io/github/go-mod/go-version/goravel/framework)](https://go.dev/)
[![Goravel](https://img.shields.io/badge/Goravel-v1.16-blue)](https://www.goravel.dev)
[![License](https://img.shields.io/github/license/goravel/framework)](./LICENSE)
[![CI](https://github.com/PanNinan/goravel-blog/actions/workflows/ci.yml/badge.svg)](https://github.com/PanNinan/goravel-blog/actions/workflows/ci.yml)

[中文](./README_zh.md) | English

</div>

---

## 📖 About

**Goravel Blog** is a community-style blog platform implemented with [Goravel](https://www.goravel.dev) framework. It covers essential features including user authentication, article management, category management, commenting system and more. The project follows Laravel-style MVC conventions, making it a great reference for Go developers familiar with PHP/Laravel.

---

## 🏗️ Architecture Overview

```
goravel-blog/
├── app/
│   ├── console/          # Artisan command definitions
│   ├── events/           # Event definitions
│   ├── grpc/             # gRPC service handlers
│   ├── http/
│   │   ├── controllers/  # HTTP request handlers
│   │   │   ├── auth_controller.go      # Login / logout / refresh token
│   │   │   ├── user_controller.go      # User CRUD + current user
│   │   │   ├── topic_controller.go     # Post management
│   │   │   ├── category_controller.go  # Category management
│   │   │   ├── reply_controller.go     # Comment management
│   │   │   └── link_controller.go      # Friendly links
│   │   └── middleware/
│   │       └── jwt.go    # JWT authentication middleware
│   ├── jobs/             # Queue job definitions
│   ├── listeners/        # Event listeners
│   ├── models/           # Eloquent-style ORM models
│   │   ├── user.go       # User model
│   │   ├── topic.go      # Post model
│   │   ├── category.go   # Category model
│   │   ├── reply.go      # Reply model
│   │   ├── link.go       # Link model
│   │   ├── notification.go
│   │   └── common/       # Shared response helpers
│   └── providers/        # Service providers
├── bootstrap/            # Application bootstrap
├── config/               # All configuration files
│   ├── app.go            # Application settings
│   ├── auth.go           # Auth guard
│   ├── cache.go          # Cache driver
│   ├── database.go       # MySQL + Redis connection
│   ├── http.go           # HTTP server settings
│   ├── jwt.go            # JWT secret & TTL
│   ├── queue.go          # Queue (sync / database / redis)
│   └── ...
├── database/
│   ├── migrations/       # Database migration files
│   └── seeders/          # Database seeders
├── resources/            # View templates (.tmpl)
├── routes/
│   ├── api.go            # REST API routes
│   ├── web.go            # Web page routes
│   └── grpc.go           # gRPC routes
├── tests/                # Feature & unit tests
├── Dockerfile            # Multi-stage Docker build
├── docker-compose.yml    # Local development compose
└── main.go               # Application entry point
```

### Technology Stack

| Layer | Technology |
|---|---|
| Language | Go 1.23+ |
| Framework | Goravel v1.16 (Laravel-style) |
| HTTP Router | Gin v1.10 |
| ORM | GORM (via Goravel ORM facade) |
| Database | MySQL 8.x |
| Cache / Queue | Redis |
| Auth | JWT (golang-jwt/jwt v5) |
| RPC | gRPC |
| Task Scheduling | Goravel Schedule (cron-based) |
| Queue Worker | Goravel Queue (sync / database / redis driver) |
| Containerization | Docker + Docker Compose |

---

## ✨ Features

- **User Authentication** — JWT-based login, logout, token refresh, current user info
- **User Management** — CRUD, profile (avatar, introduction)
- **Topic (Post) System** — Create, view, update, delete posts with category, reply count, view count, slug
- **Category Management** — Post categorization with description
- **Reply / Comment System** — Nested comments associated with topics
- **Link Management** — Friendly links module
- **Task Scheduling** — Built-in cron scheduler via `facades.Schedule()`
- **Queue Workers** — Async job processing via sync, database, or redis driver
- **gRPC Support** — Optional gRPC service endpoint
- **Graceful Shutdown** — OS signal handling (SIGINT/SIGTERM)

---

## 🚀 Getting Started

### Prerequisites

- Go 1.23+
- MySQL 8.x
- Redis (optional, for cache/queue)
- Docker & Docker Compose (optional)

### Local Development

**1. Clone the repository**

```bash
git clone https://github.com/PanNinan/goravel-blog.git
cd goravel-blog
```

**2. Install dependencies**

```bash
go mod tidy
```

**3. Configure environment**

Copy the example env file and fill in your values:

```bash
cp .env.example .env
```

Key variables:

```env
APP_NAME=GravelBlog
APP_ENV=local
APP_KEY=           # generate with: go run . artisan key:generate
APP_DEBUG=true
APP_PORT=3000

DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=goravel_blog
DB_USERNAME=root
DB_PASSWORD=secret

REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=

QUEUE_CONNECTION=sync

JWT_SECRET=        # your jwt secret
```

**4. Generate application key**

```bash
go run . artisan key:generate
```

**5. Run database migrations**

```bash
go run . artisan migrate
```

**6. Start the application**

```bash
go run .
```

The server runs on `http://localhost:3000` by default.

---

### Docker Compose

```bash
# Build and start all services
docker-compose up -d --build

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

---

## 📡 API Reference

Base URL: `http://localhost:3000`

### Authentication

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|:---:|
| POST | `/auth/login` | Login, returns JWT token | ✗ |
| GET | `/auth/info` | Get current user info | ✓ |
| POST | `/auth/logout` | Logout | ✓ |
| POST | `/auth/refresh` | Refresh JWT token | ✓ |

### Users

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|:---:|
| GET | `/users` | List all users | ✗ |
| GET | `/users/{id}` | Get user by ID | ✗ |
| POST | `/users` | Create user | ✗ |
| PUT | `/users/{id}` | Update user | ✗ |
| DELETE | `/users/{id}` | Delete user | ✗ |
| GET | `/users/current` | Get current logged-in user | ✓ |

### Categories

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/categories` | List categories |
| GET | `/categories/{id}` | Get category by ID |
| POST | `/categories` | Create category |
| PUT | `/categories/{id}` | Update category |
| DELETE | `/categories/{id}` | Delete category |

> **Note:** Request payloads use `Content-Type: application/json`. JWT token should be passed via `Authorization: Bearer <token>` header.

---

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run a specific test suite
go test -v ./tests/feature/...
```

---

## 🐳 Production Deployment

### Docker (Recommended)

The multi-stage `Dockerfile` produces a minimal Alpine-based image:

```bash
# Build image
docker build -t goravel-blog:latest .

# Run container
docker run -d \
  -p 3000:3000 \
  --env-file .env \
  goravel-blog:latest
```

### Manual Build

```bash
CGO_ENABLED=0 go build --ldflags "-s -w" -o goravel-blog .
./goravel-blog
```

### Server Deployment via SSH

For automated deployment, set the following GitHub Actions secrets:

| Secret | Description |
|--------|-------------|
| `SERVER_HOST` | SSH server hostname or IP |
| `SERVER_USER` | SSH username |
| `SERVER_SSH_KEY` | SSH private key |
| `SERVER_PORT` | SSH port (default: 22) |
| `DEPLOY_PATH` | Application directory on server |

---

## 🔄 CI/CD

This project uses GitHub Actions for automated build, test, and deployment.

- **CI workflow** (`.github/workflows/ci.yml`) — triggered on every push/PR to `main` or `develop`: runs `go vet`, `go test`, and builds the binary.
- **CD workflow** (`.github/workflows/deploy.yml`) — triggered on push to `main`: builds and pushes Docker image to registry, then deploys to server via SSH.

See [`.github/workflows/`](.github/workflows/) for full configuration.

---

## 🤝 Contributing

Pull requests and issues are welcome! Please follow the existing code style and add tests for new features.

---

## 📄 License

This project is open-sourced software licensed under the [MIT license](./LICENSE).
