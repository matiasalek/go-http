# 🏋️ Workout Management API – Golang, PostgreSQL & Docker

A clean and scalable RESTful API built in Go for managing workouts. This project was developed as part of the [Frontend Masters – Complete Go for Professional Developers](https://frontendmasters.com/courses/complete-go/) course.

## 🚀 Features

- Full CRUD for workouts (Create, Read, Update, Delete)
- User registration & login with password hashing
- Auth-protected routes using middleware
- PostgreSQL integration (Dockerized)
- Modular architecture with interfaces and dependency injection
- Unit testing for core business logic

## 📦 Tech Stack

- **Go 1.22+**
- **PostgreSQL**
- **Docker / Docker Compose**
- **Chi Router**
- **bcrypt**

## 🛠️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/matiasalek/go-http.git
cd go-http
```

### 2. Run the app using Docker Compose
```bash
docker-compose up --build
```

This will start:

    The Go backend on http://localhost:8000

    A PostgreSQL container on port 5432

### 3. 📡 API Endpoints
    Base URL: http://localhost:4000/api

    POST /register – Register a new user

    POST /login – Login and receive session

    GET /workouts – List workouts

    POST /workouts – Create new workout

    PUT /workouts/{id} – Update workout

    DELETE /workouts/{id} – Delete workout

🔐 Protected routes require a valid session cookie.

### 4. ✅ Running Tests

```bash
go test ./...
```
