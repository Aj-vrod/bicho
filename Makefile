# Makefile for Bicho Go Application

# Variables - store values we'll reuse
BINARY_NAME=bicho
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_TEST=$(GO_CMD) test
GO_RUN=$(GO_CMD) run
OUT_DIR ?= $(CURDIR)/out

# .PHONY tells Make these aren't actual files
.PHONY: all build run test clean help fmt vet

# Default target - runs when you type just "make"
all: build

# Build the application
build:
	@mkdir -p ${OUT_DIR}
	@echo "Building $(BINARY_NAME)..."
	$(GO_BUILD) -o ${OUT_DIR}/${BINARY} main.go
	@echo "Build complete!"

# Run the application
run:
	@echo "Running $(BINARY_NAME)..."
	$(GO_RUN) main.go

# Run all tests
test:
	@echo "Running tests..."
	$(GO_TEST) -v ./...

# Clean up built files
clean:
	@echo "Cleaning..."
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)
	@echo "Clean complete!"
