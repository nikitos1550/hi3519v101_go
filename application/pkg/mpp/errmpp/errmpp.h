#pragma once

#include <stdio.h>
#include <stdint.h>

#define ERR_NONE                    0
#define ERR_MPP                     1
#define ERR_GENERAL                 2

typedef struct error_in_struct {
    const char * name;
    int64_t code;
} error_in;  

typedef unsigned int error_mpp;
typedef int error_general;

#define DO_OR_RETURN_ERR_MPP(ERR, FUNC, ...) \
{ \
    unsigned int tmp_mpp_error_code; \
    tmp_mpp_error_code = FUNC(__VA_ARGS__); \
    if (tmp_mpp_error_code != HI_SUCCESS) { \
        ERR->name = #FUNC; \
        ERR->code = tmp_mpp_error_code; \
        return ERR_MPP; \
    } \
}

// printf("%s returning %u\n", #FUNC, tmp_mpp_error_code); \

#define DO_OR_RETURN_ERR_GENERAL(ERR, FUNC, ...) \
{ \
    int tmp_general_error_code = 0; \
    tmp_general_error_code = FUNC(__VA_ARGS__); \
    if (tmp_general_error_code != HI_SUCCESS) { \
        ERR->name = #FUNC; \
        ERR->code = tmp_general_error_code; \
        return ERR_GENERAL; \
    } \
}

#define RETURN_ERR_MPP(ERR, FUNC, CODE) \
    ERR->name = FUNC; \
    ERR->code = CODE; \
    return ERR_MPP

#define RETURN_ERR_GENERAL(ERR, FUNC, CODE) \
    ERR->name = FUNC; \
    ERR->code = CODE; \
    return ERR_GENERAL
