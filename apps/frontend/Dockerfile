FROM golang:1.19.2

WORKDIR /app

COPY build ./
COPY src/templates ./templates

ENTRYPOINT ["./frontend"]
