FROM golang:latest

ENV DefaultPort 8080

EXPOSE 8080:8080

LABEL maintainer="<contact@microup.ru>"

WORKDIR /stack-service

COPY . .

RUN go mod download
RUN go build -o stack-service cmd/stack-service/stack-service.go
RUN find . -name "*.go" -type f -delete

CMD ["./stack-service"]
