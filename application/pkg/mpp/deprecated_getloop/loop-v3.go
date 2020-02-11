//+build hi3516av200 hi3516cv300

package getloop

/*
#include "../include/hi3516av200_mpp.h"
#include <sys/epoll.h>
#include <pthread.h>
#include <string.h>
#include <stdlib.h>

#include "getloop.h"

#define ERR_NONE                    0
#define ERR_SYS                     2

extern void go_callback_receive_data(int venc, data_from_c * data_pointer, int num, int data_length);

//forward declarations
int mpp3_data_loop_add(unsigned int *error_code, unsigned int venc_channel_id);
void * mpp3_data_loop_thread();
void mpp3_data_loop_get_data(unsigned int venc_channel);

pthread_t tid;
int epoll_fd;

#define NUM_VENCS   10
struct st_loop_vencs {
    int id;
    int fd;
} loop_vencs[NUM_VENCS];

int mpp3_data_loop_add(unsigned int *error_code, unsigned int venc_channel_id) {
    *error_code = 0;

    int venc_fd = HI_MPI_VENC_GetFd(venc_channel_id);

    //TODO make it proper way
    int i;
    for (i=0; i<NUM_VENCS; i++) {
        if (loop_vencs[i].id == -1) {
            loop_vencs[i].id = venc_channel_id;
            loop_vencs[i].fd = venc_fd;
            break;
        }
    }

    struct epoll_event event;
    event.events = EPOLLIN; // | EPOLLET; //EPOLLIN | EPOLLPRI | EPOLLET;
    event.data.fd = venc_fd;
    if (epoll_ctl(epoll_fd, EPOLL_CTL_ADD, venc_fd, &event) < 0) {
        return ERR_SYS;
    }

    return ERR_NONE;
}

int mpp3_data_loop_init(unsigned int *error_code) {

    *error_code = 0;

    int i;
    for (i=0; i<NUM_VENCS; i++) {
        loop_vencs[i].id = -1;
        loop_vencs[i].fd = -1;
    }

    epoll_fd = epoll_create(3);
    if (epoll_fd == -1) {
        return ERR_SYS;
    }

    *error_code = pthread_create(&tid, NULL, &mpp3_data_loop_thread, NULL);
    if (*error_code != 0) return ERR_SYS;

    return ERR_NONE;
}

#define MAX_EPOLL_EVENTS    3

void * mpp3_data_loop_thread() {
    struct epoll_event events[MAX_EPOLL_EVENTS];

    while(1) {
        //printf("epoll starts waiting\n");
        int event_count = epoll_wait(epoll_fd, events, MAX_EPOLL_EVENTS, -1);

        //printf("event_count = %d\n", event_count);
        if (event_count == 0) continue;

        int i, j;
        for(i = 0; i < event_count; i++) {
            //printf("Triggered fd %d\n", events[i].data.fd);
            for(j=0; j<NUM_VENCS; j++) {
                if (loop_vencs[j].fd == events[i].data.fd) {
                    mpp3_data_loop_get_data(loop_vencs[j].id);
                    break;
                }
            }
        }
    }
}

void mpp3_data_loop_get_data(unsigned int venc_channel) {
    //printf("mpp3_data_loop_get_data processign venc channel %d\n", venc_channel);

    //VENC_CHN        VencChn = venc_channel;
    HI_S32          s32Ret;
    VENC_CHN_STAT_S stStat;
    VENC_STREAM_S   stStream;

    memset(&stStream, 0, sizeof(stStream));
    s32Ret = HI_MPI_VENC_Query(venc_channel, &stStat);
    if (HI_SUCCESS != s32Ret) {
        printf("HI_MPI_VENC_Query failed with %#x!\n", s32Ret);
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
        printf("HI_MPI_VENC_GetStream failed with %#x!\n", s32Ret);
        return; //continue;
    }

    data_from_c * st_data;
    st_data = (data_from_c *)malloc(sizeof(data_from_c) * stStream.u32PackCount);

    int i;
    int len = 0;
    for(i = 0; i < stStream.u32PackCount; i++) {
        //printf("length = %d\n", stStream.pstPack[i].u32Len);
        len += stStream.pstPack[i].u32Len;
        st_data[i].data = stStream.pstPack[i].pu8Addr;
        st_data[i].length = stStream.pstPack[i].u32Len;
    }
    //printf("NEW FRAME len = %d!\n", len);
    go_callback_receive_data(venc_channel, st_data, stStream.u32PackCount, len);

    free(st_data);

    s32Ret = HI_MPI_VENC_ReleaseStream(venc_channel, &stStream);
    if (HI_SUCCESS != s32Ret) {
        printf("failed to release stream: %#x\n", s32Ret);
        return; //continue;
    }
    free(stStream.pstPack);
}
*/
import "C"

import (
    "log"
)

func AddVenc(venc int) {
    var errorCode C.uint
    var vencChannelId C.uint
    vencChannelId = C.uint(venc)

    switch err := C.mpp3_data_loop_add(&errorCode, vencChannelId); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_data_loop_add() ok")
    case C.ERR_SYS:
        log.Println("C.mpp3_data_loop_add() SYS error")
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_data_loop_add()")
    }

    log.Println("getloop added venc channel ", venc)
}

func Init() {
    var errorCode C.uint

    switch err := C.mpp3_data_loop_init(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_data_loop_init() ok")
    case C.ERR_SYS:
        log.Println("C.mpp3_data_loop_init() SYS error")
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_data_loop_init()")
    }

}
