.PHONY: test

build:
	go build -ldflags "-s -w"

test:
	go test -v ./...
