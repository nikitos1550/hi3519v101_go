//+build arm
//+build hi3516cv500

package vi

/*

#include "../include/mpp_v4.h"
#include <string.h>

#define ERR_NONE                    0
#define ERR_HI_MPI_VI_SetDevAttr        2
#define ERR_HI_MPI_VI_EnableDev     3
#define ERR_HI_MPI_VI_SetChnAttr    4
#define ERR_HI_MPI_VI_EnableChn     5

static VI_PIPE_ATTR_S PIPE_ATTR_2592x1944_RAW12_420_3DNR_RFR =
{
    VI_PIPE_BYPASS_NONE, HI_FALSE, HI_FALSE,
    2592, 1944,
    PIXEL_FORMAT_RGB_BAYER_12BPP,
    COMPRESS_MODE_LINE,
    DATA_BITWIDTH_12,
    HI_FALSE,
    {
        PIXEL_FORMAT_YVU_SEMIPLANAR_420,
        DATA_BITWIDTH_8,
        VI_NR_REF_FROM_RFR,
        COMPRESS_MODE_NONE
    },
    HI_FALSE,                        
    {-1, -1}                
};

static VI_CHN_ATTR_S CHN_ATTR_2592x1944_420_SDR8_LINEAR =
{
    {2592, 1944},
    PIXEL_FORMAT_YVU_SEMIPLANAR_420,
    DYNAMIC_RANGE_SDR8,
    VIDEO_FORMAT_LINEAR,
    COMPRESS_MODE_NONE,
    0,      0,
    0,
    {-1, -1}
};


int mpp4_vi_init(unsigned int *error_code, void *videv, unsigned int width, unsigned int height, unsigned int fps) {
    *error_code = 0;


      //VI_StartDev
        VI_DEV_ATTR_S       stViDevAttr;
        hi_memcpy(&stViDevAttr, sizeof(VI_DEV_ATTR_S), videv, sizeof(VI_DEV_ATTR_S));
        stViDevAttr.stWDRAttr.enWDRMode = WDR_MODE_NONE;
        //stViDevAttr.enDataRate = DATA_RATE_X2; //???????????????

        *error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_SetDevAttr failed with %#x!\n", *error_code);
            return -1;
        }

        *error_code = HI_MPI_VI_EnableDev(0);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_EnableDev failed with %#x!\n", *error_code);
            return -1;
        }        

        //VI_BindPipeDev
        VI_DEV_BIND_PIPE_S  stDevBindPipe = {0};
        stDevBindPipe.u32Num = 1;

        *error_code = HI_MPI_VI_SetDevBindPipe(0, &stDevBindPipe);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_SetDevBindPipe failed with %#x!\n", *error_code);
            return -1;
        }
        //VI_StartViPipe
        VI_PIPE_ATTR_S  stPipeAttr;

        hi_memcpy(&stPipeAttr, sizeof(VI_PIPE_ATTR_S), &PIPE_ATTR_2592x1944_RAW12_420_3DNR_RFR, sizeof(VI_PIPE_ATTR_S));

        *error_code = HI_MPI_VI_CreatePipe(0, &stPipeAttr);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_CreatePipe failed with %#x!\n", *error_code);
            return -1;
        }

        *error_code = HI_MPI_VI_StartPipe(0);
        if (*error_code != HI_SUCCESS) {
            //HI_MPI_VI_DestroyPipe(ViPipe);
            printf("HI_MPI_VI_StartPipe failed with %#x!\n", *error_code);
            return -1;
        }
        //VI_StartViChn
        VI_CHN_ATTR_S       stChnAttr;

        hi_memcpy(&stChnAttr, sizeof(VI_CHN_ATTR_S), &CHN_ATTR_2592x1944_420_SDR8_LINEAR, sizeof(VI_CHN_ATTR_S));

        stChnAttr.enDynamicRange = DYNAMIC_RANGE_SDR8;
        stChnAttr.enVideoFormat  = VIDEO_FORMAT_LINEAR;
        stChnAttr.enPixelFormat  = PIXEL_FORMAT_YVU_SEMIPLANAR_420;
        stChnAttr.enCompressMode = COMPRESS_MODE_SEG;

        *error_code = HI_MPI_VI_SetChnAttr(0, 0, &stChnAttr);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_SetChnAttr failed with %#x!\n", *error_code);
            return -1;
        }

        *error_code = HI_MPI_VI_EnableChn(0, 0);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_EnableChn failed with %#x!\n", -1);
            return -1;
        }



    return ERR_NONE;
}


*/
import "C"

import (
        "application/pkg/logger"

    "application/pkg/mpp/error"
"application/pkg/mpp/cmos"
)

func Init() {
    var errorCode C.uint

    //switch err := C.mpp4_vi_init(&errorCode); err {
      switch err := C.mpp4_vi_init(&errorCode, cmos.ViDev(), C.uint(cmos.Width()), C.uint(cmos.Height()), C.uint(cmos.Fps())); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp4_vi_init() ok")
    case C.ERR_HI_MPI_VI_SetDevAttr:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_SetDevAttr()").
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp4_vi_init() error")
    case C.ERR_HI_MPI_VI_EnableDev:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_EnableDev()").
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp4_vi_init() error")

    case C.ERR_HI_MPI_VI_SetChnAttr:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_SetChnAttr()"). 
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp4_vi_init() error")
    case C.ERR_HI_MPI_VI_EnableChn:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_EnableChn()"). 
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp4_vi_init() error")
    default:
        logger.Log.Fatal().
                Int("error", int(err)).
                Msg("C.mpp4_vi_init() Unexpected return")

    }
}

