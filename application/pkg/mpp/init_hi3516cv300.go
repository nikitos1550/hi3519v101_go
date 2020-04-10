//+build arm
//+build hi3516cv300

package mpp

/*
#include "./include/mpp_v3.h"

#define ERR_NONE    0
#define ERR_MPP     1

int mpp3_sys_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_SYS_Exit();
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}

int mpp3_vb_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_VB_Exit();
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}

int mpp3_isp_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_ISP_Exit(0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/ko"
)

func systemInit() {
	//This family originally pack all reg init to sy_conf ko module
	ko.UnloadAll()
	ko.LoadAll()
}
