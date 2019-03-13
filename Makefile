# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BUILD_DIR=build/

APPNAME=GOCA
BINARY_NAME=goca
BINARY_UNIX=$(BINARY_NAME)_unix

SRC=$(wildcard goca/*.go)

all: test build

build: $(SRC)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)$(BINARY_NAME) -v $^

test:
	GOCA_TEST_SERVER="https://test.goca.io" $(GOTEST) -v -race -cover -count 1 ./...

test-local:
	GOCA_TEST_SERVER="http://localhost:5000" $(GOTEST) -v -race -cover -count 1 ./...

clean: 
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

deps:
	$(GOGET) ./...

# Cross compilation
build-linux: $(SRC)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v $^

.PHONY: all build test clean run deps build-linux