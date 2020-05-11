//+build arm
//+build hi3516av100

package cmos

/*
#include "cmos.h"
#include "cmos_hi3516av100.h"

int sensor_register_callback(void);
int sensor_register_callback_imx178_lvds(void);
int sensor_register_callback_ov4689_mipi(void);

int mpp_cmos_init(int *error_code, unsigned char cmos) {
    *error_code = 0;

    if (cmos == 0) {
        *error_code = sensor_register_callback();
        if (*error_code != HI_SUCCESS) {
            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
        }
        return ERR_NONE;
    }

    if (cmos == 1) {
        *error_code = sensor_register_callback_imx178_lvds();
        if (*error_code != HI_SUCCESS) {
            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
        }
        return ERR_NONE;
    }

    if (cmos == 2) {
        *error_code = sensor_register_callback_ov4689_mipi();
        if (*error_code != HI_SUCCESS) {
            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
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
            model: "imx290_lvds",
            modes: []cmosMode {
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    viCrop:     crop{X0: 0, Y0: 30, Width: 1920, Height: 1080,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920, 
                    height: 1080,
                    fps: 30,    
                    bitness: 12,
                    mipiLVDSAttr: unsafe.Pointer(&C.imx290_lvds_mode_0_1),
                    clock: 37.125,
                    wdr: WDRNone,
                    description: "1080p 30fps",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,}, 
                    viCrop:     crop{X0: 0, Y0: 30, Width: 1920, Height: 1080,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920, 
                    height: 1080,
                    fps: 60,    
                    bitness: 12,
                    mipiLVDSAttr: unsafe.Pointer(&C.imx290_lvds_mode_0_1),
                    clock: 37.125,
                    wdr: WDRNone,
                    description: "1080p 60fps",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,}, 
                    viCrop:     crop{X0: 0, Y0: 30, Width: 1920, Height: 1080,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920, 
                    height: 1080,
                    fps: 30,    
                    bitness: 10,
                    mipiLVDSAttr: unsafe.Pointer(&C.imx290_lvds_mode_2),
                    clock: 37.125,
                    wdr: WDR2TO1,
                    description: "1080p 30fps WDR",
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
            model: "imx178",
            modes: []cmosMode {
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 2592, Height: 1944,},
                    viCrop:     crop{X0: 0, Y0: 20, Width: 2592, Height: 1944,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 2592, Height: 1944,},
                    width: 2592,
                    height: 1944,
                    fps: 30,
                    bitness: 12,
                    mipiLVDSAttr: unsafe.Pointer(&C.imx178_lvds_mode_0),
                    clock: 25,
                    wdr: WDRNone,
                    description: "1944p 30fps",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    viCrop:     crop{X0: 0, Y0: 20, Width: 1920, Height: 1080,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920,
                    height: 1080,
                    fps: 60,
                    bitness: 12,
                    mipiLVDSAttr: unsafe.Pointer(&C.imx178_lvds_mode_1),
                    clock: 25,
                    wdr: WDRNone,
                    description: "1080p 60fps",
                },
            },
            control: cmosControl {
                bus: I2C,
                busNum: 0,
            },
            data: SubLVDS,
            bayer: RGGB,
        },
        cmos {
            vendor: "Omnivision",
            model: "ov4689",
            //TODO More modes can be added: 2304x1296@30 none, 2to1; 2048x1520@30 none, 2to1
            modes: []cmosMode { 
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    width: 2592,
                    height: 1520,
                    fps: 30,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.ov4689_mipi_mode_0),
                    clock: 24,
                    wdr: WDR2TO1FFR,
                    description: "2592x1520 30fps WDR2to1FFR",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    width: 2592,
                    height: 1520,
                    fps: 25,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.ov4689_mipi_mode_0),
                    clock: 24,
                    wdr: WDR2TO1,
                    description: "2592x1520 25fps WDR2to1",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 2592, Height: 1520,},
                    width: 2592,
                    height: 1520,
                    fps: 30,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.ov4689_mipi_mode_0),
                    clock: 24,
                    wdr: WDRNone,
                    description: "2592x1520 30fps",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920,
                    height: 1080,
                    fps: 30,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.ov4689_mipi_mode_0),
                    clock: 24,
                    wdr: WDR2TO1,
                    description: "1080p 30fps WDR2to1",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920,
                    height: 1080,
                    fps: 30,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.ov4689_mipi_mode_0),
                    clock: 24,
                    wdr: WDRNone,
                    description: "1080p 30fps",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920,
                    height: 1080,
                    fps: 60,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.ov4689_mipi_mode_0),
                    clock: 24,
                    wdr: WDRNone,
                    description: "1080p 60fps",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 1280, Height: 720,},
                    viCrop:     crop{X0: 0, Y0: 0, Width: 1280, Height: 720,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1280, Height: 720,},
                    width: 1280,
                    height: 720,
                    fps: 180,
                    bitness: 12,
                    mipiMIPIAttr: unsafe.Pointer(&C.ov4689_mipi_mode_0),
                    clock: 24,
                    wdr: WDRNone,
                    description: "1080p 180fps",
                },

            },
            control: cmosControl {
                bus: I2C,
                busNum: 0,
            },
            data: MIPI,
            bayer: BGGR,
        },
    }
)
