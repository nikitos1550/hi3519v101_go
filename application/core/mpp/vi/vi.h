#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>

typedef struct mpp_dc_sync_attrs_struct {
    unsigned char v_sync;           
    unsigned char v_sync_neg;       
    unsigned char h_sync;           
    unsigned char h_sync_neg;       
    unsigned char v_sync_valid;     
    unsigned char v_sync_valid_neg; 

    unsigned int timing_hfb;      
    unsigned int timing_act;      
    unsigned int timing_hbb;      
    unsigned int timing_vfb;      
    unsigned int timing_vact;     
    unsigned int timing_vbb;      
    unsigned int timing_vbfb;     
    unsigned int timing_vbact;    
    unsigned int timing_vbbb;     
} mpp_dc_sync_attrs;

typedef struct mpp_vi_init_in_struct {
    //void *videv;

    unsigned int pixel_bitness;
    unsigned int data_type;

    mpp_dc_sync_attrs dc_sync_attrs;
    unsigned int dc_zero_bit_offset;

    unsigned int vi_crop_x0;
    unsigned int vi_crop_y0;
    unsigned int vi_crop_width;
    unsigned int vi_crop_height;
    
    //unsigned int sync_width;
    //unsigned int sync_height;

    unsigned int width;
    unsigned int height;

    //unsigned int x0;
    //unsigned int y0;
    //unsigned int width;
    //unsigned int height;
    unsigned int cmos_fps;
    unsigned int fps;

    unsigned int wdr;

    unsigned char mirror;
    unsigned char flip;

    unsigned char ldc;
    int ldc_offset_x;
    int ldc_offset_y;
    int ldc_k;
} mpp_vi_init_in;

int mpp_vi_init(error_in *err, mpp_vi_init_in * in);

typedef struct mpp_vi_ldc_in_struct {
    int x;
    int y;
    int k;
} mpp_vi_ldc_in;

int mpp_vi_ldc_update(error_in *err, mpp_vi_ldc_in * in);
