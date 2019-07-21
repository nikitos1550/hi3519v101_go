#include "himpp3.h"

int himpp3_sys_init() {
	ret = HI_MPI_SYS_Exit();
        if (ret != HI_SUCCESS) {
                fprintf(stderr, "HI_MPI_SYS_Exit failed: ");
                resolve_mppv2_errors(ret);
                return 1;
        }

        ret = HI_MPI_VB_Exit();
        if (ret != HI_SUCCESS) {
                fprintf(stderr, "HI_MPI_VB_Exit failed: ");
                resolve_mppv2_errors(ret);
                return 1;
        }



}

int himpp3_vi_init() {


}

int himpp3_mipi_init() {


}


int himpp3_isp_init() {


}

int himpp3_vpss_init() {


}

int himpp3_venc_init() {



}

