# Makefile for tonindustries.com - Go (Gin) + TypeScript (React) web app

APP_NAME=tonindustries
CMD_PATH=./cmd/app
BINARY_NAME=$(APP_NAME)
PORT=8080

# Local dev
run:
	go run $(CMD_PATH)

# Build binary
build:
	go build -o $(BINARY_NAME) $(CMD_PATH)

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)

# Run the built binary
start:
	./$(BINARY_NAME)

# Docker commands
docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run -p $(PORT):8080 $(APP_NAME):latest

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down

# Open browser to local dev site (macOS only; comment out for Linux/Windows)
open:
	open http://localhost:$(PORT)

# Help
help:
	@echo "Available targets:"
	@echo "  run             Run the app using go run"
	@echo "  build           Build the Go binary"
	@echo "  clean           Remove built binary"
	@echo "  start           Run the built binary"
	@echo "  docker-build    Build Docker image"
	@echo

