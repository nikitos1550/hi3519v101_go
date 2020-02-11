package venc

/*
#include <sys/epoll.h>
#include <pthread.h>
#include <string.h>
#include <stdlib.h>

#include "loop.h"

#define ERR_NONE                    0
#define ERR_SYS                     2

pthread_t tid;
int epoll_fd;

#define NUM_VENCS   10
struct st_loop_vencs {
    int id;
    int fd;
} loop_vencs[NUM_VENCS];


int mpp_data_loop_add(unsigned int *error_code, unsigned int venc_channel_id) {
    *error_code = 0;

    int venc_fd = mpp_venc_getfd(venc_channel_id);

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

int mpp_data_loop_del(unsigned int *error_code, unsigned int venc_channel_id) {
    *error_code = 0;

    int venc_fd = mpp_venc_getfd(venc_channel_id);

    int i;
    for (i=0; i<NUM_VENCS; i++) {
        if (loop_vencs[i].id == venc_fd) {
            loop_vencs[i].id = -1;
            loop_vencs[i].fd = -1;
            break;
        }
    }

    struct epoll_event event;
    event.events = EPOLLIN; // | EPOLLET; //EPOLLIN | EPOLLPRI | EPOLLET;
    event.data.fd = venc_fd;
    if (epoll_ctl(epoll_fd, EPOLL_CTL_DEL, venc_fd, &event) < 0) {
        return ERR_SYS;
    }

    return ERR_NONE;
}

int mpp_data_loop_init(unsigned int *error_code) {

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

    *error_code = pthread_create(&tid, NULL, &mpp_data_loop_thread, NULL);
    if (*error_code != 0) return ERR_SYS;

    return ERR_NONE;
}

#define MAX_EPOLL_EVENTS    3

void * mpp_data_loop_thread() {
    struct epoll_event events[MAX_EPOLL_EVENTS];

    while(1) {
        //printf("epoll starts waiting\n");
        int event_count = epoll_wait(epoll_fd, events, MAX_EPOLL_EVENTS, -1);

        //printf("event_count = %d\n", event_count);
        if (event_count == 0) continue;

        int i;
        for(i = 0; i < event_count; i++) {
            //printf("Triggered fd %d\n", events[i].data.fd);
            int j;
            for(j=0; j<NUM_VENCS; j++) {
                if (loop_vencs[j].fd == events[i].data.fd) {
                    mpp_data_loop_get_data(loop_vencs[j].id);
                    break;
                }
            }
        }
    }
}
*/
import "C"

import (
    "log"

    "fmt"
    "net/http"
    "application/pkg/openapi"
)

func init() {
    openapi.AddApiRoute("serveDebugLoop", "/mpp/venc/loop", "GET", serveDebugLoop)
}

func serveDebugLoop(w http.ResponseWriter, r *http.Request) {
    log.Println("mpp.venc.serveDebugLoop")

    w.Header().Set("Content-Type", "test/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    fmt.Fprintf(w, "%s", "TODO")
}


func addVenc(venc int) {
    var errorCode C.uint
    var vencChannelId C.uint
    vencChannelId = C.uint(venc)

    switch err := C.mpp_data_loop_add(&errorCode, vencChannelId); err {
    case C.ERR_NONE:
        log.Println("C.mpp_data_loop_add() ok")
    case C.ERR_SYS:
        log.Println("C.mpp_data_loop_add() SYS error")
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp_data_loop_add()")
    }

    log.Println("getloop added venc channel ", venc)
}

func delVenc(venc int) {
    var errorCode C.uint
    var vencChannelId C.uint
    vencChannelId = C.uint(venc)

    switch err := C.mpp_data_loop_del(&errorCode, vencChannelId); err {
    case C.ERR_NONE:
        log.Println("C.mpp_data_loop_del() ok")
    case C.ERR_SYS:
        log.Println("C.mpp_data_loop_del() SYS error")
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp_data_loop_del()")
    }

    log.Println("getloop deleted venc channel ", venc)
}


func loopInit() {
    var errorCode C.uint

    switch err := C.mpp_data_loop_init(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp_data_loop_init() ok")
    case C.ERR_SYS:
        log.Println("C.mpp_data_loop_init() SYS error")
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp_data_loop_init()")
    }

}
