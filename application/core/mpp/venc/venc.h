#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#define CODEC_MJPEG     1
#define CODEC_H264      2
#define CODEC_H265      3

#define INVALID_VALUE   (0xFFFFFFFF >> 1)

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
    int id;

    int codec;
    int profile;

    int width;
    int height;

    int in_fps;
    int out_fps;

    int bitrate_control;

    int gop;

    int gop_mode;

    int i_pq_delta;
    int s_p_interval;
    int s_pq_delta;
    int bg_interval;
    int bg_qp_delta;
    int vi_qp_delta;
    int b_frm_num;
    int b_qp_delta;

    int bitrate;

    int stat_time;
    int fluctuate_level;

    int q_factor;
    int min_q_factor;
    int max_q_factor;

    int i_qp;
    int p_qp;
    int b_qp;

    int min_qp;
    int max_qp;
    int min_i_qp;

} mpp_venc_create_encoder_in;              

void invalidate_mpp_venc_create_encoder_in (mpp_venc_create_encoder_in *in);

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

typedef struct mpp_send_frame_to_encoder_in_struct {
    int                 id;
    //VIDEO_FRAME_INFO_S  *frame;
    void                *frame;
} mpp_send_frame_to_encoder_in;

int mpp_send_frame_to_encoder(error_in *err, mpp_send_frame_to_encoder_in *in, void *frame);

typedef struct mpp_venc_request_idr_in_struct {
    int id;
} mpp_venc_request_idr_in;

int mpp_venc_request_idr(error_in *err, mpp_venc_request_idr_in *in);

typedef struct mpp_venc_reset_in_struct {
    int id;
} mpp_venc_reset_in;

int mpp_venc_reset(error_in *err, mpp_venc_reset_in *in);
