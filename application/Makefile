ifndef PARAMS_FILE
    $(error 'PARAMS_FILE' variable must be defined)
endif
include $(PARAMS_FILE)

GOPATH := $(abspath ../output/go_workspace)

GOOS    := linux
GOARCH  := arm
GOARM   := $(ARMV) #FROM Makefile.params
GOENV   := GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM)

#GOTAGS  := hi3516av200
GOTAGS  := $(FAMILY)
#GOTAGS  := $(FAMILY) netgo openapi debug lua onvif

PATH    := $(PATH):$(GOPATH)/bin

APPOUT  := $(abspath ../output/application)

CGO_CFLAGS  := -I$(abspath sdk/$(FAMILY)/include)
CGO_LDFLAGS := $(abspath ../output/$(FAMILY)/lib$(FAMILY).a)

CGOPARAMS := CGO_CFLAGS='$(CFLAGS) $(CGO_CFLAGS)' CGO_CFLAGS_ALLOW='.*' CGO_ENABLED=1 CGO_LDFLAGS='$(CGO_LDFLAGS)'

EMBEDDVARS := " -X 'application/pkg/buildinfo.BuildDateTime=$(shell date +%Y-%m-%d\ %H:%M)' \
                -X 'application/pkg/buildinfo.GoVersion=$(shell $(GO) version)' \
                -X 'application/pkg/buildinfo.GccVersion=$(shell $(CC) --version)' \
                -X 'application/pkg/buildinfo.BuildTags=$(GOTAGS)' \
                -X 'application/pkg/buildinfo.BoardVendor=$(VENDOR)' \
                -X 'application/pkg/buildinfo.BoardModel=$(MODEL)' \
                -X 'application/pkg/buildinfo.Chip=$(CHIP)' \
                -X 'application/pkg/buildinfo.CmosProfile=$(CMOS)' \
                -X 'application/pkg/buildinfo.TotalRam=$(RAM_SIZE)' \
                -X 'application/pkg/buildinfo.LinuxRam=$(RAM_LINUX)' \
                -X 'application/pkg/buildinfo.MppRam=$(RAM_MPP)'"

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
	GOPATH=$(GOPATH) $(GO) generate -tags "$(GOTAGS)" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "hi3516av100" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516av200" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516cv100" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516cv200" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516cv300" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516cv500" application/pkg/koloader
	##GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516ev200" application/pkg/koloader
	##GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3519av100" application/pkg/koloader
	##GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3559av100" application/pkg/koloader

build: generate libsdk
	$(GOENV) $(CGOPARAMS) $(GO) build -a -tags "$(GOTAGS)" -ldflags $(EMBEDDVARS) -o $(APPOUT)/daemon application/cmd/daemon

deploy-dir:
	rm -rf ./distrib
	mkdir ./distrib/
	mkdir ./distrib/$(FAMILY)

build-tester: generate libsdk deploy-dir
	#$(GOENV) $(CGOPARAMS) $(GO) build -tags "$(GOTAGS)" -ldflags $(EMBEDDVARS) -o $(APPOUT)/tester application/cmd/tester
	$(GOENV) $(CGOPARAMS) $(GO) build -a -tags "$(GOTAGS)" -ldflags $(EMBEDDVARS) -o ./distrib/$(FAMILY)/application application/cmd/tester

#buid-with-vendors: generate
#	$(GOENV) $(GO) build -tags $(GOTAGS) -mod vendor -o $(APPOUT) application/cmd/daemon

clean:
	rm -f $(APPOUT)

fix-vendors:
	GOPATH=$(GOPATH) $(GO) mod vendor

#tools:
#	GOPATH=$(GOPATH) $(GO) get golang.org/x/tools/cmd/godoc

doc2:
	GOPATH=$(GOPATH) $(GOBIN)/godoc --goroot /home/nikitos1550/work/tmp/hi3519v101_go/output/boards/jvt_hi3519v101_imx274/toolchain/lib/go -v
	#cd doc; wget -r -np -N -E -p -k http://localhost:6060/pkg/
