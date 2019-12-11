# OpenHisiIpCam development facility

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

