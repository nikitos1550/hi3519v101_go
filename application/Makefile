include $(PARAMS_FILE)

GO	?= go

GOOS    ?= linux
GOARCH  ?= amd64
GOENV   := GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) 

FAMILY	?= host
CMOS	?= unknown

FAMILY_UPPERCASE    := $(shell echo '$(FAMILY)' | tr '[:lower:]' '[:upper:]')
CHIP_UPPERCASE      := $(shell echo '$(CHIP)' | tr '[:lower:]' '[:upper:]')

#GOTAGS_PROBE   	:= $(FAMILY) $(CHIP) netgo nethttpomithttp2
GOTAGS_PROBE    := $(FAMILY) netgo nethttpomithttp2
GOTAGS_CAM 	:= $(FAMILY) $(CHIP) $(CMOS) netgo nethttpomithttp2 openapi scripts debug streamerJpeg streamerFile streamerRtsp streamerPipe

CGO_CFLAGS  := -I$(abspath sdk/$(FAMILY)/include) -D$(FAMILY_UPPERCASE)
#-D$(CHIP_UPPERCASE)

#TODO rework this
ifneq ($(wildcard ./output/boards/$(BOARD)/cmos/lib_cmoses.a),)
       CMOS_TARGET := ../output/boards/$(BOARD)/cmos/lib_cmoses.a
endif

ifneq ($(wildcard ../hisilicon/$(FAMILY)/Makefile.params),) 
       LIB_TARGET := ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a
       CGO_LDFLAGS := $(abspath ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a) $(abspath $(CMOS_TARGET))
endif 


CGOPARAMS := CGO_CFLAGS='$(CFLAGS) $(CGO_CFLAGS)' CGO_CFLAGS_ALLOW='.*' CGO_ENABLED=1 CGO_LDFLAGS='$(CGO_LDFLAGS) $(FAMILY_LDFLAGS)'

EMBEDDVARS :=   -X 'application/pkg/buildinfo.BuildDateTime=$(shell date +%Y-%m-%d\ %H:%M)' \
                -X 'application/pkg/buildinfo.GoVersion=$(shell $(GO) version)' \
                -X 'application/pkg/buildinfo.GccVersion=$(shell $(CC) --version)' \
                -X 'application/pkg/buildinfo.BuildTags=$(GOTAGS)' \
                -X 'application/pkg/buildinfo.BuildUser=$(USER)' \
                -X 'application/pkg/buildinfo.BuildCommit=$(shell git log --pretty=format:"%h" -n 1)' \
                -X 'application/pkg/buildinfo.BuildBranch=$(shell git branch | grep \* | cut -d " " -f2)' \
                -X 'application/pkg/buildinfo.Family=$(FAMILY)' \
                -X 'application/pkg/buildinfo.CmosProfile=$(CMOS)'  

#NO DEBUG SYMBOLS, COMMENT IF WANT TO TURN ON DEBUG
GODEBUG := -s -w -extldflags "-static"

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

CMOSES := $(wildcard pkg/mpp/cmos/$(FAMILY)/*/)
CMOSDIRS := $(patsubst pkg/mpp/cmos/$(FAMILY)/%/, ../output/boards/$(BOARD)/cmos/%/, $(CMOSES))

CMOSSRC   := $(wildcard pkg/mpp/cmos/$(FAMILY)/*/*.c)
CMOSOBJ := $(patsubst pkg/mpp/cmos/$(FAMILY)/%.c, ../output/boards/$(BOARD)/cmos/%.o, $(CMOSSRC))

../output/boards/$(BOARD)/cmos/lib_cmoses.a: $(CMOSDIRS) $(CMOSOBJ)
	$(AR) rcs ../output/boards/$(BOARD)/cmos/lib_cmoses.a $(CMOSOBJ)

../output/boards/$(BOARD)/cmos/%/:
	[ -d $@ ] || mkdir -p $@

../output/boards/$(BOARD)/cmos/%.o: pkg/mpp/cmos/$(FAMILY)/%.c
	$(CC) $(CFLAGS) -I$(abspath sdk/$(FAMILY)/include) -c $< -o $@

generate: 
	GOPATH=$(GOPATH) $(GO) generate -tags "$(FAMILY)" application/pkg/ko
	#GOPATH=$(GOPATH) $(GO) generate -tags "$(FAMILY)" application/pkg/utils/regs

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
	if [ -f ../boards/$(BOARD)/config ]; then cp ../boards/$(BOARD)/config ./distrib/$(FAMILY)/opt/board.config; fi
	cp ../application/configs/* ./distrib/$(FAMILY)/opt/configs/
	cp -r ./init/tools ././distrib/$(FAMILY)/opt/ 

deploy-dir-probe: distrib
	rm -rf ./distrib/$(FAMILY)
	mkdir -p ./distrib/$(FAMILY)/opt/configs
	mkdir -p ./distrib/$(FAMILY)/etc/init.d
	cp ./init/S99gohisiprobe ./distrib/$(FAMILY)/etc/init.d/
	if [ -f ../boards/$(BOARD)/config ]; then cp ../boards/$(BOARD)/config ./distrib/$(FAMILY)/opt/board.config; fi
	cp ../application/configs/* ./distrib/$(FAMILY)/opt/configs/

TRIMPATH := $(abspath .) 

build-cam: generate $(LIB_TARGET) $(CMOS_TARGET) deploy-dir-cam
	$(GOENV) $(CGOPARAMS) $(GO) build -gcflags=all=-trimpath=$(TRIMPATH) -asmflags=all=-trimpath=$(TRIMPATH) -a -tags "$(GOTAGS_CAM)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" -o ./distrib/$(FAMILY)/opt/gohisicam application/cmd/gohisicam

build-probe: generate $(LIB_TARGET) deploy-dir-probe
	$(GOENV) $(CGOPARAMS) $(GO) build -gcflags=all=-trimpath=$(TRIMPATH) -asmflags=all=-trimpath=$(TRIMPATH) -a -tags "$(GOTAGS_PROBE)" -o ./distrib/$(FAMILY)/opt/gohisiprobe application/cmd/gohisiprobe

#build-object:
#	$(GOENV) $(CGOPARAMS) $(GO) tool nm ./distrib/$(FAMILY)/opt/application

#buid-with-vendors: generate
#	$(GOENV) $(CGOPARAMS) $(GO) build -a -tags $(GOTAGS) -mod vendor -o $(APPOUT) application/cmd/daemon

testvideo.h264:
	ffmpeg -f lavfi -i testsrc=size=320x340 -t 30 -pix_fmt yuv420p -c:v libx264 -tune zerolatency -profile:v main -g 25 -vbsf h264_mp4toannexb -f h264 testvideo.h264

testvideo.h265:
	ffmpeg -f lavfi -i testsrc=size=320x340 -t 30 -pix_fmt yuv420p -c:v libx265 -tune zerolatency -g 25 -f hevc testvideo.h265

testvideo-h264.mp4: testvideo.h264
	MP4Box -add testvideo.h264 -fps 25 testvideo-h264.mp4

testvideo-h265.mp4: testvideo.h265
	MP4Box -add testvideo.h265 -fps 25 testvideo-h265.mp4

clean:
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
