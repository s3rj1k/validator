GO_BIN ?= go
CURL_BIN ?= curl
MAKE_BIN ?= make
export PATH := $(PATH):/usr/local/go/bin

all: lint test

update:
	$(GO_BIN) get -u
	$(CURL_BIN) -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0
	$(GO_BIN) get -u github.com/mgechev/revive
	$(GO_BIN) mod tidy

test:
	$(GO_BIN) test -failfast ./...
	$(MAKE_BIN) test -C ./buildin

lint:
	golangci-lint run ./...
	revive -config revive.toml ./...
