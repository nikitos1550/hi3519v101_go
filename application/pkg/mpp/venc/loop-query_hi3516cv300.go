//+build arm
//+build hi3516cv300

package venc

/*
#include "../include/mpp_v3.h"

#include "../../logger/logger.h"

#include <string.h>
#include <stdlib.h>

#include "loop.h"

int mpp_venc_getfd(int venc_channel_id) {
    return HI_MPI_VENC_GetFd(venc_channel_id);
}

void mpp_data_loop_get_data(unsigned int venc_channel) {
    //printf("mpp3_data_loop_get_data processign venc channel %d\n", venc_channel);

    //VENC_CHN        VencChn = venc_channel;
    HI_S32          s32Ret;
    VENC_CHN_STAT_S stStat;
    VENC_STREAM_S   stStream;

    memset(&stStream, 0, sizeof(stStream));
    s32Ret = HI_MPI_VENC_Query(venc_channel, &stStat);
    if (HI_SUCCESS != s32Ret) {
        //printf("HI_MPI_VENC_Query failed with %#x!\n", s32Ret);
	go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_Query failed"); //TODO pass err code
        return; //continue;
    }

    //printf("venc %d stStat.u32CurPacks == %d\n", venc_channel, stStat.u32CurPacks);
    if (0 == stStat.u32CurPacks) {
        printf("stStat.u32CurPacks == 0\n");
        return; //continue;
    }

    stStream.pstPack = (VENC_PACK_S*)malloc(sizeof(VENC_PACK_S) * stStat.u32CurPacks);
    if (NULL == stStream.pstPack) {
        printf("malloc stream pack failed!\n");
        return; //continue;
    }

    stStream.u32PackCount = stStat.u32CurPacks;
    s32Ret = HI_MPI_VENC_GetStream(venc_channel, &stStream, HI_TRUE);
    if (HI_SUCCESS != s32Ret) {
        free(stStream.pstPack);
        stStream.pstPack = NULL;
        //printf("HI_MPI_VENC_GetStream failed with %#x!\n", s32Ret);
	go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_GetStream failed"); //TODO pass err code
        return; //continue;
    }

    data_from_c * st_data;
    st_data = (data_from_c *)malloc(sizeof(data_from_c) * stStream.u32PackCount);

    int i;
    //int len = 0;
    for(i = 0; i < stStream.u32PackCount; i++) {
        //printf("length = %d\n", stStream.pstPack[i].u32Len);
        //len += stStream.pstPack[i].u32Len;
        st_data[i].data = stStream.pstPack[i].pu8Addr;
        st_data[i].length = stStream.pstPack[i].u32Len;
    }
    //printf("NEW FRAME len = %d!\n", len);
    go_callback_receive_data(venc_channel, stStream.u32Seq, st_data, stStream.u32PackCount);

    free(st_data);

    s32Ret = HI_MPI_VENC_ReleaseStream(venc_channel, &stStream);
    if (HI_SUCCESS != s32Ret) {
        printf("failed to release stream: %#x\n", s32Ret);
	go_logger_venc(LOGGER_ERROR, "HI_MPI_VENC_ReleaseStream failed"); //TODO pass err code
        return; //continue;
    }
    free(stStream.pstPack);
}
*/
import "C"

