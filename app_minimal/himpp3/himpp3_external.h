#ifndef HIMPP3_EXTERNAL_H_
#define HIMPP3_EXTERNAL_H_

#include <inttypes.h>

struct jpegFramePack {
        uint32_t length;
        char * data;
        uint64_t pts;
};
struct jpegFrame {
        uint32_t seq;
        uint32_t count;
        struct jpegFramePack packs[5];

        /*
        struct jpegInfo {

        };
        struct jpegAdvancedInfo {


        };
        */
};

extern void jpegVencGetDataCallback(struct jpegFrame * newFrame);

int himpp3_ko_init();
int himpp3_sys_init();
int himpp3_vi_init();
int himpp3_mipi_isp_init();
int himpp3_vpss_init();
int himpp3_venc_init();

float gettemperature();

int himpp3_venc_mjpeg_params(unsigned int bitrate);

int himpp3_venc_max_chn_num();

//int himpp3_venc_jpeg_export_frame();

//char * himpp3_test_func(char ** buffer);

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
