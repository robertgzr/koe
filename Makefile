go_sources = $(shell go list -f '{{.Dir}}' $(1) all)

VERSION    ?= $(shell git describe --match 'v[0-9]*' --dirty='+dev' --always --tags)
REVISION   ?= $(shell git rev-parse HEAD)
BUILD_TIME ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

PKG         := github.com/robertgzr/koe
KOE_MAIN    := ./cmd/koe
KOE_BIN     := ./koe
KOE_SOURCES := $(call go_sources,$(KOE_MAIN))
KOE_LDFLAGS := \
    -X $(PKG)/version.Version=$(VERSION) \
    -X $(PKG)/version.BuildTime=$(BUILD_TIME) \
    -X $(PKG)/version.Revision=$(REVISION) \
    -X $(PKG)/version.Package=$(PKG)

all: $(KOE_BIN)

$(KOE_BIN): $(KOE_SOURCES)
	go build -ldflags "$(KOE_LDFLAGS)" -o $(KOE_BIN) $(KOE_MAIN)

.PHONY: all clean

clean:
	rm -rf $(KOE_BIN)
