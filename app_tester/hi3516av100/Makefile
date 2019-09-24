THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

AR := $(THIS_DIR)../../hi3516av100/buildroot/host/bin/arm-buildroot-linux-uclibcgnueabi-ar
CC := $(THIS_DIR)../../hi3516av100/buildroot/host/bin/arm-buildroot-linux-uclibcgnueabi-gcc

LIBS = $(wildcard lib/*.a)

libhisi.a: clean
	$(AR) rcsT libhi3516av100.a $(LIBS)

clean:
	rm -f libhi3516av100.a
