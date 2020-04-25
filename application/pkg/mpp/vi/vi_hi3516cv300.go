//+build arm
//+build hi3516cv300

package vi

/*
#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>

typedef struct hi3516cv300_vi_init_in_struct {
    void *videv;

    unsigned int cmos_width;
    unsigned int cmos_height;

    unsigned int x0;
    unsigned int y0;
    unsigned int width;
    unsigned int height;
    unsigned int cmos_fps;
    unsigned int fps;

    unsigned char mirror;
    unsigned char flip;

    //unsigned char ldc;
    //int ldc_offset_x;
    //int ldc_offset_y;
    //int ldc_k;
} hi3516cv300_vi_init_in;


static int hi3516cv300_vi_init(error_in *err, hi3516cv300_vi_init_in * in) {
    unsigned int mpp_error_code = 0;

    ISP_WDR_MODE_S stWdrMode;
    VI_DEV_ATTR_S  stViDevAttr;
    
    memset(&stViDevAttr,0,sizeof(stViDevAttr));

    memcpy(&stViDevAttr, in->videv, sizeof(stViDevAttr));

    //stViDevAttr.stDevRect.s32X = 0;
    //stViDevAttr.stDevRect.s32Y = 0;
    //stViDevAttr.stDevRect.u32Width  = in->cmos_width;
    //stViDevAttr.stDevRect.u32Height = in->cmos_height;

    mpp_error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetDevAttr, mpp_error_code);
    }

    VI_WDR_ATTR_S stWdrAttr;

    // WDR_MODE_NONE or WDR_MODE_2To1_LINE TODO
    stWdrAttr.enWDRMode = WDR_MODE_NONE;
    stWdrAttr.bCompress = HI_FALSE;

    mpp_error_code = HI_MPI_VI_SetWDRAttr(0, &stWdrAttr);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetWDRAttr, mpp_error_code);
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

    stChnAttr.enCapSel = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width = in->width;
    stChnAttr.stDestSize.u32Height =  in->height;
    stChnAttr.enPixFormat = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

    if (in->mirror == 1) {
        GO_LOG_VI(LOGGER_TRACE, "VI mirror on");
        stChnAttr.bMirror               = HI_TRUE;
    } else {
        stChnAttr.bMirror               = HI_FALSE;
    }

    if (in->flip == 1) {
        GO_LOG_VI(LOGGER_TRACE, "VI flip on");
        stChnAttr.bFlip                 = HI_TRUE;
    } else {
        stChnAttr.bFlip                 = HI_FALSE;
    }


    stChnAttr.s32SrcFrameRate = in->cmos_fps;
    stChnAttr.s32DstFrameRate = in->fps;
    //stChnAttr.enCompressMode = COMPRESS_MODE_SEG;
    stChnAttr.enCompressMode = COMPRESS_MODE_NONE;

    mpp_error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetChnAttr, mpp_error_code);
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
    //"flag"
    "application/pkg/mpp/cmos"
    "application/pkg/mpp/errmpp"
    "application/pkg/logger"
)

var(
    //ldc bool
    //ldcOffsetX int
    //ldcOffsetY int
    //ldcK  int
)

func init() {
    //flag.BoolVar(&ldc, "vi-ldc", false, "LDC enable")
    //flag.IntVar(&ldcOffsetX, "vi-ldc-offset-x", 0, "LDC x offset from center [-127;127]")
    //flag.IntVar(&ldcOffsetY, "vi-ldc-offset-y", 0, "LDC y offset from center [-127;127]")
    //flag.IntVar(&ldcK, "vi-ldc-k", 0, "LDC coefficient [-300;500]")
}

func initFamily() error {
    var inErr C.error_in
    var in C.hi3516cv300_vi_init_in

    /*
    if ldc == true {
        if ldcOffsetX < -127 || ldcOffsetX > 127 {
            logger.Log.Fatal().
                Int("ldc-offset-x", ldcOffsetX).
                Msg("vi-ldc-offset-x should be [-127;127]")
        }
        if ldcOffsetY < -127 || ldcOffsetY > 127 {
            logger.Log.Fatal().
                Int("ldc-offset-y", ldcOffsetY).
                Msg("vi-ldc-offset-y should be [-127;127]")
        }
        if ldcK < -300 || ldcK > 500 {
            logger.Log.Fatal().
                Int("ldc-k", ldcK).
                Msg("vi-ldc-k should be [-300;500]")
        }

        in.ldc = 1
        in.ldc_offset_x = C.int(ldcOffsetX)
        in.ldc_offset_y = C.int(ldcOffsetY)
        in.ldc_k = C.int(ldcK)
    }
    */

    in.videv = cmos.S.ViDev()
    in.cmos_width = C.uint(cmos.S.Width())
    in.cmos_height = C.uint(cmos.S.Height())
    in.x0 = C.uint(x0)
    in.y0 = C.uint(y0)
    in.width = C.uint(width)
    in.height = C.uint(height)
    in.cmos_fps = C.uint(cmos.S.Fps())
    in.fps = C.uint(fps)

    if flipY == true {
        in.mirror = 1
    }
    if flipX == true {
        in.flip = 1
    }

    logger.Log.Trace().
        Uint("mirror", uint(in.mirror)).
        Uint("flip", uint(in.flip)).
        Uint("cmos_width", uint(in.cmos_width)).
        Uint("cmos_height", uint(in.cmos_height)).
        Uint("x0", uint(in.x0)).
        Uint("y0", uint(in.y0)).
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("cmos_fps", uint(in.cmos_fps)).
        Uint("fps", uint(in.fps)).
        //Uint("ldc", uint(in.ldc)).
        //Int("ldc-offset-x", int(in.ldc_offset_x)).
        //Int("ldc-offset-y", int(in.ldc_offset_y)).
        //Int("ldc-k", int(in.ldc_k)).
        Msg("VI params")

    err := C.hi3516cv300_vi_init(&inErr, &in)

    if err != 0 {
        return errmpp.New(uint(inErr.f), uint(inErr.mpp))
    }

    return nil
}
