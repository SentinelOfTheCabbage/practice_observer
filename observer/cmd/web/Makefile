include ../.env

build:
	docker build -t web -f cmd/web/Dockerfile .

run:
	docker run --rm -p $(DOCKER_PORT):$(APP_PORT) web:latest

dev: build run
