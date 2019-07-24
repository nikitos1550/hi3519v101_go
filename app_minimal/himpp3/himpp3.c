#include "himpp3_internal.h"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <pthread.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/uio.h>
#include <stdint.h>
#include <sys/ioctl.h>
#include <fcntl.h>

#ifdef TEST_C_APP
int main(void) {
	himpp3_sys_init();
        himpp3_mipi_isp_init();
        himpp3_vi_init();
	himpp3_vpss_init();
        //himpp3_venc_init();
	while(1) ;;;
	return 0;
}
#endif

int himpp3_sys_init(/*int * error_func, int * error_code*/) {
        //int ret;

        int error_code = 0;
        //*error_func = HIMPP3_ERROR_FUNC_NONE;
        //*error_code = 0;

	error_code = HI_MPI_SYS_Exit();
        if (error_code != HI_SUCCESS) {
                //*error_func = HIMPP3_ERROR_FUNC_HI_MPI_SYS_Exit;
                return -1;
        }
        printf("HI_MPI_SYS_Exit ok\n");

        error_code = HI_MPI_VB_Exit();
        if (error_code != HI_SUCCESS) {
                //*error_func = HIMPP3_ERROR_FUNC_HI_MPI_VB_Exit;
                return -1;
        }
        printf("HI_MPI_VB_Exit ok\n");


        VB_CONF_S stVbConf;
        
        memset(&stVbConf, 0, sizeof(VB_CONF_S));
        stVbConf.u32MaxPoolCnt                  = 128;
        stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(3840, 64) * CEILING_2_POWER(2160, 64) * 3);
        stVbConf.astCommPool[0].u32BlkCnt       = 5;

        error_code = HI_MPI_VB_SetConf(&stVbConf);
	if(error_code != HI_SUCCESS) {
		//*error_func = HIMPP3_ERROR_FUNC_HI_MPI_VB_SetConf;
		return -1;
	}
        printf("HI_MPI_VB_SetConf ok\n");

	error_code = HI_MPI_VB_Init();
	if (error_code != HI_SUCCESS) {
		//*error_func = HIMPP3_ERROR_FUNC_???;
		return -1;
	}
        printf("HI_MPI_VB_Init ok\n");

        MPP_SYS_CONF_S	stSysConf;
	
        stSysConf.u32AlignWidth = 64;

	error_code = HI_MPI_SYS_SetConf(&stSysConf);
	if (error_code != HI_SUCCESS) {
		//*error_func = HIMPP3_ERROR_FUNC_HI_MPI_SYS_SetConf;
		return -1;
	}
        printf("HI_MPI_SYS_SetConf ok\n");

	error_code = HI_MPI_SYS_Init();
	if(error_code != HI_SUCCESS) {
		//*error_func = HIMPP3_ERROR_FUNC_HI_MPI_SYS_Init;
		return -1;
	}
        printf("HI_MPI_SYS_Init ok\n");

        return 0;
}

HI_VOID* Test_ISP_Run(HI_VOID *param){
        int error_code;
        ISP_DEV IspDev = 0;
        printf("starting HI_MPI_ISP_Run...\n");
        error_code = HI_MPI_ISP_Run(IspDev);
        printf("HI_MPI_ISP_Run %d\n", error_code);

        return HI_NULL;
}
static pthread_t gs_IspPid;

int himpp3_mipi_isp_init() {
        int error_code;

        int fd;
        combo_dev_attr_t *pstcomboDevAttr, stcomboDevAttr;

        /* mipi reset unrest */
        fd = open("/dev/hi_mipi", O_RDWR);
        if (fd < 0) {
                //printf("warning: open hi_mipi dev failed\n");
                return -1;
        }
    
 	pstcomboDevAttr = &LVDS_6lane_SENSOR_IMX274_12BIT_8M_NOWDR_ATTR;
	
        memcpy(&stcomboDevAttr, pstcomboDevAttr, sizeof(combo_dev_attr_t));
        stcomboDevAttr.devno = 0;

        /* 1.reset mipi */
        if(ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno)) {
                //printf("HI_MIPI_RESET_MIPI failed\n");
                close(fd);
                return -1;
   	}

        /* 2.reset sensor */
        if(ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno)) {
    	        //printf("HI_MIPI_RESET_SENSOR failed\n");
                close(fd);
                return -1;
        }

        if (ioctl(fd, HI_MIPI_SET_DEV_ATTR, pstcomboDevAttr)) {
                //printf("set mipi attr failed\n");
                close(fd);
                return -1;
        }

        /* 4.unreset mipi */
        if(ioctl(fd, HI_MIPI_UNRESET_MIPI, &stcomboDevAttr.devno)) {
                //printf("HI_MIPI_UNRESET_MIPI failed\n");
                close(fd);
                return -1;
        }

        /* 5.unreset sensor */
        if(ioctl(fd, HI_MIPI_UNRESET_SENSOR, &stcomboDevAttr.devno)) {
                //printf("HI_MIPI_UNRESET_SENSOR failed\n");
                close(fd);
                return -1;
        }

        close(fd);
 
        error_code = HI_MPI_ISP_Exit(0);
        if (error_code != HI_SUCCESS) {
                printf("%s: HI_MPI_ISP_Exit failed with %#x!\n", __FUNCTION__, error_code);
                return -1;
        }        
        
        ISP_DEV IspDev = 0;
    
        ISP_PUB_ATTR_S stPubAttr;
        ALG_LIB_S stLib;

	const ISP_SNS_OBJ_S *g_pstSnsObj[2] =  {&stSnsImx274Obj, HI_NULL};
	
        ALG_LIB_S stAeLib;
        ALG_LIB_S stAwbLib;
        ALG_LIB_S stAfLib;

        stAeLib.s32Id = 0;
        stAwbLib.s32Id = 0;
        stAfLib.s32Id = 0;
        strncpy(stAeLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
        strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
        strncpy(stAfLib.acLibName, HI_AF_LIB_NAME, sizeof(HI_AF_LIB_NAME)); 

	if (g_pstSnsObj[0]->pfnRegisterCallback != HI_NULL) {
                error_code = g_pstSnsObj[0]->pfnRegisterCallback(IspDev, &stAeLib, &stAwbLib);
                if (error_code != HI_SUCCESS) {
                        printf("%s: sensor_register_callback failed with %#x!\n", __FUNCTION__, error_code);
                        return -1;
                }
        } else {
                printf("%s: sensor_register_callback failed with HI_NULL!\n",  __FUNCTION__);
                return -1;
        }

        /* 2. register hisi ae lib */
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AE_LIB_NAME);
        error_code = HI_MPI_AE_Register(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("%s: HI_MPI_AE_Register failed!\n", __FUNCTION__);
                return -1;
        }

        /* 3. register hisi awb lib */
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AWB_LIB_NAME);
        error_code = HI_MPI_AWB_Register(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("%s: HI_MPI_AWB_Register failed!\n", __FUNCTION__);
                return -1;
        }

        /* 4. register hisi af lib */
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AF_LIB_NAME);
        error_code = HI_MPI_AF_Register(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("%s: HI_MPI_AF_Register failed!\n", __FUNCTION__);
                return -1;
        }

        /* 5. isp mem init */
        error_code = HI_MPI_ISP_MemInit(IspDev);
        if (error_code != HI_SUCCESS) {
                printf("%s: HI_MPI_ISP_Init failed!\n", __FUNCTION__);
                return -1;
        }

        /* 6. isp set WDR mode */
        ISP_WDR_MODE_S stWdrMode;
	stWdrMode.enWDRMode  = WDR_MODE_NONE;
	
        error_code = HI_MPI_ISP_SetWDRMode(0, &stWdrMode);    
        if (error_code != HI_SUCCESS) {
                printf("start ISP WDR failed!\n");
                return -1;
        }

	stPubAttr.enBayer		= BAYER_RGGB;
        stPubAttr.f32FrameRate          = 30;
        stPubAttr.stWndRect.s32X        = 0;
        stPubAttr.stWndRect.s32Y        = 0;
        stPubAttr.stWndRect.u32Width    = 3840;
        stPubAttr.stWndRect.u32Height   = 2160;
        stPubAttr.stSnsSize.u32Width    = 3840;
        stPubAttr.stSnsSize.u32Height   = 2160;    

        error_code = HI_MPI_ISP_SetPubAttr(IspDev, &stPubAttr);
        if (error_code != HI_SUCCESS) {
                printf("%s: HI_MPI_ISP_SetPubAttr failed with %#x!\n", __FUNCTION__, error_code);
                return -1;
        }

        /* 8. isp init */
        error_code = HI_MPI_ISP_Init(IspDev);
        if (error_code != HI_SUCCESS) {
                printf("%s: HI_MPI_ISP_Init failed!\n", __FUNCTION__);
                return -1;
        }

        if (0 != pthread_create(&gs_IspPid, 0, (void* (*)(void*))Test_ISP_Run, NULL)) {
                printf("%s: create isp running thread failed!\n", __FUNCTION__);
                return -1;
        }

	return 0;

}

int himpp3_vi_init() {
        int error_code;

        VI_DEV_ATTR_S  stViDevAttr;
    
        memset(&stViDevAttr,0,sizeof(stViDevAttr));
	memcpy(&stViDevAttr, &DEV_ATTR_LVDS_BASE, sizeof(stViDevAttr));

        stViDevAttr.stDevRect.s32Y                              = 0;
        stViDevAttr.stDevRect.u32Width                          = 3840;
        stViDevAttr.stDevRect.u32Height                         = 2160;
        stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Width    = 3840;
        stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Height   = 2160;
        stViDevAttr.stBasAttr.stSacleAttr.bCompress             = HI_FALSE;

        error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
        if (error_code != HI_SUCCESS) {
                printf("HI_MPI_VI_SetDevAttr failed with %#x!\n", error_code);
                return -1;
        }
 
        error_code = HI_MPI_VI_EnableDev(0);
        if (error_code != HI_SUCCESS) {
                printf("HI_MPI_VI_EnableDev failed with %#x!\n", error_code);
                return -1;
        }

        RECT_S stCapRect;
        SIZE_S stTargetSize;

        stCapRect.s32X          = 0;
        stCapRect.s32Y          = 0;
        stCapRect.u32Width      = 3840;
        stCapRect.u32Height     = 2160;
        stTargetSize.u32Width   = stCapRect.u32Width;
        stTargetSize.u32Height  = stCapRect.u32Height;

        VI_CHN_ATTR_S stChnAttr;

	memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));
        
        stChnAttr.enCapSel              = VI_CAPSEL_BOTH;
        stChnAttr.stDestSize.u32Width   = stTargetSize.u32Width ;
        stChnAttr.stDestSize.u32Height  = stTargetSize.u32Height ;
        stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   /* sp420 or sp422 */

        stChnAttr.bMirror       = HI_FALSE;
        stChnAttr.bFlip         = HI_FALSE;

        stChnAttr.s32SrcFrameRate       = 30;
        stChnAttr.s32DstFrameRate       = 30;
        stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;

        error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
        if (error_code != HI_SUCCESS) {
                printf("HI_MPI_VI_SetChnAttr failed with %#x!\n", error_code);
                return -1;
        }

        //#define CMOS_LDC
        //#ifdef CMOS_LDC

        VI_LDC_ATTR_S stLDCAttr;
        //First enable VI devices and VI channel.
        //Initialize LDC attributes.
        stLDCAttr.bEnable = HI_TRUE;
        stLDCAttr.stAttr.enViewType = LDC_VIEW_TYPE_ALL;
        //LDC_VIEW_TYPE_CROP;
        stLDCAttr.stAttr.s32CenterXOffset = 0;
        stLDCAttr.stAttr.s32CenterYOffset = 0;
        stLDCAttr.stAttr.s32Ratio = 58;
        stLDCAttr.stAttr.s32MinRatio = 0;
        //Set LDC attributes.
        error_code = HI_MPI_VI_SetLDCAttr(0, &stLDCAttr);
        if (error_code != HI_SUCCESS) {
                printf("Set vi LDC attr err:0x%x\n", error_code);
                return -1;
        }
        printf("HI_MPI_VI_SetLDCAttr ok\n");

        //Obtain LDC attributes.
        error_code = HI_MPI_VI_GetLDCAttr (0, &stLDCAttr);
        if (error_code != HI_SUCCESS) {
                printf("Get vi LDC attr err:0x%x\n", error_code);
                return -1;
        }
        printf("HI_MPI_VI_GetLDCAttr ok\n");
        //#endif

        error_code = HI_MPI_VI_EnableChn(0);
        if (error_code != HI_SUCCESS) {
                printf(" HI_MPI_VI_EnableChn failed with %#x!\n", error_code);
                return -1;
        }

	return 0;

}

int himpp3_vpss_init() {
        int error_code;

        VPSS_GRP VpssGrp = 0;
        VPSS_GRP_ATTR_S stVpssGrpAttr;

        VpssGrp = 0;

	stVpssGrpAttr.u32MaxW           = 3840;
	stVpssGrpAttr.u32MaxH           = 2160;
	stVpssGrpAttr.bIeEn             = HI_FALSE;
	stVpssGrpAttr.bNrEn             = HI_TRUE;
	stVpssGrpAttr.bHistEn           = HI_FALSE;
	stVpssGrpAttr.bDciEn            = HI_FALSE;
	stVpssGrpAttr.enDieMode         = VPSS_DIE_MODE_NODIE;
	stVpssGrpAttr.enPixFmt          = PIXEL_FORMAT_YUV_SEMIPLANAR_420;//SAMPLE_PIXEL_FORMAT;
        stVpssGrpAttr.bStitchBlendEn    = HI_FALSE;

        stVpssGrpAttr.stNrAttr.enNrType                         = VPSS_NR_TYPE_VIDEO;
	stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrRefSource      = VPSS_NR_REF_FROM_RFR;//VPSS_NR_REF_FROM_CHN0, VPSS_NR_REF_FROM_SRC
        stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrOutputMode     = VPSS_NR_OUTPUT_NORMAL;//VPSS_NR_OUTPUT_DELAY NORMAL
	stVpssGrpAttr.stNrAttr.u32RefFrameNum                   = 2;

        error_code = HI_MPI_VPSS_CreateGrp(VpssGrp, &stVpssGrpAttr);
        if (error_code != HI_SUCCESS) {
                printf("HI_MPI_VPSS_CreateGrp failed with %#x!\n", error_code);
                return -1;
        }

        error_code = HI_MPI_VPSS_StartGrp(VpssGrp);
        if (error_code != HI_SUCCESS) {
                printf("HI_MPI_VPSS_StartGrp failed with %#x\n", error_code);
                return -1;
        }

        MPP_CHN_S stSrcChn;
        MPP_CHN_S stDestChn;

	stSrcChn.enModId  = HI_ID_VIU;
        stSrcChn.s32DevId = 0;
        stSrcChn.s32ChnId = 0;
    
        stDestChn.enModId  = HI_ID_VPSS;
        stDestChn.s32DevId = 0;
        stDestChn.s32ChnId = 0;
    
        error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
        if (error_code != HI_SUCCESS) {
                printf("HI_MPI_SYS_Bind failed with %#x!\n", error_code);
                return -1;
        }

        /////

        VPSS_CHN VpssChn = 0;
        VPSS_CHN_ATTR_S stVpssChnAttr;
        VPSS_CHN_MODE_S stVpssChnMode;

	VpssChn = 0;
        stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
        stVpssChnMode.bDouble        = HI_FALSE;
        stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
        stVpssChnMode.u32Width       = 3840;// 
        stVpssChnMode.u32Height      = 2160;
        stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE;//COMPRESS_MODE_SEG;
    
        memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));
        stVpssChnAttr.s32SrcFrameRate = 30;
        stVpssChnAttr.s32DstFrameRate = 30;

	error_code = HI_MPI_VPSS_SetChnAttr(VpssGrp, VpssChn, &stVpssChnAttr);
	if (error_code != HI_SUCCESS) {
    	        printf("HI_MPI_VPSS_SetChnAttr failed with %#x\n", error_code);
                return -1;
        }

	error_code = HI_MPI_VPSS_SetChnMode(VpssGrp, VpssChn, &stVpssChnMode);
        if (error_code != HI_SUCCESS) {
    	        printf("%s failed with %#x\n", __FUNCTION__, error_code);
                return -1;
        }         

	error_code = HI_MPI_VPSS_EnableChn(VpssGrp, VpssChn);
        if (error_code != HI_SUCCESS) {
                printf("HI_MPI_VPSS_EnableChn failed with %#x\n", error_code);
                return -1;
        }

	return 0;

}

int himpp3_venc_init() {
        int error_code;

        VENC_CHN_ATTR_S stVencChnAttr;
        //VENC_ATTR_JPEG_S stJpegAttr;

        VENC_ATTR_MJPEG_S stMjpegAttr;
        VENC_ATTR_MJPEG_FIXQP_S stMjpegeFixQp;


        stVencChnAttr.stVeAttr.enType = PT_MJPEG;

        stMjpegAttr.u32MaxPicWidth = 3840;
        stMjpegAttr.u32MaxPicHeight = 2160;
        stMjpegAttr.u32PicWidth = 3840;
        stMjpegAttr.u32PicHeight = 2160;
        stMjpegAttr.u32BufSize = 3840 * 2160 * 3;
        stMjpegAttr.bByFrame = HI_TRUE;  /*get stream mode is field mode  or frame mode*/
        memcpy(&stVencChnAttr.stVeAttr.stAttrMjpege, &stMjpegAttr, sizeof(VENC_ATTR_MJPEG_S));

        stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;
        stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32StatTime       = 1;
        stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32SrcFrmRate      = 30;
        stVencChnAttr.stRcAttr.stAttrMjpegeCbr.fr32DstFrmRate = 1;
        stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32FluctuateLevel = 1;

        stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32BitRate = 1024 * 5;  

        /*
        stJpegAttr.u32PicWidth  = 3840;
        stJpegAttr.u32PicHeight = 2160;
        stJpegAttr.u32MaxPicWidth  = 3840;
        stJpegAttr.u32MaxPicHeight = 2160;
        stJpegAttr.u32BufSize   = 3840 * 2160 * 3;
        stJpegAttr.bByFrame     = HI_TRUE;
        stJpegAttr.bSupportDCF  = HI_FALSE;

        memcpy(&stVencChnAttr.stVeAttr.stAttrJpege, &stJpegAttr, sizeof(VENC_ATTR_JPEG_S));
        */
 	stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
        stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 0;
        



	error_code = HI_MPI_VENC_CreateChn(0, &stVencChnAttr);
        if (error_code != HI_SUCCESS) {
                printf("HI_MPI_VENC_CreateChn [%d] faild with %#x!\n", 0, error_code);
                return -1;
        }

        error_code = HI_MPI_VENC_StartRecvPic(0);
        if (error_code != HI_SUCCESS) {
                printf("HI_MPI_VENC_StartRecvPic faild with%#x!\n", error_code);
                return -1;
        }

        MPP_CHN_S stSrcChn;
        MPP_CHN_S stDestChn;

        stSrcChn.enModId = HI_ID_VPSS;
        stSrcChn.s32DevId = 0;
        stSrcChn.s32ChnId = 0;

        stDestChn.enModId = HI_ID_VENC;
        stDestChn.s32DevId = 0;
        stDestChn.s32ChnId = 0;

        error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
        if (error_code != HI_SUCCESS) {
                printf("HI_MPI_SYS_Bind failed with %#x!\n", error_code);
                return -1;
        }

        int fd =  HI_MPI_VENC_GetFd(0);

        //////////////////////////////////////////
        /*
        int fd =  HI_MPI_VENC_GetFd(0);
        printf("HI_MPI_VENC_GetFd(0) %d\n", fd);

        VENC_STREAM_S stStream;
        VENC_CHN_STAT_S stStat;

        while(1) {
                memset(&stStream, 0, sizeof(stStream));

                error_code = HI_MPI_VENC_Query(0, &stStat);
                if (error_code != HI_SUCCESS) {
    	                printf("HI_MPI_VENC_Query chn[%d] failed with %#x!\n", 0, error_code);                    
                        return 1;
                }
                
                //stStream.pstPack = (VENC_PACK_S*)malloc(sizeof(VENC_PACK_S) * stStat.u32CurPacks);
                //stStream.u32PackCount = stStat.u32CurPacks;

                error_code = HI_MPI_VENC_GetStream(0, &stStream, -1);
                if (error_code != HI_SUCCESS) {
                        printf("HI_MPI_VENC_GetStream failed with %#x!\n", error_code);
                }
                printf("got frame\n");

                error_code = HI_MPI_VENC_ReleaseStream(0, &stStream);
                if (error_code != HI_SUCCESS) {
                        printf("failed to release stream: %#x\n", error_code);
                }
                printf("frame released\n");
        }
        */
        //////////////////////////////////////////

	return fd;


}

