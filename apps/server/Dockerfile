# Dockerfile for Go server

# Stage 1: Build the Go project
FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./src

# Stage 2: Run the Go project
FROM debian:buster

WORKDIR /app

COPY --from=builder /app/server .
COPY ./migrations /app/migrations

# Install migration tool
RUN apt-get update && apt-get install -y curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin

# Run migrations
CMD ["migrate", "-path", "/app/migrations", "-database", "postgres://youruser:yourpassword@postgres:5432/yourdb?sslmode=disable", "up"]

# Expose the port the server will run on
EXPOSE 8080

# Run the Go binary
CMD ["./server"]
