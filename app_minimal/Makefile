THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GO := $(THIS_DIR)../buildroot-2019.05.1-debug/output/host/bin/go

CFLAGS := -mcpu=cortex-a7 -mfloat-abi=softfp -mfpu=neon-vfpv4 -mno-unaligned-access -fno-aggressive-loop-optimizations

test:
	cd himpp3; make
	CGO_CFLAGS='$(CFLAGS)' CGO_CFLAGS_ALLOW='.*' GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 $(GO) build ./app_minimal.go

clean:
	rm ./app_minimal
