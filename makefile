.PHONY: default dev build clean docs lint tidy test

default: dev

dev:
	@make build && air
build:
	@go build -o ./tmp/server ./cmd/server/main.go
clean:
	@rm -rf ./internal/http/docs && find ./tmp -type f ! -name '.gitkeep' -exec rm -f {} +
docs:
	@swag init -g ./cmd/server/main.go -o ./internal/http/docs
lint:
	@golangci-lint run
tidy:
	@go mod tidy
test:
	godotenv -f .env.test go test ./...
