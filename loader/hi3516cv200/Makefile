#CF := -std=c89

#test:
#	cd u-boot-2010.06; \
#		export PATH=$(abspath ../../output/hisilicon/hi3516cv200/toolchain/bin):$(PATH); \
#		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-uclibcgnueabi- distclean; \
#		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-uclibcgnueabi- hi3518ev200_config; \
#		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-uclibcgnueabi- PLATFORM_RELFLAGS="-std=gnu90" ;
#	cp u-boot-2010.06/u-boot.bin ./u-boot.bin
#	chmod -x ./u-boot.bin
#	./mkboot.sh ./reg.bin ./u-boot-ok.bin

test:
	cd u-boot-2010.06; \
        export PATH=$(abspath ./gcc-linaro-arm-none-eabi-4.8-2014.04_linux/bin):$(PATH); \
        make ARCH=arm CROSS_COMPILE=arm-none-eabi- distclean; \
        make ARCH=arm CROSS_COMPILE=arm-none-eabi- hi3518ev200_config; \
        make ARCH=arm CROSS_COMPILE=arm-none-eabi- ;
	cp u-boot-2010.06/u-boot.bin ./u-boot.bin
	chmod -x ./u-boot.bin
	./mkboot.sh ./reg.bin ./u-boot-ok.bin


#test:
#	cd u-boot-2010.06; \
#        export PATH=$(abspath /opt/hisi-linux/x86-arm/arm-hisiv510-linux/bin):$(PATH); \
#        make ARCH=arm CROSS_COMPILE=arm-hisiv510-linux-uclibcgnueabi- distclean; \
#        make ARCH=arm CROSS_COMPILE=arm-hisiv510-linux-uclibcgnueabi- hi3518ev200_config; \
#        make ARCH=arm CROSS_COMPILE=arm-hisiv510-linux-uclibcgnueabi- PLATFORM_RELFLAGS="-std=gnu90" ;
#	cp u-boot-2010.06/u-boot.bin ./u-boot.bin
#	chmod -x ./u-boot.bin
#	./mkboot.sh ./reg.bin ./u-boot-ok.bin


#test:
#	cd u-boot-2010.06; \
#        export PATH=$(abspath /opt/hisi-linux/x86-arm/arm-hisiv200-linux/bin):$(PATH); \
#        make ARCH=arm CROSS_COMPILE=arm-hisiv200-linux-gnueabi- distclean; \
#        make ARCH=arm CROSS_COMPILE=arm-hisiv200-linux-gnueabi- hi3518ev200_config; \
#        make ARCH=arm CROSS_COMPILE=arm-hisiv200-linux-gnueabi- ;
#	cp u-boot-2010.06/u-boot.bin ./u-boot.bin
#	chmod -x ./u-boot.bin
#	./mkboot.sh ./reg.bin ./u-boot-ok.bin


UBOOT_SIZE := $(shell stat -c %s ./u-boot-ok.bin)
#REST_SIZE  := $(shell expr 16777216 - $(UBOOT_SIZE))
REST_SIZE  := $(shell expr 8388608 - $(UBOOT_SIZE))

burn:
	echo $(UBOOT_SIZE)
	echo $(REST_SIZE)
	dd if=/dev/zero of=./rest.bin bs=1 count=$(REST_SIZE)
	cat ./u-boot-ok.bin ./rest.bin > ./image.bin


clean:
	rm -f ./u-boot.bin
	rm -f ./u-boot-ok.bin
