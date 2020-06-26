//+build arm
//+build hi3516ev200

package cmos

/*
#include "cmos.h"
#include "cmos_hi3516ev200.h"

int mpp_cmos_init(int *error_code, unsigned char cmos) {
    *error_code = 0;

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    strncpy(stAeLib.acLibName,  HI_AE_LIB_NAME,     sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME,    sizeof(HI_AWB_LIB_NAME));

    if (cmos == 0) {
        ISP_SNS_OBJ_S *cmos = &stSnsImx307_2l_Obj;

        if (cmos->pfnRegisterCallback != HI_NULL) {
            *error_code = cmos->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
        } else {
            return ERR_GENERAL;
        }

        ISP_SNS_COMMBUS_U uSnsBusInfo;
        //ISP_SNS_TYPE_E enBusType;

        //enBusType = ISP_SNS_I2C_TYPE;
        uSnsBusInfo.s8I2cDev = 0;

        *error_code = cmos->pfnSetBusInfo(0, uSnsBusInfo);
        if (*error_code != HI_SUCCESS) {
            printf("set sensor bus info failed with %#x!\n", *error_code);
            return ERR_GENERAL;
        }

        return ERR_NONE;
    }

    if (cmos == 1) {
        ISP_SNS_OBJ_S *cmos = &stSnsImx335Obj;
    
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
        	model: "imx307_2l",
        	modes: []cmosMode {
            	cmosMode {
                	mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,},
                	viCrop:     crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                	ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                	width: 1920,
                	height: 1080,
                	fps: 30,
                	bitness: 12,
                	mipiMIPIAttr: unsafe.Pointer(&C.imx307_mode_0),
                	clock: 0,
                	wdr: WDRNone,
                	description: "normal",
            	},
                //cmosMode { //seems hi3516ev200 chip doesn support wdr 2 to 1 frame
                //    mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,},
                //    viCrop:     crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                //    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                //    width: 1920,
                //    height: 1080,
                //    fps: 30,
                //    bitness: 10,
                //    mipiMIPIAttr: unsafe.Pointer(&C.imx307_mode_1),
                //    clock: 0,
                //    wdr: WDR2TO1L,
                //    description: "wdr",
                //},
        	},
        	control: cmosControl {
            	bus: I2C,
            	busNum: 0,
        	},
        	data: MIPI,
        	bayer: RGGB,
    	},
        cmos {   
            vendor: "Sony", 
            model: "imx335",   
            modes: []cmosMode {
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 2592, Height: 1944,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 2592, Height: 1944,},
                    width: 2592, 
                    height: 1944,
                    fps: 30,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.imx335_mode_0),
                    clock: 0,
                    wdr: WDRNone,
                    description: "5M",
                },
                //cmosMode {
                //    mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,},
                //    viCrop:     crop{X0: 0, Y0: 0, Width: 2592, Height: 1944,},
                //    ispCrop:    crop{X0: 0, Y0: 0, Width: 2592, Height: 1944,},
                //    width: 2592, 
                //    height: 1944,
                //    fps: 30,
                //    bitness: 10,
                //    mipiMIPIAttr: unsafe.Pointer(&C.imx335_mode_1),                                    
                //    clock: 0,
                //    wdr: WDR2TO1L,
                //    description: "5M wdr",
                //},
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    width: 2592, 
                    height: 1520,
                    fps: 30,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.imx335_mode_2),
                    clock: 0,
                    wdr: WDRNone,
                    description: "4M",
                },
                //cmosMode {
                //    mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,},
                //    viCrop:     crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                //    ispCrop:    crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                //    width: 2592, 
                //    height: 1520,
                //    fps: 30,
                //    bitness: 10,
                //    mipiMIPIAttr: unsafe.Pointer(&C.imx335_mode_3),                                    
                //    clock: 0,
                //    wdr: WDR2TO1L,
                //    description: "4M wdr",
                //},
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
