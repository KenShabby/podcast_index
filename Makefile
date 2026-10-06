.DEFAULT_GOAL := build

.PHONY: fmt vet build clean

# Format code
fmt:
	go fmt ./...

# Vet code
vet: fmt
	go vet ./...

# Build the binary
build: vet
	go build -o pod .

# Clean build artifacts
clean:
	rm -f pod
