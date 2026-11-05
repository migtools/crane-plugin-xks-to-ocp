.PHONY: build test

build:
	go build -o crane-plugin-xks-to-ocp .

test:
	go test ./...
