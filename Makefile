.PHONY: checker
.PHONY: web

checker:
	$(MAKE) -C observer $(MAKECMDGOALS)

web:
	$(MAKE) -C observer/cmd/web $(filter-out $@,$(MAKECMDGOALS))

format:
	black ./observer/tests

%:
	@:
