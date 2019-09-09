LDFLAGS := \
    -X main.Version=$(shell git rev-parse --short HEAD) \
    -X main.BuildTime=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)

build:
	go build -ldflags "$(LDFLAGS)" ./cmd/joe
