# Makefile for real-time-ranking project

APP_NAME := main
ENTRY_POINT := cmd/api/main.go
OAPI_CFG_DIR := docs/cfg.yaml
OAPI_DIR := docs/openapi.yaml

.PHONY: all oapi test build run

# Default: generate handler, build and test
all: oapi test build run

# ----------- OpenAPI -----------

oapi:
	@echo "Generating code from OpenAPI..."
	@oapi-codegen -config=$(OAPI_CFG_DIR) $(OAPI_DIR)

# ----------- Tests -----------

test:
	@echo "Running tests..."
	@go test ./... -v

# ----------- Build & Run -----------

build:
	@echo "Building..."
	@go build -o $(APP_NAME) $(ENTRY_POINT)

run:
	@echo "Running..."
	@./main
