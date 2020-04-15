//+build arm
//+build hi3516cv500
//+build imx327,cmos_data_mipi,cmos_control_i2c,cmos_bus_0

package cmos

/*
#include "../include/mpp_v4.h"

#include "cmos.h"

#include <string.h>

int mpp_cmos_init(int *error_code) {
    *error_code = 0;

    return ERR_NONE;
}
*/
import "C"


import (    
    "unsafe"
)

var (
    cmosItem = cmos{   
        vendor: "Sony", 
        model: "imx327",   
        modes: []cmosMode {
            cmosMode {
                width: ,     
                height: ,    
                fps: ,  
                mipi: unsafe.Pointer(&C.),                                            
                viDev: unsafe.Pointer(&C.),                  
                clock: ,  
            },
        },
        control: cmosControl {
            bus: I2C, 
            busNum: 0,
        },
        data: MIPI,
    }
)

