-include $(PARAMS_FILE)

APP	?= cam

GO	?= go

GOOS    ?= linux
GOARCH  ?= amd64
GOENV   := GOROOT_FINAL=/test/ GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) 

FAMILY	?= host

CMOS		?= unknown
CMOS_DRIVER ?= unknown

-include sdk/$(FAMILY)/Makefile.params

FAMILY_UPPERCASE    := $(shell echo '$(FAMILY)' | tr '[:lower:]' '[:upper:]')
CHIP_UPPERCASE      := $(shell echo '$(CHIP)' | tr '[:lower:]' '[:upper:]')

CGO_CFLAGS  := -I$(abspath sdk/$(FAMILY)/include) -D$(FAMILY_UPPERCASE) -DLOGGER_C
#-D$(CHIP_UPPERCASE)

#TODO rework this
ifneq ($(CMOS),unknown)
#ifneq ($(CMOS_DRIVER),)
#	CMOS := $(CMOS_DRIVER)
#endif
 	CMOS_TARGET := ../output/boards/$(BOARD)/cmos/lib_$(CMOS_DRIVER).a
endif

ifeq ($(APP),probe)
	CMOS_TARGET :=
endif

CMOS_TAGS		:= $(CMOS) cmos_data_$(CMOS_DATA) cmos_control_$(CMOS_CONTROL) cmos_bus_$(CMOS_BUS)

GOTAGS_PROBE    := gohisiprobe $(FAMILY) netgo nethttpomithttp2
#GOTAGS_CAM      := gohisicam $(FAMILY) $(CHIP) $(CMOS_TAGS) netgo nethttpomithttp2 openapi debug processing streamerJpeg streamerFile streamerRtsp streamerPipe
GOTAGS_CAM      := gohisicam $(FAMILY) $(CHIP) $(CMOS_TAGS) netgo nethttpomithttp2 openapi debug processing streamerJpeg
GOTAGS := $(GOTAGS_$(shell echo '$(APP)' | tr '[:lower:]' '[:upper:]'))

ifneq ($(wildcard ../hisilicon/$(FAMILY)/Makefile.params),) 
       LIB_TARGET := ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a
       CGO_LDFLAGS := $(abspath $(CMOS_TARGET)) $(abspath ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a)
endif 


CGOPARAMS := CGO_CFLAGS='$(CFLAGS) $(CGO_CFLAGS)' CGO_CFLAGS_ALLOW='.*' CGO_ENABLED=1 CGO_LDFLAGS='$(CGO_LDFLAGS) $(FAMILY_LDFLAGS)'

ifneq ($(strip $(shell git status --porcelain 2>/dev/null)),)
	GIT_TREE_STATE=(dirty)
endif

GIT_BRANCH	:= $(shell git branch | grep \* | cut -d " " -f2) $(GIT_TREE_STATE)
GIT_COMMIT  := $(shell git log --format="%H" -n 1) $(GIT_TREE_STATE)

SDKV 	:= $(HISDK_ARCHIVE) $(HISDK)

GOCC	:=$(shell $(GOENV) $(CGOPARAMS) $(GO) env | grep "CC=" | rev | cut -d '"' -f 2 | rev)

EMBEDDVARS :=   -X 'application/pkg/buildinfo.BuildDateTime=$(shell date --utc "+%Y-%m-%d %H:%M (UTC)")' \
                -X 'application/pkg/buildinfo.GoVersion=$(shell $(GOENV) $(GO) version)' \
                -X 'application/pkg/buildinfo.GccVersion=$(shell $(GOCC) --version | head -n 1)' \
                -X 'application/pkg/buildinfo.BuildTags=$(GOTAGS)' \
                -X 'application/pkg/buildinfo.BuildCommit=$(GIT_COMMIT)' \
                -X 'application/pkg/buildinfo.Family=$(FAMILY)' \
				-X 'application/pkg/buildinfo.SDK=$(SDKV)' \
                -X 'application/pkg/buildinfo.CmosProfile=$(CMOS)'  

#NO DEBUG SYMBOLS, COMMENT IF WANT TO TURN ON DEBUG
GODEBUG := -s -w 
#-extldflags "-static" 
#-linkmode "external"

guard:
	@echo "guard"

gover:
	$(GOENV) $(GO) version

goenv:
	$(GOENV) $(GO) env

strip-ko:
	$(STRIP) --strip-debug ./sdk/*/ko/*

LIBS := $(wildcard sdk/$(FAMILY)/lib/*.a)

../output/hisilicon/$(FAMILY)/lib$(FAMILY).a:
	$(AR) rcsT ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a $(LIBS)

#CMOS all in one archive
#CMOSES := $(wildcard pkg/mpp/cmos/$(FAMILY)/*/)
#CMOSDIRS := $(patsubst pkg/mpp/cmos/$(FAMILY)/%/, ../output/boards/$(BOARD)/cmos/%/, $(CMOSES))
#
#CMOSSRC   := $(wildcard pkg/mpp/cmos/$(FAMILY)/*/*.c)
#CMOSOBJ := $(patsubst pkg/mpp/cmos/$(FAMILY)/%.c, ../output/boards/$(BOARD)/cmos/%.o, $(CMOSSRC))
#
#../output/boards/$(BOARD)/cmos/lib_cmoses.a: $(CMOSDIRS) $(CMOSOBJ)
#	$(AR) rcs ../output/boards/$(BOARD)/cmos/lib_cmoses.a $(CMOSOBJ)
#
#../output/boards/$(BOARD)/cmos/%/:
#	[ -d $@ ] || mkdir -p $@
#
#../output/boards/$(BOARD)/cmos/%.o: pkg/mpp/cmos/$(FAMILY)/%.c
#	$(CC) $(CFLAGS) -I$(abspath sdk/$(FAMILY)/include) -c $< -o $@

#CMOS separate archive for each cmos

CMOSSRC   := $(wildcard pkg/mpp/cmos/$(FAMILY)/$(CMOS_DRIVER)/*.c)
CMOSOBJ := $(patsubst pkg/mpp/cmos/$(FAMILY)/$(CMOS_DRIVER)/%.c, ../output/boards/$(BOARD)/cmos/$(CMOS_DRIVER)/%.o, $(CMOSSRC))

../output/boards/$(BOARD)/cmos/%/:
	[ -d $@ ] || mkdir -p $@

../output/boards/$(BOARD)/cmos/$(CMOS_DRIVER)/%.o: pkg/mpp/cmos/$(FAMILY)/$(CMOS_DRIVER)/%.c
	$(CC) $(CFLAGS) -I$(abspath sdk/$(FAMILY)/include) -c $< -o $@

../output/boards/$(BOARD)/cmos/lib_$(CMOS_DRIVER).a: ../output/boards/$(BOARD)/cmos/$(CMOS_DRIVER)/ $(CMOSOBJ)
	 $(AR) rcs ../output/boards/$(BOARD)/cmos/lib_$(CMOS_DRIVER).a $(CMOSOBJ)

pkg/ko/kobin_$(APP)_$(FAMILY).go:
	GOPATH=$(GOPATH) $(GO) generate -tags "$(GOTAGS)" application/pkg/ko

generate: pkg/ko/kobin_$(APP)_$(FAMILY).go
	make -C pkg/utils/regs

distrib:
	mkdir distrib

deploy-dir-cam: distrib
	rm -rf ./distrib/$(FAMILY)
	mkdir -p ./distrib/$(FAMILY)/opt/www
	mkdir -p ./distrib/$(FAMILY)/opt/storage
	mkdir -p ./distrib/$(FAMILY)/opt/scripts
	mkdir -p ./distrib/$(FAMILY)/opt/configs
	mkdir -p ./distrib/$(FAMILY)/etc/init.d
	cp -r ./www/* ./distrib/$(FAMILY)/opt/www
	cp -r ./scripts/* ./distrib/$(FAMILY)/opt/scripts
	cp ./init/run.sh ./distrib/$(FAMILY)/opt
	cp ./init/S99gohisicam ./distrib/$(FAMILY)/etc/init.d/
	if [ -f ../boards/boards/$(BOARD)/config ]; then cp ../boards/boards/$(BOARD)/config ./distrib/$(FAMILY)/opt/board.config; fi
	cp ../application/configs/* ./distrib/$(FAMILY)/opt/configs/
	cp -r ./init/tools ././distrib/$(FAMILY)/opt/ 

deploy-dir-probe: distrib
	rm -rf ./distrib/$(FAMILY)
	mkdir -p ./distrib/$(FAMILY)/etc/init.d
	mkdir -p ./distrib/$(FAMILY)/opt
	cp ./init/S99gohisiprobe ./distrib/$(FAMILY)/etc/init.d/
	if [ -f ../boards/boards/$(BOARD)/config ]; then cp ../boards/boards/$(BOARD)/config ./distrib/$(FAMILY)/opt/board.config; fi
	#cp ../application/configs/* ./distrib/$(FAMILY)/opt/configs/

TRIMPATH := $(abspath .) 

build-cam: generate $(LIB_TARGET) $(CMOS_TARGET) deploy-dir-cam
	$(GOENV) $(CGOPARAMS) $(GO) build -gcflags=all=-trimpath=$(TRIMPATH) -asmflags=all=-trimpath=$(TRIMPATH) -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" -o ./distrib/$(FAMILY)/opt/gohisicam application/cmd/gohisicam

build-probe: generate $(LIB_TARGET) deploy-dir-probe
	$(GOENV) $(CGOPARAMS) $(GO) build -gcflags=all=-trimpath=$(TRIMPATH) -asmflags=all=-trimpath=$(TRIMPATH) -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" -o ./distrib/$(FAMILY)/opt/gohisiprobe application/cmd/gohisiprobe

build: build-$(APP)

test-host: generate
	GOPATH=$(GOPATH) $(GO) test -tags "$(GOTAGS)" ./... -v

test-target: generate
	rm -rf ./tests
	mkdir tests
	for i in `GOPATH=$(GOPATH) $(GO) list -tags "$(GOTAGS)" ./...`; do \
		echo $$i; \
		$(GOENV) $(CGOPARAMS) $(GO) test -tags "$(GOTAGS)" -c $$i -o tests/`echo $$i | rev | cut -d '/' -f 1 | rev`.test; \
	done

#build-object:
#	$(GOENV) $(CGOPARAMS) $(GO) tool nm ./tests/*

testvideo.h264:
	ffmpeg -f lavfi -i testsrc=size=320x340 -t 30 -pix_fmt yuv420p -c:v libx264 -tune zerolatency -profile:v main -g 25 -vbsf h264_mp4toannexb -f h264 testvideo.h264

testvideo.h265:
	ffmpeg -f lavfi -i testsrc=size=320x340 -t 30 -pix_fmt yuv420p -c:v libx265 -tune zerolatency -g 25 -f hevc testvideo.h265

testvideo-h264.mp4: testvideo.h264
	MP4Box -add testvideo.h264 -fps 25 testvideo-h264.mp4

testvideo-h265.mp4: testvideo.h265
	MP4Box -add testvideo.h265 -fps 25 testvideo-h265.mp4

clean:
	rm -rf ./pkg/ko/kobin_*.go
	rm -rf ./distrib
	rm -f ./testvideo.h264
	rm -f ./testvideo.h265
	rm -f ./testvideo-h264.mp4
	rm -f ./testvideo-h265.mp4

freeze-vendors:
	GOPATH=$(GOPATH) $(GO) mod vendor

install-deps:
	GOPATH=$(GOPATH) $(GO) get -u golang.org/x/tools/cmd/goyacc
	GOPATH=$(GOPATH) $(GO) get -u github.com/shuLhan/go-bindata
	go get -u github.com/shuLhan/go-bindata
