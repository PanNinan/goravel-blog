<div align="center">

<h1>Goravel Blog</h1>

<p>基于 <a href="https://github.com/goravel/goravel">Goravel</a> 构建的功能完善的博客系统 —— 一个 Laravel 风格的 Go Web 框架。</p>

[![Go Version](https://img.shields.io/github/go-mod/go-version/goravel/framework)](https://go.dev/)
[![Goravel](https://img.shields.io/badge/Goravel-v1.16-blue)](https://www.goravel.dev/zh)
[![License](https://img.shields.io/github/license/goravel/framework)](./LICENSE)
[![CI](https://github.com/PanNinan/goravel-blog/actions/workflows/ci.yml/badge.svg)](https://github.com/PanNinan/goravel-blog/actions/workflows/ci.yml)

[English](./README.md) | 中文

</div>

---

## 📖 项目介绍

**Goravel Blog** 是使用 [Goravel](https://www.goravel.dev/zh) 框架实现的社区风格博客平台，涵盖用户认证、文章管理、分类管理、评论系统等核心功能。项目遵循 Laravel 风格的 MVC 约定，是熟悉 PHP/Laravel 的开发者上手 Go Web 开发的绝佳参考示例。

---

## 🏗️ 项目架构

```
goravel-blog/
├── app/
│   ├── console/          # Artisan 命令定义
│   ├── events/           # 事件定义
│   ├── grpc/             # gRPC 服务处理器
│   ├── http/
│   │   ├── controllers/  # HTTP 请求控制器
│   │   │   ├── auth_controller.go      # 登录 / 登出 / 刷新 Token
│   │   │   ├── user_controller.go      # 用户 CRUD + 当前用户
│   │   │   ├── topic_controller.go     # 文章管理
│   │   │   ├── category_controller.go  # 分类管理
│   │   │   ├── reply_controller.go     # 评论管理
│   │   │   └── link_controller.go      # 友情链接
│   │   └── middleware/
│   │       └── jwt.go    # JWT 认证中间件
│   ├── jobs/             # 队列 Job 定义
│   ├── listeners/        # 事件监听器
│   ├── models/           # ORM 数据模型
│   │   ├── user.go       # 用户模型
│   │   ├── topic.go      # 文章模型
│   │   ├── category.go   # 分类模型
│   │   ├── reply.go      # 评论模型
│   │   ├── link.go       # 友链模型
│   │   ├── notification.go
│   │   └── common/       # 统一响应封装
│   └── providers/        # 服务提供者
├── bootstrap/            # 应用启动引导
├── config/               # 全局配置文件
│   ├── app.go            # 应用基础配置
│   ├── auth.go           # 认证守卫
│   ├── cache.go          # 缓存驱动
│   ├── database.go       # MySQL + Redis 连接
│   ├── http.go           # HTTP 服务配置
│   ├── jwt.go            # JWT 密钥与有效期
│   ├── queue.go          # 队列（sync / database / redis）
│   └── ...
├── database/
│   ├── migrations/       # 数据库迁移文件
│   └── seeders/          # 数据填充
├── resources/            # 视图模板（.tmpl）
├── routes/
│   ├── api.go            # REST API 路由
│   ├── web.go            # Web 页面路由
│   └── grpc.go           # gRPC 路由
├── tests/                # 功能测试 & 单元测试
├── Dockerfile            # 多阶段 Docker 构建
├── docker-compose.yml    # 本地开发 Compose 配置
└── main.go               # 程序入口
```

### 技术栈

| 层级 | 技术选型 |
|---|---|
| 编程语言 | Go 1.23+ |
| Web 框架 | Goravel v1.16（Laravel 风格） |
| HTTP 路由 | Gin v1.10 |
| ORM | GORM（通过 Goravel ORM Facade） |
| 数据库 | MySQL 8.x |
| 缓存 / 队列 | Redis |
| 用户认证 | JWT（golang-jwt/jwt v5） |
| 远程过程调用 | gRPC |
| 任务调度 | Goravel Schedule（基于 cron） |
| 队列工作者 | Goravel Queue（sync / database / redis 驱动） |
| 容器化 | Docker + Docker Compose |

---

## ✨ 功能特性

- **用户认证** — 基于 JWT 的登录、登出、Token 刷新、当前用户信息
- **用户管理** — 增删改查、头像与个人简介
- **文章（Topic）系统** — 文章创建、查看、修改、删除，支持分类、回复数、浏览数、Slug
- **分类管理** — 文章多分类管理
- **评论系统** — 关联文章的评论功能
- **友情链接** — 友链管理模块
- **任务调度** — 内置 cron 调度器（`facades.Schedule()`）
- **队列任务** — 异步 Job 处理，支持 sync / database / redis 驱动
- **gRPC 支持** — 可选 gRPC 服务端点
- **优雅关闭** — 处理 OS 信号（SIGINT / SIGTERM）

---

## 🚀 快速上手

### 环境要求

- Go 1.23+
- MySQL 8.x
- Redis（可选，用于缓存/队列）
- Docker & Docker Compose（可选）

### 本地开发

**1. 克隆仓库**

```bash
git clone https://github.com/PanNinan/goravel-blog.git
cd goravel-blog
```

**2. 安装依赖**

```bash
go mod tidy
```

**3. 配置环境变量**

复制示例配置文件并填写您的参数：

```bash
cp .env.example .env
```

关键配置项：

```env
APP_NAME=GravelBlog
APP_ENV=local
APP_KEY=           # 通过命令生成：go run . artisan key:generate
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

JWT_SECRET=        # 您的 JWT 密钥
```

**4. 生成应用密钥**

```bash
go run . artisan key:generate
```

**5. 执行数据库迁移**

```bash
go run . artisan migrate
```

**6. 启动应用**

```bash
go run .
```

默认访问地址：`http://localhost:3000`

---

### Docker Compose 一键启动

```bash
# 构建并启动所有服务
docker-compose up -d --build

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down
```

---

## 📡 API 接口

Base URL：`http://localhost:3000`

### 认证接口

| 方法 | 路径 | 说明 | 需要认证 |
|------|------|------|:---:|
| POST | `/auth/login` | 登录，返回 JWT Token | ✗ |
| GET | `/auth/info` | 获取当前用户信息 | ✓ |
| POST | `/auth/logout` | 登出 | ✓ |
| POST | `/auth/refresh` | 刷新 JWT Token | ✓ |

### 用户接口

| 方法 | 路径 | 说明 | 需要认证 |
|------|------|------|:---:|
| GET | `/users` | 用户列表 | ✗ |
| GET | `/users/{id}` | 用户详情 | ✗ |
| POST | `/users` | 创建用户 | ✗ |
| PUT | `/users/{id}` | 更新用户 | ✗ |
| DELETE | `/users/{id}` | 删除用户 | ✗ |
| GET | `/users/current` | 当前登录用户 | ✓ |

### 分类接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/categories` | 分类列表 |
| GET | `/categories/{id}` | 分类详情 |
| POST | `/categories` | 创建分类 |
| PUT | `/categories/{id}` | 更新分类 |
| DELETE | `/categories/{id}` | 删除分类 |

> **说明：** 请求体使用 `Content-Type: application/json`，JWT Token 通过 `Authorization: Bearer <token>` Header 传递。

---

## 🧪 运行测试

```bash
# 运行所有测试
go test ./...

# 详细输出
go test -v ./...

# 运行指定测试套件
go test -v ./tests/feature/...
```

---

## 🐳 生产部署

### Docker 部署（推荐）

项目使用多阶段 `Dockerfile` 构建最小化 Alpine 镜像：

```bash
# 构建镜像
docker build -t goravel-blog:latest .

# 运行容器
docker run -d \
  -p 3000:3000 \
  --env-file .env \
  goravel-blog:latest
```

### 手动编译部署

```bash
CGO_ENABLED=0 go build --ldflags "-s -w" -o goravel-blog .
./goravel-blog
```

### SSH 服务器部署

自动化部署需在 GitHub 仓库 Settings → Secrets 中添加以下变量：

| Secret 名称 | 说明 |
|------------|------|
| `SERVER_HOST` | 服务器主机名或 IP |
| `SERVER_USER` | SSH 登录用户名 |
| `SERVER_SSH_KEY` | SSH 私钥 |
| `SERVER_PORT` | SSH 端口（默认 22） |
| `DEPLOY_PATH` | 服务器上的应用部署目录 |

---

## 🔄 CI/CD 流程

本项目使用 GitHub Actions 实现自动化构建、测试与部署：

- **CI 工作流**（`.github/workflows/ci.yml`）— 在每次推送或 PR 到 `main` / `develop` 分支时触发，执行 `go vet`、`go test` 和二进制编译验证。
- **CD 工作流**（`.github/workflows/deploy.yml`）— 在 `main` 分支有推送时触发，构建并推送 Docker 镜像到镜像仓库，然后通过 SSH 部署到服务器。

完整配置见 [`.github/workflows/`](.github/workflows/)。

---

## 🤝 参与贡献

欢迎提交 Pull Request 和 Issue！请保持现有代码风格，并为新功能添加测试用例。

---

## 📄 开源许可

本项目基于 [MIT 许可证](./LICENSE) 开源。
