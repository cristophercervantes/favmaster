
## Makefile
```makefile
.PHONY: build install clean test

BINARY_NAME=favmaster

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) cmd/favmaster/main.go

install:
	@echo "Installing $(BINARY_NAME)..."
	go install ./cmd/favmaster

clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	go clean

test:
	@echo "Running tests..."
	go test ./...

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

all: deps build

# Cross-compilation for different platforms
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux cmd/favmaster/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows.exe cmd/favmaster/main.go

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin cmd/favmaster/main.go

# Build for all platforms
build-all: build-linux build-windows build-darwin
