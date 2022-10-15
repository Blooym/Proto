.PHONY: build build-all clean

# Variables
NAME=proto
SKIP_COMPRESS?=true

# Build a single binary for the current system.
build:
ifeq (, $(shell which goreleaser))
	@echo "WARNING: You are using an unsupported build system. Please install 'goreleaser' to build this project as intended."
	rm -rf dist
	go build -ldflags="-s -w" -o ./dist/$(NAME)
else
	SKIP_COMPRESS=$(SKIP_COMPRESS) goreleaser build --snapshot --rm-dist --single-target
endif

# Build binaries for all supported systems.
build-all:
ifeq (, $(shell which goreleaser))
	@echo "WARNING: You are using an unsupported build system. Please install 'goreleaser' to build this project as intended."
	rm -rf ./dist

	@echo "Building for linux"
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/$(NAME)_linux_amd64
	GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o ./dist/$(NAME)_linux_arm
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ./dist/$(NAME)_linux_arm64
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./dist/$(NAME)_linux_386
 else 
	SKIP_COMPRESS=$(SKIP_COMPRESS) goreleaser build --snapshot --rm-dist
endif

clean:
	rm -rf ./dist

# Run the application and pass all arguments to it.
run:
	go run main.go $(filter-out $@,$(MAKECMDGOALS))