# syntax=docker/dockerfile:1

FROM golang:alpine3.16 AS builder
WORKDIR /elo
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY src/ ./src
COPY ./*.go ./
RUN go build -o /docker-go-alpine

FROM alpine:3.16
WORKDIR /
COPY --from=builder /docker-go-alpine /docker-go-alpine
EXPOSE 9500
ENTRYPOINT ["/docker-go-alpine"]