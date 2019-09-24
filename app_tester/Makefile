THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GO := /buildroot/host/bin/go
APP := ./app_tester

GOFLAGS := -ldflags="-s -w -extldflags '-static -Wl,--gc-sections'"

GOPARAMS := GOPRIVATE=openhisiipcam.org \
            GOPATH=$(GOPATH) \
            CGO_CFLAGS='$(CFLAGS)' \
            CGO_CFLAGS_ALLOW='.*' \
            GOOS=linux \
            GOARCH=arm \
            CGO_ENABLED=1

all:
    #guard

build-hi3516cv100:
	$(GOPARAMS) GOARM=5 $(THIS_DIR)../hi3516cv100/$(GO) build $(GOFLAGS) -a -tags "hi3516cv100" -o $(APP)

build-hi3516cv200:
	$(GOPARAMS) GOARM=5 $(THIS_DIR)../hi3516cv200/$(GO) build $(GOFLAGS) -a -tags "hi3516cv200" -o $(APP)

build-hi3516cv300:
	$(GOPARAMS) GOARM=5 $(THIS_DIR)../hi3516cv300/$(GO) build $(GOFLAGS) -a -tags "hi3516cv300" -o $(APP)

build-hi3516av100:
	$(GOPARAMS) GOARM=7 $(THIS_DIR)../hi3516av100/$(GO) build $(GOFLAGS) -a -tags "hi3516av100" -o $(APP)

build-hi3516av200:
	$(GOPARAMS) GOARM=7 $(THIS_DIR)../hi3516av200/$(GO) build $(GOFLAGS) -a -tags "hi3516av200" -o $(APP)

clean:
	rm -f ./app_tester
