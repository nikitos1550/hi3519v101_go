ifndef PARAMS_FILE
    $(error 'PARAMS_FILE' variable must be defined)
endif
include $(PARAMS_FILE)

GOPATH := $(abspath ../output/go_workspace)

#GOTAGS := "$(FAMILY) netgo openapi debug lua onvif"

GOOS    := linux
GOARCH  := arm
GOARM   := $(ARMV) #FROM Makefile.params
GOENV   := GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM)

FAMILY  := hi3516av200

GOTAGS  := "hi3516av200"

PATH    := $(PATH):$(GOPATH)/bin

APPOUT     := $(abspath ../output/application/a.out)

guard:
	@echo "guard"

gover:
	$(GOENV) $(GO) version

goenv:
	$(GOENV) $(GO) env

generate:
	cd src; GOPATH=$(GOPATH) $(GO) generate -tags $(GOTAGS) application/koloader

build: generate
	cd src; $(GOENV) $(GO) build -tags $(GOTAGS) -o $(APPOUT)

buid-with-vendors: generate
	cd src; $(GOENV) $(GO) build -tags $(GOTAGS) -mod vendor -o $(APPOUT)

clean:
	rm -f $(APPOUT)

fix-vendors:
	GOPATH=$(GOPATH) $(GO) mod vendor
