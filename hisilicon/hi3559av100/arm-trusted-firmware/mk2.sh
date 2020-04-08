export CROSS_COMPILE=/home/nikitos1550/work/venc/hi3519v101_go/output/hisilicon/hi3559av100/toolchain/bin/aarch64-buildroot-linux-gnu-

make clean PLAT=hi3559av100 DEBUG=1
make distclean PLAT=hi3559av100 DEBUG=1
make PLAT=hi3559av100 SPD=none BL33=/home/nikitos1550/work/venc/hi3519v101_go/output/boards/hisilicon_dembverc_hi3559av100_imx334/kernel/uImage CCI_UP=0 DEBUG=1 BL33_SEC=0 HISILICON=1 fip 
