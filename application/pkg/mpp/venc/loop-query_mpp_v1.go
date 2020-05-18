//+build arm
//+build hi3516cv100

package venc

/*
#include "../include/mpp.h"
#include "../../logger/logger.h"
#include "loop.h"

#include <string.h>
#include <stdlib.h>

int mpp_venc_getfd(int venc_channel_id) {
    return HI_MPI_VENC_GetFd(venc_channel_id);
}

void mpp_data_loop_get_data(unsigned int venc_channel) {
    HI_S32          s32Ret;
    VENC_CHN_STAT_S stStat;
    VENC_STREAM_S   stStream;

    memset(&stStream, 0, sizeof(stStream));
    s32Ret = HI_MPI_VENC_Query(venc_channel, &stStat);
    if (HI_SUCCESS != s32Ret) {
	    go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_Query failed"); //TODO pass err code
        return; 
    }

    if (0 == stStat.u32CurPacks) {
        go_logger_venc(LOGGER_PANIC, "stStat.u32CurPacks"); 
        return;
    }

    stStream.pstPack = (VENC_PACK_S*)malloc(sizeof(VENC_PACK_S) * stStat.u32CurPacks);
    if (NULL == stStream.pstPack) {
        go_logger_venc(LOGGER_PANIC, "malloc stream pack failed!");
        return;
    }

    stStream.u32PackCount = stStat.u32CurPacks;
    s32Ret = HI_MPI_VENC_GetStream(venc_channel, &stStream, HI_TRUE);
    if (HI_SUCCESS != s32Ret) {
        free(stStream.pstPack);
        stStream.pstPack = NULL;
	    go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_GetStream failed"); //TODO pass err code
        return; 
    }

    data_from_c * st_data;
    st_data = (data_from_c *)malloc(sizeof(data_from_c) * stStream.u32PackCount * 2);

    int i = 0;
    int j = 0;
    for(i = 0; i < stStream.u32PackCount; i++) {
        st_data[j].data = stStream.pstPack[i].pu8Addr[0];
        st_data[j].length = stStream.pstPack[i].u32Len[0];
        j++;
        if (stStream.pstPack[i].u32Len[1] > 0) {
            st_data[j].data = stStream.pstPack[i].pu8Addr[1];
            st_data[j].length = stStream.pstPack[i].u32Len[1];    
            j++;
        }
    }
    go_callback_receive_data(venc_channel, stStream.u32Seq, st_data, j);
    //go_callback_receive_data(venc_channel, stStream.u32Seq, st_data, stStream.u32PackCount);

    free(st_data);

    s32Ret = HI_MPI_VENC_ReleaseStream(venc_channel, &stStream);
    if (HI_SUCCESS != s32Ret) {
	    go_logger_venc(LOGGER_ERROR, "HI_MPI_VENC_ReleaseStream failed"); //TODO pass err code
        return; //continue;
    }
    free(stStream.pstPack);
}
*/
import "C"

