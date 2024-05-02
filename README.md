# JWT Authentication in Golang using Gin and Gorm

This repository contains working example of implementing JWT(JSON Web Token) Authentication in Golang using Gin as the webframework and GORM as the ORM

## Features
- User Registration
- User Login and JWT token generation
- Authenticated routes using JWT middleware

## Requirements
- Go     1.22.1
- Gin    (https://github.com/gin-gonic/gin)
- Gorm   (https://github.com/go-gorm/gorm)
- JWT-go (https://github.com/dgrijalva/jwt-go)
- PostgreSQL

## Installation
- Clone the repository from the terminal using git clone https://github.com/jos-h/jwt-gin.git
- Navigate to the project directory
- Install dependencies by running **go mod tidy**
- Setup your database and provide the database connections strings in the .env(Create a .env file locally inside the project directory)
- Run the application by running **go run main.go**
