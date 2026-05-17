.PHONY: build run migrate-up migrate-down sqlc-gen lint test tidy docker-up docker-down

build:
	go build -o bin/app/goku ./cmd/...

run:
	go run ./cmd/...

migrateup:
	migrate -path platform/database/migrations -database "postgres://postgres:postgres@localhost:5448/goku?sslmode=disable" -verbose up

migrate-down:
	go run ./cmd/... migrate down

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
