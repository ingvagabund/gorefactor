.PHONY: build clean test install vet

BINARY_NAME=gorefactor
CMD_PATH=./cmd/gorefactor
BIN_DIR=bin

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

.DEFAULT_GOAL := build
