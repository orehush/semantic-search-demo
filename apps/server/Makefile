# Makefile for Go server

# Variables
DB_URL=postgres://youruser:yourpassword@localhost:5432/yourdb?sslmode=disable
MIGRATE_CMD=migrate -path ./migrations -database "$(DB_URL)"

# Run the server locally
run:
	go run ./main.go

# Run unit tests locally
test:
	go test ./src/...

# Run migrations
migrate:
	$(MIGRATE_CMD) up

# Build and run the server in Docker
docker_run:
	docker build -t semantic-search-server ./ && docker run --rm -e DATABASE_URL=$(DB_URL) -p 8080:8080 semantic-search-server

# Run tests in Docker container
docker_test:
	docker build -t semantic-search-server ./ && docker run --rm -e DATABASE_URL=$(DB_URL) semantic-search-server make test
