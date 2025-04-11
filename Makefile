# Makefile for real-time-ranking project

APP_NAME := main
ENTRY_POINT := cmd/api/main.go
SWAG_DIR := cmd/api,internal/server,internal/model
SWAG_OUTPUT := docs

.PHONY: all swagger test build run

# Default: generate swagger, build and test
all: swagger test build run

# ----------- Build & Run -----------

build:
	@echo "🔨 Building..."
	@go build -o $(APP_NAME) $(ENTRY_POINT)

run:
	@echo "🚀 Running..."
	@./main

# ----------- Swagger Docs -----------

swagger:
	@echo "📄 Generating Swagger docs..."
	@swag init --dir $(SWAG_DIR) --output $(SWAG_OUTPUT)

# ----------- Tests -----------

test:
	@echo "🧪 Running tests..."
	@go test ./... -v
