THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

include $(THIS_DIR)/../../../../$(FAMILY)/Makefile.params

SOURCES = $(wildcard ./*.ko)
OBJECTS = $(patsubst ./%.ko, ./%.o, $(SOURCES))

convert: strip $(OBJECTS)

strip:
	$(STRIP) --strip-debug *.ko

./%.o: ./%.ko
	$(LD) -r -b binary $< -o $@

clean:
	rm -f ./*.o
	rm -f include.h
