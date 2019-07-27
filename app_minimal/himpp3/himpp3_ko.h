#ifndef HIMPP3_KO_H_
#define HIMPP3_KO_H_

#include <stdio.h>
#include <stdlib.h>

struct komodules {
	char *name;
	char *start;
	char *end;
	char *params;
};

extern char _binary_hi_osal_ko_start[], _binary_hi_osal_ko_end[];
extern char _binary_hi3519v101_base_ko_start[], _binary_hi3519v101_base_ko_end[];
extern char _binary_hi3519v101_sys_ko_start[], _binary_hi3519v101_sys_ko_end[];
extern char _binary_hi3519v101_tde_ko_start[], _binary_hi3519v101_tde_ko_end[];
extern char _binary_hi3519v101_region_ko_start[], _binary_hi3519v101_region_ko_end[];
extern char _binary_hi3519v101_fisheye_ko_start[], _binary_hi3519v101_fisheye_ko_end[];
extern char _binary_hi3519v101_vgs_ko_start[], _binary_hi3519v101_vgs_ko_end[];
extern char _binary_hi3519v101_isp_ko_start[], _binary_hi3519v101_isp_ko_end[];
extern char _binary_hi3519v101_viu_ko_start[], _binary_hi3519v101_viu_ko_end[];
extern char _binary_hi3519v101_vpss_ko_start[], _binary_hi3519v101_vpss_ko_end[];
extern char _binary_hi3519v101_vou_ko_start[], _binary_hi3519v101_vou_ko_end[];
extern char _binary_hifb_ko_start[], _binary_hifb_ko_end[];
extern char _binary_hi3519v101_rc_ko_start[], _binary_hi3519v101_rc_ko_end[];
extern char _binary_hi3519v101_venc_ko_start[], _binary_hi3519v101_venc_ko_end[];
extern char _binary_hi3519v101_chnl_ko_start[], _binary_hi3519v101_chnl_ko_end[];
extern char _binary_hi3519v101_vedu_ko_start[], _binary_hi3519v101_vedu_ko_end[];
extern char _binary_hi3519v101_h264e_ko_start[], _binary_hi3519v101_h264e_ko_end[];
extern char _binary_hi3519v101_h265e_ko_start[], _binary_hi3519v101_h265e_ko_end[];
extern char _binary_hi3519v101_jpege_ko_start[], _binary_hi3519v101_jpege_ko_end[];
extern char _binary_hi3519v101_ive_ko_start[], _binary_hi3519v101_ive_ko_end[];
extern char _binary_hi3519v101_photo_ko_start[], _binary_hi3519v101_photo_ko_end[];
extern char _binary_hi_sensor_i2c_ko_start[], _binary_hi_sensor_i2c_ko_end[];
extern char _binary_hi_pwm_ko_start[], _binary_hi_pwm_ko_end[];
extern char _binary_hi_piris_ko_start[], _binary_hi_piris_ko_end[];
extern char _binary_hi_sil9136_ko_start[], _binary_hi_sil9136_ko_end[];
extern char _binary_gyro_bosch_ko_start[], _binary_gyro_bosch_ko_end[];
extern char _binary_hi3519v101_aio_ko_start[], _binary_hi3519v101_aio_ko_end[];
extern char _binary_hi3519v101_ai_ko_start[], _binary_hi3519v101_ai_ko_end[];
extern char _binary_hi3519v101_ao_ko_start[], _binary_hi3519v101_ao_ko_end[];
extern char _binary_hi3519v101_aenc_ko_start[], _binary_hi3519v101_aenc_ko_end[];
extern char _binary_hi3519v101_adec_ko_start[], _binary_hi3519v101_adec_ko_end[];
extern char _binary_hi_acodec_ko_start[], _binary_hi_acodec_ko_end[];
extern char _binary_hi_tlv320aic31_ko_start[], _binary_hi_tlv320aic31_ko_end[];
extern char _binary_hi_mipi_ko_start[], _binary_hi_mipi_ko_end[];
extern char _binary_hi_user_ko_start[], _binary_hi_user_ko_end[];
extern char _binary_hi_ssp_sony_ko_start[], _binary_hi_ssp_sony_ko_end[];


struct komodules modules[] = {
{"hi_osal.ko",		    
    _binary_hi_osal_ko_start,	
    _binary_hi_osal_ko_end, 	
    "mmz=anonymous,0,0x90000000,256M anony=1"},
{"hi3519v101_base.ko",
    _binary_hi3519v101_base_ko_start,
    _binary_hi3519v101_base_ko_end,
    ""},
{"hi3519v101_sys.ko", 	
    _binary_hi3519v101_sys_ko_start, 				
    _binary_hi3519v101_sys_ko_end, 				
    "vi_vpss_online=0 sensor=imx274,NULL mem_total=512"},
/*
{"hi3519v101_tde.ko", 	
    _binary_hi3519v101_tde_ko_start,
    _binary_hi3519v101_tde_ko_end,
    ""},
*/
{"hi3519v101_region.ko",
    _binary_hi3519v101_region_ko_start, 
    _binary_hi3519v101_region_ko_end, 
    ""},

{"hi3519v101_fisheye.ko",	
    _binary_hi3519v101_fisheye_ko_start, 
    _binary_hi3519v101_fisheye_ko_end, 
    ""},

{"hi3519v101_vgs.ko", 
    _binary_hi3519v101_vgs_ko_start, 
    _binary_hi3519v101_vgs_ko_end, 
    ""},
{"hi3519v101_isp.ko", 
    _binary_hi3519v101_isp_ko_start, 
    _binary_hi3519v101_isp_ko_end, 
    "proc_param=30"},
{"hi3519v101_viu.ko", 
    _binary_hi3519v101_viu_ko_start, 
    _binary_hi3519v101_viu_ko_end, 
    "detect_err_frame=10"},
{"hi3519v101_vpss.ko", 
    _binary_hi3519v101_vpss_ko_start, 
    _binary_hi3519v101_vpss_ko_end, 
    ""},
{"hi3519v101_vou.ko", 
    _binary_hi3519v101_vou_ko_start, 
    _binary_hi3519v101_vou_ko_end, 
    ""},

/*
{"hifb.ko", 
    _binary_hifb_ko_start, 
    _binary_hifb_ko_end, 
    "video='hifb:vram0_size:1620'"},
*/
{"hi3519v101_rc.ko", 
    _binary_hi3519v101_rc_ko_start, 
    _binary_hi3519v101_rc_ko_end, 
    ""},
{"hi3519v101_venc.ko", 
    _binary_hi3519v101_venc_ko_start, 
    _binary_hi3519v101_venc_ko_end, 
    ""},
{"hi3519v101_chnl.ko", 
    _binary_hi3519v101_chnl_ko_start, 
    _binary_hi3519v101_chnl_ko_end, 
    ""},
{"hi3519v101_vedu.ko", 
    _binary_hi3519v101_vedu_ko_start, 
    _binary_hi3519v101_vedu_ko_end, 
    ""},
{"hi3519v101_h264e.ko", 
    _binary_hi3519v101_h264e_ko_start, 
    _binary_hi3519v101_h264e_ko_end, 
    ""},
{"hi3519v101_h265e.ko",
    _binary_hi3519v101_h265e_ko_start, 
    _binary_hi3519v101_h265e_ko_end, 
    ""},
{"hi3519v101_jpege.ko", 
    _binary_hi3519v101_jpege_ko_start, 
    _binary_hi3519v101_jpege_ko_end, 
    ""},
/*
{"hi3519v101_ive.ko", 
    _binary_hi3519v101_ive_ko_start, 
    _binary_hi3519v101_ive_ko_end, 
    "save_power=1"},
*/
/*
{"hi3519v101_photo.ko", 
    _binary_hi3519v101_photo_ko_start, 
    _binary_hi3519v101_photo_ko_end, 
    ""},
*/
/*
{"hi_sensor_i2c.ko", 
    _binary_hi_sensor_i2c_ko_start, 
    _binary_hi_sensor_i2c_ko_end, 
    ""},
*/

{"hi_pwm.ko", 
    _binary_hi_pwm_ko_start, 
    _binary_hi_pwm_ko_end, 
    ""},

/*
{"hi_piris.ko", 
    _binary_hi_piris_ko_start, 
    _binary_hi_piris_ko_end, 
    ""},
*/
/*
{"hi_sil9136.ko", 
    _binary_hi_sil9136_ko_start, 
    _binary_hi_sil9136_ko_end, 
    "norm=12"},
*/
/*
{"gyro_bosch.ko", 
    _binary_gyro_bosch_ko_start, 
    _binary_gyro_bosch_ko_end, 
    ""},
*/
/*
{"hi3519v101_aio.ko", 
    _binary_hi3519v101_aio_ko_start, 
    _binary_hi3519v101_aio_ko_end, 
    ""},
{"hi3519v101_ai.ko", 
    _binary_hi3519v101_ai_ko_start, 
    _binary_hi3519v101_ai_ko_end, 
    ""},
{"hi3519v101_ao.ko", 
    _binary_hi3519v101_ao_ko_start, 
    _binary_hi3519v101_ao_ko_end, 
    ""},
{"hi3519v101_aenc.ko", 
    _binary_hi3519v101_aenc_ko_start, 
    _binary_hi3519v101_aenc_ko_end, 
    ""},
{"hi3519v101_adec.ko", 
    _binary_hi3519v101_adec_ko_start, 
    _binary_hi3519v101_adec_ko_end, 
    ""},
{"hi_acodec.ko", 
    _binary_hi_acodec_ko_start, 
    _binary_hi_acodec_ko_end, 
    ""},
*/
/*
{"hi_tlv320aic31.ko", 
    _binary_hi_tlv320aic31_ko_start, 
    _binary_hi_tlv320aic31_ko_end, 
    ""},
*/
{"hi_mipi.ko", 
    _binary_hi_mipi_ko_start, 
    _binary_hi_mipi_ko_end, 
    ""},
{"hi_user.ko", 
    _binary_hi_user_ko_start, 
    _binary_hi_user_ko_end, 
    ""},
{NULL, NULL, NULL, NULL}
};

#endif
