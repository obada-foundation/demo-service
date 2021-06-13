SERVICE = obada/admin-application
COMMIT_BRANCH ?= develop
SERVICE_IMAGE = $(SERVICE):$(COMMIT_BRANCH)
SERVICE_IMAGE_RELEASE_IMAGE = $(SERVICE):master
SERVICE_IMAGE_TAG_IMAGE = $(SERVICE):$(COMMIT_TAG)

SHELL := /bin/bash
.DEFAULT_GOAL := help

run:
	cd ./src && go run ./main.go

test:
	cd ./src && go test ./...

vendor:
	cd ./src && go mod tidy && go mod vendor

build-docker:
	docker build -t $(SERVICE_IMAGE) -f docker/Dockerfile .

deploy:
	@echo "$$DEPLOY_KEY" > id_rsa
	chmod 600 id_rsa
	@echo "$$ANSIBLE_INVENTORY" > hosts
	@echo "tag=$$CI_COMMIT_REF_SLUG" >> hosts

	docker run -t --rm \
		-w /home/ansible/deployment \
		-v $$(pwd)/deployment:/home/ansible/deployment \
		-v $$(pwd)/hosts:/etc/ansible/hosts \
		-v $$(pwd)/id_rsa:/home/ansible/.ssh/id_rsa \
		securityrobot/ansible-alpine:2.9.1 \
		ansible-playbook playbook.yml --limit gateway.obada.io
