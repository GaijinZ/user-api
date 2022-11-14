# syntax=docker/dockerfile:1

FROM golang:alpine3.16 AS builder
WORKDIR /userapi
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY src/ ./src
COPY ./*.go ./
RUN go build -o /userapi

FROM alpine:3.16
WORKDIR /
COPY --from=builder /userapi /userapi
EXPOSE 9500
ENTRYPOINT ["/userapi"]