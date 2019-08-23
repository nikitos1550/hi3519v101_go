THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GO := $(THIS_DIR)../buildroot-2019.05.1-debug/output/host/bin/go
#GO := $(THIS_DIR)../buildroot-2019.05.1-musl/output/host/bin/go

CFLAGS := -std=gnu11 -mcpu=cortex-a7 -mfloat-abi=softfp -mfpu=neon-vfpv4 -mno-unaligned-access -fno-aggressive-loop-optimizations

#DNODEBUG -s -w
BUILDFLAGS := -ldflags="-extldflags '-static' -X 'main.BuildTime=__________1231231231'"
#BUILDFLAGS := -ldflags="-s -w -extldflags '-static'"

test:
	cd himpp3; make clean; make
	CGO_CFLAGS='$(CFLAGS)' CGO_CFLAGS_ALLOW='.*' GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 $(GO) build -a $(BUILDFLAGS) ./app_minimal.go

godeps:
	$(GO) get github.com/drone/routes

clean:
	cd himpp3; make clean
	rm -f ./app_minimal
