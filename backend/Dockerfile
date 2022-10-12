# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

# Add bash to execute wait-for-it.sh
RUN apk update && apk add bash

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./
RUN go build -o /docker-gs-ping ./cmd/docker/main.go

EXPOSE 8080

RUN chmod +x ./wait-for-it.sh ./docker-entrypoint.sh
ENTRYPOINT [ "./docker-entrypoint.sh" ]
CMD ["/docker-gs-ping"]
