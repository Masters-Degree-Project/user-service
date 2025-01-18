# Go Authentication Service

A robust authentication service built with Go, featuring JWT-based authentication, user management, and secure password handling.

## Getting Started

### Prerequisites

- Go 1.19 or higher
- PostgreSQL

### Installation & Setup

1. Clone and install dependencies:
```
git clone https://github.com/Masters-Degree-Project/user-service
cd https://github.com/Masters-Degree-Project/user-service
go mod download
```

2. Configure environment:
   - Copy `.env.dist` to `.env`
   - Update the variables in `.env` file

### Running the Application

Start the server:
```
go run cmd/main.go
```

## Environment Variables

Required environment variables:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_database
JWT_SECRET=your_jwt_secret
```

## API Endpoints

### Authentication
- `POST /api/v1/login`
  - Login with email and password
  - Request body: `{"email": "user@example.com", "password": "yourpassword"}`
