/*
    Hey man, this file is used for auto link list making, be carefull!
*/
#ifndef HI3516AV200_KO_H_
#define HI3516AV200_KO_H_

#include <stddef.h>

#define KOLIST(F, F_NO) \
    F(hi_osal,              "mmz=anonymous,0,0x90000000,256M anony=1") \
    F(hi3519v101_base,      "") \
    F(hi3519v101_sys,       "vi_vpss_online=0 sensor=imx274,NULL mem_total=512") \
    F_NO(hi3519v101_tde,    "") \
    F(hi3519v101_region,    "") \
    F(hi3519v101_fisheye,   "") \
    F(hi3519v101_vgs,       "") \
    F(hi3519v101_isp,       "proc_param=30") \
    F(hi3519v101_viu,       "detect_err_frame=10") \
    F(hi3519v101_vpss,      "") \
    F(hi3519v101_vou,       "") \
    F_NO(hifb,              "video='hifb:vram0_size:1620'") \
    F(hi3519v101_rc,        "") \
    F(hi3519v101_venc,      "") \
    F(hi3519v101_chnl,      "") \
    F(hi3519v101_vedu,      "") \
    F(hi3519v101_h264e,     "") \
    F(hi3519v101_h265e,     "") \
    F(hi3519v101_jpege,     "") \
    F_NO(hi3519v101_ive,    "save_power=1") \
    F_NO(hi3519v101_photo,  "") \
    F_NO(hi_sensor_i2c,     "") \
    F(hi_pwm,               "") \
    F_NO(hi_piris,          "") \
    F_NO(hi_sil9136,        "norm=12") \
    F_NO(gyro_bosch,        "") \
    F_NO(hi3519v101_aio,    "") \
    F_NO(hi3519v101_ai,     "") \
    F_NO(hi3519v101_ao,     "") \
    F_NO(hi3519v101_aenc,   "") \
    F_NO(hi3519v101_adec,   "") \
    F_NO(hi_acodec,         "") \
    F_NO(hi_tlv320aic31,    "") \
    F(hi_mipi,              "") \
    F(hi_user,              "")

#define EXTERN(NAME, PARAMS) extern char _binary_ ## NAME ## _ko_start[], _binary_ ## NAME ## _ko_end[];
#define STRUCT(NAME, PARAMS) { #NAME , _binary_ ## NAME ## _ko_start, _binary_ ## NAME ## _ko_end, PARAMS},
#define EXTERN_NO(NAME, PARAMS) /* extern char _binary_ ## NAME ## _ko_start[], _binary_ ## NAME ## _ko_end[]; */
#define STRUCT_NO(NAME, PARAMS) /* { #NAME , _binary_ ## NAME ## _ko_start, _binary_ ## NAME ## _ko_end, PARAMS}, */


////////////////////////////////////////////////////////////////////////////////

KOLIST(EXTERN, EXTERN_NO)
EXTERN(hi_ssp_sony, "")

struct ko_modules {
    char *name;
    char *start;
    char *end;
    char *default_params;
};

struct ko_modules modules[] = {
    KOLIST(STRUCT, STRUCT_NO)
    {NULL, NULL, NULL, NULL}
};

#endif //HI3516AV200_KO_H_
