include .env

build:
	docker build -t checker -f cmd/checker/Dockerfile .

run:
	docker run --rm -p $(DOCKER_PORT):$(APP_PORT) checker:latest

dev: build run
