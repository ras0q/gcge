SHELL := /bin/bash

build:
	@go mod tidy
	@go install

.PHONY: run
run:
	@go run github.com/cosmtrek/air@latest

.PHONY: lint
link:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

add-cmd:
	@read -p "Command Name > " CMD_NAME && \
	go run github.com/spf13/cobra/cobra@latest add $$CMD_NAME --config ./.cobra.yml
