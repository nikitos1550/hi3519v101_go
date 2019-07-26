THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

BUILDROOT := buildroot-2019.05.1-debug
TOOLCHAIN := arm-buildroot-linux-uclibcgnueabi
#TOOLCHAIN := arm-buildroot-linux-musleabi

LD := $(THIS_DIR)/../../../$(BUILDROOT)/output/host/bin/$(TOOLCHAIN)-ld

#all: $(objects)

SOURCES = $(wildcard ./*.ko)
OBJECTS = $(patsubst ./%.ko, ./%.o, $(SOURCES))

test: $(OBJECTS)
	echo $(SOURCES)
	echo $(OBJECTS)

./%.o: ./%.ko
	$(LD) -r -b binary $< -o $@

#list: ./*.ko
#	$(LD) -r -b binary $@ -o $@.o

#.PHONY: *.ko

clean:
	rm ./*.o
