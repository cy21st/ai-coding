# Metadata Management System

A backend management system for metadata management of a big data tracking system.


```
最开始的prompt:

设计一个管理后台，大数据埋点系统的元数据管理；
后台用户表，元数据表的事件表、属性表、以及他们的关联表在sql文件夹中
使用golang的web框架gin来开发，使用gorm的ORM框架，数据库用mysql，缓存用redis
身份验证使用jwt来，根据api.md的接口要求生成对应的接口代码，并且目录结构清晰易懂


除了获取添加、删除、修改、登录统一用POST接口其他一律用GET
##后台的用户管理模块接口
-用户信息
-用户添加、删除、修改
-用户列表
-用户登录

除了获取添加、删除、修改统一用POST接口其他一律用GET
##元数据管理接口
-事件的添加修改删除
-添加属性并关联事件、属性的修改和删除
-事件列表
-属性列表
-事件的属性列表
-事件信息
-属性信息
```


## Features

- User Management
  - User authentication with JWT
  - User CRUD operations
  - Role-based access control

- Metadata Management
  - Event management (CRUD)
  - Attribute management (CRUD)
  - Event-Attribute relationship management

## Tech Stack

- Go 1.21+
- Gin Web Framework
- GORM ORM
- MySQL Database
- Redis Cache
- JWT Authentication

## Project Structure

```
.
├── config/         # Configuration files
├── database/       # Database connection and initialization
├── handlers/       # HTTP request handlers
├── middleware/     # Middleware components (JWT, etc.)
├── models/         # Data models
├── sql/           # SQL schema files
├── go.mod         # Go module file
├── main.go        # Application entry point
└── README.md      # Project documentation
```

## API Endpoints

### User Management

- POST /api/login - User login
- POST /api/users - Create user
- GET /api/users - Get user list
- GET /api/users/info - Get current user info
- POST /api/users/:id - Update user
- POST /api/users/:id/delete - Delete user

### Event Management

- POST /api/events - Create event
- POST /api/events/:id - Update event
- POST /api/events/:id/delete - Delete event
- GET /api/events - Get event list
- GET /api/events/:id - Get event info
- GET /api/events/:id/attributes - Get event attributes

### Attribute Management

- POST /api/attributes - Create attribute
- POST /api/attributes/:id - Update attribute
- POST /api/attributes/:id/delete - Delete attribute
- GET /api/attributes - Get attribute list
- GET /api/attributes/:id - Get attribute info

## Setup

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Configure the database and Redis connection in `config/config.go`

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on port 8080.

## Authentication

All endpoints except `/api/login` require JWT authentication. Include the JWT token in the Authorization header:

```
Authorization: Bearer <your-token>
``` 