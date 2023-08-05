# Password Manager (API)

## Technologies and Libraries

- [Go](https://go.dev/)
- [Fiber](https://docs.gofiber.io/)
- [Docker](https://docker.com/)
- [Swagger](https://swagger.io/)
- [Google UUID](https://github.com/google/uuid/)
- [Validator](https://github.com/go-playground/validator/)
- [Godotenv](https://github.com/joho/godotenv/)

## Running locally

### Requirements

- [Go](https://go.dev/)
- [git](https://git-scm.com/)
- [make](https://github.com/wkusnierczyk/make/)
- [Swag](https://github.com/swaggo/swag/)
- [Air](https://github.com/cosmtrek/air/)

### Start API

```shell
# Clone the project
$ git clone https://github.com/danielmesquitta/password-manager-api

# Access the folder
cd password-manager-api

# Install dependencies
$ go get ./cmd/server

# Create .env file
$ cp .env.example .env

# Start
$ make dev
```
