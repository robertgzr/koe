go_sources = $(shell go list -f '{{.Dir}}' $(1) all)

KOE_MAIN    := ./cmd/koe
KOE_BIN     := ./koe
KOE_SOURCES := $(call go_sources,$(KOE_MAIN))
KOE_LDFLAGS := \
    -X main.Version=$(shell git rev-parse --short HEAD) \
    -X main.BuildTime=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)

$(KOE_BIN): $(KOE_SOURCES)
	go build -ldflags "$(KOE_LDFLAGS)" -o $(KOE_BIN) $(KOE_MAIN)
