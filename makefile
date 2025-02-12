NAME = proxy-core
COMMIT = $(shell git rev-parse --short HEAD)

TAGS_GO120 = with_gvisor,with_dhcp,with_wireguard,with_reality_server,with_clash_api,with_quic,with_utls
TAGS_GO121 = with_ech
TAGS ?= $(TAGS_GO118),$(TAGS_GO120),$(TAGS_GO121)
TAGS_TEST ?= with_gvisor,with_quic,with_wireguard,with_grpc,with_ech,with_utls,with_reality_server

GOHOSTOS = $(shell go env GOHOSTOS)
GOHOSTARCH = $(shell go env GOHOSTARCH)
VERSION=$(shell CGO_ENABLED=0 GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) go run ./cmd/internal/read_tag)

PARAMS = -v -trimpath -ldflags "-X 'proxy-core/constant.Version=$(VERSION)' -s -w -buildid="
MAIN_PARAMS = $(PARAMS) -tags $(TAGS)
MAIN = ./cmd/proxy-core

ifeq ($(GOHOSTOS),windows)
    EXT = .exe
else
    EXT =
endif

OUTPUT_DIR ?= bin
OUTPUT_FILE = $(OUTPUT_DIR)/$(NAME)$(EXT)

.PHONY: test release docs build clean

build:
	mkdir $(OUTPUT_DIR)
	go build -o $(OUTPUT_FILE) $(MAIN_PARAMS) $(MAIN)

clean:
	rm -rf $(OUTPUT_DIR)
