#pragma once

#ifdef HI3516CV100 //Family includes
    #define HI_MPP_V1
    #include "../include/hi3516cv100_mpp.h"
#endif

#ifdef HI3516AV100 //Familiy includes
    #define HI_MPP_V2
    #include "../include/hi3516av100_mpp.h"
#endif

#ifdef HI3516CV200 //Family includes
    #define HI_MPP_V2
    #include "../include/hi3516cv200_mpp.h"
#endif

#ifdef HI3516CV300 //Family includes
    #define HI_MPP_V3
    #include "../include/hi3516cv300_mpp.h"
#endif

#ifdef HI3516AV200 //Family includes
    #define HI_MPP_V3
    #include "../include/hi3516av200_mpp.h"
#endif

#ifdef HI3516CV500 //Family includes
    #define HI_MPP_V4
    #include "../include/hi3516cv500_mpp.h"
#endif

#ifdef HI3516EV200 //Family includes
    #define HI_MPP_V4
    #include "../include/hi3516ev200_mpp.h"
#endif

#ifdef HI3519AV100 //Family includes
    #define HI_MPP_V4
    #include "../include/hi3519av100.h"
#endif

#ifdef HI3559AV100 //Family includes
    #define HI_MPP_V4
    #include "../include/hi3559av100.h"
#endif
