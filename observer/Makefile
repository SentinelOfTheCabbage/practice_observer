include .env

APP_NAME := myapp

build: 
	docker build -t $(APP_NAME) -f ./cmd/web/Dockerfile .

run:
	@echo "Check service on http://localhost:$(DOCKER_PORT)\n"
	docker run --rm -p $(DOCKER_PORT):$(APP_PORT) $(APP_NAME):latest

dev: build run
