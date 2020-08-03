VERSION := 1.0.0.0
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := go-pangu
GOBASE := $(shell pwd)
GODIST := $(GOBASE)/dist
GOPATH := $(shell echo $(HOME)/go)
GOARCH ?= $(shell go env GOARCH)
GOOS ?= $(shell go env GOOS)
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`
CC=gcc

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -w -s"
GOBUILD=GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build $(LDFLAGS)
# The -w and -s flags reduce binary sizes by excluding unnecessary symbols and debug info
SUBDIRS := c/

all: debug release

debug:
	$(GOBUILD) -tags debug -o $(GODIST)/$(PROJECTNAME)-$(GOARCH)-debug-$(GOOS)

release:
	$(GOBUILD) -tags release -o $(GODIST)/$(PROJECTNAME)-$(GOARCH)-release-$(GOOS)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GODIST=$(GODIST) go get $(get)

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

clean:
	@-$(MAKE) clean -C c/
	@-$(MAKE) go-clean
.PHONY: clean

go-clean:
	@GOPATH=$(GOPATH) GODIST=$(GODIST) go clean
.PHONY: go-clean

test:
	@-ginkgo -skip="Pressure test"
.PHONY: test

ptest:
	@-ginkgo -focus="Pressure test"
