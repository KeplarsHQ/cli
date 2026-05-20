.PHONY: build install clean test run help

BINARY_NAME=keplars
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME).exe
VERSION=$(shell cat ../VERSION)

help:
	@echo "Available targets:"
	@echo "  build          Build the CLI binary"
	@echo "  install        Install the CLI globally"
	@echo "  clean          Remove build artifacts"
	@echo "  test           Run tests"
	@echo "  run            Run the CLI (use ARGS to pass arguments)"
	@echo "  build-all      Build for all platforms"
	@echo "  help           Show this help message"

build:
	@echo "Building $(BINARY_NAME)..."
	go build -ldflags "-X github.com/Swing-Technologies/keplars-email-cli/cmd.Version=$(VERSION)" -o $(BINARY_NAME) .

install:
	@echo "Installing $(BINARY_NAME)..."
	go install -ldflags "-X github.com/Swing-Technologies/keplars-email-cli/cmd.Version=$(VERSION)" .

clean:
	@echo "Cleaning..."
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_WINDOWS)

test:
	@echo "Running tests..."
	go test -v ./...

run:
	@go run . $(ARGS)

build-all: build-linux build-darwin build-windows
	@echo "Built all platform binaries"

build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/Swing-Technologies/keplars-email-cli/cmd.Version=$(VERSION)" -o $(BINARY_UNIX) .

build-darwin:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X github.com/Swing-Technologies/keplars-email-cli/cmd.Version=$(VERSION)" -o $(BINARY_NAME)_darwin .

build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -ldflags "-X github.com/Swing-Technologies/keplars-email-cli/cmd.Version=$(VERSION)" -o $(BINARY_WINDOWS) .
