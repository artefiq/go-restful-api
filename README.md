# Go RESTful API with Gin and MongoDB

## Overview
This project is a RESTful API built using the [Gin](https://github.com/gin-gonic/gin) framework in Go, with MongoDB as the database. The API includes authentication, user management, and protected routes using JWT-based authentication.

## Features
- User registration and authentication
- JWT-based authentication middleware
- CRUD operations for user profiles
- Swagger API documentation

## Technologies Used
- **Go** (Golang)
- **Gin** - Web framework
- **MongoDB** - NoSQL database
- **JWT** - Authentication
- **Swagger** - API documentation

## Installation

### Prerequisites
- Go installed on your system ([Download](https://go.dev/dl/))
- MongoDB running locally or in a cloud environment
- Git installed

### Clone the Repository
```sh
git clone https://github.com/artefiq/go-restful-api.git
cd go-restful-api
```

### Install Dependencies
```sh
go mod tidy
```

### Set Up Environment Variables
Create a `.env` file in the project root and configure the following:
```
MONGO_URI=mongodb://localhost:27017
JWT_SECRET=your_secret_key
```

### Run the Server
```sh
go run main.go
```
The server will start on `http://localhost:8080`

## API Endpoints
### Authentication
- **POST** `/api/v1/users/register` - Register a new user
- **POST** `/api/v1/users/login` - Login and receive a JWT token

### User Management (Protected)
- **GET** `/api/v1/users` - Get all users
- **GET** `/api/v1/users/:id` - Get user by ID
- **PUT** `/api/v1/users/:id` - Update user details
- **DELETE** `/api/v1/users/:id` - Delete user

## Swagger Documentation
Swagger UI is available at:
```
http://localhost:8080/api/v1/swagger/index.html
```

## Contributing
Feel free to fork the repository, submit issues, or contribute improvements through pull requests.

## License
This project is licensed under the MIT License.

