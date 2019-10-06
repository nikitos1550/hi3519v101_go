THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

#FAMILY  =hi3516av100
#FAMILY  =hi3516av200
#FAMILY  =hi3516cv100
FAMILY  =hi3516cv200
#FAMILY  =hi3516cv300


GOPATH := $(THIS_DIR)go

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

guard:
	@echo "USAGE: make FAMILY=hi3516av200 build"
	#@echo "GOPATH=$(GOPATH)"

build: hisi ko
	rm -rf ./distrib/$(FAMILY)
	test -e distrib || mkdir distrib
	test -e distrib/$(FAMILY) || mkdir distrib/$(FAMILY)
	test ! -e ./putonrootfs || cp -r ./putonrootfs/* distrib/$(FAMILY)
	test ! -e ./$(FAMILY)/putonrootfs || cp -r ./$(FAMILY)/putonrootfs/* distrib/$(FAMILY)
	test -e distrib/$(FAMILY)/opt || mkdir distrib/$(FAMILY)/opt
	$(GOPARAMS) GOARM=$(ARMV) $(GO) build $(GOFLAGS) -a -tags "$(FAMILY)" -o distrib/$(FAMILY)/opt/$(APP)

hisi:
	$(AR) rcsT $(FAMILY)/lib$(FAMILY).a $(LIBS)

ko:
	$(GOPATH)/bin/go-bindata -nomemcopy -nocompress -o kobin.go -prefix "$(FAMILY)/ko/" $(FAMILY)/ko

godeps:
	GOPATH=$(GOPATH) $(GO) get -u golang.org/x/sys/unix
	GOPATH=$(GOPATH) $(GO) get -u github.com/shuLhan/go-bindata/...
	GOPATH=$(GOPATH) $(GO) get -u github.com/omeid/go-resources

buildall: clean
	make FAMILY=hi3516av100 build
	make FAMILY=hi3516av200 build
	make FAMILY=hi3516cv100 build
	make FAMILY=hi3516cv200 build
	make FAMILY=hi3516cv300 build
	make clean

clean:
	rm -rf ./distrib/*
	rm -f $(APP)
	rm -f ./*/libhi35*.a
	rm -f ./kobin.go
