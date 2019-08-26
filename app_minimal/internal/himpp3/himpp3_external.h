#ifndef HIMPP3_EXTERNAL_H_
#define HIMPP3_EXTERNAL_H_

#include <inttypes.h>

//char chipFamily[] = "hi3519av200";

struct jpegFramePack {
        uint32_t length;
        char * data;
        uint64_t pts;
};
struct jpegFrame {
        uint32_t seq;
        uint32_t count;
        struct jpegFramePack packs[5];
};

extern void jpegVencGetDataCallback(struct jpegFrame * newFrame);

int himpp3_ko_init();
int himpp3_sys_init();
int himpp3_vi_init();
int himpp3_mipi_isp_init();
int himpp3_vpss_init();
int himpp3_venc_init();

float gettemperature();

char * getChipFamily();

//experimental funcs
int himpp3_venc_mjpeg_params(unsigned int bitrate);
int himpp3_venc_max_chn_num();

//venc structs and funcs
enum enc_type       {enc_type_none, enc_type_jpeg, enc_type_mjpeg, enc_type_h264, enc_type_h265};
enum rc_type        {rc_type_none, rc_type_cbr, rc_type_vbr, rc_type_fixqp, rc_type_avbr, rc_type_qvbr, rc_type_qmap};
enum h264_profile   {h264_profile_none, h264_profile_baseline, h264_profile_main, h264_profile_high, h264_profile_svct};
enum h265_profile   {h265_profile_none, h265_profile_main, h265_profile_main10};

struct encoder {
    unsigned int    id;
    unsigned char   enabled;
    enum enc_type   etype;//enc type (jpeg|mjpeg|h264|h265)
    unsigned int    source;//source (now only vpss channel 0)
};

struct mjpeg_static {
    unsigned int    width;
    unsigned int    height;
    unsigned int    fps;
};

struct mjpeg_dynamic_cbr {
    unsigned int    bitrate;
    unsigned int    stattime;
    unsigned int    fluctuate;
};

struct mjpeg_dynamic_vbr {
    unsigned int    maxbitrate;
    unsigned int    maxqp;
    unsigned int    minqp;
    unsigned int    stattime;
};

struct mjpeg_dynamic_fixqp {
    unsigned int    qp;
};

struct h264_static {
    unsigned int    width;
    unsigned int    height;
    unsigned int    fps;
    //profile
    //rc
};

struct h264_dynamic_cbr {
    unsigned int    bitrate;
    unsigned int    gop;
    unsigned int    stattime;
    unsigned int    fluctuate;
};

struct h264_dynamic_vbr {
    unsigned int    maxbitrate;
    unsigned int    gop;
    unsigned int    maxqp;
    unsigned int    minqp;
    unsigned int    stattime;
};

struct h264_dynamic_fixqp {
    unsigned int    gop;
    unsigned int    iqp;
    unsigned int    pqp;
    unsigned int    stattime;
};

struct h265_static {
    unsigned int    width;
    unsigned int    height;
    unsigned int    fps;
    //profile
    //rc
};

struct h265_dynamic_cbr {
    unsigned int    bitrate;
    unsigned int    gop;
    unsigned int    stattime;
    unsigned int    fluctuate;
};

struct h265_dynamic_vbr {
    unsigned int    maxbitrate;
    unsigned int    gop;
    unsigned int    maxqp;
    unsigned int    minqp;
    unsigned int    stattime;
};

struct h265_dynamic_fixqp {
    unsigned int    gop;
    unsigned int    iqp;
    unsigned int    pqp;
    unsigned int    stattime;
};

int himpp3_venc_info_chn(unsigned int chn, struct encoder * e);

int himpp3_venc_create_mjpeg_cbr        (struct mjpeg_static    * s,    struct mjpeg_dynamic_cbr    * d);
int himpp3_venc_create_mjpeg_vbr        (struct mjpeg_static    * s,    struct mjpeg_dynamic_vbr    * d);
int himpp3_venc_create_mjpeg_fixqp      (struct mjpeg_static    * s,    struct mjpeg_dynamic_fixqp  * d);

int himpp3_venc_create_h264_cbr         (struct h264_static     * s,    struct h264_dynamic_cbr     * d);
int himpp3_venc_create_h264_vbr         (struct h264_static     * s,    struct h264_dynamic_vbr     * d);
int himpp3_venc_create_h264_fixqp       (struct h264_static     * s,    struct h264_dynamic_fixqp   * d);

int himpp3_venc_create_h265_cbr         (struct h265_static     * s,    struct h265_dynamic_cbr     * d);
int himpp3_venc_create_h265_vbr         (struct h265_static     * s,    struct h265_dynamic_vbr     * d);
int himpp3_venc_create_h265_fixqp       (struct h265_static     * s,    struct h265_dynamic_fixqp   * d);

int himpp3_venc_info_mjpeg_cbr          (struct mjpeg_static    * s,    struct mjpeg_dynamic_cbr    * d);
int himpp3_venc_info_mjpeg_vbr          (struct mjpeg_static    * s,    struct mjpeg_dynamic_vbr    * d);
int himpp3_venc_info_mjpeg_fixqp        (struct mjpeg_static    * s,    struct mjpeg_dynamic_fixqp  * d);

int himpp3_venc_info_h264_cbr           (struct h264_static     * s,    struct h264_dynamic_cbr     * d);
int himpp3_venc_info_h264_vbr           (struct h264_static     * s,    struct h264_dynamic_vbr     * d);
int himpp3_venc_info_h264_fixqp         (struct h264_static     * s,    struct h264_dynamic_fixqp   * d);

int himpp3_venc_info_h265_cbr           (struct h265_static     * s,    struct h265_dynamic_cbr     * d);
int himpp3_venc_info_h265_vbr           (struct h265_static     * s,    struct h265_dynamic_vbr     * d);
int himpp3_venc_info_h265_fixqp         (struct h265_static     * s,    struct h265_dynamic_fixqp   * d);

int himpp3_venc_update_mjpeg_cbr        (struct mjpeg_dynamic_cbr   * d);
int himpp3_venc_update_mjpeg_vbr        (struct mjpeg_dynamic_vbr   * d);
int himpp3_venc_update_mjpeg_fixqp      (struct mjpeg_dynamic_fixqp * d);

int himpp3_venc_update_h264_cbr         (struct h264_dynamic_cbr    * d);
int himpp3_venc_update_h264_vbr         (struct h264_dynamic_vbr    * d);
int himpp3_venc_update_h264_fixqp       (struct h264_dynamic_fixqp  * d);

int himpp3_venc_update_h265_cbr         (struct h265_dynamic_cbr    * d);
int himpp3_venc_update_h265_vbr         (struct h265_dynamic_vbr    * d);
int himpp3_venc_update_h265_fixqp       (struct h265_dynamic_fixqp  * d);

int himpp3_venc_delete                  (unsigned int chn);

/////////////////////////////////////////////////

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
