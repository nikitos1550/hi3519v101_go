//+build arm
//+build hi3516av200

package cmos

/*
#include "cmos.h"
#include "cmos_hi3516av200.h"

int mpp_cmos_init(int *error_code, unsigned char cmos) {
    *error_code = 0;

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    strncpy(stAeLib.acLibName,  HI_AE_LIB_NAME,     sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME,    sizeof(HI_AWB_LIB_NAME));

    if (cmos == 0) {
        ISP_SNS_OBJ_S *cmos = &stSnsImx274Obj;

        if (cmos->pfnRegisterCallback != HI_NULL) {
            *error_code = cmos->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
        } else {
            return ERR_GENERAL;
        }

        return ERR_NONE;
    }

    if (cmos == 1) {
        ISP_SNS_OBJ_S *cmos = &stSnsImx226Obj;

        if (cmos->pfnRegisterCallback != HI_NULL) {
            *error_code = cmos->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
        } else {
            return ERR_GENERAL;
        }

        return ERR_NONE;
    }

    if (cmos == 2) {
        ISP_SNS_OBJ_S *cmos = &stSnsOs08a10Obj;

        if (cmos->pfnRegisterCallback != HI_NULL) {
            *error_code = cmos->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
        } else {
            return ERR_GENERAL;
        }

        return ERR_NONE;
    }

    *error_code = 999;
    return ERR_GENERAL;
}
*/
import "C"

import (
    "unsafe"
)

var (
    cmosItems = []cmos {
        cmos{
            vendor: "Sony",
            model: "imx274",
            modes: []cmosMode {
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    width: 3840,
                    height: 2160,
                    fps: 30,
                    bitness: 12,
                    mipiLVDSAttr: unsafe.Pointer(&C.imx274_mode_0),
                    clock: 72,
                    wdr: WDRNone,
                    description: "normal",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    width: 3840,
                    height: 2160,
                    fps: 30,
                    bitness: 10,
                    mipiLVDSAttr: unsafe.Pointer(&C.imx274_mode_1),
                    clock: 72,
                    wdr: WDR2TO1,
                    description: "wdr",
                },
            },
            control: cmosControl {
                bus: SPI,
                busNum: 0,
            },
            data: LVDS,
            bayer: RGGB,
        },
        cmos {
            vendor: "Sony",
            model: "imx226",
            modes: []cmosMode {
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    width: 3840,
                    height: 2160,
                    fps: 30,
                    bitness: 12,
                    mipiLVDSAttr: unsafe.Pointer(&C.imx226_mode_0),
                    clock: 72,
                    wdr: WDRNone,
                    description: "normal",
                },
            },
            control: cmosControl {
                bus: SPI,
                busNum: 0,
            },
            data: LVDS,
            bayer: RGGB,
        },
        cmos {
            vendor: "OmniVision",
            model: "os08a10",
            modes: []cmosMode {
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    width: 3840,
                    height: 2160,
                    fps: 30,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.os08a10_mode_0),
                    clock: 24,
                    wdr: WDRNone,
                    description: "normal",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 3840, Height: 2160,},
                    width: 3840,
                    height: 2160,
                    fps: 30,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.os08a10_mode_1),
                    clock: 24,
                    wdr: WDR2TO1,
                    description: "wdr",
                },
            },
            control: cmosControl {
                bus: I2C,
                busNum: 0,
            },
            data: MIPI,
            bayer: RGGB,
        },
    }
)

