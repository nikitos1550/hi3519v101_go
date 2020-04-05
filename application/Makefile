ifndef PARAMS_FILE
    $(error 'PARAMS_FILE' variable must be defined)
endif
include $(PARAMS_FILE)

GOPATH := $(abspath ../output/go_workspace)

GOOS    := linux
GOARCH  := arm
GOARM   := $(ARMV) #FROM Makefile.params
GOENV   := GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) 
#GOROOT_FINAL="/test/"

APIDOC          := tools/apidoc/bin/apidoc
APIDOCOPENAPI   := tools/apidoc-openapi/bin/apidoc-openapi
REDOC           := ../tools/redoc/cli/redoc-cli

FAMILY_UPPERCASE    := $(shell echo '$(FAMILY)' | tr '[:lower:]' '[:upper:]')
CHIP_UPPERCASE      := $(shell echo '$(CHIP)' | tr '[:lower:]' '[:upper:]')

GOTAGS_PROBE   := $(FAMILY) $(CHIP) gohisiprobe netgo nethttpomithttp2
GOTAGS_DAEMON  := $(FAMILY) $(CHIP) netgo nethttpomithttp2 imx274 openapi scripts debug streamerJpeg streamerFile streamerWebrtc streamerRtsp
#GOTAGS_DAEMON  := $(FAMILY) $(CHIP) netgo nethttpomithttp2 imx274 openapi debug streamerJpeg 
#GOTAGS_DAEMON  := $(FAMILY) $(CHIP) imx274 openapi debug netgo streamerJpeg
#GOTAGS_DAEMON  := $(FAMILY) imx274 debug
#GOTAGS  := $(FAMILY) netgo openapi debug lua onvif

PATH    := $(PATH):$(GOPATH)/bin

APPOUT  := $(abspath ../output/application)

CGO_CFLAGS  := -I$(abspath sdk/$(FAMILY)/include) -D$(FAMILY_UPPERCASE) 
#-D$(CHIP_UPPERCASE)
CGO_LDFLAGS_DAEMON :=   $(abspath ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a) \
                        $(abspath ../output/boards/$(BOARD)/cmos/lib_cmoses.a)

CGO_LDFLAGS_TESTER :=   $(abspath ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a)

#CGO_LDFLAGS := $(abspath ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a) $(abspath ../output/boards/$(BOARD)/cmos/lib$(CMOS).a) -ldl -lm

CGOPARAMS_PROBE  := CGO_CFLAGS='$(CFLAGS) $(CGO_CFLAGS)' CGO_CFLAGS_ALLOW='.*' CGO_ENABLED=1 CGO_LDFLAGS='$(CGO_LDFLAGS_TESTER)'
CGOPARAMS_DAEMON := CGO_CFLAGS='$(CFLAGS) $(CGO_CFLAGS)' CGO_CFLAGS_ALLOW='.*' CGO_ENABLED=1 CGO_LDFLAGS='$(CGO_LDFLAGS_DAEMON)'

EMBEDDVARS :=   -X 'application/pkg/buildinfo.BuildDateTime=$(shell date +%Y-%m-%d\ %H:%M)' \
                -X 'application/pkg/buildinfo.GoVersion=$(shell $(GO) version)' \
                -X 'application/pkg/buildinfo.GccVersion=$(shell $(CC) --version)' \
                -X 'application/pkg/buildinfo.BuildTags=$(GOTAGS)' \
                -X 'application/pkg/buildinfo.BuildUser=$(USER)' \
                -X 'application/pkg/buildinfo.BuildCommit=$(shell git log --pretty=format:"%h" -n 1)' \
                -X 'application/pkg/buildinfo.BuildBranch=$(shell git branch | grep \* | cut -d " " -f2)' \
				-X 'application/pkg/buildinfo.BoardProfile=$(BOARD)' \
                -X 'application/pkg/buildinfo.BoardVendor=$(VENDOR)' \
                -X 'application/pkg/buildinfo.BoardModel=$(MODEL)' \
                -X 'application/pkg/buildinfo.Chip=$(CHIP)' \
                -X 'application/pkg/buildinfo.CmosProfile=$(CMOS)' \
                -X 'application/pkg/buildinfo.TotalRam=$(RAM_SIZE)' \
                -X 'application/pkg/buildinfo.LinuxRam=$(RAM_LINUX)' \
                -X 'application/pkg/buildinfo.MppRam=$(RAM_MPP)'

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
libsdk:
	$(AR) rcsT ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a $(LIBS)

CMOSES := $(wildcard pkg/mpp/cmos/$(FAMILY)/*/)
CMOSDIRS := $(patsubst pkg/mpp/cmos/$(FAMILY)/%/, ../output/boards/$(BOARD)/cmos/%/, $(CMOSES))

CMOSSRC   := $(wildcard pkg/mpp/cmos/$(FAMILY)/*/*.c)
CMOSOBJ := $(patsubst pkg/mpp/cmos/$(FAMILY)/%.c, ../output/boards/$(BOARD)/cmos/%.o, $(CMOSSRC))

#CMOSSRC := $(wildcard pkg/mpp/cmos/$(FAMILY)/*/*.c)
#CMOSOBJ := $(patsubst pkg/mpp/cmos/$(FAMILY)/$(CMOS)/%.c, ../output/boards/$(BOARD)/%/%.o, $(CMOSSRC))

../output/boards/$(BOARD)/cmos/lib_cmoses.a: $(CMOSDIRS) $(CMOSOBJ)
	$(AR) rcs ../output/boards/$(BOARD)/cmos/lib_cmoses.a $(CMOSOBJ)

../output/boards/$(BOARD)/cmos/%/:
	[ -d $@ ] || mkdir -p $@

../output/boards/$(BOARD)/cmos/%.o: pkg/mpp/cmos/$(FAMILY)/%.c
	$(CC) $(CFLAGS) -I$(abspath sdk/$(FAMILY)/include) -c $< -o $@

#libcmos: $(CMOSOBJ)
#	@echo $(CMOSOBJ)
#	$(AR) rcs ../output/boards/$(BOARD)/cmos/lib$(CMOS).a $(CMOSOBJ)

#../output/boards/$(BOARD)/cmos/%.o: pkg/mpp/cmos/$(FAMILY)/$(CMOS)/%.c
#	[ -d ../output/boards/$(BOARD)/cmos ] || mkdir -p ../output/boards/$(BOARD)/cmos
#	$(CC) $(CFLAGS) -I$(abspath sdk/$(FAMILY)/include) -c $< -o $@

generate: 
	GOPATH=$(GOPATH) $(GO) generate -tags "$(FAMILY)" application/pkg/ko
	GOPATH=$(GOPATH) $(GO) generate -tags "$(FAMILY)" application/pkg/utils/regs
	#GOPATH=$(GOPATH) $(GO) generate -tags "$(FAMILY) openapi" application/pkg/openapi
	#GOPATH=$(GOPATH) $(GO) generate -tags "hi3516av100" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516av200" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516cv100" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516cv200" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516cv300" application/pkg/koloader
	#GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516cv500" application/pkg/koloader
	##GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3516ev200" application/pkg/koloader
	##GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3519av100" application/pkg/koloader
	##GOPATH=$(GOPATH) $(GO) generate -tags "generate hi3559av100" application/pkg/koloader

deploy-dir-daemon:
	rm -rf ./distrib
	mkdir ./distrib/
	mkdir ./distrib/$(FAMILY)
	mkdir ./distrib/$(FAMILY)/opt
	mkdir ./distrib/$(FAMILY)/opt/www
	mkdir ./distrib/$(FAMILY)/opt/storage
	mkdir ./distrib/$(FAMILY)/opt/scripts
	cp -r ./www/* ./distrib/$(FAMILY)/opt/www
	cp -r ./scripts/* ./distrib/$(FAMILY)/opt/scripts
	mkdir ./distrib/$(FAMILY)/etc
	cp ./init/run.sh ./distrib/$(FAMILY)/opt
	cp ./testvideo.h264 ././distrib/$(FAMILY)/opt
	cp ./testvideo.mp4 ././distrib/$(FAMILY)/opt
	mkdir ./distrib/$(FAMILY)/etc/init.d
	cp ./init/S99daemon ./distrib/$(FAMILY)/etc/init.d/
	cp ../boards/$(BOARD)/config ./distrib/$(FAMILY)/opt/board.config
	mkdir ./distrib/$(FAMILY)/opt/configs
	cp ../application/configs/* ./distrib/$(FAMILY)/opt/configs/

deploy-dir-probe:
	rm -rf ./distrib
	mkdir ./distrib/
	mkdir ./distrib/$(FAMILY)
	mkdir ./distrib/$(FAMILY)/opt
	mkdir ./distrib/$(FAMILY)/etc
	mkdir ./distrib/$(FAMILY)/etc/init.d
	cp ./init/S99gohisiprobe ./distrib/$(FAMILY)/etc/init.d/
	cp ../boards/$(BOARD)/config ./distrib/$(FAMILY)/opt/board.config
	mkdir ./distrib/$(FAMILY)/opt/configs
	cp ../application/configs/* ./distrib/$(FAMILY)/opt/configs/

TRIMPATH :=/home/nikitos1550/work/hi3519v101_go/application

build-daemon: generate libsdk ../output/boards/$(BOARD)/cmos/lib_cmoses.a testvideo.mp4 testvideo.h264 deploy-dir-daemon
	$(eval GOTAGS := $(GOTAGS_DAEMON))
	#$(GOENV) $(CGOPARAMS) $(GO) build -a -tags "$(GOTAGS_DAEMON)" -ldflags $(EMBEDDVARS) -o $(APPOUT)/daemon application/cmd/daemon
	$(GOENV) $(CGOPARAMS_DAEMON) $(GO) build -gcflags=all=-trimpath=$(TRIMPATH) -asmflags=all=-trimpath=$(TRIMPATH) -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" -o ./distrib/$(FAMILY)/opt/application application/cmd/daemon
	#$(GOENV) $(CGOPARAMS_DAEMON) $(GO) tool nm -size ./distrib/$(FAMILY)/opt/application | grep runtime.pclntab

build-probe: generate libsdk deploy-dir-probe
	$(eval GOTAGS := $(GOTAGS_PROBE))
	#$(GOENV) $(CGOPARAMS) $(GO) build -tags "$(GOTAGS_TESTER)" -ldflags $(EMBEDDVARS) -o $(APPOUT)/tester application/cmd/tester
	#$(GOENV) $(CGOPARAMS_TESTER) $(GO) build -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" -o ./distrib/$(FAMILY)/opt/application application/cmd/tester
	$(GOENV) $(CGOPARAMS_PROBE) $(GO) build -gcflags=all=-trimpath=$(TRIMPATH) -asmflags=all=-trimpath=$(TRIMPATH) -a -tags "$(GOTAGS)" -o ./distrib/$(FAMILY)/opt/gohisiprobe application/cmd/gohisiprobe

build-object:
	#$(eval GOTAGS := $(GOTAGS_DAEMON))
	#$(GOENV) $(CGOPARAMS_DAEMON) $(GO) build -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" \
    #    -o ./distrib/$(FAMILY)/_cache/cmos.a \
    #    application/pkg/mpp/cmos/hi3516cv300/imx290
	$(GOENV) $(CGOPARAMS_DAEMON) $(GO) tool nm ./distrib/$(FAMILY)/opt/application

#buid-with-vendors: generate
#	$(GOENV) $(CGOPARAMS) $(GO) build -a -tags $(GOTAGS) -mod vendor -o $(APPOUT) application/cmd/daemon

testvideo.h264:
	ffmpeg -f lavfi -i testsrc=size=320x340 -t 30 -pix_fmt yuv420p -c:v libx264 -tune zerolatency -profile:v main -g 25 -vbsf h264_mp4toannexb -f h264 testvideo.h264

testvideo.h265:
	ffmpeg -f lavfi -i testsrc=size=320x340 -t 30 -pix_fmt yuv420p -c:v libx265 -tune zerolatency -g 25 -f hevc testvideo.h265

testvideo.mp4: testvideo.h264
	MP4Box -add testvideo.h264 -fps 25 testvideo.mp4

clean:
	rm -f $(APPOUT)
	rm ./testvideo.h264

freeze-vendors:
	GOPATH=$(GOPATH) $(GO) mod vendor

install-deps:
	GOPATH=$(GOPATH) $(GO) get -u golang.org/x/tools/cmd/goyacc
	GOPATH=$(GOPATH) $(GO) get -u github.com/shuLhan/go-bindata
	go get -u github.com/shuLhan/go-bindata

apidoc:
	$(APIDOCOPENAPI) -p ./api/apidoc.json -s ./pkg -o ./api/openapi_extracted.json -v
	cd api; ./openapi.sh
	$(APIDOC) -c ./api/apidoc.json -i ./pkg -o ./api/docs-openapi
	$(REDOC) bundle ./api/openapi.json --options ./api/redoc.conf
	mv redoc-static.html ./api/
	$(REDOC) bundle ./api/openhisiipcam.yaml --options ./api/redoc.conf
	mv redoc-static.html ./api/redoc-static2.html

tinygo:
	GOPATH=$(GOPATH) $(GO) get -d -u github.com/tinygo-org/tinygo
