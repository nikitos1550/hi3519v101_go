#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#define CODEC_MJPEG 1
#define CODEC_H264  2
#define CODEC_H265  3

//forward declarations
int mpp_data_loop_add(unsigned int *error_code, unsigned int venc_channel_id, unsigned int codec);
int mpp_data_loop_del(unsigned int *error_code, unsigned int venc_channel_id);
int mpp_data_loop_init(unsigned int *error_code);
void * mpp_data_loop_thread();

void mpp_data_loop_get_data(unsigned int id);
int mpp_venc_getfd(int venc_channel_id);

typedef struct data_from_c_struct {
    unsigned char   *data;
    int             length;
} data_from_c;

typedef struct info_from_c_struct {
   unsigned int         seq; 
   unsigned long long   pts;
   unsigned int         q_factor;
   unsigned int         ref_type;
} info_from_c;

void go_callback_receive_data(int venc, info_from_c *info_pointer, data_from_c *data_pointer, int num);

int mpp3_venc_sample_mjpeg(unsigned int *error_code, int width, int height, int bitrate, int channelId);
int mpp3_venc_sample_h264(unsigned int *error_code, int width, int height, int bitrate, int channelId);
int mpp3_venc_sample_h265(unsigned int *error_code, int width, int height, int bitrate, int channelId);
int mpp3_venc_delete_encoder(unsigned int *error_code, int channelId);

typedef struct mpp_venc_create_encoder_in_struct {              
    unsigned int    id;

    unsigned int    codec;
    unsigned int    profile;

    unsigned int    width;
    unsigned int    height;

    int             in_fps;
    int             out_fps;

    unsigned int    bitrate_control;

    unsigned int    gop;

    unsigned int    gop_mode;

    int             i_pq_delta;
    unsigned int    s_p_interval;
    int             s_pq_delta;
    unsigned int    bg_interval;
    int             bg_qp_delta;
    int             vi_qp_delta;
    unsigned int    b_frm_num;
    int             b_qp_delta;

    unsigned int    bitrate;

    unsigned int    stat_time;
    unsigned int    fluctuate_level;

    unsigned int    q_factor;
    unsigned int    min_q_factor;
    unsigned int    max_q_factor;

    unsigned int    i_qp;
    unsigned int    p_qp;
    unsigned int    b_qp;

    unsigned int    min_qp;
    unsigned int    max_qp;
    unsigned int    min_i_qp;

} mpp_venc_create_encoder_in;              

typedef struct mpp_venc_destroy_encoder_in_struct {              
    unsigned int id;
} mpp_venc_destroy_encoder_in; 

typedef struct mpp_venc_start_encoder_in_struct {              
    unsigned int id;
} mpp_venc_start_encoder_in; 

typedef struct mpp_venc_stop_encoder_in_struct {              
    unsigned int id;
} mpp_venc_stop_encoder_in; 

int mpp_venc_create_encoder(error_in *err, mpp_venc_create_encoder_in *in);
int mpp_venc_start_encoder(error_in *err, mpp_venc_start_encoder_in *in);                   
int mpp_venc_stop_encoder(error_in *err, mpp_venc_stop_encoder_in *in);                     
int mpp_venc_update_encoder(error_in *err, mpp_venc_create_encoder_in *in);
int mpp_venc_destroy_encoder(error_in *err, mpp_venc_destroy_encoder_in *in); 
