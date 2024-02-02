GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=lzw-testdata
BINARY_DIR=bin
CMD_DIR=cmd
PREFIX=/usr/local
INSTALL_DIR=$(PREFIX)/bin

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

install:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_DIR)/$(BINARY_NAME) $(INSTALL_DIR)

uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

.PHONY: all build test clean run deps install uninstall
