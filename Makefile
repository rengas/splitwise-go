.PHONY: run migrate-up migrate-down migrate-new install-migrate docker-up docker-down

# Get database URL from .env file
DATABASE_URL := $(shell grep DATABASE_URL .env | cut -d '=' -f2-)

run:
	@echo "Starting server..."
	go run main.go

install-migrate:
	@echo "Installing golang-migrate..."
	go install -tags 'postgres' -v github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Docker commands
docker-up:
	@echo "Starting docker containers..."
	docker compose up -d

docker-down:
	@echo "Stopping docker containers..."
	docker compose down

# Migration commands
migrate-new:
	@echo "Creating new migration: $(name)"
	@migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	@echo "Running 'up' migrations..."
	@migrate -path migrations -database "$(DATABASE_URL)" -verbose up

migrate-down:
	@echo "Running 'down' migrations..."
	@migrate -path migrations -database "$(DATABASE_URL)" -verbose down 1

