# hi3519v101_go
Hi3519v101+IMX274 Golang enviroiment

## Chips families information

| chips                                                 | shortcode     |
|
| hi3516av100, hi3516dv100                              | hi3516av100   |
| hi3519v101,  hi3516av200                              | hi3516av200   |
| hi3516cv100, hi3518cv100, hi3518ev100                 | hi3516cv100   |
| hi3516cv200, hi3518ev200, hi3518ev201                 | hi3516cv200   |
| hi3516cv300, hi3516ev100                              | hi3516cv300   |
| hi3516cv500, hi3516dv300, hi3516av300                 | hi3516cv500   |
| hi3519av100                                           | hi3519av100   |
| hi3559av100                                           | hi3559av100   |
| hi3516ev300, hi3516ev200, hi3516dv200, hi3518ev300    | hi3516ev200   |

## Repo structure

* **app_minimal** - target application, development happens here
* **app_sample** - sample application
* **buildroot-2019.05.1-debug** - toolchain, rootfs build dir
* **burner** - tool for automated upload software to camera
* **docs** - documentation files, image storage used by github wiki
* **facility** - files that are used on server to organize enviroiment
* **kernel** - kernel src
* **mpp_hi3519_v101** - HiMPP from SDK
* **putonrootfs-debug** - rootfs overlay
* **reg-tools** - himm utility, used for register configuration. From SDK.

## How to deploy debug enviroiment

```make enviroiment-deploy-debug```

This will do following:

1. unpack buildroot
2. create separate dir for sources that will downloaded
2. set buildroot settings according premade config
3. build buildroot (this may take some time...)

### Kernel

Kernel is already prebuilt, you can find it in burner/images (uBoot file).

If you want to rebuild it do following:

1. ```cd kernel```
2. ```make unpack```
3. ```make build```


## How to deploy software to camera attached to server

Before you start, please edit Makefile. In the bottom of file you will find:

```
#THIS IS YOUR CUSTOM SETTINGS
APP := app_sample       # TARGET APPLICATION
#APP := app_minimal
CAMERA := 1             # NUMBER OF CAMERA ATTACHED TO SERVER TO TEST ON

CAMERA_LOCAL_LOAD_IP := 192.168.0.200 #ONLY FOR LOCAL USAGE, SERVER DOESN'T USE IT
```

Set application name accroding that one you want to use.

Set camera number according free one, to prevent interference with other user.

Finally run ```make app-deploy-debug-server```

This command will do following:
1. Build application
2. Copy application to rootfs overlay (putonrootfs-debug dir)
3. Pack everything with buildroot
4. Copy built rootfs image to burner dir
5. Run burner.py to load software to camera
6. Start serial console attached to camera

## FAQ

Q: My rootfs.squashfs becomes too big!

A: Check *putonrootfs-debug/opt*, seems there are several apps. Clean dir. Check *buildroot-2019.05.1-debug/output/target/opt*, do same.

---

Q: How to exit ```screen``` serial terminal?

A: Ctrl+A ky (Press **Ctrl + A**, then press **k**, then press **y**)

---

Q: I'm tired of starting the application typing same commands each time.

A: Check *putonrootfs-debug/etc/init.d/S90debug*

```
...
start() {
	#gdbserver :2345 /opt/mpp_version &
	#/opt/app_sample & <--- UNCOMMENT THIS LINE
}
...
```

## Notes

* Don`t forget about git-lfs to clone this repo fully
* For clean ubuntu server 19.04 I installed mc build-essential make cmake u-boot-tools python libncurses5-dev packages
* hisi-build webserver basic auth - test/hisilicon
