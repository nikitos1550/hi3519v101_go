THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

CC := $(THIS_DIR)/../../buildroot-2019.05.1-debug/output/host/bin/arm-buildroot-linux-uclibcgnueabi-gcc
AR := $(THIS_DIR)/../../buildroot-2019.05.1-debug/output/host/bin/arm-buildroot-linux-uclibcgnueabi-ar
LD := $(THIS_DIR)/../../buildroot-2019.05.1-debug/output/host/bin/arm-buildroot-linux-uclibcgnueabi-ld
RANLIB := $(THIS_DIR)/../../buildroot-2019.05.1-debug/output/host/bin/arm-buildroot-linux-uclibcgnueabi-ranlib

GO := $(THIS_DIR)/../../buildroot-2019.05.1-debug/output/host/bin/go

CFLAGS := -fPIC -mcpu=cortex-a7 -mfloat-abi=softfp -mfpu=neon-vfpv4 -mno-unaligned-access -fno-aggressive-loop-optimizations
#CFLAGS := -fPIC

all: libhimpp3 himpp3

libhimpp3:
	$(CC) $(CFLAGS) -c ./himpp3.c -o himpp3.o
	$(AR) rc libhimpp3.a himpp3.o
	rm ./himpp3.o
	$(RANLIB) libhimpp3.a

himpp3:
	CGO_CFLAGS='$(CFLAGS)' CGO_CFLAGS_ALLOW='.*' GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 $(GO) build -ldflags="-s -w" himpp3.go

LIBS=   ${THIS_DIR}/../../mpp_hi3519_v101/lib/libsns_imx274.a \
		${THIS_DIR}/../../mpp_hi3519_v101/lib/libmpi.a \
        ${THIS_DIR}/../../mpp_hi3519_v101/lib/libVoiceEngine.a \
        ${THIS_DIR}/../../mpp_hi3519_v101/lib/libisp.a \
        ${THIS_DIR}/../../mpp_hi3519_v101/lib/lib_hiae.a \
        ${THIS_DIR}/../../mpp_hi3519_v101/lib/lib_hiawb.a \
        ${THIS_DIR}/../../mpp_hi3519_v101/lib/lib_hiaf.a \
        ${THIS_DIR}/../../mpp_hi3519_v101/lib/libupvqe.a \
        ${THIS_DIR}/../../mpp_hi3519_v101/lib/libdnvqe.a \
        ${THIS_DIR}/../../mpp_hi3519_v101/lib/lib_hidefog.a

test:
	$(CC) -DTEST_C_APP $(CFLAGS) ./himpp3.c -o ./himpp3-test-c-app $(LIBS)

clean:
	rm ./himpp3-test-c-app
	rm ./libhimpp3.a
	rm ./himpp3.o
