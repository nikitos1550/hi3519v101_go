//+build arm
//+build hi3516av100

package vi

/*
#include "../include/mpp.h"
#include "../../logger/logger.h"

#include <string.h>

#define ERR_NONE                    0
#define ERR_MPP                     1
#define ERR_GENERAL                 2

typedef struct hi3516av100_vi_init_in_struct {
    void *videv;
    unsigned int x0;
    unsigned int y0;
    unsigned int width;
    unsigned int height;
    unsigned int cmos_fps;
    unsigned int fps;

    unsigned char mirror;
    unsigned char flip;

    unsigned char ldc;
    int ldc_offset_x;
    int ldc_offset_y;
    int ldc_k;
} hi3516av100_vi_init_in;

static int hi3516av100_vi_init(int *error_code, hi3516av100_vi_init_in *in) {
    unsigned int mpp_error_code = 0;

    VI_DEV_ATTR_S  stViDevAttr;

    memset(&stViDevAttr, 0, sizeof(stViDevAttr));
    memcpy(&stViDevAttr, in->videv, sizeof(stViDevAttr));
    stViDevAttr.stDevRect.s32X = in->x0;
    stViDevAttr.stDevRect.s32Y = in->y0;
    stViDevAttr.stDevRect.u32Width  = in->width;
    stViDevAttr.stDevRect.u32Height = in->height;

    mpp_error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetDevAttr, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VI_EnableDev(0);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_EnableDev, mpp_error_code);
    }

    RECT_S stCapRect;

    stCapRect.s32X = in->x0;
    stCapRect.s32Y = in->y0;
    stCapRect.u32Width  = in->width;
    stCapRect.u32Height = in->height;

    VI_CHN_ATTR_S stChnAttr;

	memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));
    stChnAttr.enCapSel              = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width   = in->width;
    stChnAttr.stDestSize.u32Height  = in->height;
    stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

    if (in->mirror == 1) {
        stChnAttr.bMirror = HI_TRUE;
    } else {
        stChnAttr.bMirror = HI_FALSE;
    }

    if (in->flip == 1) {
        stChnAttr.bFlip = HI_TRUE;
    } else {
        stChnAttr.bFlip = HI_FALSE;
    }

    stChnAttr.s32SrcFrameRate = in->fps;
    stChnAttr.s32DstFrameRate = in->fps;

    stChnAttr.enCompressMode = COMPRESS_MODE_NONE;
    //TODO check family mpp datasheet
    //if (in->ldc == 1) {
    //    stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;
    //} else {
    //    stChnAttr.enCompressMode        = COMPRESS_MODE_SEG;
    //}

    mpp_error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetChnAttr, mpp_error_code);
    }

    if (in.ldc == 1) {
        VI_LDC_ATTR_S stLDCAttr;
        
        stLDCAttr.bEnable = HI_TRUE;
        stLDCAttr.stAttr.enViewType =   LDC_VIEW_TYPE_ALL;
                                        //LDC_VIEW_TYPE_CROP;
        stLDCAttr.stAttr.s32CenterXOffset = ldc_offset_x;
        stLDCAttr.stAttr.s32CenterYOffset = ldc_offset_y;
        stLDCAttr.stAttr.s32Ratio = ldc_k;

        mpp_error_code = HI_MPI_VI_SetLDCAttr(0, &stLDCAttr);
        if (mpp_error_code != HI_SUCCESS) {
            RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetLDCAttr, mpp_error_code);
        }

        //s32Ret = HI_MPI_VI_GetLDCAttr (ViChn, &stLDCAttr);
        //if (HI_SUCCESS != s32ret)
        //{
        //printf("Get vi LDC attr err:0x%x\n", s32ret);
        //return s32Ret;
        //}
    }

    mpp_error_code = HI_MPI_VI_EnableChn(0);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_EnableChn, mpp_error_code);
    }

	return ERR_NONE;
}


*/
import "C"

import (
    "flag"
    "application/pkg/mpp/cmos"
    "application/pkg/mpp/errmpp"
    "application/pkg/logger"
)

var(
    ldc bool
    ldcOffsetX int
    ldcOffsetY int
    ldcK  int
)

/*
When the resolution of the captured VI picture is not greater than D1, the value range of s32Ratio is [0, 480].
When the resolution of the captured VI picture is greater than D1 but not greater than 720p, the value range of s32Ratio is [0, 433].
When the resolution of the captured VI picture is greater than 720p but not greater than 1080p, the value range of s32Ratio is [0, 400].
When the resolution of the captured VI picture is greater than 1080p but not greater than 2304 x 1536, the value range of s32Ratio is [0, 300].
When the resolution of the captured VI picture is greater than 2304 x 1536 but not greater than 5 megapixels, the value range of s32Ratio is [0, 168].
*/

func init() {
    flag.BoolVar(&ldc, "vi-ldc", false, "LDC enable")
    flag.IntVar(&ldcOffsetX, "vi-ldc-offset-x", 0, "LDC x offset from center [-75;75]")
    flag.IntVar(&ldcOffsetY, "vi-ldc-offset-y", 0, "LDC y offset from center [-75;75]")
    flag.IntVar(&ldcK, "vi-ldc-k", 0, "LDC coefficient [0;168]")
}

func initFamily() error {
    var errorCode C.uint
    var in C.hi3516av100_vi_init_in

    if ldc == true {
        if ldcOffsetX < -75 || ldcOffsetX > 75 {
            logger.Log.Fatal().
                Int("ldc-offset-x", ldcOffsetX).
                Msg("vi-ldc-offset-x should be [-75;75]")
        }
        if ldcOffsetY < -75 || ldcOffsetY > 75 {
            logger.Log.Fatal().
                Int("ldc-offset-y", ldcOffsetY).
                Msg("vi-ldc-offset-y should be [-75;75]")
        }
        if ldcK < 0 || ldcK > 168 {
            logger.Log.Fatal().
                Int("ldc-k", ldcK).
                Msg("vi-ldc-k should be [0;168]")
        }

        in.ldc = 1
        in.ldc_offset_x = C.int(ldcOffsetX)
        in.ldc_offset_y = C.int(ldcOffsetY)
        in.ldc_k = C.int(ldcK)
    }

    in.videv = cmos.S.ViDev()
    in.width = C.uint(cmos.S.Width())
    in.height = C.uint(cmos.S.Height())
    in.cmos_fps = C.uint(cmos.S.Fps())
    in.fps = C.uint(cmos.S.Fps())

    if flipY == true {
        in.mirror = 1
    }
    if flipX == true {
        in.flip = 1
    }

    logger.Log.Trace().
        Uint("mirror", uint(in.mirror)).
        Uint("flip", uint(in.flip)).
        Uint("x0", uint(in.x0)).
        Uint("y0", uint(in.y0)).
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("cmos_fps", uint(in.cmos_fps)).
        Uint("fps", uint(in.fps)).
        Uint("ldc", uint(in.ldc)).
        Int("ldc-offset-x", int(in.ldc_offset_x)).
        Int("ldc-offset-y", int(in.ldc_offset_y)).
        Int("ldc-k", int(in.ldc_k)).
        Msg("VI params")

    err := C.hi3516av100_vi_init(&errorCode, &in)

    if err != 0 {
        return errmpp.New("funcname", int64(errorCode))
    }

    return nil
}

