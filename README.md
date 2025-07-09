# Go Backend RESTful API Template

This is a simple **Go (Golang)** backend template for building a **RESTful API**. It uses **GORM** for MySQL database interaction and **JWT-based Authentication** to secure your routes. The template is structured with clean architecture principles, making it easy to extend and maintain.

## Features
- **JWT Authentication**: Secure your routes with JWT tokens.
- **GORM**: Database interactions are handled using the GORM ORM (supports MySQL).
- **Gin**: High-performance web framework for Go.
- **Clean Architecture**: Organized project structure for maintainability.
- **User Registration & Login**: Example implementation of user registration and authentication with JWT.

## Requirements

- **Go**: Version 1.16 or higher
- **MySQL**: You should have a MySQL database installed or use a Dockerized MySQL instance.

## Project Structure
```/go-project
├── /cmd
│ └── /main.go # Entry point of the application
├── /config # Configuration for DB, JWT secret key, etc.
├── /controllers # Logic for handling requests and interacting with services
├── /middlewares # Middleware like JWT validation
├── /models # Database schema (User, Post, etc.)
├── /routes # API routing
├── /services # Business logic like register, login, etc.
├── /utils # Helper functions (e.g., hash password)
├── /database # DB connection functions
└── /tests # Unit and integration tests
```
## Run Locally

Clone the project

```bash
  git clone https://github.com/puwadon-swpt/go-project.git
```

Go to the project directory

```bash
  cd go-project
```

Install dependencies

```bash
  go mod tidy
```

Start docker

```bash
  docker-compose up -d --build
```


## API Reference

#### POST Login

```http
  POST /login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**.|
| `password` | `string` | **Required**.|

#### Response:
***200 OK***: Return the JWT token.

***401 Unauthorized***: If credentials are incorrect or missing.

#### POST Register

```http
  POST /register
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `username` | `string` | **Required**.|
| `password` | `string` | **Required**.|
| `email` | `string` | **Required**.|

#### Response:

**201 Created**: Successfully registered the user.

**400 Bad Request**: If the user already exists or the input is invalid.
#### GET Profile

```http
  GET /api/profile
```
Authorization: Required. Bearer token for JWT authentication.
#### Response:

***200 OK***: Return the user's profile information.

***401 Unauthorized***: If the JWT token is missing or invalid.
| Parameter | Type     |
| :-------- | :------- |
| `id` | 1 |
| `username` | `newuser`|
| `email` | `user@example.com`|
| `created_at` | `2021-10-01T00:00:00Z`|
| `updated_at` | `2021-10-01T00:00:00Z` |





## License

[MIT](https://choosealicense.com/licenses/mit/)

