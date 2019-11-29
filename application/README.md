# Application notes

## How to build
* ```make PARAMS_FILE=../output/boards/jvt_hi3519v101_imx274/Makefile.params build-[tester|daemon]```
* but better user ```make build-app``` from root Makefile!

## Configs, flags and scripting

Videopipeline can`t be setuped by config file, becasue config file cant handle command order,
fixing commands execution order will make system not flexible, so tasks for initial videopipeline
configuration are on lua scripts.

### Config apply order

1. Cmd params
2. Config file param
3. Default hardcoded values