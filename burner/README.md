# Tool for burn firmware for embedded hisilicon device

```
usage: burner.py [-h] --ip IP [--skip SKIP] [--uboot UBOOT] [--uimage UIMAGE]
                 [--rootfs ROOTFS] [--memory MEMORY] [--initrd INITRD]
                 [--lserial LSERIAL] [--port PORT] [--speed SPEED]
                 [--iface IFACE] [--duplex DUPLEX] [--mc21 MC21] [--force]
                 action

Load/burn firmware to hisilicon ip device.

positional arguments:
  action             Required action load|burn|uboot|mac

optional arguments:
  -h, --help         show this help message and exit
  --ip IP            Ip address for device
  --skip SKIP        Uboot size to skip in kb (default 512kb)
  --uboot UBOOT      Uboot image file name
  --uimage UIMAGE    Kernel uImage file name
  --rootfs ROOTFS    RootFS file name
  --memory MEMORY    Amount of RAM for Linux in MB
  --initrd INITRD    Amount of RAM for Linux Initrd (only for load action)
  --lserial LSERIAL  Linux load serial config
  --port PORT        Serial port dev name
  --speed SPEED      Serial port speed
  --iface IFACE      Network interface name
  --duplex DUPLEX    Full (default) or Half duplex mode
  --mc21 MC21        To remove
  --force            Skip usefull checks
```

## Examples:
To burn uboot ```sudo ./burner.py uboot --ip 192.168.0.200 --uboot ../loader/u-boot-50M-3s.bin```

To load linux ```sudo ./burner.py load --uimage ./images/uImage --rootfs ./images/rootfs.romfs --ip 192.168.0.200 --speed 9600 --skip 1024 --memory 96```

To burn linux ```sudo ./burner.py burn --uimage ./images/uImage --rootfs ./images/rootfs.romfs --ip 192.168.0.200 --speed 9600 --skip 1024 --memory 96```






