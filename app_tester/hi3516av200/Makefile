THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

AR := $(THIS_DIR)../../hi3516av200/buildroot/host/bin/arm-buildroot-linux-uclibcgnueabi-ar

LIBS = $(wildcard lib/*.a)

libhisi.a: clean
	$(AR) rcsT libhi3516av200.a $(LIBS)

clean:
	rm -f libhi3516av200.a
