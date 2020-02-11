ifndef PARAMS_FILE
    $(error 'PARAMS_FILE' variable must be defined)
endif
include $(PARAMS_FILE)

GOPATH := $(abspath ../output/go_workspace)

GOOS    := linux
GOARCH  := arm
GOARM   := $(ARMV) #FROM Makefile.params
GOENV   := GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM)

GOTAGS_TESTER  := $(FAMILY) tester
GOTAGS_DAEMON  := $(FAMILY) imx274 openapi scripts debug netgo
#GOTAGS_DAEMON  := $(FAMILY) imx274 debug
#GOTAGS  := $(FAMILY) netgo openapi debug lua onvif

PATH    := $(PATH):$(GOPATH)/bin

APPOUT  := $(abspath ../output/application)

CGO_CFLAGS  := -I$(abspath sdk/$(FAMILY)/include)
CGO_LDFLAGS_DAEMON :=   $(abspath ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a) \
                        $(abspath ../output/boards/$(BOARD)/cmos/lib$(CMOS).a)

CGO_LDFLAGS_TESTER :=   $(abspath ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a)

#CGO_LDFLAGS := $(abspath ../output/hisilicon/$(FAMILY)/lib$(FAMILY).a) $(abspath ../output/boards/$(BOARD)/cmos/lib$(CMOS).a) -ldl -lm

CGOPARAMS_TESTER := CGO_CFLAGS='$(CFLAGS) $(CGO_CFLAGS)' CGO_CFLAGS_ALLOW='.*' CGO_ENABLED=1 CGO_LDFLAGS='$(CGO_LDFLAGS_TESTER)'
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

CMOSSRC := $(wildcard pkg/mpp/cmos/$(FAMILY)/$(CMOS)/*.c)
CMOSOBJ := $(patsubst pkg/mpp/cmos/$(FAMILY)/$(CMOS)/%.c, ../output/boards/$(BOARD)/cmos/%.o, $(CMOSSRC))

libcmos: $(CMOSOBJ)
	@echo $(CMOSOBJ)
	$(AR) rcs ../output/boards/$(BOARD)/cmos/lib$(CMOS).a $(CMOSOBJ)

../output/boards/$(BOARD)/cmos/%.o: pkg/mpp/cmos/$(FAMILY)/$(CMOS)/%.c
	[ -d ../output/boards/$(BOARD)/cmos ] || mkdir -p ../output/boards/$(BOARD)/cmos
	$(CC) $(CFLAGS) -I$(abspath sdk/$(FAMILY)/include) -c $< -o $@

generate:
	GOPATH=$(GOPATH) $(GO) generate -tags "$(FAMILY)" application/pkg/koloader
	GOPATH=$(GOPATH) $(GO) generate -tags "$(FAMILY)" application/pkg/utils/regs
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
	mkdir ./distrib/$(FAMILY)/opt/scripts
	cp -r ./www/* ./distrib/$(FAMILY)/opt/www
	cp -r ./scripts/* ./distrib/$(FAMILY)/opt/scripts
	mkdir ./distrib/$(FAMILY)/etc
	cp ./init/run.sh ./distrib/$(FAMILY)/opt
	cp ./testvideo.h264 ././distrib/$(FAMILY)/opt
	mkdir ./distrib/$(FAMILY)/etc/init.d
	cp ./init/S99daemon ./distrib/$(FAMILY)/etc/init.d/

deploy-dir-tester:
	rm -rf ./distrib
	mkdir ./distrib/
	mkdir ./distrib/$(FAMILY)
	mkdir ./distrib/$(FAMILY)/opt
	mkdir ./distrib/$(FAMILY)/etc
	mkdir ./distrib/$(FAMILY)/etc/init.d
	cp ./init/S99tester ./distrib/$(FAMILY)/etc/init.d/

build-daemon: generate libsdk libcmos testvideo.h264 deploy-dir-daemon
	$(eval GOTAGS := $(GOTAGS_DAEMON))
	#$(GOENV) $(CGOPARAMS) $(GO) build -a -tags "$(GOTAGS_DAEMON)" -ldflags $(EMBEDDVARS) -o $(APPOUT)/daemon application/cmd/daemon
	$(GOENV) $(CGOPARAMS_DAEMON) $(GO) build -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" -o ./distrib/$(FAMILY)/opt/application application/cmd/daemon

build-tester: generate libsdk deploy-dir-tester
	$(eval GOTAGS := $(GOTAGS_TESTER))
	#$(GOENV) $(CGOPARAMS) $(GO) build -tags "$(GOTAGS_TESTER)" -ldflags $(EMBEDDVARS) -o $(APPOUT)/tester application/cmd/tester
	$(GOENV) $(CGOPARAMS_TESTER) $(GO) build -a -tags "$(GOTAGS)" -ldflags "$(GODEBUG) $(EMBEDDVARS)" -o ./distrib/$(FAMILY)/opt/application application/cmd/tester

#buid-with-vendors: generate
#	$(GOENV) $(CGOPARAMS) $(GO) build -a -tags $(GOTAGS) -mod vendor -o $(APPOUT) application/cmd/daemon

testvideo.h264:
	ffmpeg -f lavfi -i testsrc=size=320x340 -t 30 -pix_fmt yuv420p -c:v libx264 -tune zerolatency -profile:v main -g 25 -vbsf h264_mp4toannexb -f h264 testvideo.h264

clean:
	rm -f $(APPOUT)
	rm ./testvideo.h264

freeze-vendors:
	GOPATH=$(GOPATH) $(GO) mod vendor

install-deps:
	GOPATH=$(GOPATH) $(GO) get -u golang.org/x/tools/cmd/goyacc
