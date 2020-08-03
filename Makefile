VERSION := 1.0.0.0
# $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := go-pangu
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOPATH := $(shell echo $(HOME)/go)
GOARCH ?= $(shell go env GOARCH)
GOOS ?= $(shell go env GOOS)
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`
CC=gcc

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -w -s"
GOBUILD=go build $(LDFLAGS)
# The -w and -s flags reduce binary sizes by excluding unnecessary symbols and debug info
TOOLCHAIN=${ANDROID_NDK_HOME}/toolchains/llvm/prebuilt/darwin-x86_64/bin
SUBDIRS := c/

all: linux

linux:
	GOARCH=amd64 GOOS=linux $(GOBUILD) --tags linux -o $(GOBIN)/$(PROJECTNAME)-$(VERSION)-$(GOARCH)-$(GOOS)

linux-cli: arm amd64 386

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(get)

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

clean:
	@-$(MAKE) clean -C c/
	@-$(MAKE) go-clean
.PHONY: clean

go-clean:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean
.PHONY: go-clean

test:
	@-ginkgo -skip="Pressure test"
.PHONY: test

ptest:
	@-ginkgo -focus="Pressure test"
