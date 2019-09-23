NAME := koala-pos
DESC := A web-based point of sale system
VERSION ?= $(shell git describe --tags --always --dirty)
GOVERSION ?= $(shell go version)
BUILDTIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDER ?= $(shell echo "`git config user.name` <`git config user.email`>")
GOBIN := $(PWD)/bin
GOPATH ?= $(HOME)/go
export CGO_ENABLED ?= 0

ifeq ($(shell uname -s), Cygwin)
PWD := $(shell cygpath -w -a `pwd`)
GOBIN := $(PWD)\bin
endif

LDFLAGS := -X 'main.version=$(VERSION)' \
			-X 'main.buildTime=$(BUILDTIME)' \
			-X 'main.builder=$(BUILDER)' \
			-X 'main.goversion=$(GOVERSION)'

.PHONY: all fmt test benchmark lint build clean

all: test build

build:
	go build -o bin/koala-pos -v -ldflags "$(LDFLAGS)" -tags '$(BUILDTAGS)' ./cmd/koala-pos

build-in-docker:
	docker run --rm \
		-e GOPATH=/go \
		-e BUILDTIME \
		-e BUILDER \
		-v "$(PWD)/..":/usr/src/koala-pos \
		-v $(GOPATH)/pkg/mod:/go/pkg/mod \
		-w /usr/src/koala-pos/backend \
		golang:1.13 make build

# development tasks
fmt:
	@gofmt -s -l -d ./src/*

test: lint
	@go test ./src/...

benchmark:
	@go test -bench=. $$(go list ./src/...)

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	@golint ./src/...

clean:
	rm -rf ./bin/*
	rm -rf ./logs/*
	rm -rf ./sessions/*
