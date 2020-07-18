//+build arm
//+build hi3516cv100

package cmos

/*
#include "cmos.h"

int sensor_register_callback_ar0130(void);
int sensor_register_callback_imx122(void);

int mpp_cmos_init(int *error_code, unsigned char cmos) {
    *error_code = 0;
    if (cmos == 0) {
        *error_code = sensor_register_callback_ar0130();
        if (*error_code != HI_SUCCESS) {
            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
        }
        return ERR_NONE;
    }

    if (cmos == 1) {
        *error_code = sensor_register_callback_imx122();    
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


var (
    cmosItems = []cmos {
        cmos {
            vendor: "Aptina",
            model: "ar0130",
            modes: []cmosMode {
                cmosMode {
                    //mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
                    //viCrop:     crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
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
                        TimingAct:      1280,
                        TimingHbb:      0,
                        TimingVfb:      0,
                        TimingVact:     720,
                        TimingVbb:      0,
                        TimingVbfb:     0,
                        TimingVbact:    0,
                        TimingVbbb:     0,
                    },
                    //mipi: nil,
                    clock: 27,
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
        cmos {
            vendor: "Sony",
            model: "imx122",
            modes: []cmosMode {
                cmosMode {
                    //mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
                    //viCrop:     crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
                    ispCrop:    crop{X0: 200, Y0: 18, Width: 1920, Height: 1080,},
                    width: 1920,
                    height: 1080,
                    fps: 30,
                    bitness: 12,
                    data: DC,
                    dcSync: dcSyncAttr {
                        VSync:          DCVSyncPulse,
                        VSyncNeg:       DCVSyncNegHigh,
                        HSync:          DCHSyncSignal,
                        HSyncNeg:       DCHSyncNegHigh,
                        VSyncValid:     DCVSyncValidPulse,
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
                    //mipi: nil,
                    clock: 37.125,
                    wdr: WDRNone,
                    description: "normal",
                },
                cmosMode {
                    //mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
                    //viCrop:     crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
                    ispCrop:    crop{X0: 200, Y0: 18, Width: 1280, Height: 720,},
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
                        VSyncValid:     DCVSyncValidPulse,
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
                    //mipi: nil,
                    clock: 37.125,
                    wdr: WDRNone,
                    description: "720p 30fps",
                },
                cmosMode {
                    //mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
                    //viCrop:     crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
                    ispCrop:    crop{X0: 200, Y0: 18, Width: 1280, Height: 720,},
                    width: 1280,
                    height: 720,
                    fps: 60,
                    bitness: 12,
                    data: DC,
                    dcSync: dcSyncAttr {
                        VSync:          DCVSyncPulse,
                        VSyncNeg:       DCVSyncNegHigh,
                        HSync:          DCHSyncSignal,
                        HSyncNeg:       DCHSyncNegHigh,
                        VSyncValid:     DCVSyncValidPulse,
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
                    //mipi: nil,
                    clock: 37.125,
                    wdr: WDRNone,
                    description: "720p 60fps",
                },
            },
            control: cmosControl {
                bus: SPI,
                busNum: 0,
            },
            data: DC,
            bayer: RGGB,
        },
    }
)

