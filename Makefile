
.PHONY: all
all: # Build the project
	@echo "=================== Running false-negative ==============="
	$(MAKE) -C false-negative
	@echo "=================== Running false-positive ==============="
	$(MAKE) -C false-positive

# default goal
.DEFAULT_GOAL := all
