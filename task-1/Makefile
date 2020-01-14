APP_NAME=anz-test
COV_FILE=coverage

fmt:
	go fmt ./...

deps:
	GOPROXY=https://proxy.golang.org go mod download

test:
	go test -coverprofile=$(COV_FILE).out -covermode=count  ./...
	go tool cover -html=$(COV_FILE).out -o $(COV_FILE).html
	open $(COV_FILE).html

gosec:
	gosec .

run:
	docker run -p 8080:8080 -t $(APP_NAME)

build:
	./build.sh

build-run: build run

local-ci: fmt deps test gosec build run

remote-build:
	gcloud builds submit .. --config=../cloudbuild.yaml