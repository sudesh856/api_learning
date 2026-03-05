# Todo API

A lightweight, production-ready REST API for managing tasks with user authentication. Built with Go and PostgreSQL, this project demonstrates clean architecture principles with proper separation of concerns.

> **Note**: This project was built following a YouTube tutorial. While the architecture and concepts were learned from that source, every line of code was written from scratch during implementation—no copy-pasting shortcuts here.

## Overview

This API provides a complete task management system where users can register, authenticate, and manage their TODO items. It features JWT-based authentication, a secure password handling system, and a well-structured codebase that's easy to extend and maintain.

## Tech Stack

- **Language**: Go 1.25.4
- **Web Framework**: Gin (high-performance HTTP web framework)
- **Database**: PostgreSQL with pgx driver
- **Authentication**: JWT (JSON Web Tokens) with bcrypt password hashing
- **Configuration**: Environment variables via godotenv

## Project Structure

```
.
├── cmd/api/                    # Application entry point
│   └── main.go
├── internal/
│   ├── config/                 # Configuration management
│   ├── database/               # Database connection setup
│   ├── handlers/               # HTTP request handlers
│   ├── middleware/             # Authentication middleware
│   ├── models/                 # Data models (User, Todo)
│   └── repository/             # Data access layer
├── migrations/                 # SQL migration files
├── scripts/                    # Utility scripts
└── go.mod                      # Go module definition
```

## Features

✓ **User Authentication** - Register and login with secure password hashing  
✓ **JWT Authorization** - Token-based authentication for protected endpoints  
✓ **Todo Management** - Create, read, update, and delete tasks  
✓ **User-Scoped Data** - Each user sees only their own todos  
✓ **Database Migrations** - Version-controlled schema changes  
✓ **Clean Architecture** - Handlers, repository, and middleware separation  

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- Go 1.25.4 or higher
- PostgreSQL 12 or higher
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd api_learning
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the project root:
   ```env
   DATABASE_URL=postgres://username:password@localhost:5432/todo_api
   PORT=8080
   JWT_SECRET=your_super_secret_jwt_key_change_this
   ```

4. **Create the database**
   ```bash
   createdb todo_api
   ```

5. **Run migrations**
   ```powershell
   .\scripts\migrate.ps1
   ```
   
   Or manually with migrate CLI:
   ```bash
   migrate -path migrations -database "$DATABASE_URL" up
   ```

### Running the Application

Start the API server:

```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Public Routes

**Register a new user**
```
POST /auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secure_password"
}
```

**Login**
```
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secure_password"
}
Response:
{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### Protected Routes

All todo endpoints require a valid JWT token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

**Create a todo**
```
POST /todos
Content-Type: application/json

{
  "title": "Buy groceries",
  "completed": false
}
```

**Get all todos**
```
GET /todos
```

**Get a specific todo**
```
GET /todos/:id
```

**Update a todo**
```
PUT /todos/:id
Content-Type: application/json

{
  "title": "Updated title",
  "completed": true
}
```

**Delete a todo**
```
DELETE /todos/:id
```

## Database Schema

### Users Table
- `id` - UUID (Primary Key)
- `email` - VARCHAR (Unique)
- `password` - VARCHAR (bcrypt hashed)
- `created_at` - TIMESTAMP
- `updated_at` - TIMESTAMP

### Todos Table
- `id` - SERIAL (Primary Key)
- `title` - VARCHAR(255)
- `completed` - BOOLEAN
- `user_id` - UUID (Foreign Key)
- `created_at` - TIMESTAMP
- `updated_at` - TIMESTAMP

## Development

### Code Standards

The codebase follows Go best practices:
- Package organization by feature (handlers, models, repository)
- Clear separation of concerns
- Repository pattern for data access
- Middleware for cross-cutting concerns

### Adding New Features

1. Define your model in `internal/models/`
2. Create repository methods in `internal/repository/`
3. Implement handlers in `internal/handlers/`
4. Wire up routes in `cmd/api/main.go`
5. Create database migrations for schema changes

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `DATABASE_URL` | PostgreSQL connection string | Yes |
| `PORT` | Server port (default: 8080) | No |
| `JWT_SECRET` | Secret key for JWT token signing | Yes |

## Troubleshooting

**Connection refused error**
- Ensure PostgreSQL is running
- Verify DATABASE_URL is correct
- Check that the database exists

**JWT token errors**
- Confirm JWT_SECRET is set in .env
- Ensure token is passed in Authorization header with "Bearer " prefix
- Check that token hasn't expired

**Migration failures**
- Verify database_url format is correct
- Ensure you have proper permissions on the database
- Check that migration files exist in the migrations directory

## Future Enhancements

- [ ] Pagination for todo lists
- [ ] Filtering and sorting options
- [ ] Rate limiting
- [ ] Request logging
- [ ] Unit and integration tests
- [ ] Docker containerization
- [ ] API documentation with Swagger

## Contributing

Feel free to submit issues and enhancement requests. When contributing:

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Submit a pull request

## License

This project is open source and available under the MIT License.

## Support

For questions or issues, please open an issue on the repository or contact the development team.
