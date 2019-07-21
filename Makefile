THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

#THIS IS YOUR CUSTOM SETTINGS
APP := app_sample       # TARGET APPLICATION
#APP := app_minimal
CAMERA := 1             # NUMBER OF CAMERA ATTACHED TO SERVER TO TEST ON

CAMERA_LOCAL_LOAD_IP := 192.168.0.200 #ONLY FOR LOCAL USAGE, SERVER DOESN'T USE IT
########################################################################

#buildroot-2019.05.1-toolchain:
#	tar -xzf buildroot-2019.05.1.tar.gz -C $(THIS_DIR)
#	mv buildroot-2019.05.1 buildroot-2019.05.1-toolchain
#	cd buildroot-2019.05.1-toolchain; rm dl; ln -s ../buildroot-dl dl
#	cd buildroot-2019.05.1-toolchain; make defconfig BR2_DEFCONFIG=$(THIS_DIR)/defconfig-toolchain.buildroot
#	cp -r ./buildroot-2019.05.1-patch/* ./buildroot-2019.05.1-debug

enviroiment-buildroot-2019.05.1-debug:
	tar -xzf buildroot-2019.05.1.tar.gz -C $(THIS_DIR)
	mv buildroot-2019.05.1 buildroot-2019.05.1-debug
	test -e buildroot-dl || makdir buildroot-dl
	cd buildroot-2019.05.1-debug; ln -s ../buildroot-dl dl
	cd buildroot-2019.05.1-debug; make defconfig BR2_DEFCONFIG=$(THIS_DIR)/defconfig-debug.buildroot
	cp -r ./buildroot-2019.05.1-patch/* ./buildroot-2019.05.1-debug

#buildroot-2019.05.1-release:
#	test -e buildroot-2019.05.1-release || tar -xzf buildroot-2019.05.1.tar.gz -C $(THIS_DIR)
#	mv buildroot-2019.05.1 buildroot-2019.05.1-release
#	cd buildroot-2019.05.1-release; rm dl; ln -s ../buildroot-dl dl
#	cd buildroot-2019.05.1-release; make defconfig BR2_DEFCONFIG=$(THIS_DIR)/defconfig-release.buildroot
#	cp -r ./buildroot-2019.05.1-patch/* ./buildroot-2019.05.1-release

enviroiment-deploy-debug: enviroiment-buildroot-2019.05.1-debug
	cd buildroot-2019.05.1-debug; make clean; make

########################################################################

app-build-debug:
	cd $(APP); make
	cd $(APP); cp ./$(APP) ../putonrootfs-debug/opt
	cd buildroot-2019.05.1-debug; make
	cp buildroot-2019.05.1-debug/output/images/rootfs.squashfs ./burner/images

app-deploy-debug-server: app-build-debug
	cd burner; \
		authbind --deep ./burner.py load --uimage ./images/uImage --rootfs ./images/rootfs.squashfs --ip 192.169.0.10$(CAMERA) --skip 1024 --memory 96 --servercamera $(CAMERA)
	screen /dev/ttyCAM$(CAMERA) 115200

app-deploy-debug-local: app-build-debug
	cd burner; \
		sudo ./burner.py load --uimage ./images/uImage --rootfs ./images/rootfs.squashfs --ip $(CAMERA_LOCAL_LOAD_IP) --skip 1024 --memory 96
