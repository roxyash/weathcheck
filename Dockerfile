FROM golang:1.18-alpine AS builder
WORKDIR /weatherchecker
COPY . .
RUN go build -o .bin/main cmd/main.go
ENTRYPOINT [ "/weatherchecker/main" ]