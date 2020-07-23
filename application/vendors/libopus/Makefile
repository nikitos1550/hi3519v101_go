download: opus-1.3.1

opus-1.3.1.tar.gz:
	wget https://archive.mozilla.org/pub/opus/opus-1.3.1.tar.gz

opus-1.3.1: opus-1.3.1.tar.gz
	tar -xzf opus-1.3.1.tar.gz

build:
	cd opus-1.3.1; \
		export PATH=/home/nikitos1550/work/testing/hi3519v101_go/output/hisilicon/hi3516av200/toolchain/bin:$(PATH); \
		./configure --host=arm-buildroot-linux-gnueabi --prefix=/home/nikitos1550/work/testing/hi3519v101_go/application/vendors/libopus/output; \
		make clean; \
		make -j8; \
		make install
