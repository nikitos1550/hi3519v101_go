//+build arm
//+build hi3516av100 hi3516av200 hi3516cv100 hi3516cv200 hi3516cv300 hi3516cv500 hi3516ev200 hi3519av100 hi3559av100

package venc

/*
#include "../include/mpp.h"
#include "../../logger/logger.h"

#include <sys/epoll.h>
#include <pthread.h>
#include <string.h>
#include <stdlib.h>

#include "venc.h"
//#include "loop.h"

#define ERR_NONE                    0
#define ERR_SYS                     2

pthread_t tid;
int epoll_fd;
pthread_mutex_t lock;

#define NUM_VENCS   VENC_MAX_CHN_NUM

struct st_loop_vencs {
    int id;
    int fd;
} loop_vencs[NUM_VENCS];

int loop_vencs2[NUM_VENCS]; //TODO change struct above to this array, use array index as venc channel id

//mpp_data_loop_add can run at any time, go space will garantee that several copies will not run simultaniously

int mpp_data_loop_add(unsigned int *error_code, unsigned int venc_channel_id) {
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
        return ERR_SYS; //TODO create return code for this case
    }

    loop_vencs[item].id = venc_channel_id;
    loop_vencs[item].fd = venc_fd;

    struct epoll_event event;
    event.events = EPOLLIN; // | EPOLLET; //EPOLLIN | EPOLLPRI | EPOLLET;
    event.data.fd = venc_fd;
    if (epoll_ctl(epoll_fd, EPOLL_CTL_ADD, venc_fd, &event) < 0) {
        return ERR_SYS;
    }

    return ERR_NONE;
}

//mpp_data_loop_del can`t run during loop operation on same venc channel
//go space will garantee that several copies will not run simultaniously<Paste>

int mpp_data_loop_del(unsigned int *error_code, unsigned int venc_channel_id) {
    *error_code = 0;

    int venc_fd = mpp_venc_getfd(venc_channel_id);

    struct epoll_event event;
    event.events = EPOLLIN; // | EPOLLET; //EPOLLIN | EPOLLPRI | EPOLLET;
    event.data.fd = venc_fd;

    int i, item;
    item = -1;
    for (i=0; i<NUM_VENCS; i++) {
        if (loop_vencs[i].id == venc_fd) {
            item = i;
            break;
        }
    }

    if (item == -1) {
        return ERR_SYS; //TODO create new return code for such case
    }

    pthread_mutex_lock(&lock);

    loop_vencs[item].id = -1;
    loop_vencs[item].fd = -1;

    if (epoll_ctl(epoll_fd, EPOLL_CTL_DEL, venc_fd, &event) < 0) {
        pthread_mutex_unlock(&lock);
        return ERR_SYS;
    }
    pthread_mutex_unlock(&lock);

    return ERR_NONE;
}

int mpp_data_loop_init(unsigned int *error_code) {

    *error_code = 0;

    int i;
    for (i=0; i<NUM_VENCS; i++) {
        loop_vencs[i].id = -1;
        loop_vencs[i].fd = -1;
    }

    if (pthread_mutex_init(&lock, NULL) != 0) {
        //printf("\n mutex init failed\n");
        return ERR_SYS;
    }

    epoll_fd = epoll_create(3); //WTF? 3?
    if (epoll_fd == -1) {
        return ERR_SYS;
    }

    *error_code = pthread_create(&tid, NULL, &mpp_data_loop_thread, NULL);
    if (*error_code != 0) return ERR_SYS;

    return ERR_NONE;
}

#define MAX_EPOLL_EVENTS    NUM_VENCS

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
                    mpp_data_loop_get_data(loop_vencs[j].id);	//mpp_data_loop_del made delation bettwen 
                    break;										//epoll_wait unblock and this lock of this area
                }                                               //https://stackoverflow.com/questions/3652056/how-efficient-is-locking-an-unlocked-mutex-what-is-the-cost-of-a-mutex
            }
        }
        pthread_mutex_unlock(&lock);    //Here we unlock
    }
    GO_LOG_VENC(LOGGER_ERROR, "LOOP failed");
}
*/
import "C"

import (
	"application/pkg/logger"


    "fmt"
    "net/http"
    "application/pkg/openapi"

    "sync"
)

var mutex = &sync.Mutex{}

func init() {
    openapi.AddApiRoute("serveDebugLoop", "/mpp/venc/loop", "GET", serveDebugLoop)
}

func serveDebugLoop(w http.ResponseWriter, r *http.Request) {
    logger.Log.Trace().
	    Msg("mpp.venc.serveDebugLoop")

    w.Header().Set("Content-Type", "test/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    fmt.Fprintf(w, "%s", "TODO")
}

//Rules:
//addVenc/delVenc functions should operate one in a time.
//There should be some sync for them run:
//1) exlusive mutex
//2) query (based obviosly on go channels)

func addVenc(venc int) {
    mutex.Lock()
	defer mutex.Unlock()

    var errorCode C.uint
    var vencChannelId C.uint
    vencChannelId = C.uint(venc)

    switch err := C.mpp_data_loop_add(&errorCode, vencChannelId); err {
    case C.ERR_NONE:
	    logger.Log.Debug().
		    Msg("C.mpp_data_loop_add() ok")
    case C.ERR_SYS:
	    logger.Log.Fatal().
		    Msg("C.mpp_data_loop_add() SYS error")
    default:
	    logger.Log.Fatal().
		    Int("error", int(err)).
		    Msg("C.mpp_data_loop_add() Unexpected return")
    }

    logger.Log.Debug().
	    Int("channel", venc).
	    Msg("VENC channel added to loop")
}

func delVenc(venc int) {
    mutex.Lock()
	defer mutex.Unlock()

    var errorCode C.uint
    var vencChannelId C.uint
    vencChannelId = C.uint(venc)

    switch err := C.mpp_data_loop_del(&errorCode, vencChannelId); err {
    case C.ERR_NONE:
	    logger.Log.Debug().
		    Msg("C.mpp_data_loop_del() ok")
    case C.ERR_SYS:
	    logger.Log.Fatal().
		    Msg("C.mpp_data_loop_del() SYS error")
    default:
	    logger.Log.Fatal().
		    Int("error", int(err)).
		    Msg("C.mpp_data_loop_del() Unexpected return")
    }

    logger.Log.Debug().
        Int("channel", venc).
        Msg("VENC channel deleted from loop")
}


func loopInit() {
    var errorCode C.uint

    switch err := C.mpp_data_loop_init(&errorCode); err {
    case C.ERR_NONE:
	    logger.Log.Debug().
		    Msg("C.mpp_data_loop_init() ok")
    case C.ERR_SYS:
	    logger.Log.Fatal().
		    Msg("C.mpp_data_loop_init() SYS error")
    default:
	    logger.Log.Fatal().
		    Int("error", int(err)).
		    Msg("C.mpp_data_loop_init() Unexpected return")
    }

}
