# All the targets have been tested on ubuntu system
#-------------------------------------------------------------------------------------------
# ENV VARIABLES
#-------------------------------------------------------------------------------------------

export APP ?= base-lib

#-------------------------------------------------------------------------------------------
# APP VARIABLES 
#-------------------------------------------------------------------------------------------
.DEFAULT_GOAL = help

# Find repository root (multiple methods)
GIT_ROOT := $(shell git rev-parse --show-toplevel 2>/dev/null)
GO_MOD_ROOT := $(shell cd .. && while [ ! -f go.mod ] && [ "$$PWD" != "/" ]; do cd ..; done && pwd)
REPO_ROOT := $(if $(GIT_ROOT),$(GIT_ROOT),$(if $(GO_MOD_ROOT),$(GO_MOD_ROOT),$(shell cd ../.. && pwd)))
PARENT_DIR := $(shell dirname $(REPO_ROOT))
# Set CONFIG_PATH relative to the repository root
CONFIG_PATH ?= $(PARENT_DIR)/core-config

#-------------------------------------------------------------------------------------------
# USE COMMON FILE TARGETS
#-------------------------------------------------------------------------------------------

ifneq (,$(wildcard $(CONFIG_PATH)/make/common.mk))
%: force
	@make -s -f $(CONFIG_PATH)/make/common.mk $@
force: ;
endif	

# Print repository root (for debugging)
print-repo-root:
	@echo "Repository root: $(REPO_ROOT)"
	@echo "CONFIG_PATH: $(CONFIG_PATH)"
	@echo "PARENT_DIR: $(PARENT_DIR)"
