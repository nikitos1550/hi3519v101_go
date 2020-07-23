#include "venc.h"

#include <string.h>
#include <sys/epoll.h>
#include <pthread.h>
#include <string.h>
#include <stdlib.h>

pthread_t tid;
int epoll_fd;
pthread_mutex_t lock;

#define NUM_VENCS   VENC_MAX_CHN_NUM

struct st_loop_vencs {
    int id;
    int fd;
    int codec;
} loop_vencs[NUM_VENCS];

//int loop_vencs[NUM_VENCS]; //TODO change struct above to this array, use array index as venc channel id

#define MAX_EPOLL_EVENTS    NUM_VENCS

int mpp_venc_getfd(int venc_channel_id) {
    return HI_MPI_VENC_GetFd(venc_channel_id);
}

int mpp_venc_closefd(int venc_channel_id) {
    #if     HI_MPP == 1
        //Not implemented
        return ERR_NONE;
    #elif   HI_MPP == 2 \
            || HI_MPP == 3 \
            || HI_MPP == 4
        return HI_MPI_VENC_CloseFd(venc_channel_id);
    #endif
}

unsigned long long lastPts = 0;

void mpp_data_loop_get_data(unsigned int id) {
    HI_S32          s32Ret;

    unsigned int venc_channel_id = loop_vencs[id].id;

    #if HI_MPP == 1 || \
        HI_MPP == 2 || \
        HI_MPP == 3
        VENC_CHN_STAT_S stStat;
    #endif
    #if HI_MPP == 4
        VENC_CHN_STATUS_S stStat;
    #endif

    VENC_STREAM_S   stStream;

    //#if HI_MPP == 4 //TODO do we really need this call?
    //	VENC_STREAM_BUF_INFO_S stStreamBufInfo;
    //
    //	s32Ret = HI_MPI_VENC_GetStreamBufInfo (venc_channel_id, &stStreamBufInfo);
    //	if (HI_SUCCESS != s32Ret) {
    //    	go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_GetStreamBufInfo failed");
    //    	return;
    //	}
	//#endif

    memset(&stStream, 0, sizeof(stStream));

    #if HI_MPP == 1 || \
        HI_MPP == 2 || \
        HI_MPP == 3 
    	s32Ret = HI_MPI_VENC_Query(venc_channel_id, &stStat);
    	if (HI_SUCCESS != s32Ret) {
        	go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_Query failed"); //TODO pass err code
        	return; //continue;
    	}
    #endif

    #if HI_MPP == 4
    	s32Ret = HI_MPI_VENC_QueryStatus(venc_channel_id, &stStat);
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
    s32Ret = HI_MPI_VENC_GetStream(venc_channel_id, &stStream, HI_TRUE);
    if (HI_SUCCESS != s32Ret) {
        free(stStream.pstPack);
        stStream.pstPack = NULL;
        go_logger_venc(LOGGER_PANIC, "HI_MPI_VENC_GetStream failed"); //TODO pass err code
        return; //continue;
    }

    info_from_c info;
    data_from_c *data;

    #if HI_MPP == 1

        info.seq = stStream.u32Seq;

        data = (data_from_c *)malloc(sizeof(data_from_c) * stStream.u32PackCount * 2);

        int i = 0;
        int j = 0;
        for(i = 0; i < stStream.u32PackCount; i++) {
            //TODO VENC_PACK_S->U64PTS
            info.pts = stStream.pstPack[i].u64PTS;
            //printf("VENC %llu\n", stStream.pstPack[i].u64PTS);
            
            data[j].data = stStream.pstPack[i].pu8Addr[0];
            data[j].length = stStream.pstPack[i].u32Len[0];
            j++;
            if (stStream.pstPack[i].u32Len[1] > 0) {
                //TODO VENC_PACK_S->U64PTS
                data[j].data = stStream.pstPack[i].pu8Addr[1];
                data[j].length = stStream.pstPack[i].u32Len[1];    
                j++;   
            } 
        }
        //go_callback_receive_data(venc_channel_id, stStream.u32Seq, st_data, j);
        go_callback_receive_data(venc_channel_id, &info, data, j);
    #endif
    #if HI_MPP == 2 || \
        HI_MPP == 3 || \
		HI_MPP == 4

        info.seq = stStream.u32Seq;

        switch (loop_vencs[id].codec) {
            case CODEC_MJPEG:
                info.q_factor = stStream.stJpegInfo.u32Qfactor;
                break;
            case CODEC_H264:
                info.ref_type = stStream.stH264Info.enRefType;
                break;
            case CODEC_H265:
                info.ref_type = stStream.stH265Info.enRefType;
                break;
            default:
                ;;;//TODO error
        }

        data = (data_from_c *)malloc(sizeof(data_from_c) * stStream.u32PackCount);

        //////////////unsigned int sizeTmp = 0;

        int i;
        for(i = 0; i < stStream.u32PackCount; i++) {
            data[i].data = stStream.pstPack[i].pu8Addr;
            data[i].length = stStream.pstPack[i].u32Len;

            /////sizeTmp += stStream.pstPack[i].u32Len;
            /////for(int j=0;j<=8;j++) {
            /////    printf("0x%x ", stStream.pstPack[i].pu8Addr[j]);
            /////}
            /////printf("\n");

            //TODO VENC_PACK_S->U64PTS
            info.pts = stStream.pstPack[i].u64PTS;
            //printf("VENC %llu\n", stStream.pstPack[i].u64PTS);
        }

        //if (id == 1) {
        //    printf("%d venc seq %d (%d) count %d, delta: %lld\n", id, stStream.u32Seq, stStream.stH264Info.enRefType, i, (stStream.pstPack[0].u64PTS-lastPts));
        //    lastPts = stStream.pstPack[0].u64PTS;
        //}

        //go_callback_receive_data(venc_channel_id, stStream.u32Seq, st_data, stStream.u32PackCount);
        go_callback_receive_data(venc_channel_id, &info, data, stStream.u32PackCount);
    #endif

    free(data);

    s32Ret = HI_MPI_VENC_ReleaseStream(venc_channel_id, &stStream);
    if (HI_SUCCESS != s32Ret) {
        go_logger_venc(LOGGER_ERROR, "HI_MPI_VENC_ReleaseStream failed"); //TODO pass err code
        return; //continue;
    }
    free(stStream.pstPack);
}



//mpp_data_loop_add can run at any time, go space will garantee that several copies will not run simultaniously
int mpp_data_loop_add(unsigned int *error_code, unsigned int venc_channel_id, unsigned int codec) {
    *error_code = 0;

    int venc_fd = mpp_venc_getfd(venc_channel_id);

    int i, item;
    item = -1;
    for (i=0; i<NUM_VENCS; i++) {
        if (loop_vencs[i].id == -1) {
            item = i;
            break;
        }
    }

    if (item == -1) {
        return ERR_GENERAL; //TODO create return code for this case
    }

    loop_vencs[item].id     = venc_channel_id;
    loop_vencs[item].fd     = venc_fd;
    loop_vencs[item].codec  = codec;

    struct epoll_event event;
    event.events = EPOLLIN; // | EPOLLET; //EPOLLIN | EPOLLPRI | EPOLLET;
    event.data.fd = venc_fd;
    if (epoll_ctl(epoll_fd, EPOLL_CTL_ADD, venc_fd, &event) < 0) {
        return ERR_GENERAL;
    }

    return ERR_NONE;
}

//mpp_data_loop_del can`t run during loop operation on same venc channel
//go space will garantee that several copies will not run simultaniously
int mpp_data_loop_del(unsigned int *error_code, unsigned int venc_channel_id) {
    *error_code = 0;

    int venc_fd = mpp_venc_getfd(venc_channel_id);

    struct epoll_event event;
    event.events = EPOLLIN; // | EPOLLET; //EPOLLIN | EPOLLPRI | EPOLLET;
    event.data.fd = venc_fd;

    int i, item;
    item = -1;
    for (i=0; i<NUM_VENCS; i++) {
        if (loop_vencs[i].id == venc_channel_id) {
            item = i;
            break;
        }
    }

    if (item == -1) {
        return ERR_GENERAL; //TODO create new return code for such case
    }

    pthread_mutex_lock(&lock);

    loop_vencs[item].id     = -1;
    loop_vencs[item].fd     = -1;
    loop_vencs[item].codec  = -1;

    if (epoll_ctl(epoll_fd, EPOLL_CTL_DEL, venc_fd, &event) < 0) {
        pthread_mutex_unlock(&lock);
        return ERR_GENERAL;
    }
    pthread_mutex_unlock(&lock);

    mpp_venc_closefd(venc_channel_id);

    return ERR_NONE;
}

int mpp_data_loop_init(unsigned int *error_code) {

    *error_code = 0;

    int i;
    for (i=0; i<NUM_VENCS; i++) {
        loop_vencs[i].id    = -1;
        loop_vencs[i].fd    = -1;
        loop_vencs[i].codec = -1;
    }

    if (pthread_mutex_init(&lock, NULL) != 0) {
        return ERR_GENERAL;
    }

    epoll_fd = epoll_create(3); //WTF? 3?
    if (epoll_fd == -1) {
        return ERR_GENERAL;
    }

    *error_code = pthread_create(&tid, NULL, &mpp_data_loop_thread, NULL);
    if (*error_code != 0) return ERR_GENERAL;

    return ERR_NONE;
}

//mpp_data_loop_thread can`t run venc channel action if 
void * mpp_data_loop_thread() {
    struct epoll_event events[MAX_EPOLL_EVENTS];

    while(1) {
        //GO_LOG_VENC(LOGGER_TRACE, "LOOP epoll starts waiting");
        int event_count = epoll_wait(epoll_fd, events, MAX_EPOLL_EVENTS, -1);

        if (event_count == 0) {
            GO_LOG_VENC(LOGGER_TRACE, "LOOP invoked with event_counter==0");
            continue;
        }

        int i;
        pthread_mutex_lock(&lock);      //Here should be lock, durint this lock mpp_data_loop_del can`t run
        for(i = 0; i < event_count; i++) {
            //printf("Triggered fd %d\n", events[i].data.fd);
            int j;
            for(j=0; j<NUM_VENCS; j++) {                        //this for will not also find proper venc id and fd relation
                if (loop_vencs[j].fd == events[i].data.fd) {    //but also will protect us from situation when
                    mpp_data_loop_get_data(j);
                    //mpp_data_loop_get_data(loop_vencs[j].id);	//mpp_data_loop_del made delation bettwen 
                    break;										//epoll_wait unblock and this lock of this area
                }                                               //https://stackoverflow.com/questions/3652056/how-efficient-is-locking-an-unlocked-mutex-what-is-the-cost-of-a-mutex
            }
        }
        pthread_mutex_unlock(&lock);    //Here we unlock
    }
    GO_LOG_VENC(LOGGER_ERROR, "LOOP failed");
}

