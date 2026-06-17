APP_NAME := myapp
DOCKER_PORT := 8000
APP_PORT := 4000

build: 
	docker build -t $(APP_NAME) -f ./observer/cmd/web/Dockerfile ./observer

run:
	echo "Check service on http://localhost:$(DOCKER_PORT)"
	docker run -p $(DOCKER_PORT):$(APP_PORT) $(APP_NAME):latest

dev: build run
