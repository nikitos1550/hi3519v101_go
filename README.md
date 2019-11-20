# OpenHisiIpCam development facility

## Chips families information

| chips                                                 | shortcode     | status       |
|-------------------------------------------------------|---------------|--------------|
| hi3516av100, hi3516dv100                              | hi3516av100   | ++           |
| hi3519v101,  hi3516av200                              | hi3516av200   | ++++         |
| hi3516cv100, hi3518cv100, hi3518ev100                 | hi3516cv100   | +            |
| hi3516cv200, hi3518ev200, hi3518ev201                 | hi3516cv200   | +            |
| hi3516cv300, hi3516ev100                              | hi3516cv300   | +            |
| hi3516cv500, hi3516dv300, hi3516av300                 | hi3516cv500   | +            |
| hi3516ev300, hi3516ev200, hi3516dv200, hi3518ev300    | hi3516ev200   | n/a          |
| hi3519av100                                           | hi3519av100   | n/a          |
| hi3559av100                                           | hi3559av100   | 0            |

## SDK information

|family     |kernel |uboot  |MPP    |
|-----------|-------|-------|-------|
|hi3516av100|3.4.?  |       |v2
|hi3516av200|3.18.20|       |v3
|hi3516cv100|3.0.?  |       |v1/v2?
|hi3516cv200|3.4.?  |       |v2
|hi3516cv300|3.18.20|       |v3
|hi3516cv500|4.9.37 |       |v4
|hi3516ev200|4.9.37 |       |v4
|hi3519av100|4.9.37 |       |v4
|hi3559av100|4.9.37 |       |v4


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
* **hi35XXxvXXX** - chip family specific toolchain, default rootfs and kernel build env
* **putonrootfs** - default rootfs overlay

## Boards

Current build system is targeted to specific boards to make it easy checking software on different devices.

| board name            | status                  |
|-----------------------|-------------------------|
| jvt_hi3519v101_imx274 | app_tester, app_minimal |
| xm_hi3516av100_imx178 | app_tester              |
| other                 | n/a                     |

## Build model
### Terms:
* family - shortname for chips set with same cpu architecture and same videopipeline implementation, shared SDK
* chip - exact SoC model

### Rootfs layers (from top to bottom)

1. application
   1. rootfs overlay (```app_xxx/putonrootfs```)
   2. family dependent rootfs overlay (```app_xxx/hi35XXxvXXX/putonrootfs```)
2. board rootfs overlay (```boards/XXX/putonrootfs```)
3. shared rootfs overlay (```putonrootfs```)
4. generic family rootfs (```hi35XXxvXXX/rootfs```)

### Kernel customization

1. apply board patch for kernel source tree (should be used only for DTS customization)
2. use custom board kernel.config if exist or further
3. use custom board kernel.config.patch for chip generic config if exist or go further
2. use chip generic kernel config as default option

### Sample board config

* ```config``` file
* ```kernel``` dir
  * ```kernel.config``` file, should replace generic chip kernel config
  * ```kernel.config.patch``` file, should be merged with generic chip kernel config (TODO)
  * ```patch``` dir, used for custom dts
* ```putonrootfs``` dir, rootfs overlay

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

1. ```cp Makefile.user.params.example Makefile.user.params``` copy default config and edit it.

Here you can point to default camera you are working with.

2. ```make enviroiment-setup```

This will unpack buildroot and setup toolchains and default rootfses for all chips.
Also it will prepare kernel build env for all chips.
If you want to save time comment lines in *enviroiment-setup* target to exclude chips you are not working with.


### Kernel

Kernel should be built for each board. 
```make kernl-build``` command will build kernel according your config file.

## How to deploy software to camera attached to server
* ```make deploy-app``` to build and deploy app according your config file

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


