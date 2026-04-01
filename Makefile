.PHONY: build clean test install vet lint tools

BINARY_NAME=gorefactor
CMD_PATH=./cmd/gorefactor
BIN_DIR=bin

PROJECT_ROOT=$(shell git rev-parse --show-toplevel)
LOCAL_BIN=$(PROJECT_ROOT)/tmp/bin

export PATH := $(LOCAL_BIN):$(PATH)

# NOTE: please keep this list sorted so that it can be easily searched
TOOLS = golangci-lint

.PHONY: tools
tools:
	./hack/tools.sh

$(TOOLS):
	./hack/tools.sh $@

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(CMD_PATH)

clean:
	rm -rf $(BIN_DIR)

test:
	go test -race -v ./...

vet:
	go vet ./...

install:
	go install $(CMD_PATH)

lint: golangci-lint
	$(LOCAL_BIN)/golangci-lint run ./...

.DEFAULT_GOAL := build
