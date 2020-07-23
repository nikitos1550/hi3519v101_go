#pragma once

#ifdef HI3516CV100 //Family includes
    #define LVDS_LANE_NUM   0
    #define VI_MODE_LVDS    0
    #define VI_MODE_MIPI    0

    #define INPUT_MODE_CMOS     0
    #define INPUT_MODE_CMOS_33V 0
    #define INPUT_MODE_HISPI    0
    #define INPUT_MODE_LVDS     0
    #define INPUT_MODE_MIPI     0
    #define INPUT_MODE_SUBLVDS  0

    #define VI_MODE_HISPI 0
    #define WDR_MODE_2To1_FRAME_FULL_RATE 0
    #define WDR_MODE_2To1_FRAME 0

    #define PT_H265                 10000
    #define VENC_RC_MODE_H264AVBR   10000
    #define VENC_RC_MODE_H264CVBR   10001
    #define VENC_RC_MODE_H264QVBR   10002
    #define VENC_RC_MODE_H265AVBR   10003
    #define VENC_RC_MODE_H265CBR    10004
    #define VENC_RC_MODE_H265CVBR   10005
    #define VENC_RC_MODE_H265FIXQP  10006
    #define VENC_RC_MODE_H265QVBR   10007
    #define VENC_RC_MODE_H265VBR    10008

    #define VENC_GOPMODE_ADVSMARTP      10000
    #define VENC_GOPMODE_BIPREDB        10001
    #define VENC_GOPMODE_DUALP          10002
    #define VENC_GOPMODE_NORMALP        10003
    #define VENC_GOPMODE_SMARTP         10004
    #define VENC_GOPMODE_INTRAREFRESH   10005

    #define HI_MPP_V1
    #define HI_MPP 1

    #include "../include/hi3516cv100_mpp.h"
#endif

#ifdef HI3516AV100 //Familiy includes
    #define INPUT_MODE_CMOS 0

    #define VENC_RC_MODE_H264CVBR   1000
    #define VENC_RC_MODE_H264QVBR   2000
    #define VENC_RC_MODE_H265CVBR   3000
    #define VENC_RC_MODE_H265QVBR   4000

    #define VENC_GOPMODE_ADVSMARTP      10000
    #define VENC_GOPMODE_BIPREDB        10001
    #define VENC_GOPMODE_DUALP          10002
    #define VENC_GOPMODE_NORMALP        10003
    #define VENC_GOPMODE_SMARTP         10004
    #define VENC_GOPMODE_INTRAREFRESH   10005

    #define HI_MPP_V2
    #define HI_MPP 2
    #include "../include/hi3516av100_mpp.h"
#endif

#ifdef HI3516CV200 //Family includes
    #define INPUT_MODE_CMOS 0

    #define VENC_RC_MODE_H264CVBR   1000
    #define VENC_RC_MODE_H264QVBR   1001
    #define VENC_RC_MODE_H265CVBR   1000
    #define VENC_RC_MODE_H265QVBR   1001

    #define VENC_GOPMODE_ADVSMARTP      10000
    #define VENC_GOPMODE_BIPREDB        10001
    #define VENC_GOPMODE_DUALP          10002
    #define VENC_GOPMODE_NORMALP        10003
    #define VENC_GOPMODE_SMARTP         10004
    #define VENC_GOPMODE_INTRAREFRESH   10005 

    #define HI_MPP_V2
    #define HI_MPP 2
    #include "../include/hi3516cv200_mpp.h"
#endif

#ifdef HI3516CV300 //Family includes
    #define INPUT_MODE_CMOS_33V 0

    #define VENC_RC_MODE_H264CVBR   1000
    #define VENC_RC_MODE_H265CVBR   1000

    #define VENC_GOPMODE_ADVSMARTP      10000
    #define VENC_GOPMODE_INTRAREFRESH   10001

    #define HI_MPP_V3
    #define HI_MPP 3
    #include "../include/hi3516cv300_mpp.h"
#endif

#ifdef HI3516AV200 //Family includes
    #define INPUT_MODE_CMOS_33V 0

    #define VENC_RC_MODE_H264CVBR   1000
    #define VENC_RC_MODE_H265CVBR   1000

    #define VENC_GOPMODE_ADVSMARTP      10000
    #define VENC_GOPMODE_INTRAREFRESH   10001

    #define HI_MPP_V3
    #define HI_MPP 3
    #include "../include/hi3516av200_mpp.h"
#endif

#ifdef HI3516CV500 //Family includes
    #define INPUT_MODE_CMOS_33V 0

    #define VENC_GOPMODE_INTRAREFRESH   10000

    #define HI_MPP_V4
    #define HI_MPP 4
    #include "../include/hi3516cv500_mpp.h"
#endif

#ifdef HI3516EV200 //Family includes
    #define INPUT_MODE_CMOS_33V 0

    #define VENC_GOPMODE_INTRAREFRESH   10000

    #define HI_MPP_V4
    #define HI_MPP 4
    #include "../include/hi3516ev200_mpp.h"
#endif

#ifdef HI3519AV100 //Family includes
    #define HI_MPP_V4
    #define HI_MPP 4
    #include "../include/hi3519av100.h"
#endif

#ifdef HI3559AV100 //Family includes
    #define HI_MPP_V4
    #define HI_MPP 4
    #include "../include/hi3559av100.h"
#endif
