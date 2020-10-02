# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:latest
LABEL maintainer="Jesús López <jeslopcru@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build cmd/checkout-api/main.go
RUN go test -c cmd/checkout-api/main_test.go cmd/checkout-api/main.go
RUN go test -c functional_test.go
EXPOSE 8080
CMD ["./main"]