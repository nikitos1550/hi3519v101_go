include ../../Makefile.params

SOURCES = $(wildcard *.c) $(wildcard hi3516av200/*.c)
OBJECTS = $(patsubst %.c, %.o, $(SOURCES))
LIBS    = $(wildcard hi3516av200/lib/*.a)

#KO = $(wildcard hi3516av200/ko/*.o)
KO = $(shell ./ko.sh) #TODO make it pretty way & inline

libhisi.a: clean $(OBJECTS) ko
	$(AR) rcsT libhisi.a $(OBJECTS) $(KO) $(LIBS)

./%.o: ./%.c
	$(CC) $(CFLAGS) -c $< -o $@

ko:
	make -C hi3516av200/ko

testapps: libhisi.a
	$(CC) $(CFLAGS) test/test_hisi.c libhisi.a -o test_c.out
	CGO_CFLAGS='$(CFLAGS)' \
    CGO_CFLAGS_ALLOW='.*' \
    GOOS=linux \
    GOARCH=arm \
    GOARM=7 \
    CGO_ENABLED=1 \
    $(GO) build -o test_go.out \
        -a $(GOFLAGS) \
        test/test_hisi.go

clean:
	rm -f *.out
	rm -f *.o
	rm -f *.a
	rm -f */*.o
	make -C */ko clean
