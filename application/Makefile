include $(PARAMS_FILE)

APP ?= cam
GO  ?= go
OUTDIR ?= ../output

GOOS    ?= linux
GOARCH  ?= amd64
GOENV   := GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) 

FAMILY	?= host

CMOS		?= unknown
CMOS_DRIVER ?= unknown

-include sdk/$(FAMILY)/Makefile.params

FAMILY_UPPERCASE    := $(shell echo '$(FAMILY)' | tr '[:lower:]' '[:upper:]')
CHIP_UPPERCASE      := $(shell echo '$(CHIP)' | tr '[:lower:]' '[:upper:]')

CFLAGS += -fPIC

CGO_CFLAGS  := 	-Wall \
				-Wno-comment \
				-I$(abspath sdk/$(FAMILY)/include) \
				-I$(abspath $(STAGING_DIR)/usr/include/opus) \
				-D$(FAMILY_UPPERCASE) \
				-DLOGGER_C

LIB_CMOSES := $(abspath $(OUTDIR)/boards/$(BOARD)/cmos/lib_cmoses.a)
LIB_FAMILY := $(abspath $(OUTDIR)/hisilicon/$(FAMILY)/lib$(FAMILY).a)

#-I$(abspath ../output/hisilicon/$(FAMILY)/.br-rootfs/staging/usr/include/fdk-aac) \

#-D$(CHIP_UPPERCASE)

#CGO_CXXFLAGS := -Wall -I$(abspath ../experimental/tf/tensorflow) \
#				-I$(abspath ../experimental/tf/tensorflow/tensorflow/lite/tools/make/downloads/absl) \
#				-I$(abspath ../experimental/tf/tensorflow/tensorflow/lite/tools/make/downloads/flatbuffers/include)



#TODO rework this
ifneq ($(CMOS),unknown)
    CMOS_TARGET := $(OUTDIR)/boards/$(BOARD)/cmos/lib_$(CMOS_DRIVER).a
endif
ifeq ($(APP),probe)
    CMOS_TARGET :=
endif


CMOS_TAGS     := $(CMOS) cmos_data_$(CMOS_DATA) cmos_control_$(CMOS_CONTROL) cmos_bus_$(CMOS_BUS)

GOTAGS_PROBE  := gohisiprobe $(FAMILY) netgo nethttpomithttp2
 
#GOTAGS_CAM      := gohisicam $(FAMILY) $(CHIP) $(CMOS_TAGS) netgo nethttpomithttp2 openapi debug processing streamerJpeg streamerFile streamerRtsp streamerPipe streamerYuv
#GOTAGS_CAM      := gohisicam $(FAMILY) netgo nethttpomithttp2 openapi debug processing streamerJpeg koEmbed scripts streamerFile streamerWebrtc streamerRtsp streamerYuv
GOTAGS_CAM      := gohisicam $(FAMILY) netgo nethttpomithttp2 openapi debug processing streamerJpeg koEmbed scripts streamerWebrtc

#GOTAGS_CAM      := gohisicam $(FAMILY) netgo nethttpomithttp2 openapi debug processing streamerJpeg streamerFile streamerRtsp streamerPipe streamerWebrtc streamerYuv

GOTAGS := $(GOTAGS_$(shell echo '$(APP)' | tr '[:lower:]' '[:upper:]'))


CGO_LDFLAGS := \
    $(LIB_CMOSES) \
    $(LIB_FAMILY) \
    $(abspath $(STAGING_DIR)/usr/lib/libopus.a)


#GODEBUG=cgocheck=0
CGOPARAMS := CGO_CFLAGS='$(CFLAGS) $(CGO_CFLAGS)' CGO_CXXFLAGS='$(CXXFLAGS) $(CGO_CXXFLAGS)' CGO_CFLAGS_ALLOW='.*' CGO_ENABLED=1 CGO_LDFLAGS='$(CGO_LDFLAGS) $(FAMILY_LDFLAGS)'

ifneq ($(strip $(shell git status --porcelain 2>/dev/null)),)
	GIT_TREE_STATE=(dirty)
endif

GIT_BRANCH	:= $(shell git branch | grep \* | cut -d " " -f2) $(GIT_TREE_STATE)
GIT_COMMIT  := $(shell git log --format="%H" -n 1) $(GIT_TREE_STATE)

SDKV 	:= $(HISDK_ARCHIVE) $(HISDK)

GOCC	:=$(shell $(GOENV) $(CGOPARAMS) $(GO) env | grep "CC=" | rev | cut -d '"' -f 2 | rev)

#ELFSIZE:=1234567
ELFSIZE:=1234567
#ELFSIZE:=$(shell stat -c %s $(abspath ./distrib/hi3516ev200/opt/gohisicam))

EMBEDDVARS  =   -X 'application/core/compiletime.BuildDateTime=$(shell date --utc "+%Y-%m-%d %H:%M (UTC)")' \
                -X 'application/core/compiletime.GoVersion=$(shell $(GOENV) $(GO) version)' \
                -X 'application/core/compiletime.GccVersion=$(shell $(GOCC) --version | head -n 1)' \
                -X 'application/core/compiletime.BuildTags=$(GOTAGS)' \
                -X 'application/core/compiletime.BuildCommit=$(GIT_COMMIT)' \
                -X 'application/core/compiletime.Family=$(FAMILY)' \
				-X 'application/core/compiletime.SDK=$(SDKV)' \
                -X 'application/core/compiletime.CmosProfile=$(CMOS)' \
				-X 'application/core/ko.ElfSizeStr=$(ELFSIZE)'

#NO DEBUG SYMBOLS, COMMENT IF WANT TO TURN ON DEBUG
GODEBUG := -s -w -extldflags "-static" 
#-linkmode "external"
#GODEBUG := -s -w

guard:
	@echo "guard"

gover:
	$(GOENV) $(GO) version

goenv:
	$(GOENV) $(GO) env

strip-ko:
	$(STRIP) --strip-debug ./sdk/*/ko/*

LIBS := $(wildcard sdk/$(FAMILY)/lib/*.a)

$(LIB_FAMILY):
	mkdir -p $(@D)
	$(AR) rcsT $@ $(abspath $(LIBS))

#CMOS all in one archive
CMOSES   := $(wildcard core/mpp/cmos/$(FAMILY)/*/)
CMOSSRC  := $(wildcard core/mpp/cmos/$(FAMILY)/*/*.c)
CMOSDIRS := $(patsubst core/mpp/cmos/$(FAMILY)/%/, $(OUTDIR)/boards/$(BOARD)/cmos/%/, $(CMOSES))
CMOSOBJ  := $(patsubst core/mpp/cmos/$(FAMILY)/%.c, $(OUTDIR)/boards/$(BOARD)/cmos/%.o, $(CMOSSRC))

$(LIB_CMOSES): $(CMOSDIRS) $(CMOSOBJ)
	$(AR) rcs $@ $(CMOSOBJ)

$(OUTDIR)/boards/$(BOARD)/cmos/%/:
	[ -d $@ ] || mkdir -p $@

$(OUTDIR)/boards/$(BOARD)/cmos/%.o: core/mpp/cmos/$(FAMILY)/%.c
	$(CC) $(CFLAGS) -I$(abspath sdk/$(FAMILY)/include) -c $< -o $@


core/ko/embed_$(FAMILY).go:
	GOPATH=$(GOPATH) $(GO) generate -tags "$(GOTAGS) generate" application/core/ko

generate: core/ko/embed_$(FAMILY).go
	make -C core/utils/regs

distrib:
	mkdir distrib

deploy-dir-cam: distrib
	rm -rf ./distrib/$(FAMILY)
	mkdir -p ./distrib/$(FAMILY)/opt/www
	mkdir -p ./distrib/$(FAMILY)/opt/storage
	mkdir -p ./distrib/$(FAMILY)/opt/scripts
	#mkdir -p ./distrib/$(FAMILY)/opt/configs
	mkdir -p ./distrib/$(FAMILY)/etc/init.d
	mkdir -p ./distrib/$(FAMILY)/opt/cgi-bin
	cp -r ./www/* ./distrib/$(FAMILY)/opt/www
	#cp -r ./scripts/* ./distrib/$(FAMILY)/opt/scripts
	cp ./init/run.sh ./distrib/$(FAMILY)/opt
	cp ./init/S99gohisicam ./distrib/$(FAMILY)/etc/init.d/
	cp ./init/ldd.sh ./distrib/$(FAMILY)/opt
	cp ./init/uhttpd.sh ./distrib/$(FAMILY)/opt
	if [ -f ../boards/boards/$(BOARD)/config ]; then cp ../boards/boards/$(BOARD)/config ./distrib/$(FAMILY)/opt/board.config; fi
	#cp ../application/configs/* ./distrib/$(FAMILY)/opt/configs/
	cp -r ./init/tools ././distrib/$(FAMILY)/opt/ 
	mkdir -p ./distrib/$(FAMILY)/opt/svp
	#cp -r ./svp/* ./distrib/$(FAMILY)/opt/svp
	mkdir -p ./distrib/$(FAMILY)/opt/nfs

#deploy-dir-probe: distrib
#	rm -rf ./distrib/$(FAMILY)
#	mkdir -p ./distrib/$(FAMILY)/etc/init.d
#	mkdir -p ./distrib/$(FAMILY)/opt
#	cp ./init/S99gohisiprobe ./distrib/$(FAMILY)/etc/init.d/
#	if [ -f ../boards/boards/$(BOARD)/config ]; then cp ../boards/boards/$(BOARD)/config ./distrib/$(FAMILY)/opt/board.config; fi
#	#cp ../application/configs/* ./distrib/$(FAMILY)/opt/configs/

TRIMPATH := $(abspath .) 

#build-cam: generate $(LIB_TARGET) $(CMOS_TARGET) deploy-dir-cam
build-cam: generate $(LIB_FAMILY) $(LIB_CMOSES) deploy-dir-cam vendors-build
	#$(GOENV) $(CGOPARAMS) $(GO) build -work -gcflags=all="-trimpath=$(TRIMPATH) -l -B -C" -asmflags=all=-trimpath=$(TRIMPATH) -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" -o ./gohisicam application/cmd/gohisicam
	#$(GOENV) $(CGOPARAMS) $(GO) install -work -gcflags=all=-trimpath=$(TRIMPATH) -asmflags=all=-trimpath=$(TRIMPATH) -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" ./pkg/...
	#stat -c %s ./gohisicam
	#echo $(shell du -b ./gohisicam | cut -f 1 2>&1)
	#echo $(shell stat -c %s ./gohisicam 2>&1)
	#$ ( eval ELFSIZE := $(shell stat -c %s ./gohisicam 2>&1))
	#echo $(ELFSIZE)
	#rm ./gohisicam
	$(GOENV) $(CGOPARAMS) $(GO) build -work -gcflags=all="-trimpath=$(TRIMPATH) -l -B -C" -asmflags=all=-trimpath=$(TRIMPATH) -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" -o ./distrib/$(FAMILY)/opt/gohisicam application/cmd/gohisicam
	#stat -c %s ./distrib/$(FAMILY)/opt/gohisicam
	#rm ./gohisicam
	#cat pkg/ko/append_$(FAMILY).bin >> ./distrib/$(FAMILY)/opt/gohisicam

build-cam2: generate $(LIB_FAMILY) $(LIB_CMOSES) deploy-dir-cam vendors-build
	$(GOENV) $(CGOPARAMS) $(GO) build -work -gcflags=all="-trimpath=$(TRIMPATH) -l -B -C" -asmflags=all=-trimpath=$(TRIMPATH) -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" -o ./distrib/$(FAMILY)/opt/gohisicam application/cmd/recorder

vendors-build:
	cd vendors/quirc; make; cd quirc; make clean; make CC=$(CC) libquirc.a

build-probe: generate $(LIB_FAMILY) deploy-dir-probe
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
	rm -rf ./core/ko/embed_*.go
	rm -rf ./core/ko/append_*.bin
	rm -rf ./core/ko/append_*.go
	rm -rf ./distrib
	rm -f ./testvideo.h264
	rm -f ./testvideo.h265
	rm -f ./testvideo-h264.mp4
	rm -f ./testvideo-h265.mp4
	rm -f $(LIB_CMOSES)
	rm -f $(LIB_FAMILY)


freeze-vendors:
	GOPATH=$(GOPATH) $(GO) mod vendor

install-deps:
	GOPATH=$(GOPATH) $(GO) get -u golang.org/x/tools/cmd/goyacc
	GOPATH=$(GOPATH) $(GO) get -u github.com/shuLhan/go-bindata
	go get -u github.com/shuLhan/go-bindata

