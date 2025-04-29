FROM golang:1.24-alpine

WORKDIR /app

COPY ./go.mod ./
RUN go mod download

COPY ./ ./

WORKDIR /app/cmd/api-server
EXPOSE 8080
CMD [ "go","run","main.go" ]