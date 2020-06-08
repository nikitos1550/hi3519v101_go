#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

#include <pthread.h>

#include <stdint.h>

int pthread_setname_np(pthread_t thread, const char *name);

int mpp_ai_test(error_in *err);
int mpp_ai_config_inner(error_in *err);
int mpp_ao_test(error_in *err);

typedef struct audio_data_from_c_struct {
    unsigned char   *data;
    int             length;
} audio_data_from_c;

typedef struct audio_info_from_c_struct {
   unsigned int     seq;
   uint64_t         timestamp;
} audio_info_from_c;

void go_callback_raw_tmp(audio_info_from_c *info_pointer, audio_data_from_c *data_pointer);
void go_callback_opus_tmp(audio_info_from_c *info_pointer, audio_data_from_c *data_pointer);
