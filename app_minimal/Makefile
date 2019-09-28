THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GOPATH := $(THIS_DIR)go

FAMILY := hi3516av200

include $(THIS_DIR)/../$(FAMILY)/Makefile.params

GOTAGS := "$(FAMILY) netgo openapi debug lua onvif"

NOW=$(shell date)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

VERSIONFILE=$(GOPATH)/src/openhisiipcam.org/buildinfo/buildinfo.go

APP := ./app_minimal

guard:
	@echo "USAGE: make FAMILY=hi3516av200 build"
	@echo "GOPATH $(GOPATH)"

build: clean genversion
	rm -rf ./distrib/$(FAMILY)
	test -e distrib || mkdir distrib
	test -e distrib/$(FAMILY) || mkdir distrib/$(FAMILY)
	test ! -e ./putonrootfs || cp -r ./putonrootfs/* distrib/$(FAMILY)
	test ! -e ./$(FAMILY)/putonrootfs || cp -r ./$(FAMILY)/putonrootfs/* distrib/$(FAMILY)
	test -e distrib/$(FAMILY)/opt || mkdir distrib/$(FAMILY)/opt
	make FAMILY=$(FAMILY) -C libhisi -j4
	GOPRIVATE=openhisiipcam.org \
	GOPATH=$(GOPATH) \
    CGO_CFLAGS='$(CFLAGS)' \
    CGO_CFLAGS_ALLOW='.*' \
    GOOS=linux \
    GOARCH=arm \
    GOARM=7 \
    CGO_ENABLED=1 \
    $(GO) build \
        -gcflags "all=-l" \
        -tags $(GOTAGS) \
        $(GOFLAGS) \
        -o distrib/$(FAMILY)/opt/$(APP) \
        ./cmd

genversion:
	rm -f $(VERSIONFILE)
	@echo "#THIS FILE WAS GENERATED AUTOMATICLY. PLEASE DO NOT EDIT!" > $(VERSIONFILE)
	@echo "package buildinfo" > $(VERSIONFILE)
	@echo "" >> $(VERSIONFILE)
	@echo "const (" >> $(VERSIONFILE)
	@echo "  Date = \"$(NOW)\"" >> $(VERSIONFILE)
	@echo "  Branch = \"$(BRANCH)\"" >> $(VERSIONFILE)
	@echo "  User = \"$(USER)\"" >> $(VERSIONFILE)
	@echo "  ChipFamily = \"hi3516av200\"" >> $(VERSIONFILE) #TODO cross chip build
	@echo ")" >> $(VERSIONFILE)

godeps:
	GOPATH=$(GOPATH) $(GO) get -u github.com/gorilla/mux
	GOPATH=$(GOPATH) $(GO) get -u github.com/yuin/gopher-lua
	GOPATH=$(GOPATH) $(GO) get -u github.com/divan/expvarmon
	GOPATH=$(GOPATH) $(GO) get -u golang.org/x/lint/golint
	GOPATH=$(GOPATH) $(GO) get -u github.com/paultag/go-modprobe
	GOPATH=$(GOPATH) $(GO) get -u github.com/freman/goterm
	GOPATH=$(GOPATH) $(GO) get -u github.com/pion/webrtc

gover:
	$(GO) version

clean:
	make -C libhisi clean -j4
	rm -f ./app_minimal
