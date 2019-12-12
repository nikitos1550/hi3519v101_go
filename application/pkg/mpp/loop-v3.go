//+build hi3516av200 hi3516cv300

package mpp

/*
#include "include/hi3516av200_mpp.h"

#define ERR_NONE                    0

extern void go_callback_receive_data();

int mpp3_data_loop(unsigned int *error_code) {
    *error_code = 0;

    int fd = HI_MPI_VENC_GetFd(0);

    go_callback_receive_data();
}


*/
import "C"

import (
    "log"
)

func StartLoop() {
    var errorCode C.uint

    switch err := C.mpp3_data_loop(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_data_loop() ok")
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_data_loop()")
    }

}
