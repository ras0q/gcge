SHELL   := /bin/bash

build:
	@go build -v ./...

install:
	@go mod tidy
	@go install
i: install

.PHONY: run
run:
	@go run github.com/cosmtrek/air@latest

.PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

add-cmd:
	@read -p "Command Name > " CMD_NAME && \
	go run github.com/spf13/cobra/cobra@latest add $$CMD_NAME --config ./.cobra.yml

gen:
	@go generate ./...
