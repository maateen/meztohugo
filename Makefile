# Go parameters
GOCMD=$(shell which go)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=meztohugo
GOBIN=bin
GOOS=darwin linux windows
GOARCH=amd64 386
VERSION=v0.1-alpha

all: prepare build
prepare:
		$(GOCMD) mod tidy
		$(GOCMD) mod verify
build:
		mkdir -p $(GOBIN); \
		for goos in $(GOOS); do \
			for goarch in $(GOARCH); do \
				GOOS=$$goos GOARCH=$$goarch $(GOBUILD) cmd/meztohugo/main.go; \
				if [ -f main.exe ]; then \
					mv main.exe $(GOBIN)/$(BINARY_NAME)-$$goos-$$goarch-$(VERSION).exe; \
				else \
					mv main $(GOBIN)/$(BINARY_NAME)-$$goos-$$goarch-$(VERSION); \
				fi \
			done \
		done
clean: 
		$(GOCLEAN) cmd/meztohugo/main.go
		rm -rf $(GOBIN);