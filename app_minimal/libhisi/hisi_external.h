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

// Configurates two streams 0 - mjpeg 1280x720@1 and 1 - h.264 1280x720@25fps
int hisi_configure();

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

