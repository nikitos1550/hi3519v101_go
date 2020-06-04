#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

#include <pthread.h>

int pthread_setname_np(pthread_t thread, const char *name);

int mpp_ai_test(error_in *err);
int mpp_ai_config_inner(error_in *err);
int mpp_ao_test(error_in *err);
