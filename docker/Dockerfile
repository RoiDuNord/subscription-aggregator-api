FROM golang:1.25-alpine AS builder
RUN apk add --no-cache tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/main.go

FROM alpine:latest
RUN apk add --no-cache tzdata

WORKDIR /docker_app

ENV TZ=Europe/Moscow

COPY --from=builder /app/main /docker_app/main
COPY --from=builder /app/db /docker_app/db

EXPOSE 8080

CMD ["/docker_app/main"]