//+build processing

package processing

/*
#include "processing.h"
int bindInit(unsigned int channelId, unsigned int encoderId) {
    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId    = HI_ID_VPSS;
    stSrcChn.s32DevId   = 0;
    stSrcChn.s32ChnId   = channelId;
    stDestChn.enModId   = HI_ID_VENC;
    stDestChn.s32DevId  = 0;
    stDestChn.s32ChnId  = encoderId;

    return HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
}

int bindDestroy(unsigned int channelId, unsigned int encoderId) {
    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId    = HI_ID_VPSS;
    stSrcChn.s32DevId   = 0;
    stSrcChn.s32ChnId   = channelId;
    stDestChn.enModId   = HI_ID_VENC;
    stDestChn.s32DevId  = 0;
    stDestChn.s32ChnId  = encoderId;

    return HI_MPI_SYS_UnBind(&stSrcChn, &stDestChn);
}
*/
import "C"

import (
	"log"    
	"strconv"    
	"unsafe"
	"application/pkg/common"
)

type bind struct {
	Name string
	Id int
	EncoderId int
	ChannelId int
}

func (p bind) GetName() string {
	return p.Name
}

func (p bind) GetId() int {
	return p.Id
}

func (p bind) Create(id int, params map[string][]string) (common.Processing,int,string) {
	encoderIdString, ok := params["encoderId"]
	if (!ok || len(encoderIdString) <= 0) {
		return nil, 0, "encoderId not specified"
	}	

	encoderId, err := strconv.Atoi(encoderIdString[0])
	if err != nil {
		return nil, 0, "encoderId must be number"
	}

	channelIdString, ok := params["channelId"]
	if (!ok || len(channelIdString) <= 0) {
		return nil, 0, "channelId not specified"
	}	

	channelId, err := strconv.Atoi(channelIdString[0])
	if err != nil {
		return nil, 0, "channelId must be number"
	}

	bindErr := C.bindInit(C.uint(channelId), C.uint(encoderId))
	if (bindErr != 0){
		return nil, int(bindErr), "failed to bind"
	}

	var v bind
	v.Name = "bind"
	v.Id = id
	v.EncoderId = encoderId
	v.ChannelId = channelId
	return v,id,""
}

func (p bind) Destroy() {
	bindErr := C.bindDestroy(C.uint(p.ChannelId), C.uint(p.EncoderId))
	if (bindErr != 0){
		log.Println("failed to destroy bind processing", int(bindErr))
	}

	log.Println("bind processing was destroyed")
}

func (p bind) Callback(data unsafe.Pointer) {
}

func init() {
	var v bind
	v.Name = "bind"
	v.Id = -1
	v.EncoderId = -1
	v.ChannelId = -1
	register(v)
}
