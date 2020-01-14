BINARY_NAME=anz-test
COV_FILE=coverage

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

fmt:
	go fmt ./...

clean:
	rm -rf coverage* && rm -rf vendor*
	go clean ./...

vendor:
	go mod vendor

test: vendor
	go test -coverprofile=$(COV_FILE).out -covermode=count  ./...
	go tool cover -html=$(COV_FILE).out -o $(COV_FILE).html

build:
	go build -o bin/$(BINARY_NAME) .