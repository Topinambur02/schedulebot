APP_NAME=bot
BIN_DIR=bin
GO_VERSION=1.24
VERSION=$(shell git describe --tags 2>/dev/null || echo "v0.0.0")
BUILD_TIME=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

all: deps build

build:
	@mkdir -p $(BIN_DIR)
	go build $(LDFLAGS) -o $(BIN_DIR)/$(APP_NAME) ./main.go

run:
	go run $(LDFLAGS) ./main.go

clean:
	rm -rf $(BIN_DIR)

deps:
	go mod download
	go mod verify

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

docker-build:
	docker compose up -d

install-tools:
	@echo "Установка golangci-lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

help:
	@echo "Доступные цели:"
	@echo "  all          - запустить deps и build"
	@echo "  build        - собрать приложение"
	@echo "  run          - запустить приложение"
	@echo "  clean        - удалить собранные файлы"
	@echo "  deps         - загрузить зависимости"
	@echo "  lint         - запустить линтер"
	@echo "  fmt          - отформатировать код"
	@echo "  vet          - запустить go vet"
	@echo "  docker-build - собрать Docker-образ"
	@echo "  install-tools - установить инструменты"

.PHONY: all build run clean deps lint fmt vet docker-build install-tools help