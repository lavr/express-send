VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)

.PHONY: build
build:
	go build -ldflags="-s -w -X main.version=$(VERSION)" -o express-botx .
