FROM golang:1.19.2 AS builder

WORKDIR /app

COPY build ./
COPY templates ./templates

ENTRYPOINT ["./frontend"]
