.PHONY: default dev build clean docs lint tidy test

default: dev

dev:
	@air
build:
	@go build -o ./tmp/server ./cmd/server/main.go
clean:
	@rm -rf ./internal/docs && find ./tmp -type f ! -name '.gitkeep' -exec rm -f {} +
docs:
	@swag init -g ./cmd/server/main.go -o ./internal/docs
lint:
	@golangci-lint run
tidy:
	@go mod tidy
test:
	@GO_ENV=test && go test ./...
