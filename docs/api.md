# Subscription Management API Documentation

Base URL: `http://localhost:3000/api/v1`

## Authentication

The API uses JWT (JSON Web Token) for authentication. Access tokens are valid for 24 hours, and refresh tokens are valid for 7 days.

### Authentication Endpoints

#### Register New User
```http
POST /auth/register
```

Request Body:
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
}
```

Response (201 Created):
```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 1640995200
}
```

#### Login
```http
POST /auth/login
```

Request Body:
```json
{
    "email": "john@example.com",
    "password": "password123"
}
```

Response (200 OK):
```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 1640995200
}
```

## Users

### User Endpoints

#### Get All Users
```http
GET /users
```

Response (200 OK):
```json
[
    {
        "id": 1,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z",
        "name": "John Doe",
        "email": "john@example.com",
        "subscriptions": [...]
    }
]
```

#### Get User by ID
```http
GET /users/:id
```

Response (200 OK):
```json
{
    "id": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z",
    "name": "John Doe",
    "email": "john@example.com",
    "subscriptions": [...]
}
```

#### Create User
```http
POST /users
```

Request Body:
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
}
```

Response (201 Created):
```json
{
    "id": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z",
    "name": "John Doe",
    "email": "john@example.com"
}
```

#### Update User
```http
PUT /users/:id
```

Request Body:
```json
{
    "name": "John Updated",
    "email": "john.updated@example.com"
}
```

Response (200 OK):
```json
{
    "id": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z",
    "name": "John Updated",
    "email": "john.updated@example.com"
}
```

#### Delete User
```http
DELETE /users/:id
```

Response (200 OK):
```json
{
    "message": "User deleted successfully"
}
```

## Subscription Plans

### Plan Endpoints

#### Get All Plans
```http
GET /plans
```

Response (200 OK):
```json
{
    "success": true,
    "data": [
        {
            "id": 1,
            "created_at": "2024-01-01T00:00:00Z",
            "updated_at": "2024-01-01T00:00:00Z",
            "name": "Premium Plan",
            "description": "Premium features included",
            "price": 29.99,
            "duration": 30
        }
    ]
}
```

#### Get Plan by ID
```http
GET /plans/:id
```

Response (200 OK):
```json
{
    "success": true,
    "data": {
        "id": 1,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z",
        "name": "Premium Plan",
        "description": "Premium features included",
        "price": 29.99,
        "duration": 30
    }
}
```

#### Create Plan
```http
POST /plans
```

Request Body:
```json
{
    "name": "Premium Plan",
    "description": "Premium features included",
    "price": 29.99,
    "duration": 30
}
```

Response (201 Created):
```json
{
    "success": true,
    "data": {
        "id": 1,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z",
        "name": "Premium Plan",
        "description": "Premium features included",
        "price": 29.99,
        "duration": 30
    }
}
```

## Subscriptions

### Subscription Endpoints

#### Get All Subscriptions
```http
GET /subscriptions
```

Response (200 OK):
```json
{
    "status": "success",
    "data": [
        {
            "id": 1,
            "created_at": "2024-01-01T00:00:00Z",
            "updated_at": "2024-01-01T00:00:00Z",
            "user_id": 1,
            "plan_id": 1,
            "status": "active",
            "start_date": "2024-01-01T00:00:00Z",
            "expires_at": "2024-02-01T00:00:00Z",
            "active": true
        }
    ]
}
```

#### Get User's Subscriptions
```http
GET /subscriptions/user/:userId
```

Response (200 OK):
```json
[
    {
        "id": 1,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z",
        "user_id": 1,
        "plan_id": 1,
        "status": "active",
        "start_date": "2024-01-01T00:00:00Z",
        "expires_at": "2024-02-01T00:00:00Z",
        "active": true
    }
]
```

#### Subscribe User to Plan
```http
POST /subscriptions/subscribe
```

Request Body:
```json
{
    "user_id": 1,
    "plan_id": 1
}
```

Response (201 Created):
```json
{
    "success": true,
    "data": {
        "id": 1,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z",
        "user_id": 1,
        "plan_id": 1,
        "status": "active",
        "start_date": "2024-01-01T00:00:00Z",
        "expires_at": "2024-02-01T00:00:00Z",
        "active": true
    }
}
```

#### Get Subscription Statistics
```http
GET /subscriptions/stats
```

Response (200 OK):
```json
{
    "status": "success",
    "data": {
        "total_subscriptions": 100,
        "total_amount": 2999.00,
        "monthly_spending": 2999.00
    }
}
```

## Error Responses

The API uses standard HTTP status codes and returns errors in the following format:

```json
{
    "error": "Error message description"
}
```

Common status codes:
- 400: Bad Request
- 401: Unauthorized
- 404: Not Found
- 409: Conflict
- 500: Internal Server Error

## Rate Limiting

Currently, there are no rate limits implemented on the API endpoints.

## Data Models

### User
```json
{
    "id": "uint",
    "created_at": "timestamp",
    "updated_at": "timestamp",
    "deleted_at": "timestamp",
    "name": "string",
    "email": "string",
    "password": "string (hashed)",
    "subscriptions": "Subscription[]"
}
```

### Plan
```json
{
    "id": "uint",
    "created_at": "timestamp",
    "updated_at": "timestamp",
    "deleted_at": "timestamp",
    "name": "string",
    "description": "string",
    "price": "float64",
    "duration": "int"
}
```

### Subscription
```json
{
    "id": "uint",
    "created_at": "timestamp",
    "updated_at": "timestamp",
    "deleted_at": "timestamp",
    "user_id": "uint",
    "plan_id": "uint",
    "status": "string",
    "start_date": "timestamp",
    "expires_at": "timestamp",
    "active": "boolean"
}
```