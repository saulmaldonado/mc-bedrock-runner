SHELL = /bin/bash
BINARY = mc-bedrock-runner
BUILD_FLAGS ?= -v
ARCH ?= amd64

-include .env
export

.PHONY: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build -o build/$(BINARY) $(BUILD_FLAGS) .

clean:
	@rm -rf build
