//+build arm
//+build hi3516cv500

package mipi

/*
#include "../include/mpp_v4.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

#define ERR_NONE    0
#define ERR_GENERAL 1


int mpp4_mipi_init(int *error_code, void *mipi) {
    *error_code = 0;

     VI_VPSS_MODE_S      stVIVPSSMode;

    *error_code = HI_MPI_SYS_GetVIVPSSMode(&stVIVPSSMode);
    if (*error_code != HI_SUCCESS) return ERR_GENERAL;

    stVIVPSSMode.aenMode[0] = VI_OFFLINE_VPSS_OFFLINE;

    *error_code = HI_MPI_SYS_SetVIVPSSMode(&stVIVPSSMode);
    if (*error_code != HI_SUCCESS) return ERR_GENERAL; 
    
    //ISP_CTRL_PARAM_S    stIspCtrlParam;
    //HI_U32              u32FrameRate;

    //*error_code = HI_MPI_ISP_GetCtrlParam(0, &stIspCtrlParam);
    //if (*error_code != HI_SUCCESS) return ERR_GENERAL; 

    //u32FrameRate = 30;
    //stIspCtrlParam.u32StatIntvl  = u32FrameRate/30;

    //*error_code = HI_MPI_ISP_SetCtrlParam(0, &stIspCtrlParam);
    //if (*error_code != HI_SUCCESS) return ERR_GENERAL; 

    //VI_StartMIPI
    lane_divide_mode_t lane_divide_mode = LANE_DIVIDE_MODE_0;

    int fd;
    #define MIPI_DEV_NODE       "/dev/hi_mipi"
    fd = open(MIPI_DEV_NODE, O_RDWR);
    if (fd < 0) return ERR_GENERAL;

    *error_code = ioctl(fd, HI_MIPI_SET_HS_MODE, &lane_divide_mode);
    if (*error_code != HI_SUCCESS) {
        close(fd);  
        return ERR_GENERAL; 
    }

    combo_dev_t devno = 0;
    *error_code = ioctl(fd, HI_MIPI_ENABLE_MIPI_CLOCK, &devno);
    if (*error_code != HI_SUCCESS) {
        close(fd);
        return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_RESET_MIPI, &devno);
    if (*error_code != HI_SUCCESS) {
        close(fd);
        return ERR_GENERAL; 
    }

    sns_clk_source_t       SnsDev = 0;
    *error_code = ioctl(fd, HI_MIPI_ENABLE_SENSOR_CLOCK, &SnsDev);
    if (*error_code != HI_SUCCESS) {
        close(fd);
        return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_RESET_SENSOR, &SnsDev);
    if (*error_code != HI_SUCCESS) {
		close(fd);
		return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_SET_DEV_ATTR, mipi); //&MIPI_4lane_CHN0_SENSOR_IMX327_12BIT_2M_NOWDR_ATTR);
	if (*error_code != HI_SUCCESS) {
		close(fd);
		return ERR_GENERAL; 
	}

    *error_code = ioctl(fd, HI_MIPI_UNRESET_MIPI, &devno);
	if (*error_code != HI_SUCCESS) {
		close(fd);
		return ERR_GENERAL; 
	}

    *error_code = ioctl(fd, HI_MIPI_UNRESET_SENSOR, &SnsDev);
	if (*error_code != HI_SUCCESS) {
		close(fd);
		return ERR_GENERAL; 
	}

    close(fd);
	return ERR_NONE;

}

*/
import "C"

import (
        "application/pkg/logger"

        "application/pkg/mpp/cmos"
)

func Init() {
    var errorCode C.int

    switch err := C.mpp4_mipi_init(&errorCode, cmos.Mipi() ); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp4_mipi_init() ok")
    default:
        logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp4_mipi_init()")
    }

}

