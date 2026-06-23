include .env

REBUILD := $(findstring R,$(MAKEFLAGS))

ifeq ($(REBUILD),R)
    BUILD_CMD := $(MAKE) build
else
    BUILD_CMD := @echo "⏭️  Skipping rebuild (use make -R to force rebuild)"
endif

build: 
	docker build -t $(APP_NAME) -f ./cmd/web/Dockerfile .

run:
	@echo "Check service on http://localhost:$(DOCKER_PORT)\n"
	docker run --rm -p $(DOCKER_PORT):$(APP_PORT) $(APP_NAME):latest

dev: build run

test:
	mkdir -p ../logs
	go test ./tests > ../logs/tests.log
