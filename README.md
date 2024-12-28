# Subscription Management API

A robust subscription management system built with Go, featuring user authentication, plan management, and subscription tracking.

## 🚀 Features

- **User Management**
  - User registration and authentication
  - JWT-based authentication with access and refresh tokens
  - Password encryption using bcrypt

- **Plan Management**
  - Create and manage subscription plans
  - Flexible plan duration and pricing
  - Plan details and description management

- **Subscription Handling**
  - Subscribe users to plans
  - Track subscription status and expiration
  - Automatic subscription expiration management
  - Subscription statistics and analytics

## 🛠️ Technology Stack

- **Backend Framework**: [Fiber](https://github.com/gofiber/fiber) (Go web framework)
- **Database**: PostgreSQL
- **ORM**: [GORM](https://gorm.io/)
- **Authentication**: JWT (JSON Web Tokens)
- **Configuration**: Environment variables via godotenv
- **API Documentation**: Markdown

## 📋 Prerequisites

- Go 1.23.4 or higher
- PostgreSQL 12 or higher
- Git

## 🔧 Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/chandra-devs/subscription_app.git
   cd subscription_app
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   ```
   Update the `.env` file with your configuration:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database
   ```

4. **Create the database**
   ```sql
   CREATE DATABASE your_database;
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

   The server will start at `http://localhost:3000`

## 📁 Project Structure

```
SUBSCRIPTION_APP/
├── config/              # Configuration files
│   ├── database.go     # Database connection setup
│   └── jwt.go          # JWT configuration
├── controllers/         # Request handlers
│   ├── auth_controller.go
│   ├── plan_controller.go
│   ├── subscription_controller.go
│   └── user_controller.go
├── handlers/           # Business logic
│   └── subscription_handler.go
├── models/             # Database models
│   ├── subscription.go
│   ├── swagger_types.go
│   └── user.go
├── routes/             # Route definitions
│   └── setup.go
└── main.go            # Application entry point
```

## 🔒 Authentication

The API uses JWT for authentication with two types of tokens:
- **Access Token**: Valid for 24 hours
- **Refresh Token**: Valid for 7 days

Include the access token in the Authorization header:
```
Authorization: Bearer <your_access_token>
```

## 📝 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - User login

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create new user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Plans
- `GET /api/v1/plans` - Get all plans
- `GET /api/v1/plans/:id` - Get plan by ID
- `POST /api/v1/plans` - Create new plan

### Subscriptions
- `GET /api/v1/subscriptions` - Get all subscriptions
- `GET /api/v1/subscriptions/user/:userId` - Get user subscriptions
- `POST /api/v1/subscriptions/subscribe` - Subscribe user to plan
- `GET /api/v1/subscriptions/stats` - Get subscription statistics

For detailed API documentation, see [API Documentation](docs/api.md)

## 🧪 Running Tests

```bash
go test ./...
```

## 📈 Database Schema

### Users
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

### Plans
```sql
CREATE TABLE plans (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    duration INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

### Subscriptions
```sql
CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    plan_id INTEGER REFERENCES plans(id),
    status VARCHAR(50) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

## 🔐 Security

- Passwords are hashed using bcrypt
- JWT tokens for authentication
- PostgreSQL SSL mode can be enabled
- Input validation on all endpoints
- CORS configuration available

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- Chandra Devs - [GitHub](https://github.com/chandra-devs)

## 🙏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://gorm.io/)
- [JWT-Go](https://github.com/golang-jwt/jwt)