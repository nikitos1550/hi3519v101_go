#ifndef ERRMPP_H_
#define ERRMPP_H_

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

#define RETURN_ERR_MPP(x, y) \
    err->f = x; \
    err->mpp = y; \
    return ERR_MPP

#include "functions.h"

#endif // ERRMPP_H_
