THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

FAMILY  =hi3516av200
#CHIP    =hi3519v101

include ../$(FAMILY)/Makefile.params

APP := ./app_tester

GOFLAGS := -ldflags="-s -w -extldflags '-static -Wl,--gc-sections'"

GOPARAMS := GOPRIVATE=openhisiipcam.org \
            GOPATH=$(GOPATH) \
            CGO_CFLAGS='$(CFLAGS)' \
            CGO_CFLAGS_ALLOW='.*' \
            GOOS=linux \
            GOARCH=arm \
            CGO_ENABLED=1

LIBS = $(wildcard $(FAMILY)/lib/*.a)

all:
	@echo "USAGE: make FAMILY=hi3516av200 build"

build: hisi
	test -e distrib || mkdir distrib
	test -e distrib/opt || mkdir distrib/opt
	$(GOPARAMS) GOARM=$(ARMV) $(GO) build $(GOFLAGS) -a -tags "$(FAMILY)" -o distrib/opt/$(APP)

hisi:
	$(AR) rcsT $(FAMILY)/lib$(FAMILY).a $(LIBS)

clean:
	rm -f distrib/$(APP)
	rm -f $(APP)
	rm -f ./*/libhi35*.a
