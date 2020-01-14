APP_NAME=anz-test
COV_FILE=coverage

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

fmt:
	go fmt ./...

test:
	go test -coverprofile=$(COV_FILE).out -covermode=count  ./...
	go tool cover -html=$(COV_FILE).out -o $(COV_FILE).html
	open $(COV_FILE).html

run:
	docker run -p 8080:8080 -t $(APP_NAME)

local-build:
	./build.sh

remote-build:
	gcloud builds submit . --config=cloudbuild.yaml