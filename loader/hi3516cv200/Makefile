test:
	cd u-boot-2010.06; \
		export PATH=$(abspath ../../output/hisilicon/hi3516cv200/toolchain/bin):$(PATH); \
		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-uclibcgnueabi- clean; \
		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-uclibcgnueabi- hi3518ev201_config; \
		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-uclibcgnueabi- ;


#make ARCH=arm CROSS_COMPILE=arm-hisiv510-linux- hi3518ev201_config 
#make ARCH=arm CROSS_COMPILE=arm-hisiv510-linux- 
