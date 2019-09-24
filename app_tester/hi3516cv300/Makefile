THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

AR := $(THIS_DIR)../../hi3516cv300/buildroot/host/bin/arm-buildroot-linux-uclibcgnueabi-ar

LIBS = $(wildcard lib/*.a)

libhisi.a: clean
	$(AR) rcsT libhi3516cv300.a $(LIBS)

clean:
	rm -f libhi3516cv300.a
