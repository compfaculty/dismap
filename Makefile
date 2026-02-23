# Dismap - Asset discovery and identification tool
# https://github.com/zhzyker/dismap

BINARY     := dismap
MAIN       := ./cmd/dismap
VERSION    ?= 0.3
GO         := go
GOFLAGS    := -v
LDFLAGS    := -s -w
BUILD_FLAGS := -trimpath

# Optimized flags for production: smaller binary, no debug info
OPT_LDFLAGS  := -s -w
OPT_FLAGS    := -trimpath

# Linux Kali build (amd64, static, optimized)
KALI_GOOS   := linux
KALI_GOARCH := amd64
KALI_SUFFIX := -$(KALI_GOOS)-$(KALI_GOARCH)

.PHONY: all build build-kali install clean test help run lint fmt

all: build

## Build: default (native)
build:
	$(GO) build $(GOFLAGS) $(BUILD_FLAGS) -ldflags="$(LDFLAGS)" -o $(BINARY) $(MAIN)

## Build: optimized for Linux Kali (amd64, CGO disabled, stripped)
build-kali:
	CGO_ENABLED=0 GOOS=$(KALI_GOOS) GOARCH=$(KALI_GOARCH) $(GO) build \
		$(GOFLAGS) $(OPT_FLAGS) \
		-ldflags="$(OPT_LDFLAGS)" \
		-o $(BINARY)$(KALI_SUFFIX) $(MAIN)

## Build: release binaries for linux, darwin, windows (amd64)
build-release:
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 $(GO) build $(OPT_FLAGS) -ldflags="$(OPT_LDFLAGS)" -o $(BINARY)-0.3-linux-amd64   $(MAIN)
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 $(GO) build $(OPT_FLAGS) -ldflags="$(OPT_LDFLAGS)" -o $(BINARY)-0.3-darwin-amd64  $(MAIN)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build $(OPT_FLAGS) -ldflags="$(OPT_LDFLAGS)" -o $(BINARY)-0.3-windows-amd64.exe $(MAIN)

## Install binary to $GOPATH/bin or $GOBIN
install: build
	$(GO) install $(BUILD_FLAGS) -ldflags="$(LDFLAGS)" $(MAIN)

## Run dismap
run: build
	./$(BINARY) $(ARGS)

## Run tests
test:
	$(GO) test ./...

## Tidy and verify module
mod:
	$(GO) mod tidy
	$(GO) mod verify

## Format code
fmt:
	$(GO) fmt ./...

## Run linter (golangci-lint if installed)
lint:
	@which golangci-lint >/dev/null 2>&1 && golangci-lint run ./... || $(GO) vet ./...

## Clean build artifacts
clean:
	rm -f $(BINARY) $(BINARY)-*
	$(GO) clean -cache -testcache 2>/dev/null || true

## Show help
help:
	@echo "Dismap Makefile targets:"
	@echo "  make build         - Build native binary"
	@echo "  make build-kali    - Build optimized binary for Linux Kali (amd64)"
	@echo "  make build-release - Build release binaries (linux, darwin, windows)"
	@echo "  make install       - Install to GOPATH/bin"
	@echo "  make run ARGS=...  - Build and run with optional args"
	@echo "  make test          - Run tests"
	@echo "  make mod           - Tidy and verify go.mod"
	@echo "  make fmt           - Format code"
	@echo "  make lint          - Run vet/linter"
	@echo "  make clean         - Remove build artifacts"
