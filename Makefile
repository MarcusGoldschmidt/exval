GO ?= go
GOPATH ?= $(shell go env GOPATH)

api: gen-swagger build-api

migration: build-api

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

install-deps:
	go mod download;

gen-swagger:
	swag init -g cmd/api/api.go -o cmd/api/docs

build-api:
	CGO_ENABLED=0 $(GO) build -v ./cmd/api/api.go

build-migration:
	CGO_ENABLED=0 $(GO) build -v ./cmd/migration/migration.go

