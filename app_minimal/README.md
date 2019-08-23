# Minimal application

## About
* / - get hello world
* /image.jpeg - get jpeg from camera
* /t - internal SoC temperature sensor
* /experimental/hidebug[/(filename).(raw|json)] - HiMPP /proc/umap debug
* /experimental/himpp3/bitrate/{value} - test method set bitrate for mjpeg channel

## Notes

App can`t rerun, you have to reset camera. Need improve ko module loading and sensor register to fix it.
