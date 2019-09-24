include Makefile.params
include Makefile.tags

THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GOPATH := $(THIS_DIR)go

NOW=$(shell date)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

VERSIONFILE=$(GOPATH)/src/openhisiipcam.org/buildinfo/buildinfo.go

build: clean genversion
	make -C libhisi -j4
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
        -o app_minimal \
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

#gohi3516av200:
#	GOPRIVATE=openhisiipcam.org \
#    GOPATH=$(GOPATH) \
#    CGO_CFLAGS='$(CFLAGS)' \
#    CGO_CFLAGS_ALLOW='.*' \
#    GOOS=linux \
#    GOARCH=arm \
#    GOARM=7 \
#    CGO_ENABLED=1 \
#	$(GO) install \
#        -gcflags "all=-l" \
#        -tags $(GOTAGS) \
#        -a \
#        -ldflags="-extldflags '-static -Wl,--gc-sections'" \
#        openhisiipcam.org/hi3516av200

#gofmt:
	#@GOPATH=$(GOPATH) $(GOFMT) -l cmd/*.go
	#@GOPATH=$(GOPATH) $(GOFMT) -l $(GOPATH)/src/openhisiipcam.org/*/*.go
	#$(GOPATH)/bin/golint cmd
	#$(GOPATH)/bin/golint $(GOPATH)/src/openhisiipcam.org/*
	#GOPATH=$(GOPATH) $(GO) vet $(GOPATH)/src/openhisiipcam.org/*
	#GOPATH=$(GOPATH) $(GO) vet cmd/*

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
