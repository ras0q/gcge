SHELL   := /bin/bash

BLUE    := \033[34m
RESET   := \033[0m
USE_MA_MODE := echo -e "$(BLUE)INFO: $(RESET)Running in module-aware mode"

install:
	@go mod tidy
	@go install
i: install

.PHONY: run
run:
	@go run github.com/cosmtrek/air@latest

.PHONY: lint
lint:
	@golangci-lint run --fix \
	|| $(USE_MA_MODE)\
	&& go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

add-cmd:
	@read -p "Command Name > " CMD_NAME && \
	go run github.com/spf13/cobra/cobra@latest add $$CMD_NAME --config ./.cobra.yml
