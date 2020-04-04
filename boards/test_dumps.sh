#!/bin/sh

make CAMERA=1 BOARD=jvt_s274h19v-l29_hi3519v101_imx274 PROG=1 test_dump_tmp
make CAMERA=2 BOARD=xm_ivg-85hf20pya-s_hi3516ev200_imx307 PROG=2 test_dump_tmp
make CAMERA=3 BOARD=xm_53h20-s_hi3516cv100_imx122 PROG=3 test_dump_tmp
make CAMERA=4 BOARD=xm_ivg-hp203y-se_hi3516cv300_imx291 PROG=4 test_dump_tmp
make CAMERA=5 BOARD=xm_ivg-hp201y-se_hi3516cv300_imx323 PROG=5 test_dump_tmp 
make CAMERA=6 BOARD=jvt_s323h16xf_hi3516cv300_imx323 PROG=6 test_dump_tmp
make CAMERA=7 BOARD=ruision_rs-h622qm-b0_hi3516cv300_imx323 PROG=7 test_dump_tmp
make CAMERA=8 BOARD=xm_ivg-85hg50pya-s_hi3516ev300_imx335 PROG=8 test_dump_tmp
make CAMERA=9 BOARD=xm_ipg-83h50p-b_hi3516av100_imx178 PROG=9 test_dump_tmp
make CAMERA=10 BOARD=xm_ipg-83he20py-s_hi3516ev100_imx323 PROG=10 test_dump_tmp

#make CAMERA=11 BOARD=xm_ivg-83h80nv-be_hi3516av200_os08a10 PROG=no test_dump_tmp

make CAMERA=12 BOARD=ssqvision_unknown_hi3516av300_imx334 PROG=12 test_dump_tmp
make CAMERA=13 BOARD=ssqvision_on335h16d_hi3516dv300_imx335 PROG=13 test_dump_tmp

#make BOARD=jvt_s226h19v-l29_hi3519v101_imx226 PROG=no test_dump_tmp
#make BOARD=xm_53h20-ae_hi3516cv100_imx222 PROG=no test_dump_tmp
#make BOARD=xm_83h40pl-b_hi3516av100_ov4689 PROG=no test_dump_tmp
#make BOARD=ssqvision_unknown_hi3519v101_imx326 PROG=no test_dump_tmp
#make BOARD=ssqvision_on290h16d_hi3516dv100_imx290 PROG=no test_dump_tmp
#make BOARD=hisilicon_demb_hi3516dv300_imx290 PROG=no test_dump_tmp
#make BOARD=hisilicon_dembverc_hi3559av100_imx334 PROG=no test_dump_tmp
