#pragma once

#define ERR_NONE                    0
#define ERR_MPP                     1
#define ERR_GENERAL                 2

typedef struct error_in_struct {
    unsigned int f;
    unsigned int mpp; 
    int general;
} error_in;  

typedef unsigned int error_mpp;
typedef int error_general;

#define DO_OR_RETURN_MPP(FUNC, ...) \
    mpp_error_code = FUNC(__VA_ARGS__); \
    if (mpp_error_code != HI_SUCCESS) { \
        err->f = ERR_F_##FUNC; \
        err->mpp = mpp_error_code; \
        return ERR_MPP; \
    }   

#define RETURN_ERR_MPP(x, y) \
    err->f = x; \
    err->mpp = y; \
    return ERR_MPP

#include "functions.h"

