THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

#THIS IS YOUR CUSTOM SETTINGS
#APP := app_sample       # TARGET APPLICATION
APP := app_minimal
CAMERA := 1             # NUMBER OF CAMERA ATTACHED TO SERVER TO TEST ON

CAMERA_LOCAL_LOAD_IP := 192.168.0.200 #ONLY FOR LOCAL USAGE, SERVER DOESN'T USE IT
########################################################################

-include ./boards/$(BOARD)/config

ttest:
	@[ "$(BOARD)" ] && echo "all good" || ( echo "var is not set"; exit 1 )
	@echo "FAMILY = $(FAMILY)"
	@echo "CAMERA FACILITY ID = $(CAMERA)"
	test -e boards/$(BOARD)/build || mkdir boards/$(BOARD)/build
	cp $(FAMILY)/kernel/uImage boards/$(BOARD)/build/uImage
	rm -f boards/$(BOARD)/build/rootfs.squashfs
	rm -rf boards/$(BOARD)/build/rootfs.tmp; mkdir boards/$(BOARD)/build/rootfs.tmp
	cp -r $(FAMILY)/rootfs/target/* boards/$(BOARD)/build/rootfs.tmp/
	test ! -e boards/$(BOARD)/putonrootfs || cp -r boards/$(BOARD)/putonrootfs/* boards/$(BOARD)/build/rootfs.tmp/
	mksquashfs  boards/$(BOARD)/build/rootfs.tmp \
                boards/$(BOARD)/build/rootfs.squashfs \
                -all-root
                #-info
	cp boards/$(BOARD)/build/uImage burner/images/uImage
	cp boards/$(BOARD)/build/rootfs.squashfs burner/images/rootfs.squashfs
	#cp $(FAMILY)/rootfs/images/rootfs.squashfs burner/images/rootfs.squashfs
	exit
	cd burner; \
        authbind --deep ./burner.py \
            load \
            --port /dev/ttyCAM$(CAMERA) \
            --uimage ./images/uImage \
            --rootfs ./images/rootfs.squashfs \
            --ip 192.169.0.10$(CAMERA) \
            --skip 1024 \
            --initrd 4 \
            --memory $(MEM_LINUX) \
            --servercamera $(CAMERA)
	screen -L /dev/ttyCAM$(CAMERA) 115200


########################################################################
enviroiment-buildroot-2019.08:
	tar -xzf buildroot-2019.08.tar.gz -C $(THIS_DIR)
	cp -r ./buildroot-2019.08-patch/* ./buildroot-2019.08

########################################################################

app-build-debug:
	#cd $(APP); make
	#cd $(APP); cp ./$(APP) ../putonrootfs/opt
	cd buildroot-2019.05.1-debug; make
	cp buildroot-2019.05.1-debug/output/images/rootfs.squashfs ./burner/images

app-deploy-debug-server: app-build-debug deploy-no-build

app-deploy-debug-local: app-build-debug
	cd burner; \
		sudo ./burner.py load --uimage ./images/uImage --rootfs ./images/rootfs.squashfs --ip $(CAMERA_LOCAL_LOAD_IP) --skip 1024 --memory 96

deploy-no-build:
	cd burner; \
		authbind --deep ./burner.py \
			load \
			--uimage ./images/uImage \
			--rootfs ./images/rootfs.squashfs \
			--ip 192.169.0.10$(CAMERA) \
			--skip 1024 \
			--memory 96 \
			--servercamera $(CAMERA)
	screen /dev/ttyCAM$(CAMERA) 115200

########################################################################
camera-serial:
	screen -L /dev/ttyCAM$(CAMERA) 115200

camera-serial-7:
	screen -L /dev/ttyCAM7 115200
########################################################################


#APP
#BOARD

#newbuild:
