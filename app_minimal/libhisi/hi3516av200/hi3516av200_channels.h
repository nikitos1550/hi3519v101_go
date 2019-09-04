#ifndef HI3516AV200_CHANNELS_H_
#define HI3516AV200_CHANNLES_H_

#define CHANNEL_DISABLED    0
#define CHANNEL_ENABLED     1

extern int channels_enable[VPSS_MAX_PHY_CHN_NUM];

struct vpss_setup { //Temporary structure to save configured vpss state
    unsigned int width;
    unsigned int height;
    unsigned int fps;
};

extern struct vpss_setup vpss;


#endif //HI3516AV200_CHANNELS_H_

