MAIN_PACKAGE := ./cmd/
BINARY_NAME := thunder-api

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## build: build the production code
.PHONY: build
build:
	@echo "Building binary..."
	@go build -o bin/$(BINARY_NAME) $(MAIN_PACKAGE)

## run: run the production code
.PHONY: run
run:
	@echo "Running binary..."
	@go run $(MAIN_PACKAGE)

## dev: run the code development environment
.PHONY: dev
dev:
	@echo "Running development environment..."
	@go run $(MAIN_PACKAGE)


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	@echo "Tidying up..."
	@go fmt ./...
	@go mod tidy -v

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'