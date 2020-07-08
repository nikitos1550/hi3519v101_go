//+build arm
//+build hi3516cv300

package cmos

/*
#include "cmos.h"
#include "cmos_hi3516cv300.h"

int sensor_register_callback_imx290_lvds(void);
int sensor_register_callback_imx323_spi(void);
int sensor_register_callback_imx323_i2c(void);

int mpp_cmos_init(int *error_code, unsigned char cmos) {
    *error_code = 0;

    if (cmos == 0) {
        *error_code = sensor_register_callback_imx290_lvds();
        if (*error_code != HI_SUCCESS) {
            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
        }
        return ERR_NONE;
    }

    if (cmos == 1) {
        *error_code = sensor_register_callback_imx323_spi();
        if (*error_code != HI_SUCCESS) {
            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
        }
        return ERR_NONE;
    }

    if (cmos == 2) {
        *error_code = sensor_register_callback_imx323_i2c();    
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
                    viCrop:     crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920,
                    height: 1080,
                    fps: 30,
                    bitness: 12,
                    mipiLVDSAttr: unsafe.Pointer(&C.imx290_lvds_mode_0),
                    clock: 37.125,
                    wdr: WDRNone,
                    description: "normal",
                },
            },
            control: cmosControl {
                bus: I2C,
                busNum: 0,
            },
            data: LVDS,
            bayer: GBRG, //TODO
        },
        cmos {
            vendor: "Sony",
            model: "imx323_spi",
            dcZeroBitOffset:  4, //typical for hisi
            modes: []cmosMode {
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //???
                    viCrop:     crop{X0: 200, Y0: 20, Width: 1920, Height: 1080,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920,
                    height: 1080,
                    fps: 30,
                    bitness: 12,
                    //data: DC,
                    dcSync: dcSyncAttr {
                        VSync:          DCVSyncPulse,
                        VSyncNeg:       DCVSyncNegHigh,
                        HSync:          DCHSyncSignal,
                        HSyncNeg:       DCHSyncNegHigh,
                        VSyncValid:     DCVSyncValidSignal,
                        VSyncValidNeg:  DCVSyncValidNegHigh,
                        TimingHfb:      0,
                        TimingAct:      1920,
                        TimingHbb:      0,
                        TimingVfb:      0,
                        TimingVact:     1080,
                        TimingVbb:      0,
                        TimingVbfb:     0,
                        TimingVbact:    0,
                        TimingVbbb:     0,
                    },
                    //mipi: unsafe.Pointer(&C.MIPI_CMOS323_ATTR),
                    clock: 37.125,
                    wdr: WDRNone,
                    description: "normal",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //???
                    viCrop:     crop{X0: 200, Y0: 20, Width: 1280, Height: 720,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1280, Height: 720,},
                    width: 1280,
                    height: 720,
                    fps: 30,
                    bitness: 12,
                    data: DC,
                    dcSync: dcSyncAttr {
                        VSync:          DCVSyncPulse,
                        VSyncNeg:       DCVSyncNegHigh,
                        HSync:          DCHSyncSignal,
                        HSyncNeg:       DCHSyncNegHigh,
                        VSyncValid:     DCVSyncValidSignal,
                        VSyncValidNeg:  DCVSyncValidNegHigh,
                        TimingHfb:      0,
                        TimingAct:      1920,
                        TimingHbb:      0,
                        TimingVfb:      0,
                        TimingVact:     1080,
                        TimingVbb:      0,
                        TimingVbfb:     0,
                        TimingVbact:    0,
                        TimingVbbb:     0,
                    },
                    //mipi: unsafe.Pointer(&C.MIPI_CMOS323_ATTR),
                    clock: 37.125,
                    wdr: WDRNone,
                    description: "720p 30fps 12bit",
                },
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //???
                    viCrop:     crop{X0: 200, Y0: 20, Width: 1280, Height: 720,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1280, Height: 720,},
                    width: 1280,
                    height: 720,
                    fps: 60,
                    bitness: 10,
                    data: DC,
                    dcSync: dcSyncAttr {
                        VSync:          DCVSyncPulse,
                        VSyncNeg:       DCVSyncNegHigh,
                        HSync:          DCHSyncSignal,
                        HSyncNeg:       DCHSyncNegHigh,
                        VSyncValid:     DCVSyncValidSignal,
                        VSyncValidNeg:  DCVSyncValidNegHigh,
                        TimingHfb:      0,
                        TimingAct:      1280,
                        TimingHbb:      0,
                        TimingVfb:      0,
                        TimingVact:     720,
                        TimingVbb:      0,
                        TimingVbfb:     0,
                        TimingVbact:    0,
                        TimingVbbb:     0,
                    },
                    //mipi: unsafe.Pointer(&C.MIPI_CMOS323_ATTR),
                    clock: 37.125,
                    wdr: WDRNone,
                    description: "720p 60fps 10bit",
                },
            },
            control: cmosControl {
                bus: SPI,
                busNum: 0,
            },
            data: DC,
            bayer: RGGB,
        },
        cmos {
            vendor: "Sony",
            model: "imx323_i2c",
            dcZeroBitOffset:  4, //typical for hisi
            modes: []cmosMode {
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //???
                    viCrop:     crop{X0: 200, Y0: 20, Width: 1920, Height: 1080,},
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920,
                    height: 1080,
                    fps: 30,
                    bitness: 12,
                    //data: DC,
                    dcSync: dcSyncAttr {
                        VSync:          DCVSyncPulse,
                        VSyncNeg:       DCVSyncNegHigh,
                        HSync:          DCHSyncSignal,
                        HSyncNeg:       DCHSyncNegHigh,
                        VSyncValid:     DCVSyncValidSignal,
                        VSyncValidNeg:  DCVSyncValidNegHigh,
                        TimingHfb:      0,
                        TimingAct:      1920,
                        TimingHbb:      0,
                        TimingVfb:      0,
                        TimingVact:     1080,
                        TimingVbb:      0,
                        TimingVbfb:     0,
                        TimingVbact:    0,
                        TimingVbbb:     0,
                    },
                    //mipi: unsafe.Pointer(&C.MIPI_CMOS323_ATTR),
                    clock: 37.125,
                    wdr: WDRNone,
                    description: "normal",
                },
            },
            control: cmosControl {
                bus: I2C,
                busNum: 0,
            },
            data: DC,
            bayer: RGGB,
        },
    }
)

