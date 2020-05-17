#CF := -std=c89

test:
	cd u-boot-2016.11; \
		export PATH=$(abspath ../../output/hisilicon/hi3519av100/toolchain/bin):$(PATH); \
		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-gnueabi- ;
	cd u-boot-2016.11; \
        export PATH=$(abspath ../../output/hisilicon/hi3519av100/toolchain/bin):$(PATH); \
		make ARCH=arm CROSS_COMPILE=arm-buildroot-linux-gnueabi- u-boot-z.bin;
	#cp u-boot-2016.11/u-boot.bin ./u-boot.bin
	#chmod -x ./u-boot.bin
	#./mkboot.sh ./reg.bin ./u-boot-ok.bin

clean:
	rm -f ./u-boot.bin
	rm -f ./u-boot-ok.bin
