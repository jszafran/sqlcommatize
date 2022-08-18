# credits for this multi-platform Makefile: https://www.codershaven.com/multi-platform-makefile-for-go/

EXECUTABLE=go_sql_commas
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=0.0.1
DIST_DIR=dist
.PHONY: all test clean

all: test build ## Build and run tests

test: ## Run unit tests
	go test ./...

build: windows linux darwin ## Build binaries
	@echo version: $(VERSION)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(DIST_DIR)/$(WINDOWS) ./cmd/cli/main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(DIST_DIR)/$(LINUX) ./cmd/cli/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o $(DIST_DIR)/$(DARWIN) ./cmd/cli/main.go

clean: ## Remove previous build
	rm -f $(DIST_DIR)/$(WINDOWS) $(DIST_DIR)/$(LINUX) $(DIST_DIR)/$(DARWIN)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
