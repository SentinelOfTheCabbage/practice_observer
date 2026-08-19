include .env

.PHONY: checker
.PHONY: web

REBUILD := $(findstring R,$(MAKEFLAGS))

ifeq ($(REBUILD),R)
    BUILD_CMD := $(MAKE) build
else
    BUILD_CMD := @echo "⏭️  Skipping rebuild (use make -R to force rebuild)"
endif


%:
	@:

checker:
	$(MAKE) -f cmd/checker/Makefile $(filter-out $@,$(MAKECMDGOALS))
web:
	$(MAKE) -f cmd/web/Makefile $(filter-out $@,$(MAKECMDGOALS))

test:
	mkdir -p ../logs
	go test ./tests > ../logs/tests.log
