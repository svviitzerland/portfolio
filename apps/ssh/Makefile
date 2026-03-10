.PHONY: build run dev clean docker docker-run

APP_NAME := ssh-portfolio
PORT := 2222

## build: Build the binary
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) .
	@echo "Done! Binary: ./$(APP_NAME)"

## run: Build and run
run: build
	@./$(APP_NAME)

## dev: Run with go run (development)
dev:
	@go run .

## clean: Remove binary and SSH keys
clean:
	@rm -f $(APP_NAME)
	@rm -rf .ssh/
	@echo "Cleaned."

## docker: Build Docker image
docker:
	@docker build -t $(APP_NAME) -f Dockerfile ../..

## docker-run: Run with Docker
docker-run: docker
	@docker run -p $(PORT):$(PORT) $(APP_NAME)

## docker-compose: Run with docker-compose
docker-compose:
	@docker compose up --build

## help: Show this help
help:
	@echo ""
	@echo "  SSH Portfolio - Farhan Aulianda"
	@echo "  ────────────────────────────────"
	@echo ""
	@sed -n 's/^## //p' $(MAKEFILE_LIST) | column -t -s ':' | sed 's/^/  /'
	@echo ""
	@echo "  Connect: ssh -p $(PORT) localhost"
	@echo ""
