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

_start_test:
	docker run --rm -d -p $(DOCKER_PORT):$(APP_PORT) --name $(TEST_CONTAINER) $(APP_NAME):latest

_stop_test:
	docker rm -f $(TEST_CONTAINER)

test:
	$(MAKE) _stop_test
	$(BUILD_CMD)
	$(MAKE) _start_test
	pytest tests
	$(MAKE) _stop_test

_github_test_runner:
	mkdir ../logs
	pytest tests --tb=line > ../logs/pytest_output.txt
