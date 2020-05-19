#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

//forward declarations
int mpp_data_loop_add(unsigned int *error_code, unsigned int venc_channel_id);
int mpp_data_loop_del(unsigned int *error_code, unsigned int venc_channel_id);
void * mpp_data_loop_thread();

//external funcs
void mpp_data_loop_get_data(unsigned int venc_channel);
int mpp_venc_getfd(int venc_channel_id);

typedef struct st_data_from_c {
    char            *data;
    int             length;
} data_from_c;

void go_callback_receive_data(int venc, unsigned int seq, data_from_c * data_pointer, int num);

//#define ERR_NONE                0
//#define ERR_MPP                 2

int mpp3_venc_sample_mjpeg(unsigned int *error_code, int width, int height, int bitrate, int channelId);
int mpp3_venc_sample_h264(unsigned int *error_code, int width, int height, int bitrate, int channelId);
int mpp3_venc_sample_h265(unsigned int *error_code, int width, int height, int bitrate, int channelId);
int mpp3_venc_delete_encoder(unsigned int *error_code, int channelId);



typedef struct mpp_venc_create_in_struct {              
    unsigned int venc_id;
    unsigned int width;
    unsigned int height;
    unsigned int bitrate;

    unsigned int in_fps;
    unsigned int out_fps;
} mpp_venc_create_in;              


