THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

#BUILDROOT := buildroot-2019.05.1-debug
#TOOLCHAIN := arm-buildroot-linux-uclibcgnueabi
#STRIP := $(THIS_DIR)../buildroot-2019.05.1-debug/output/host/bin/$(TOOLCHAIN)-strip

GO := $(THIS_DIR)../buildroot-2019.05.1-debug/output/host/bin/go
GOPATH := $(THIS_DIR)go

NOW=$(shell date)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

VERSIONFILE=$(THIS_DIR)internal/info/info.go

CFLAGS := -std=gnu11 -mcpu=cortex-a7 -mfloat-abi=softfp -mfpu=neon-vfpv4 -mno-unaligned-access -fno-aggressive-loop-optimizations

#DEBUG
#BUILDFLAGS := -ldflags="-extldflags '-static'"

#NODEBUG
BUILDFLAGS := -ldflags="-s -w -extldflags '-static'"

build: clean genversion
	cd internal/himpp3; make clean; make
	GOPATH=$(GOPATH) \
    CGO_CFLAGS='$(CFLAGS)' \
    CGO_CFLAGS_ALLOW='.*' \
    GOOS=linux \
    GOARCH=arm \
    GOARM=7 \
    CGO_ENABLED=1 \
    $(GO) build \
        -a $(BUILDFLAGS) \
        app_minimal.go

#strip:
#	$(STRIP) ./app_minimal

genversion:
	rm -f $(VERSIONFILE)
	@echo "package info" > $(VERSIONFILE)
	@echo "const (" >> $(VERSIONFILE)
	@echo "  DATE = \"$(NOW)\"" >> $(VERSIONFILE)
	@echo "  BRANCH = \"$(BRANCH)\"" >> $(VERSIONFILE)
	@echo "  USER = \"$(USER)\"" >> $(VERSIONFILE)
	@echo ")" >> $(VERSIONFILE)

goenv:
	GOPATH=$(GOPATH) $(GO) env

godeps:
	GOPATH=$(GOPATH) $(GO) get github.com/gorilla/mux

clean:
	cd internal/himpp3; make clean
	rm -f ./app_minimal
