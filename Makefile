PROJECT_NAME="stack-service"

.PHONY: all dep build test lint 

all: lint test race build

init: 
	go mod tidy
	go mod vendor

fmt:
	go fmt ./... 

lint: 
	go vet ./...
	golangci-lint run -v ./...

test: 
	go test -v ./...

race: dep 
	go test -race -v ./...

build: 
	go build -o build/$(PROJECT_NAME) cmd/$(PROJECT_NAME)/$(PROJECT_NAME).go

docker-create:
	docker build --tag $(PROJECT_NAME) . -f Dockerfile

docker-run:
	docker run --restart=always -p 8080:8080 $(PROJECT_NAME) 

