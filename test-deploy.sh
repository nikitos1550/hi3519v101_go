#!/bin/sh

# Devices list for 5 october 2019
#1.  hi3519v101+imx274 jvt OK
#make BOARD=jvt_hi3519v101_imx274        CAMERA=1    APP=app_tester build-kernel
#make BOARD=jvt_hi3519v101_imx274        CAMERA=1    APP=app_tester deploy-app

#2.  hi3519v101+imx274 jvt OK
#make BOARD=jvt_hi3519v101_imx274        CAMERA=2    APP=app_tester build-kernel
#make BOARD=jvt_hi3519v101_imx274        CAMERA=2    APP=app_tester deploy-app

#3.  hi3516ev100 KERNEL NO OUTPUT, MAYBE SOME CN UBOOT LOCK
##make BOARD=generic_hi3516cv300_unknown  CAMERA=3    APP=app_tester build-kernel
##make BOARD=generic_hi3516cv300_unknown  CAMERA=3    APP=app_tester deploy-app

#4.  hi3516cv200 NOT WORKING - NO FEEDBACK FROM SERIAL
##make BOARD=generic_hi3516cv200_unknown  CAMERA=4    APP=app_tester build-kernel
##make BOARD=generic_hi3516cv200_unknown  CAMERA=4    APP=app_tester deploy-app

#5.  hi3518ev201 NOT WORKING
##make BOARD=generic_hi3516cv200_unknown  CAMERA=5    APP=app_tester build-kernel
##make BOARD=generic_hi3516cv200_unknown  CAMERA=5    APP=app_tester deploy-app

#6.  hi3516cv300+imx290 OK
#make BOARD=generic_hi3516cv300_unknown  CAMERA=6    APP=app_tester build-kernel
#make BOARD=generic_hi3516cv300_unknown  CAMERA=6    APP=app_tester deploy-app

#7.  hi3518ev100 OK
#make BOARD=generic_hi3518ev100_unknown  CAMERA=7    APP=app_tester build-kernel
#make BOARD=generic_hi3518ev100_unknown  CAMERA=7    APP=app_tester deploy-app

#8.  hi3516av100+imx178 OK
#make BOARD=xm_hi3516av100_imx178        CAMERA=8    APP=app_tester build-kernel
#make BOARD=xm_hi3516av100_imx178        CAMERA=8    APP=app_tester deploy-app

#9.  hi3518ev200 NO BOOT LOG
#make BOARD=generic_hi3516cv200_unknown  CAMERA=9    APP=app_tester build-kernel
#make BOARD=generic_hi3516cv200_unknown  CAMERA=9    APP=app_tester deploy-app

#10. hi3518ev200 OK
#make BOARD=generic_hi3518ev200_unknown  CAMERA=10   APP=app_tester build-kernel
#make BOARD=generic_hi3518ev200_unknown  CAMERA=10   APP=app_tester deploy-app

#11. hi3516ev100 CAN`T GET UBOOT CONTROL
#make BOARD=generic_hi3516cv300_unknown  CAMERA=11   APP=app_tester build-kernel
#make BOARD=generic_hi3516cv300_unknown  CAMERA=11   APP=app_tester deploy-app

#12. hi3516cv100+imx122 SOME PROBLEM WITH SERIAL
##make BOARD=generic_hi3516cv100_unknown  CAMERA=12   APP=app_tester build-kernel
##make BOARD=generic_hi3516cv100_unknown  CAMERA=12   APP=app_tester deploy-app

#13. ??? NOT WORKING - NOT CONNECTED
#make BOARD=??? CAMERA=13 APP=app_tester deploy-app

#14. hi35109v101+imx226
make BOARD=jvt_hi3519v101_imx226        CAMERA=14   APP=app_tester build-kernel
make BOARD=jvt_hi3519v101_imx226        CAMERA=14   APP=app_tester deploy-app

#15. hi3516dv300+imx290 hisilicon evb
# NOT IMPLEMENTED

#16. hi3559av100 hisilicon evb
# NOT IMPLEMENTED

