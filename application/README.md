# Application
Application is wrote mainly in golang with a bit C (via cgo).

## Build targets
Build target is version of application for some special purpose.

At the moment there are two targets:
* tester
* daemon

### Tester
Version of application that shares same codebase as other targets for smoke test purpose.
Application will start up simple web server on 80 port and serve same json answer for all requests.

Example answer:
```json
{
  "appName": "tester",
  "chipDetectedReg": "hi3519v101",
  "chipDetectedMpp": "hi3519v101",
  "mppVersion": "HI_VERSION=Hi3519V101_MPP_V2.0.5.0 B040 Release",
  "chipIdReg": 890831105,
  "chipIdMpp": 890831105,
  "temperature": 52.940453,
  "temperatureHW": "availible",
  "buildInfo": {
    "goversion": "go version go1.13.4 linux/amd64",
    "gccversion": "arm-buildroot-linux-uclibcgnueabi-gcc.br_real (Buildroot 2019.08-g1aead48-dirty) 7.4.0 Copyright (C) 2017 Free Software Foundation, Inc. This is free software; see the source for copying conditions.  There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.",
    "builddatetime": "2019-12-03 17:06",
    "buildtags": "hi3516av200",
    "builduser": "nikitos1550",
    "buildcommit": "c19625a",
    "buildbranch": "daemon_jvt_hi3519v101_imx274",
    "boardprofile": "jvt_hi3519v101_imx274",
    "boardvendor": "JVT",
    "boardmodel": "unknown",
    "chip": "hi3519v101",
    "cmos": "imx274",
    "totalram": "512",
    "linuxram": "256",
    "mppram": "256"
  }
}
```

Tester build depends only on hisilicon family tag, no other tags should be provided.

Successful run means that board is properly connected to facility, kernel mpp ko and a libs are consistent,
toolchain settings are correct and overall build system id working as expected.

### Daemon

## Code structure

## MPP backend

## HTTP API

## Lua API

## Debug notes

## Build notes
###Invoke build from application dir
```console
foo@build-hisi:~/hi3519v101_go/application$ make PARAMS_FILE=../output/boards/jvt_hi3519v101_imx274/Makefile.params build-tester
```
or
```console
foo@build-hisi:~/hi3519v101_go/application$ make PARAMS_FILE=../output/boards/jvt_hi3519v101_imx274/Makefile.params build-daemon
```

