#ifndef HIMPP3_EXTERNAL_H_
#define HIMPP3_EXTERNAL_H_

int himpp3_ko_init();
int himpp3_sys_init();
int himpp3_vi_init();
int himpp3_mipi_isp_init();
int himpp3_vpss_init();
int himpp3_venc_init();

int himpp3_venc_jpeg_export_frame();

char * himpp3_test_func(char ** buffer);

#define HIMPP3_ERROR_FUNC_NONE                  0
#define HIMPP3_ERROR_FUNC_HI_MPI_SYS_Exit       1
#define HIMPP3_ERROR_FUNC_HI_MPI_VB_Exit        2
#define HIMPP3_ERROR_FUNC_HI_MPI_VB_SetConf     3
#define HIMPP3_ERROR_FUNC_HI_MPI_SYS_SetConf    4
#define HIMPP3_ERROR_FUNC_HI_MPI_SYS_Init       5
#define HIMPP3_ERROR_FUNC_
#define HIMPP3_ERROR_FUNC_
#define HIMPP3_ERROR_FUNC_
#define HIMPP3_ERROR_FUNC_

#endif
