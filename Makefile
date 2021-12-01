NAME:=bellman
DOCKER_REPOSITORY:=ghcr.io/sgettys
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')

build:
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/sgettys/bellman/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/bellman ./cmd/bellman/*

build-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .