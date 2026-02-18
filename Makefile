.PHONY: build clean test install

BINARY_NAME=gorefactor
CMD_PATH=./cmd/gorefactor

build:
	go build -o $(BINARY_NAME) $(CMD_PATH)

clean:
	rm -f $(BINARY_NAME)

test:
	go test -v ./...

install:
	go install $(CMD_PATH)

.DEFAULT_GOAL := build
