include .env

CONN_STR=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

run:
	go run cmd/main.go

migrate-up:
	migrate -path migrations -database "${CONN_STR}" up

migrate-down:
	migrate -path migrations -database "${CONN_STR}" down

.PHONY: run migrate-up migrate-down