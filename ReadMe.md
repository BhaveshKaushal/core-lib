# base-lib

Base Library for all the golang related projects.


## Packages

### Logger

A structured logging framework built on top of logrus. For detailed documentation, see the [logger package README](pkg/logger/README.md).

## Contributing

Feel free to submit issues and enhancement requests.

# Core-lib Makefile

This project uses shared Makefile targets from core-config and can include local targets.

## Location
`core-lib/makefile`

## Structure
```makefile
CONFIG_PATH ?= ../core-config
ifneq (,$(wildcard $(CONFIG_PATH)/make/common.mk))
include $(CONFIG_PATH)/make/common.mk
endif

# Local targets can be added here
```

## Usage
1. Common targets from core-config/make/common.mk are available
2. Local targets can be added directly in this Makefile
3. Run `make help` to see all available targets

## Target Categories

### Common Targets (from common.mk)
- Development targets (build, test, etc.)
- Service connections (redis, postgres, etc.)
- Information commands (help, info)

### Local Targets
Add your project-specific targets in the local Makefile:
```makefile
local-build:
    @echo "Building core-lib locally"
    @go build ./...
```

## Setup
1. Ensure core-config is available at the correct path
2. Run `make help` to verify available targets
3. Use common targets or add local ones as needed

## Notes
- Infrastructure setup targets are in core-config/local/makefile
- Service deployment targets are in core-config/local/makefile
- Common development targets are in core-config/make/common.mk

