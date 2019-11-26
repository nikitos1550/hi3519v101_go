# Tester application

Tester application features:
* load minimal required ko modules 
* starts http server
* serves information json answer on any request

Application confirms consistency of kernel, ko and libs.

## Example output fot jvt_hi3519v101_imx274 board
```
{
  "appName": "tester",
  "chipDetectedReg": "hi3519v101",
  "chipDetectedMpp": "hi3519v101",
  "mppVersion": "HI_VERSION=Hi3519V101_MPP_V2.0.5.0 B040 Release",
  "chipIdReg": 890831105,
  "chipIdMpp": 890831105,
  "temperature": 57.853592,
  "temperatureHW": "availible",
  "buildInfo": {
    "GoVersion": "go version go1.13.4 linux/amd64",
    "GccVersion": "arm-buildroot-linux-uclibcgnueabi-gcc.br_real (Buildroot 2019.08-ga3ce976-dirty) 7.4.0 Copyright (C) 2017 Free Software Foundation, Inc. This is free software; see the source for copying conditions.  There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.",
    "BuildDateTime": "2019-11-26 01:01",
    "BuildTags": "hi3516av200",
    "BuildUser": "nikitos1550",
    "BuildCommit": "5a2527f",
    "BuildBranch": "make-refactoring",
    "BoardVendor": "JVT",
    "BoardModel": "unknown",
    "Chip": "hi3519v101",
    "CmosProfile": "imx274",
    "TotalRam": "512",
    "LinuxRam": "256",
    "MppRam": "256"
  }
}
```