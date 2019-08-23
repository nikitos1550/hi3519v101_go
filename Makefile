THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

#THIS IS YOUR CUSTOM SETTINGS
#APP := app_sample       # TARGET APPLICATION
APP := app_minimal
CAMERA := 1             # NUMBER OF CAMERA ATTACHED TO SERVER TO TEST ON

#ROOTFS := romfs
ROOTFS := squashfs

CAMERA_LOCAL_LOAD_IP := 192.169.0.102 #ONLY FOR LOCAL USAGE, SERVER DOESN'T USE IT
########################################################################
BUILDROOT := buildroot-2019.05.1-debug
#BUILDROOT := buildroot-2019.05.1-musl
########################################################################
#buildroot-2019.05.1-toolchain:
#	tar -xzf buildroot-2019.05.1.tar.gz -C $(THIS_DIR)
#	mv buildroot-2019.05.1 buildroot-2019.05.1-toolchain
#	cd buildroot-2019.05.1-toolchain; rm dl; ln -s ../buildroot-dl dl
#	cd buildroot-2019.05.1-toolchain; make defconfig BR2_DEFCONFIG=$(THIS_DIR)/defconfig-toolchain.buildroot
#	cp -r ./buildroot-2019.05.1-patch/* ./buildroot-2019.05.1-debug

environment-buildroot-2019.05.1-debug:
	tar -xzf buildroot-2019.05.1.tar.gz -C $(THIS_DIR)
	mv buildroot-2019.05.1 buildroot-2019.05.1-debug
	test -e buildroot-dl || mkdir buildroot-dl
	cd buildroot-2019.05.1-debug; ln -s ../buildroot-dl dl
	cd buildroot-2019.05.1-debug; make defconfig BR2_DEFCONFIG=$(THIS_DIR)/defconfig-debug.buildroot
	cp -r ./buildroot-2019.05.1-patch/* ./buildroot-2019.05.1-debug

enviroment-buildroot-2019.05.1-debug-update-config:
	cd buildroot-2019.05.1-debug; make defconfig BR2_DEFCONFIG=$(THIS_DIR)/defconfig-debug.buildroot

#buildroot-2019.05.1-release:
#	test -e buildroot-2019.05.1-release || tar -xzf buildroot-2019.05.1.tar.gz -C $(THIS_DIR)
#	mv buildroot-2019.05.1 buildroot-2019.05.1-release
#	cd buildroot-2019.05.1-release; rm dl; ln -s ../buildroot-dl dl
#	cd buildroot-2019.05.1-release; make defconfig BR2_DEFCONFIG=$(THIS_DIR)/defconfig-release.buildroot
#	cp -r ./buildroot-2019.05.1-patch/* ./buildroot-2019.05.1-release

environment-deploy-debug: environment-buildroot-2019.05.1-debug
	cd buildroot-2019.05.1-debug; make clean; make

########################################################################

app-build-debug:
	cd $(APP); make
	cd $(APP); cp ./$(APP) ../putonrootfs-debug/opt
	cd $(BUILDROOT); make rootfs-$(ROOTFS)
	cp $(BUILDROOT)/output/images/rootfs.$(ROOTFS) ./burner/images

app-deploy-debug-server: app-build-debug deploy-no-build

app-deploy-debug-local: app-build-debug
	cd burner; \
		sudo ./burner.py \
				load \
				--port /dev/ttyCAM$(CAMERA) \
				--uimage ./images/uImage \
				--rootfs ./images/rootfs.$(ROOTFS) \
				--ip $(CAMERA_LOCAL_LOAD_IP) \
				--skip 1024 \
				--memory 96

deploy-no-build:
	cd burner; \
		authbind --deep ./burner.py \
			load \
			--port /dev/ttyCAM$(CAMERA) \
			--uimage ./images/uImage \
			--rootfs ./images/rootfs.$(ROOTFS) \
			--ip 192.169.0.10$(CAMERA) \
			--skip 1024 \
			--memory 96 \
			--servercamera $(CAMERA)
	screen -L /dev/ttyCAM$(CAMERA) 115200

########################################################################
camera-serial:
	screen -L /dev/ttyCAM$(CAMERA) 115200

camera-serial-1:
	screen -L /dev/ttyCAM1 115200
camera-serial-2:
	screen -L /dev/ttyCAM2 115200

