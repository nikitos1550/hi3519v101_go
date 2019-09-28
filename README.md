# OpenHisiIpCam development facility


## Chips families information

| chips                                                 | shortcode     | status |
|-------------------------------------------------------|---------------|--------------|
| hi3516av100, hi3516dv100                              | hi3516av100   | ++           |
| hi3519v101,  hi3516av200                              | hi3516av200   | ++++         |
| hi3516cv100, hi3518cv100, hi3518ev100                 | hi3516cv100   | +            |
| hi3516cv200, hi3518ev200, hi3518ev201                 | hi3516cv200   | +            |
| hi3516cv300, hi3516ev100                              | hi3516cv300   | +            |
| hi3516cv500, hi3516dv300, hi3516av300                 | hi3516cv500   | +            |
| hi3519av100                                           | hi3519av100   | n/a          |
| hi3559av100                                           | hi3559av100   | 0            |
| hi3516ev300, hi3516ev200, hi3516dv200, hi3518ev300    | hi3516ev200   | n/a          |

## Network structure

```
  213.141.129.12                                                                   
  :2223 ssh to build-hisi                                                          
  :1194 openvpn (internal 192.168.11.X)                                            
                                                                                   
    +---------------------+                                                        
    |                     |                                                        
    |    ROUTER Tp-Link   |                                                        
    |                     |                                                        
    +---------------------+                                                        
                         ---------nikita-home-computer                             
  192.168.10.X ---------/         192.168.10.3                                     
     |      -----\                                                                 
     |            ---------\                      AC/DC 220v -> 12v 250W           
     |                      ------build-hisi                |                      
  Camera slot #X                --server -\                 |                      
  192.168.10.1XX              -/           --\              |                      
     |                      -/                --            |                      
     /                    -/     ttyAMA0 arduino relay power resetter              
    |                   -/                                                         
    |                 -/                                                           
 pl2303 usb-uart adapter                                                           
```



## Repo structure

* **api**
* **app_minimal** - target application, development happens here
* **app_tester** - sample application
* **boards** - known devices configurations
* **burner** - tool for automated upload software to device
* **docs** - documentation files, image storage used by github wiki
* **facility** - files that are used on server to organize enviroiment
* **hi35XXXvXXX** - chip family specific toolchain, default rootfs and kernel build env
* **putonrootfs** - default rootfs overlay

## Boards

Current build system is targeted to specific boards to make it easy checking software on different devices.

| board name            | status                  |
|-----------------------|-------------------------|
| jvt_hi3519v101_imx274 | app_tester, app_minimal |
| xm_hi3516av100_imx178 | app_tester              |
| other                 | n/a                     |

### Sample board config

* config file
* kernel dir
  * kernel.config
  * patch

```
VENDOR      =JVT
MODEL       =unknown
FAMILY      =hi3516av200
CHIP        =hi3519v101
RAM_SIZE    =512
RAM_LINUX   =256
RAM_MPP     =256
ROM_SIZE    =16
CMOS        =imx274
UBOOT_SIZE  =1024
INITRD_TMP  =16
```


## How to deploy debug enviroiment

1. ```cp Makefile.params.example Makefile.params``` copy default config and edit it.

Here you can point to default camera you are working with.

2. ```make enviroiment-setup```

This will unpack buildroot and setup toolchains and default rootfses for all chips.
Also it will prepare kernel build env for all chips.
If you want to save time comment lines in *enviroiment-setup* target to exclude chips you are not working with.


### Kernel

Kernel should be built for each board. 
```make kernl-build``` command will build kernel for your board.


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

