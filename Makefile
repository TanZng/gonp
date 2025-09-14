# Variables
BINARY_NAME := gonp.out
PACKAGE := ./...

.PHONY: all test run build clean

all: test build

# Run all tests recursively with verbose output
test:
	go test -v $(PACKAGE)

# Run the app
run:
	go run main.go

# Build the app binary
build:
	go build -o $(BINARY_NAME) main.go

# Clean up binary
clean:
	rm -f $(BINARY_NAME)
