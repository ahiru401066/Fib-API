FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY ./go.mod ./
RUN go mod download
COPY ./ ./
RUN go build -o /app/main /app/cmd/api-server/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/
EXPOSE 8080
CMD [ "/app/main" ]