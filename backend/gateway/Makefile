# Makefile for migration commands

MIGRATOR_PATH=./cmd/migrator/main.go

.PHONY: migrate-up migrate-down migrate-status migrate-version

migrate-up:
	go run $(MIGRATOR_PATH) up

migrate-down:
	go run $(MIGRATOR_PATH) down

migrate-status:
	go run $(MIGRATOR_PATH) status

migrate-version:
	go run $(MIGRATOR_PATH) version

up: migrate-up
down: migrate-down
status: migrate-status
version: migrate-version
