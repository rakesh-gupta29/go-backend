APP_NAME := portal
MAIN_FILE := cmd/server/main.go

# Default target
.PHONY: all
all: run

# Run the Go app
.PHONY: run
run:
	go run $(MAIN_FILE)

# Build the Go app
.PHONY: build
build:
	go build -o $(APP_NAME) $(MAIN_FILE)

# Run tests
.PHONY: test
test:
	go test ./...

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Clean up binary
.PHONY: clean
clean:
	rm -f $(APP_NAME)
