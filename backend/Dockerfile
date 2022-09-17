# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# COPY *.go ./
COPY ./ ./

RUN go build -o /docker-gs-ping ./cmd/main.go

EXPOSE 8080

CMD ["/docker-gs-ping"]