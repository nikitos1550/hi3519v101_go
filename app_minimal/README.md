# Minimal application

## Files structure
|Path|Description|
|----|-------|
|/cmd|main GO package|
|/go|GOPATH|
|/go/src/openhisiipcam.org/*|internal GO packages|
|/api|openapi3 api specs|
|/www|static web files|
|/lua|LUA scripts|
|libhisi|C hw abstraction layer|

## Conditioanal compilation
See Makefile.tags for details

## About
* / - get hello world
* /image.jpeg - get jpeg from camera
* /t - internal SoC temperature sensor
* /experimental/hidebug[/(filename).(raw|json)] - HiMPP /proc/umap debug
* /experimental/himpp3/bitrate/{value} - test method set bitrate for mjpeg channel

## Notes

* ~~App can`t rerun, you have to reset camera. Need improve ko module loading and sensor register to fix it.~~ fixed.
* First time run ```app_minimal/himpp3/ko make```, even it is envoked in app_minimal Makefile, it is not working. Need fixing.
