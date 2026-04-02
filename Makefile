GOOSE_DRIVE ?= mysql
GOOSE_DBSTRING ?= "root:root1234@tcp(127.0.0.1:3306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema

# name app
APP_NAME = server

docker_build:
	docker compose up -d --build
	docker compose ps

docker_stop:
	docker compose down

docker_logs:
	docker compose logs -f

docker_up:
	docker compose up -d

dev:
	go run ./cmd/$(APP_NAME)/main.go

upse:
	@GOOSE_DRIVE=$(GOOSE_DRIVE) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVE=$(GOOSE_DRIVE) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVE=$(GOOSE_DRIVE) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) reset

sqlgen:
	sqlc generate


.PHONY: dev downse upse resetse docker_build docker_stop docker_up

.PHONY: air