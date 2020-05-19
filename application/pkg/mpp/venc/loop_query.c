#include "venc.h"

#include <string.h>

int mpp_venc_getfd(int venc_channel_id) {
    return HI_MPI_VENC_GetFd(venc_channel_id);
}

#if 0 //HI_MPP == 1
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
#endif

#if HI_MPP == 1 || \
    HI_MPP == 2 || \
    HI_MPP == 3 || \
    HI_MPP == 4
void mpp_data_loop_get_data(unsigned int venc_channel) {
    HI_S32          s32Ret;

    #if HI_MPP == 1 || \
        HI_MPP == 2 || \
        HI_MPP == 3
        VENC_CHN_STAT_S stStat;
    #endif
    #if HI_MPP == 4
        VENC_CHN_STATUS_S stStat;
    #endif

    VENC_STREAM_S   stStream;

    #if HI_MPP == 4 //TODO do we really need this call?
    	VENC_STREAM_BUF_INFO_S stStreamBufInfo;
    
    	s32Ret = HI_MPI_VENC_GetStreamBufInfo (venc_channel, &stStreamBufInfo);
    	if (HI_SUCCESS != s32Ret) {
        	go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_GetStreamBufInfo failed");
        	return;
    	}
	#endif

    memset(&stStream, 0, sizeof(stStream));

    #if HI_MPP == 1 || \
        HI_MPP == 2 || \
        HI_MPP == 3 
    	s32Ret = HI_MPI_VENC_Query(venc_channel, &stStat);
    	if (HI_SUCCESS != s32Ret) {
        	go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_Query failed"); //TODO pass err code
        	return; //continue;
    	}
    #endif

    #if HI_MPP == 4
    	s32Ret = HI_MPI_VENC_QueryStatus(venc_channel, &stStat);
    	if (HI_SUCCESS != s32Ret) {
        	go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_QueryStatus failed");
        	return;
    	}
    #endif

    if (0 == stStat.u32CurPacks) {
        go_logger_venc(LOGGER_PANIC, "stStat.u32CurPacks"); 
        return; //continue;
    }

    stStream.pstPack = (VENC_PACK_S*)malloc(sizeof(VENC_PACK_S) * stStat.u32CurPacks);
    if (NULL == stStream.pstPack) {
        go_logger_venc(LOGGER_PANIC, "malloc stream pack failed!");
        return; //continue;
    }

    stStream.u32PackCount = stStat.u32CurPacks;
    s32Ret = HI_MPI_VENC_GetStream(venc_channel, &stStream, HI_TRUE);
    if (HI_SUCCESS != s32Ret) {
        free(stStream.pstPack);
        stStream.pstPack = NULL;
        go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_GetStream failed"); //TODO pass err code
        return; //continue;
    }

    data_from_c * st_data;

    #if HI_MPP == 1
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
    #endif
    #if HI_MPP == 2 || \
        HI_MPP == 3 || \
		HI_MPP == 4
        st_data = (data_from_c *)malloc(sizeof(data_from_c) * stStream.u32PackCount);

        int i;
        for(i = 0; i < stStream.u32PackCount; i++) {
            st_data[i].data = stStream.pstPack[i].pu8Addr;
            st_data[i].length = stStream.pstPack[i].u32Len;
        }
        go_callback_receive_data(venc_channel, stStream.u32Seq, st_data, stStream.u32PackCount);
    #endif

    free(st_data);

    s32Ret = HI_MPI_VENC_ReleaseStream(venc_channel, &stStream);
    if (HI_SUCCESS != s32Ret) {
        go_logger_venc(LOGGER_ERROR, "HI_MPI_VENC_ReleaseStream failed"); //TODO pass err code
        return; //continue;
    }
    free(stStream.pstPack);
}
#endif

#if 0 //HI_MPP == 4
void mpp_data_loop_get_data(unsigned int venc_channel) {
    HI_S32          s32Ret;
    VENC_CHN_STATUS_S stStat;
    VENC_STREAM_S   stStream;

    VENC_STREAM_BUF_INFO_S stStreamBufInfo;
    
    s32Ret = HI_MPI_VENC_GetStreamBufInfo (venc_channel, &stStreamBufInfo);
    if (HI_SUCCESS != s32Ret) {
        go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_GetStreamBufInfo failed");
        return;
    }

    memset(&stStream, 0, sizeof(stStream));

    s32Ret = HI_MPI_VENC_QueryStatus(venc_channel, &stStat);
    if (HI_SUCCESS != s32Ret) {
        go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_QueryStatus failed");
        return;
    }

    //printf("venc %d stStat.u32CurPacks == %d\n", venc_channel, stStat.u32CurPacks);
    if (0 == stStat.u32CurPacks) {
        go_logger_venc(LOGGER_PANIC, "stStat.u32CurPacks"); 
        return; //continue;
    }

    stStream.pstPack = (VENC_PACK_S*)malloc(sizeof(VENC_PACK_S) * stStat.u32CurPacks);
    if (NULL == stStream.pstPack) {
        go_logger_venc(LOGGER_PANIC, "malloc stream pack failed!");
        return; //continue;
    }

    stStream.u32PackCount = stStat.u32CurPacks;
    s32Ret = HI_MPI_VENC_GetStream(venc_channel, &stStream, HI_TRUE);
    if (HI_SUCCESS != s32Ret) {
        free(stStream.pstPack);
        stStream.pstPack = NULL;
	    go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_GetStream failed"); //TODO pass err code
        return; //continue;
    }

    data_from_c * st_data;
    st_data = (data_from_c *)malloc(sizeof(data_from_c) * stStream.u32PackCount);

    int i;
    for(i = 0; i < stStream.u32PackCount; i++) {
        st_data[i].data = stStream.pstPack[i].pu8Addr;
        st_data[i].length = stStream.pstPack[i].u32Len;
    }
    go_callback_receive_data(venc_channel, stStream.u32Seq, st_data, stStream.u32PackCount);

    free(st_data);

    s32Ret = HI_MPI_VENC_ReleaseStream(venc_channel, &stStream);
    if (HI_SUCCESS != s32Ret) {
	    go_logger_venc(LOGGER_ERROR, "HI_MPI_VENC_ReleaseStream failed"); //TODO pass err code
        return; //continue;
    }
    free(stStream.pstPack);
}

#endif
