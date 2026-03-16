VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)

.PHONY: build
build:
	mkdir -p dist/
	go build -ldflags="-s -w -X main.version=$(VERSION)" -o dist/express-botx .
