PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
PWD := $(shell pwd)

ifndef VERSION
	VERSION := $(shell git describe --tags --abbrev=0)
endif

COMMIT_HASH :=$(shell git rev-parse --short HEAD)
DEV_VERSION := dev-${COMMIT_HASH}

USERID := $(shell id -u $$USER)
GROUPID:= $(shell id -g $$USER)

all: install-generator install generate check

.PHONY: install
install:
	go mod tidy
	go mod vendor

install-generator:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/moov-io/xsd2go/cli/moovio_xsd2go@latest

generate:
	./scripts/generate.sh

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	@echo "Skipping checks on Windows, currently unsupported."
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=60.0 ./lint-project.sh
endif
