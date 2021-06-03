SERVICE = obada/demo-service
COMMIT_BRANCH ?= develop
SERVICE_IMAGE = $(SERVICE):$(COMMIT_BRANCH)
SERVICE_IMAGE_RELEASE_IMAGE = $(SERVICE):master
SERVICE_IMAGE_TAG_IMAGE = $(SERVICE):$(COMMIT_TAG)

SHELL := /bin/bash
.DEFAULT_GOAL := help

run:
	go run ./src/main.go

vendor:
	cd ./src && go mod tidy && go mod vendor

build-docker:
	docker build -t $(SERVICE_IMAGE) -f docker/Dockerfile .
