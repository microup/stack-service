# LIFO Stack Service Containerization Example

This project provides a simple example of containerizing a LIFO (Last In, First Out) stack service, which has three basic API calls:

- curl "http://localhost:8080/push?value=foo"
- curl "http://localhost:8080/pop"
- curl "http://localhost:8080/top"

## Building Docker Image

To build a Docker image for this project, use the following command:

```make docker-create```

This command will create a Docker image using the Dockerfile included in the project.

## Running Docker Container

To run a Docker container based on the previously built Docker image, use the following command:

```make docker-run```

This command will run a Docker container that maps the container's port 8080 to the host's port 8080 and restarts the container automatically in case of a failure.

## Configuration

The configuration for this project is stored in the config.yaml file, which sets the default port to 8080.

## Stack Implementation

The stack implementation is provided in the stack package, which includes the following methods:

- New() *Stack: Creates a new stack.
- Push(value string): Pushes a value onto the stack.
- Pop() (string, bool): Pops a value from the stack.
- Top() (string, bool): Returns the top value of the stack without removing it.
- Save() error: Saves the stack data to a JSON file.
- LoadStack() error: Loads the stack data from a JSON file.

The stack data is stored in a file named stack.json.

## Dockerfile

The Dockerfile is based on the latest Golang image and sets the default port to 8080. It copies the project files to the container's working directory, downloads the dependencies, builds the binary, and then deletes the source code files.

```shell
FROM golang:latest

ENV DefaultPort 8080

EXPOSE 8080

LABEL maintainer="<contact@microup.ru>"

WORKDIR /stack-service

COPY . .

RUN go mod download
RUN go build -o stack-service cmd/stack-service/stack-service.go
RUN find . -name "*.go" -type f -delete

CMD ["./stack-service"]
```
