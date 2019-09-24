THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

AR := $(THIS_DIR)../../hi3516cv100/buildroot/host/bin/arm-buildroot-linux-uclibcgnueabi-ar

LIBS = $(wildcard lib/*.a)

libhisi.a: clean
	$(AR) rcsT libhi3516cv100.a $(LIBS)

clean:
	rm -f libhi3516cv100.a
