# Makefile with default command shows all commands along with the comments as the description

# Default command
all: help

# Help command
#
help:
	@echo "Usage: make [command]"
	@echo ""
	@echo "Commands:"
	@echo "  help          Show this help message"
	@echo "  build         Build the application"
	@echo "  run           Run the application"
	@echo ""

build:
	@echo "Building the application"
	@go build -o bin/main main.go

run:
	@echo "Running the application"
	@go run main.go
