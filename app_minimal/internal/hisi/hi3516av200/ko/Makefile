include ../../../../Makefile.params

SOURCES = $(wildcard ./*.ko)
OBJECTS = $(patsubst ./%.ko, ./%.o, $(SOURCES))

convert: $(OBJECTS)

./%.o: ./%.ko
	$(LD) -r -b binary $< -o $@

clean:
	rm -f ./*.o
	rm -f include.h
