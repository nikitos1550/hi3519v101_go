ifndef PARAMS_FILE
    $(error 'PARAMS_FILE' variable must be defined)
endif
include $(PARAMS_FILE)

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
