#CF := -std=c89
#PLATFORM_RELFLAGS="-std=gnu90"

#test:
#	cp ./reg.bin ./u-boot-2016.11/.reg
#	cd u-boot-2016.11; \
#        export PATH=$(abspath ../../output/hisilicon/hi3516cv500/toolchain/bin):$(PATH); \
#        make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-gnueabi- distclean; \
#        make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-gnueabi- hi3516av300_defconfig; \
#        make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-gnueabi- ; \
#        make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-gnueabi- u-boot-z.bin;
#	cp u-boot-2016.11/u-boot-hi3516av300.bin ./u-boot-ok.bin

test:
	cp ./reg.bin ./u-boot-2016.11/.reg
	cd u-boot-2016.11; \
		export PATH=$(abspath ../../output/hisilicon/hi3516ev200/toolchain/bin):$(PATH); \
		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-uclibcgnueabi- distclean; \
		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-uclibcgnueabi- hi3516ev200_defconfig; \
		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-uclibcgnueabi- ; \
		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-uclibcgnueabi- u-boot-z.bin;
	cp u-boot-2016.11/u-boot-hi3516ev200.bin ./u-boot-ok.bin

UBOOT_SIZE := $(shell stat -c %s ./u-boot-ok.bin)
#REST_SIZE  := $(shell expr 16777216 - $(UBOOT_SIZE))
REST_SIZE  := $(shell expr 8388608 - $(UBOOT_SIZE))

burn:
	echo $(UBOOT_SIZE)
	echo $(REST_SIZE)
	dd if=/dev/zero of=./rest.bin bs=1 count=$(REST_SIZE)
	cat ./u-boot-ok.bin ./rest.bin > ./image.bin
	#../../scripts/power2.py --num 8 off
	#flashrom -p serprog:dev=/dev/ttyPROG8:2000000 -c W25Q128.V --layout ./layout --image uboot -w ./image.bin
	#flashrom -p serprog:dev=/dev/ttyPROG8:2000000 -c W25Q128.V -w ./image.bin

	#flashrom -p serprog:dev=/dev/ttyPROG2:2000000 -c W25Q64.V --layout ./layout --image uboot -w ./image.bin
	#../../scripts/power2.py --num 8 on


clean:
	rm -f ./u-boot.bin
	rm -f ./u-boot-ok.bin
