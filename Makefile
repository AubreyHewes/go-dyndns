.PHONY: clean checks test build image e2e fmt

export GO111MODULE=on

SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

APP_NAME := go-dyndns
TAG_NAME := $(shell git tag -l --contains HEAD)
VERSION_TAG := $(shell git describe --tags)
VERSION_SHA := $(shell git rev-parse --short HEAD)
VERSION := $(VERSION_TAG)

DOCKER_IMAGE := aubreyhewes/$(APP_NAME)

APP_NAME_CLI := $(APP_NAME)-cli
MAIN_DIRECTORY_CLI := ./cmd/$(APP_NAME)/
ifeq (${GOOS}, windows)
    BIN_OUTPUT_CLI := $(APP_NAME).exe
else
    BIN_OUTPUT_CLI := $(APP_NAME)
endif

APP_NAME_UI := $(APP_NAME)-ui
MAIN_DIRECTORY_UI := ./ui/
ifeq (${GOOS}, windows)
    BIN_OUTPUT_UI := $(APP_NAME_UI).exe
else
    BIN_OUTPUT_UI := $(APP_NAME_UI)
endif

default: clean generate-dns checks test build dist

clean:
	rm -rf dist/ builds/ cover.out

build: clean build-cli build-ui

build-ui: clean
	@echo Name: $(APP_NAME_UI)
	@echo Version: $(VERSION)
	@echo GOROOT: $(GOROOT)
	@echo GOPATH: $(GOPATH)
	go build -v -ldflags '-s -w -X "main.name=${APP_NAME_UI}" -X "main.version=${VERSION}"' -o builds/${BIN_OUTPUT_UI} ${MAIN_DIRECTORY_UI}

build-cli: clean
	@echo Name: $(APP_NAME_CLI)
	@echo Version: $(VERSION)
	@echo GOROOT: $(GOROOT)
	@echo GOPATH: $(GOPATH)
	go build -v -ldflags '-s -w -X "main.name=${APP_NAME}" -X "main.version=${VERSION}"' -o builds/${BIN_OUTPUT_CLI} ${MAIN_DIRECTORY_CLI}

dist: clean dist-cli dist-ui image

UPX_FLAGS := --color --ultra-brute
dist-cli: clean build-cli
	mkdir dist
	upx $(UPX_FLAGS) builds/${BIN_OUTPUT_CLI} -o dist/${BIN_OUTPUT_CLI}

dist-ui: clean build-ui
	mkdir dist
	upx $(UPX_FLAGS) builds/${BIN_OUTPUT_UI} -o dist/${BIN_OUTPUT_UI}

image:
	@echo Name: $(DOCKER_IMAGE)
	@echo Version: $(VERSION)
	docker build -t $(DOCKER_IMAGE) .

test: clean
	go test -v -cover ./...

e2e: clean
	E2E_TESTS=local go test -count=1 -v ./e2e/...

checks:
	golangci-lint run

fmt:
	gofmt -s -l -w $(SRCS)

# Release helper
.PHONY: patch minor major detach

patch:
	go run internal/release.go release -m patch

minor:
	go run internal/release.go release -m minor

major:
	go run internal/release.go release -m major

detach:
	go run internal/release.go detach

# Docs
.PHONY: docs-build docs-serve docs-themes

docs-build: generate-dns
	@make -C ./docs hugo-build

docs-serve: generate-dns
	@make -C ./docs hugo

docs-themes:
	@make -C ./docs hugo-themes

# DNS Documentation
.PHONY: generate-dns validate-doc

generate-dns:
	go generate ./...

validate-doc: generate-dns
ifneq ($(shell git status --porcelain -- ./docs/ ./cmd/ 2>/dev/null),)
	@echo 'The documentation must be regenerated, please use `make generate-dns`.'
	@git status --porcelain -- ./docs/ ./cmd/ 2>/dev/null
	@exit 2
else
	@echo 'All documentation changes are done the right way.'
endif
