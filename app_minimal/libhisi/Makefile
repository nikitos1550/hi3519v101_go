ifndef PARAMS_FILE
    $(error 'PARAMS_FILE' variable must be defined)
endif
include $(PARAMS_FILE)

#CHIPFAMILY = hi3516av200
#CHIPFAMILY = hi3516cv100 #TODO
#CHIPFAMILY = hi3516av100 #TODO

SOURCES = $(wildcard *.c) $(wildcard $(FAMILY)/*.c) $(wildcard $(FAMILY)/cmos/*/*.c)
OBJECTS = $(patsubst %.c, %.o, $(SOURCES))
LIBS    = $(wildcard $(FAMILY)/lib/*.a)

#KO = $(wildcard hi3516av200/ko/*.o)
KO = $(shell FAMILY=$(FAMILY) ./ko.sh) #TODO make it pretty way & inline

libhisi.a: clean $(OBJECTS) ko
	$(AR) rcsT libhisi.a $(OBJECTS) $(KO) $(LIBS)

./%.o: ./%.c
	$(CC) $(CFLAGS) -I$(FAMILY)/include -c $< -o $@

ko:
	make -C $(FAMILY)/ko -j4

#showko:
#	@echo $(KO)

doxygen:
	rm -rf ./docs
	doxygen

clean:
	rm -f *.out
	rm -f *.o
	rm -f *.a
	rm -f */*.o
	rm -f */cmos/*/*.o
	make -C hi3516av200/ko clean -j4
	#make -C hi3516cv100/ko clean
	#make -C hi3516av100/ko clean
