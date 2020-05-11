//+build arm
//+build hi3516cv200

package cmos

/*
#include "cmos.h"
#include "cmos_hi3516cv200.h"

int mpp_cmos_init(int *error_code, unsigned char cmos) {
    *error_code = 0;

    if (cmos == 0) {
        *error_code = sensor_register_callback();
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

//import (
//    _"unsafe"
//)
var (
    cmosItems = []cmos {
        cmos {
            vendor: "JFX",
            model: "f22",   
            dcZeroBitOffset:  6, //WTF?
            modes: []cmosMode {
                cmosMode {
                    mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
                    viCrop:     crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,}, 
                    ispCrop:    crop{X0: 0, Y0: 0, Width: 1920, Height: 1080,},
                    width: 1920,
                    height: 1080,
                    fps: 30,    
                    bitness: 10,
                    data: DC,
                    dcSync: dcSyncAttr {
                        VSync:          DCVSyncField,
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
            bayer: BGGR,
        },
    }
)
