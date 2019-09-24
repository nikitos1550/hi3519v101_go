//Exports described here
//All functions should be implemented by all chip family implementations

//Convention:   funcs returns only error codes, all data goes via pointer arguments
//Convention:   params that have 0 value (NO_CHANGE defines) are not changed during update

#ifndef HISI_EXTERNAL_H_
#define HISI_EXTERNAL_H_

////////////////////////////////////////////////////////////////////////////////

struct capture_params {
    int fps;
    int x0;
    int y0;
    int width;
    int height;
    //int undistortion;
    //int undistortion_x;
    //int undistortion_y;
    //int undistortion_ratio;
    //int undistortion_minratio;
};

int hisi_init(unsigned int cid, struct capture_params * cp);

struct cmos {
    int width;
    int height;
    int fps;
};

int hisi_cmos(struct cmos * c);

////////////////////////////////////////////////////////////////////////////////

int hisi_get_chipid(unsigned int * chip);
int hisi_get_temperature(float * t);

////////////////////////////////////////////////////////////////////////////////

/*
struct gpio_state {
    unsigned int configured;
    unsigned int value;
};

int hisi_gpio_state();
int hisi_gpio_configure();
int hisi_gpio_on();
int hisi_gpio_off();
*/

////////////////////////////////////////////////////////////////////////////////

/*
struct cmos {
    unsigned int        id;
    char                *name;
    char                *description;
};

int hisi_cmos_fetch(unsigned int id, struct cmos * c);
int hisi_cmos_status(unsigned int * id);
int hisi_cmos_init(unsigned int id);
*/

////////////////////////////////////////////////////////////////////////////////

int hisi_channels_max_num   (unsigned int * num);
int hisi_channels_min_width (unsigned int * w);
int hisi_channels_max_width (unsigned int * w);
int hisi_channels_min_height(unsigned int * h);
int hisi_channels_max_height(unsigned int * h);
int hisi_channels_max_fps   (unsigned int * fps);

struct channel_params {
    //int crop_x0;
    //int crop_y0;
    //int crop_width;
    //int crop_height;
    int width;
    int height;
    int fps;
};

int hisi_channel_fetch  (unsigned int id, struct channel_params * chn);
int hisi_channel_enable (unsigned int id, struct channel_params * chn);
int hisi_channel_disable(unsigned int id);

////////////////////////////////////////////////////////////////////////////////

int hisi_encoders_max_num(unsigned int * num);

#define CODEC_JPEG  1
#define CODEC_MJPEG 2
#define CODEC_H264  3
#define CODEC_H265  4

#define RC_CBR      1
#define RC_VBR      2
#define RC_FIXQP    3
#define RC_AVBR     4
#define RC_QVBR     5
#define RC_QMAP     6

#define PROFILE_JPEG_BASELINE   10
#define PROFILE_MJPEG_BASELINE  20
#define PROFILE_H264_BASELINE   30
#define PROFILE_H264_MAIN       31
#define PROFILE_H264_HIGH       32
#define PROFILE_H265_MAIN       40
#define PROFILE_H265_MAIN10     41

struct encoder_static_params {
    int codec;
    int rc;
    int profile;
    int width;
    int height;
    int fps;
    int channel;
};

//checks encoder channel id, fills sparams structure or return error in libhisi namespace
int hisi_encoder_fetch  (unsigned int id,
                        struct encoder_static_params * sparams);

int hisi_encoder_delete (unsigned int id);

struct encoder_mjpeg_cbr_params {
    int bitrate;
    int stattime;
    int fluctuate;
};

struct encoder_mjpeg_vbr_params {
    int maxbitrate;
    int stattime;
    int maxqfactor;
    int minqfactor;
};

struct encoder_mjpeg_fixqp_params {
    int qfactor;
};

struct encoder_h26x_cbr_params {
    int gop;
    int bitrate;
    int stattime;
    int fluctuate;
};

struct encoder_h26x_vbr_params {
    int gop;
    int stattime;
    int maxbitrate;
    int maxqp;
    int minqp;
    int miniqp;
};

struct encoder_h26x_fixqp_params {
    int gop;
    int iqp;
    int pqp;
    int bqp;
};

int hisi_hi3516av200_encoder_mjpeg_cbr_fetch    (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_mjpeg_cbr_params * dparams);
int hisi_hi3516av200_encoder_mjpeg_cbr_create   (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_mjpeg_cbr_params * dparams);
int hisi_hi3516av200_encoder_mjpeg_cbr_update   (unsigned int id,
                                                 struct encoder_mjpeg_cbr_params * dparams);

int hisi_hi3516av200_encoder_mjpeg_vbr_fetch    (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_mjpeg_vbr_params * dparams);
int hisi_hi3516av200_encoder_mjpeg_vbr_create   (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_mjpeg_vbr_params * dparams);
int hisi_hi3516av200_encoder_mjpeg_vbr_update   (unsigned int id,
                                                 struct encoder_mjpeg_vbr_params * dparams);

int hisi_hi3516av200_encoder_mjpeg_fixqp_fetch  (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_mjpeg_fixqp_params * dparams);
int hisi_hi3516av200_encoder_mjpeg_fixqp_create (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_mjpeg_fixqp_params * dparams);
int hisi_hi3516av200_encoder_mjpeg_fixqp_update (unsigned int id,
                                                 struct encoder_mjpeg_fixqp_params * dparams);

int hisi_hi3516av200_encoder_h264_cbr_fetch     (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_cbr_params * dparams);
int hisi_hi3516av200_encoder_h264_cbr_create    (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_cbr_params * dparams);
int hisi_hi3516av200_encoder_h264_cbr_update    (unsigned int id,
                                                 struct encoder_h26x_cbr_params * dparams);

int hisi_hi3516av200_encoder_h264_vbr_fetch     (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_vbr_params * dparams);
int hisi_hi3516av200_encoder_h264_vbr_create    (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_vbr_params * dparams);
int hisi_hi3516av200_encoder_h264_vbr_update    (unsigned int id,
                                                 struct encoder_h26x_vbr_params * dparams);

int hisi_hi3516av200_encoder_h264_fixqp_fetch   (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_fixqp_params * dparams);
int hisi_hi3516av200_encoder_h264_fixqp_create  (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_fixqp_params * dparams);
int hisi_hi3516av200_encoder_h264_fixqp_update  (unsigned int id,
                                                 struct encoder_h26x_fixqp_params * dparams);


int hisi_hi3516av200_encoder_h265_cbr_fetch     (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_cbr_params * dparams);
int hisi_hi3516av200_encoder_h265_cbr_create    (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_cbr_params * dparams);
int hisi_hi3516av200_encoder_h265_cbr_update    (unsigned int id,
                                                 struct encoder_h26x_cbr_params * dparams);

int hisi_hi3516av200_encoder_h265_vbr_fetch     (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_vbr_params * dparams);
int hisi_hi3516av200_encoder_h265_vbr_create    (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_vbr_params * dparams);
int hisi_hi3516av200_encoder_h265_vbr_update    (unsigned int id,
                                                 struct encoder_h26x_vbr_params * dparams);

int hisi_hi3516av200_encoder_h265_fixqp_fetch   (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_fixqp_params * dparams);
int hisi_hi3516av200_encoder_h265_fixqp_create  (unsigned int id,
                                                 struct encoder_static_params * sparams,
                                                 struct encoder_h26x_fixqp_params * dparams);
int hisi_hi3516av200_encoder_h265_fixqp_update  (unsigned int id,
                                                 struct encoder_h26x_fixqp_params * dparams);



////////////////////////////////////////////////////////////////////////////////

struct encoderData {
    //TODO
};
extern int goDataCallback(unsigned int encId, struct encoderData * encData);    //data callback in gospace

////////////////////////////////////////////////////////////////////////////////

#define     PARAM_NO_CHANGE          0  //INPUT
#define     PARAM_BAD_VALUE         -1  //OUTPUT means param value is not valid (out of range for example)

#define     ERR_NONE                 0  //Success
#define     ERR_GENERAL             -1  //temporary, should be replaced with specific error code
#define     ERR_OBJECT_NOT_FOUND    -2  //id out of range
#define     ERR_NOT_ALLOWED         -3  //NOT ALLOWED AT THE MOMENT
#define     ERR_BAD_PARAMS          -4  //BAD INPUT PARAMS
#define     ERR_NOT_IMPLEMENTED     -5  //func NOT IMPLEMENTED at all
#define     ERR_NOT_SUPPORTED       -6  //func NOT SUPPORTED at all
#define     ERR_MPP                 -7  //should not happen, if so check and debug code
#define     ERR_INTERNAL            -8  //should not happen, if so check and debug code
#define     ERR_DISABLED            -9

////////////////////////////////////////////////////////////////////////////////

#define CHIP_FAMILY_HI3516CV100 10
#define CHIP_HI3516CV100        11
#define CHIP_HI3518CV100        12
#define CHIP_HI3518EV100        13

#define CHIP_FAMILY_HI3516CV200 20
//TODO

#define CHIP_FAMILY_HI3516CV300 30
#define CHIP_HI3516CV300        31
#define CHIP_HI3516EV100        32

#define CHIP_FAMILY_HI3516AV100 40
#define CHIP_HI3516DV100        41
#define CHIP_HI3516AV100        42

#define CHIP_FAMILY_HI3516AV200 50
#define CHIP_HI3516AV200        51
#define CHIP_HI3519V101         52

#define CHIP_FAMILY_HI3516CV500 60
#define CHIP_HI3516CV500        61
#define CHIP_HI3516DV300        62

#define CHIP_FAMILY_HI3519AV100 70
//TODO

#define CHIP_FAMILY_HI3559AV100 80
//TODO

#define CHIP_FAMILY_HI3518EV300 90
//TODO

#endif //HISI_EXTERNAL_H_

