# syntax=docker/dockerfile:1

## Build
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /command cmd/migration/migration.go

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /command /command

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/command"]