THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

ifeq ("$(wildcard Makefile.user.params)","")
 $(error cp Makefile.user.params.example to Makefile.user.params) 
endif

include $(THIS_DIR)Makefile.user.params

BR             := buildroot-2019.08
BUILDROOT_DIR  := $(abspath ./buildroot-2019.08)
BOARD_OUTDIR   := $(abspath ./output/boards/$(BOARD))
CAMERA_TTY     := /dev/ttyCAM$(CAMERA)
CAMERA_IP      := 192.168.10.1$(shell printf '%02d' $(CAMERA))
TELNET_PORT    := 453$(shell printf '%02d' $(CAMERA))

########################################################################

CAMSTORE       := $(THIS_DIR)/facility/camstore/control.sh client

########################################################################

APP             := application
APP_TARGET      ?= tester   #default target will be tester, daemon build on request durin it`s early dev stage

-include ./boards/$(BOARD)/config

.PHONY: $(APP)/distrib/$(FAMILY) help prepare cleanall

# -----------------------------------------------------------------------------------------------------------

help:
	@echo -e "Help:\n" \
		" - make prepare                          - prepare; MUST be done once before anything\n"\
		" - make deploy-app                       - build&deploy application onto particular board\n"\
		" - make deploy-app-control-[uart|telnet] - build&deploy application, then attach control console onto particular board\n"\
		" - make control-[uart|telnet]            - attach control console onto particular board\n"\
		" - make rootfs.squashfs                  - build application and pack it within RootFS image\n"\
		" - make kernel                           - build board kernel\n"\
		" - make cleanall                         - remove all artifacts"

prepare: $(BUILDROOT_DIR)
	@echo "All prepared"

$(BUILDROOT_DIR):
	tar -xzf $(BR).tar.gz -C $(THIS_DIR)
	cp -r ./$(BR)-patch/* $(BUILDROOT_DIR)/

cleanall:
	if [ -d ./output ]; then chmod --recursive 777 ./output; fi
	rm -rf ./output $(BUILDROOT_DIR)

# -----------------------------------------------------------------------------------------------------------

rootfs.squashfs: $(BOARD_OUTDIR)/rootfs+app.squashfs
	@echo "--- RootFS image is ready: $^"

kernel: $(BOARD_OUTDIR)/kernel/uImage
	@echo "--- Kernel uImage is ready: $^"

$(BOARD_OUTDIR)/rootfs+app.squashfs: $(BOARD_OUTDIR)/rootfs+app
	rm -f $@
	mksquashfs $< $@ -all-root

$(BOARD_OUTDIR)/rootfs+app: $(BOARD_OUTDIR)/rootfs $(APP)/distrib/$(FAMILY)
	rm -rf $@; mkdir -p $@
	cp -r $(BOARD_OUTDIR)/rootfs/* $@/
	cp -r $(APP)/distrib/$(FAMILY)/* $@/

$(APP)/distrib/$(FAMILY): $(BOARD_OUTDIR)/Makefile.params
	rm -rf $@
	make -C $(APP) PARAMS_FILE=$< build-$(APP_TARGET)

# -----------------------------------------------------------------------------------------------------------
# Board's artifacts

$(BOARD_OUTDIR)/rootfs:
	make -C ./boards BOARD_OUTDIR=$(BOARD_OUTDIR) BOARD=$(BOARD) rootfs

$(BOARD_OUTDIR)/Makefile.params:
	make -C ./boards BOARD_OUTDIR=$(BOARD_OUTDIR) BOARD=$(BOARD) toolchain

$(BOARD_OUTDIR)/kernel/uImage:
	make -C ./boards BOARD_OUTDIR=$(BOARD_OUTDIR) BOARD=$(BOARD) kernel

# ====================================================================================================================

build-app: $(APP)/distrib/$(FAMILY)

pack-app: $(BOARD_OUTDIR)/rootfs+app.squashfs $(BOARD_OUTDIR)/kernel/uImage

deploy-app: pack-app
	cd burner; authbind --deep ./burner2.py \
        --log-level DEBUG \
		--mode camstore \
		--port /dev/ttyCAM$(CAMERA) \
		--reset-power "./power2.py --num $(CAMERA) reset" \
		load \
		--target-ip $(CAMERA_IP) --iface enp3s0 \
		--uimage $(BOARD_OUTDIR)/kernel/uImage \
		--rootfs $(BOARD_OUTDIR)/rootfs+app.squashfs \
		--initrd-size 16M --memory-size 256M

deploy-app-control-uart: deploy-app control-uart

deploy-app-control-telnet: deploy-app
	@echo "waiting for 10s"
	@sleep 3
	@echo "7s more..."
	@sleep 5
	@echo "be patient, 2s more"
	telnet $(CAMERA_IP)


########################################################################

deprecated-control-uart:
	miniterm $(CAMERA_TTY) 115200

control-uart:
	#telnet localhost $(TELNET_PORT)
	$(CAMSTORE) forward_serial $(CAMERA_TTY)

deprecated-control-uart-%:
	miniterm /dev/ttyCAM$(subst control-uart-,,$@) 115200

control-uart-%:
	#telnet localhost 453$(shell printf '%02d' $(subst control-uart-,,$@))
	$(CAMSTORE) forward_serial /dev/ttyCAM$(subst control-uart-,,$@)

control-telnet:
	telnet $(CAMERA_IP)

control-telnet-%:
	telnet 192.168.10.1$(shell printf '%02d' $(subst control-telnet-,,$@))

