GO_BIN ?= go

export PATH := $(PATH):/usr/local/go/bin

all: lint test

update:
	$(GO_BIN) get -u
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0
	$(GO_BIN) get -u github.com/mgechev/revive
	$(GO_BIN) mod tidy

test:
	$(GO_BIN) test -failfast ./...

lint:
	golangci-lint run ./...
	revive -config revive.toml ./...
