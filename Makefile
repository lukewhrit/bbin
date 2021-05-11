.PHONY: clean

all: build

# Build executable for Sojourner program
build:
	go mod download
	go build --ldflags "-s -w" -o bin/sojourner ./cmd/sojourner/main.go

# Build and execute Sojourner program
start: build
	./bin/sojourner

# Format Sojourner source code with Go toolchain
format:
	go fmt ./...

# Clean up binary output folder
clean:
	rm -rf bin/
