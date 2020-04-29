#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <stdint.h>
#include <string.h>

#define MAX_CHANNELS VPSS_MAX_PHY_CHN_NUM
//VIDEO_FRAME_INFO_S channelFrames[MAX_CHANNELS];

typedef struct mpp_vpss_init_in_struct {
    unsigned int width;
    unsigned int height;
    unsigned char nr;
    unsigned char nr_frames;
} mpp_vpss_init_in;

typedef struct mpp_vpss_create_channel_in_struct {
    unsigned int channel_id;
    unsigned int width;
    unsigned int height;
    unsigned int vi_fps;
    unsigned int fps;
} mpp_vpss_create_channel_in;

typedef struct mpp_vpss_destroy_channel_in_struct {
    unsigned int channel_id;
} mpp_vpss_destroy_channel_in;

typedef struct mpp_receive_frame_out_struct {

} mpp_receive_frame_out;

int mpp_vpss_init(error_in *err, mpp_vpss_init_in *in);
int mpp_vpss_create_channel(error_in *err, mpp_vpss_create_channel_in * in);
int mpp_vpss_destroy_channel(error_in * err, mpp_vpss_destroy_channel_in *in);
int mpp_receive_frame(error_in *err, unsigned int channel_id, void** frame);
int mpp_release_frame(error_in *err, unsigned int channel_id);
