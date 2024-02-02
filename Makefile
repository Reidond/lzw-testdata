GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=lzw-testdata
BINARY_DIR=bin
CMD_DIR=cmd

all: test build

build:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v $(CMD_DIR)/*

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_DIR)/$(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v $(CMD_DIR)/*
	./$(BINARY_DIR)/$(BINARY_NAME)

deps:
	$(GOGET) -v ./...

.PHONY: all build test clean run deps
