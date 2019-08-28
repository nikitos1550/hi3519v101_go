//Exports described here
//All functions should be implemented by all chip family implementations

//Convention:   funcs returns only error codes, all data goes via pointer arguments
//Convention:   params that have 0 value (NO_CHANGE defines) are not changed during update

#ifndef HISI_H_
#define HISI_H_

//#define CMOS_WDR_NONE   1
//#define CMOS_WDR_ON     2

#define CMOS_UNDISTORTION_OFF   1
#define CMOS_UNDISTORTION_ON    2

struct cmos_params {
    unsigned int fps;
    unsigned int x0;
    unsigned int y0;
    unsigned int width;
    unsigned int height;
    //unsigned int wdr;
    unsigned int undistortion;
    unsigned int undistortion_x;
    unsigned int undistortion_y;
    unsigned int undistortion_ratio;
    unsigned int undistortion_minratio;
};

int hisi_init(struct cmos_params * cp); //All-in-one init for hi3519v101+imx174 at the moment

////////////////////////////////////////////////////////////////////////////////

int hisi_get_cmos_info(unsigned int * width, unsigned int * height, unsigned int * fps);

////////////////////////////////////////////////////////////////////////////////

int hisi_get_temperature(float * t);

////////////////////////////////////////////////////////////////////////////////

int hisi_channels_max_num(unsigned int * num);

struct channel_params {
    //unsigned int    id;
    //unsigned int    enabled; //TODO How to determine???
    unsigned int    width;
    unsigned int    height;
    unsigned int    fps;
};

int hisi_channel_info(unsigned int id, struct channel_params * chn);
int hisi_channel_enable(unsigned int id, struct channel_params * chn);
int hisi_channel_disable(unsigned int id);

////////////////////////////////////////////////////////////////////////////////

int hisi_encoders_max_num(unsigned int * num);

//int hisi_encoders_max_width(unsigned int codec, unsigned int * width);
//int hisi_encoders_max_height(unsigned int codec, unsigned int * height);

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

#define JPEG_PROFILE_BASELINE   1

#define MJPEG_PROFILE_BASELINE  1

#define H264_PROFILE_BASELINE   1
#define H264_PROFILE_MAIN       2
#define H264_PROFILE_HIGH       3

#define H265_PROFILE_MAIN       1
//#define H265_PROFILE_MAIN10     2

struct encoder_static_params {
    unsigned char   codec;
    unsigned char   rc;
    unsigned char   profile;
    int             width;
    int             height;
    int             fps;
    unsigned int    channel_id;
};

struct encoder_dynamic_params {
    struct cbr_params {
        int    bitrate;
        int    gop;
        int    stattime;
        int    fluctuate;
    } cbr;

    struct vbr_params {
        int    maxbitrate;
        int    gop;
        int    maxqp;
        int    minqp;
        int    stattime;
    } vbr;

    struct fixqp_params {
        int    gop;
        int    iqp;
        int    pqp;
        //int    stattime;
    } fixqp;
};

int hisi_encoder_info   (unsigned int id,
                        struct encoder_static_params * sparams,
                        struct encoder_dynamic_params * dparams);

int hisi_encoder_create (unsigned int id,
                        struct encoder_static_params * sparams,
                        struct encoder_dynamic_params * dparams);

int hisi_encoder_update (unsigned int id,
                        struct encoder_dynamic_params * dparams);

int hisi_encoder_delete (unsigned int id);

////////////////////////////////////////////////////////////////////////////////

struct encoderData {
    //TODO
};
extern int goDataCallback(unsigned int encId, struct encoderData * encData);    //data callback in gospace

////////////////////////////////////////////////////////////////////////////////

#define     PARAM_NO_CHANGE          0  //INPUT
#define     PARAM_NO_VALUE           0  //OUTPUT means at current state object doesn`t have such param
#define     PARAM_NOT_SET           -1  //OUTPUT means not set, but required, seems can be combined with PARAM_BAD_VALUE
#define     PARAM_NOT_APPLICABLE    -2  //OUTPUT means param value is valid, but can`t be used
#define     PARAM_BAD_VALUE         -3  //OUTPUT means param value is not valid (out of range for example)
#define     PARAM_NOT_SUPPORTED     -4  //OUTPUT means current implementation doesn`t support param

#define     ERR_NONE                 0  //Success
#define     ERR_GENERAL             -1  //temporary, should be replaced with specific error code
#define     ERR_OBJECT_NOT_FOUNT    -2  //id out of range
#define     ERR_NOT_ALLOWED         -3  //NOT ALLOWED AT THE MOMENT
#define     ERR_BAD_PARAMS          -4  //BAD INPUT PARAMS
#define     ERR_NOT_IMPLEMENTED     -5  //func NOT IMPLEMENTED at all
#define     ERR_NOT_SUPPORTED       -6  //func NOT SUPPORTED at all
#define     ERR_MPP                 -7  //should not happen, if so check and debug code
#define     ERR_INTERNAL            -8  //should not happen, if so check and debug code

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

#endif //HISI_H_

