SHELL   := /bin/bash

build:
	@go build -o gcg .

install:
	@go mod tidy
	@go install
i: install

.PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

add-cmd:
	@read -p "Command Name > " CMD_NAME && \
	go run github.com/spf13/cobra/cobra@latest add $$CMD_NAME --config ./.cobra.yml

go-gen: build
	@go generate ./...
	@./gcg gen example/struct.go -o example/gcg_gen.go
	@./gcg gen internal/model/file.go -o internal/model/file_cst.go
