.PHONY: observer

observer:
	$(MAKE) -C observer $(filter-out $@,$(MAKECMDGOALS))

%:
	@: