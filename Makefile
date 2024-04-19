SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all build

BINARY_NAME=tgtherapist
IMAGE_TAG=$(shell git describe --tags --always)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
ORG_PREFIX := loqutus

get:
	go get -v ./...
tidy:
	go mod tidy

build:
	GOARCH=arm64 GOOS=darwin go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-deploy-darwin-arm64 cmd/tgtherapist/main.go
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-deploy-linux-amd64 cmd/tgtherapist/main.go
	GOARCH=arm64 GOOS=linux go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-client-darwin-arm64 cmd/tgtherapist/main.go
	chmod +x bin/*

docker:
	docker build -t ${ORG_PREFIX}/${BINARY_NAME}:${IMAGE_TAG} -f build/Dockerfile .