#ifndef HIMPP3_H_
#define HIMPP3_H_

#include "hisi3.h"

int himpp3_sys_init();
int himpp3_vi_init();
int himpp3_mipi_init();
int himpp3_isp_init();
int himpp3_vpss_init();
int himpp3_venc_init();

int himpp3_venc_jpeg_get_frame();

#endif // HIMPP3_H_
