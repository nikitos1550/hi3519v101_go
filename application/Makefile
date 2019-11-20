ifndef PARAMS_FILE
    $(error 'PARAMS_FILE' variable must be defined)
endif
include $(PARAMS_FILE)

GOPATH := $(abspath ../output/go_workspace)

GOOS    := linux
GOARCH  := arm
GOARM   := $(ARMV) #FROM Makefile.params
GOENV   := GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM)

#GOTAGS  := "hi3516av200"
GOTAGS  := "$(FAMILY)"
#GOTAGS  := "$(FAMILY) netgo openapi debug lua onvif"

PATH    := $(PATH):$(GOPATH)/bin

APPOUT  := $(abspath ../output/application/a.out)

CGO_CFLAGS  := -I$(abspath sdk/$(FAMILY)/include)
CGO_LDFLAGS := $(abspath ../output/$(FAMILY)/lib$(FAMILY).a)

CGOPARAMS := CGO_CFLAGS='$(CFLAGS) $(CGO_CFLAGS)' CGO_CFLAGS_ALLOW='.*' CGO_ENABLED=1 CGO_LDFLAGS='$(CGO_LDFLAGS)'



guard:
	@echo "guard"

gover:
	$(GOENV) $(GO) version

goenv:
	$(GOENV) $(GO) env

strip-ko:
	$(STRIP) --strip-debug ./sdk/*/ko/*

LIBS := $(wildcard sdk/$(FAMILY)/lib/*.a)
libsdk:
	$(AR) rcsT ../output/$(FAMILY)/lib$(FAMILY).a $(LIBS)

generate:
	GOPATH=$(GOPATH) $(GO) generate -tags $(GOTAGS) application/pkg/koloader

build: generate
	$(GOENV) $(CGOPARAMS) $(GO) build -tags $(GOTAGS) -o $(APPOUT) application/cmd/daemon

build-tester: generate
	$(GOENV) $(CGOPARAMS) $(GO) build -tags $(GOTAGS) -o $(APPOUT) application/cmd/tester

#buid-with-vendors: generate
#	$(GOENV) $(GO) build -tags $(GOTAGS) -mod vendor -o $(APPOUT) application/cmd/daemon

clean:
	rm -f $(APPOUT)

fix-vendors:
	GOPATH=$(GOPATH) $(GO) mod vendor
