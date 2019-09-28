#ifndef HI3516AV200_CMOS_H_
#define HI3516AV200_CMOS_H_

#include "hi3516av200_mpp.h"

struct hi3516av200_cmos {
    unsigned int        id;
    char                *name;
    char                *description;
    unsigned int        width;
    unsigned int        height;
    unsigned int        fps;
    combo_dev_attr_t    *mipidev;
    VI_DEV_ATTR_S       *videv;
    ISP_BAYER_FORMAT_E  bayer;
    ISP_SNS_OBJ_S       *snsobj;
};

extern struct hi3516av200_cmos hi3516av200_cmoses[];

#endif // HI3516AV200_CMOS_H_

