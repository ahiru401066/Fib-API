FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

WORKDIR /app/cmd/api-server