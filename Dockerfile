FROM golang:1.24-alpine

WORKDIR /backend

COPY ./backend/go.mod ./
RUN go mod download

WORKDIR /backend/cmd/api-server