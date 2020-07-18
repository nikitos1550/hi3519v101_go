#include "utils.h"

unsigned int SyncPts (HI_U64 pts) {
    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        return HI_MPI_SYS_SyncPts(pts);
    #elif HI_MPP == 4
        return HI_MPI_SYS_SyncPTS(pts);
    #endif
}

unsigned int InitPtsBase (HI_U64 pts) {
    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        return HI_MPI_SYS_InitPtsBase(pts);
    #elif HI_MPP == 4
        return HI_MPI_SYS_InitPTSBase(pts);
    #endif
}

int mpp_bind_vpss_venc(error_in *err, mpp_bind_vpss_venc_in *in) {
    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId  = HI_ID_VPSS;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = in->vpss_id;

    #if HI_MPP == 1
        stDestChn.enModId  = HI_ID_GROUP;
        stDestChn.s32DevId = in->venc_id;
        stDestChn.s32ChnId = 0;
    #elif HI_MPP == 2 \
        || HI_MPP == 3 \
        || HI_MPP == 4
        stDestChn.enModId  = HI_ID_VENC;
        stDestChn.s32DevId = 0;
        stDestChn.s32ChnId = in->venc_id;
    #endif

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Bind, &stSrcChn, &stDestChn);

    return ERR_NONE;
}

int mpp_unbind_vpss_venc(error_in *err, mpp_bind_vpss_venc_in *in) {
    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId  = HI_ID_VPSS;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = in->vpss_id;

    #if HI_MPP == 1
        stDestChn.enModId  = HI_ID_GROUP;
        stDestChn.s32DevId = in->venc_id;
        stDestChn.s32ChnId = 0;  
    #elif HI_MPP == 2 \
        || HI_MPP == 3 \
        || HI_MPP == 4
        stDestChn.enModId  = HI_ID_VENC;
        stDestChn.s32DevId = 0;
        stDestChn.s32ChnId = in->venc_id;
    #endif
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_UnBind, &stSrcChn, &stDestChn);

    return ERR_NONE;
}

