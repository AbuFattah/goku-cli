DB_URL = postgres://postgres:postgres@localhost:5448/goku?sslmode=disable
MIGRATIONS_PATH = internal/platform/database/migrations

.PHONY: build run migrateup migratedown sqlc-gen lint test tidy docker-up docker-down

build:
	go build -o bin/app/goku ./cmd/...

buildcmd:
	go build -o /usr/local/bin/goku ./cmd/...
	
run:
	go run ./cmd/...

migrateup:
	migrate -path $(MIGRATIONS_PATH) -database $(DB_URL) -verbose up

migratedown:
	go run ./cmd/... migrate -verbose down

sqlc-gen:
	sqlc generate

lint:
	golangci-lint run

test:
	go test ./... -race -count=1

tidy:
	go mod tidy

docker-up:
	docker compose up -d

docker-down:
	docker compose down
