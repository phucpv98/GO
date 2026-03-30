# name app
APP_NAME = server

dev:
	go run ./cmd/$(APP_NAME)/main.go

.PHONY: air