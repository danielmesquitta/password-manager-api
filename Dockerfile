FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go get ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server ./cmd/server/main.go
RUN touch .env

FROM scratch

WORKDIR /

COPY --from=builder /app/server /server
COPY --from=builder /app/.env /.env

EXPOSE 3000

ENTRYPOINT ["/server"]
