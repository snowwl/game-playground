GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY=playground

all: test build run

build:
	@echo "Building..."
	@mkdir -p ./bin/
	$(GOBUILD) -o ./bin/$(BINARY) cmd/playground/main.go

clean:
	@echo "Cleaning bin/"
	@rm -rf ./bin/*

test:
	$(GOTEST) -v ./...

run:
	@echo "Running playground"
	@./bin/$(BINARY)

