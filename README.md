# Password Manager (API)

## Technologies and Libraries

- [Go](https://go.dev/)
- [Fiber](https://docs.gofiber.io/)
- [Docker](https://docker.com/)
- [Swagger](https://swagger.io/)
- [Google UUID](https://github.com/google/uuid/)
- [Validator](https://github.com/go-playground/validator/)
- [Godotenv](https://github.com/joho/godotenv/)

## Live application (Deployed in AWS)

- [Swagger Docs](https://passmanager-api.danielmesquitta.com/docs/index.html)

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

## Running in production

### Requirements

- [git](https://git-scm.com/)
- [Docker](https://docker.com/)

### Start API

```shell
# Clone the project
$ git clone https://github.com/danielmesquitta/password-manager-api

# Access the folder
cd password-manager-api

# Create .env file
$ cp .env.example .env

# Build and run docker container
$ docker compose up -d
```
