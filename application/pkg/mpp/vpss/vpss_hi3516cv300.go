//+build arm
//+build hi3516cv300

package vpss

/*
#include "../include/mpp_v3.h"

#include <string.h>

#define ERR_NONE                    0
#define ERR_MPP                     2
#define ERR_HI_MPI_VPSS_CreateGrp   3
#define ERR_HI_MPI_VPSS_StartGrp    4
#define ERR_HI_MPI_SYS_Bind         5

int mpp3_vpss_init(unsigned int *error_code) {
    *error_code = 0;

  VPSS_GRP VpssGrp = 0;
    //VPSS_CHN VpssChn = 0;
    VPSS_GRP_ATTR_S stVpssGrpAttr = {0};
   // VPSS_CHN_ATTR_S stVpssChnAttr = {0};
   // VPSS_CHN_MODE_S stVpssChnMode;


    VpssGrp = 0;
        stVpssGrpAttr.u32MaxW = 1920;
        stVpssGrpAttr.u32MaxH = 1080;
        stVpssGrpAttr.bIeEn = HI_FALSE;
        stVpssGrpAttr.bNrEn = HI_TRUE;
        stVpssGrpAttr.bHistEn = HI_FALSE;
        stVpssGrpAttr.enDieMode = VPSS_DIE_MODE_NODIE;
        stVpssGrpAttr.enPixFmt = PIXEL_FORMAT_YUV_SEMIPLANAR_420;

    *error_code = HI_MPI_VPSS_CreateGrp(VpssGrp, &stVpssGrpAttr);
    if (*error_code != HI_SUCCESS)
    {
        printf("HI_MPI_VPSS_CreateGrp failed with %#x!\n", *error_code);
        return ERR_MPP;
    }

    *error_code = HI_MPI_VPSS_StartGrp(VpssGrp);
    if (*error_code != HI_SUCCESS)
    {
        printf("HI_MPI_VPSS_StartGrp failed with %#x\n", *error_code);
        return ERR_MPP;
    }

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

     stSrcChn.enModId  = HI_ID_VIU;
        stSrcChn.s32DevId = 0;
        stSrcChn.s32ChnId = 0;
    
        stDestChn.enModId  = HI_ID_VPSS;
        stDestChn.s32DevId = 0;
        stDestChn.s32ChnId = 0;
    
        *error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
        if (*error_code != HI_SUCCESS)
        {
            printf("failed with %#x!\n", *error_code);
            return ERR_MPP;
        }



    return ERR_NONE;
}

int mpp3_vpss_sample_channel0(unsigned int *error_code) {
    *error_code = 0;

      VPSS_GRP VpssGrp = 0;
    VPSS_CHN VpssChn = 0;
    VPSS_GRP_ATTR_S stVpssGrpAttr = {0};
    VPSS_CHN_ATTR_S stVpssChnAttr = {0};
    VPSS_CHN_MODE_S stVpssChnMode;


 VpssChn = 0;
    stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
    stVpssChnMode.bDouble        = HI_FALSE;
    stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    stVpssChnMode.u32Width       = 1920;//3840;//4000;
    stVpssChnMode.u32Height      = 1080;//2160;//3000;
    stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE;//COMPRESS_MODE_SEG;
    memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));
    stVpssChnAttr.s32SrcFrameRate = 30;
    stVpssChnAttr.s32DstFrameRate = 30;

        *error_code = HI_MPI_VPSS_SetChnAttr(VpssGrp, VpssChn, &stVpssChnAttr);
        if (*error_code != HI_SUCCESS)
        {
            printf("HI_MPI_VPSS_SetChnAttr failed with %#x\n", *error_code);
            return ERR_MPP;
        }

         *error_code = HI_MPI_VPSS_SetChnMode(VpssGrp, VpssChn, &stVpssChnMode);
        if (*error_code != HI_SUCCESS)
        {
            printf("%s failed with %#x\n", __FUNCTION__, *error_code);
            return ERR_MPP;
        }     

     *error_code = HI_MPI_VPSS_EnableChn(VpssGrp, VpssChn);
    if (*error_code != HI_SUCCESS)
    {
        printf("HI_MPI_VPSS_EnableChn failed with %#x\n", *error_code);
        return ERR_MPP;
    }



    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/mpp/error"
	"log"
)

func Init() {
	var errorCode C.uint

	switch err := C.mpp3_vpss_init(&errorCode); err {
	case C.ERR_NONE:
		log.Println("C.mpp3_vpss_init() ok")
	case C.ERR_HI_MPI_VPSS_CreateGrp:
		log.Fatal("C.mpp3_vpss_init() HI_MPI_VPSS_CreateGrp() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_VPSS_StartGrp:
		log.Fatal("C.mpp3_vpss_init() HI_MPI_VPSS_StartGrp() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_SYS_Bind:
		log.Fatal("C.mpp3_vpss_init() HI_MPI_SYS_Bind() error ", error.Resolve(int64(errorCode)))
	default:
		log.Fatal("Unexpected return ", err, " of C.mpp1_vpss_init()")
	}
}

func SampleChannel0() {
	var errorCode C.uint

	switch err := C.mpp3_vpss_sample_channel0(&errorCode); err {
	case C.ERR_NONE:
		log.Println("C.mpp3_vpss_sample_channel0() ok")
	case C.ERR_MPP:
		log.Fatal("C.mpp3_vpss_sample_channel0() MPP error ", error.Resolve(int64(errorCode)))
	default:
		log.Fatal("Unexpected return ", err, " of C.mpp3_vpss_sample_channel0()")
	}

}

func CreateChannel(channel Channel) {}

func DestroyChannel(channel Channel) {}

