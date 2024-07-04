# Makefile for Go project

# Variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
MAIN_FILE=main.go
BINARY_NAME=bin/main

# Main targets
all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE) && ./$(BINARY_NAME)

.PHONY: all build test clean run
