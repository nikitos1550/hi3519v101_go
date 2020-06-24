//+build processing

package processing

/*
#include "processing.h"

unsigned long long getTimestamp(void* data) {
	VIDEO_FRAME_INFO_S* stFrame = data;

    //printf("getTimestamp %llu \n", stFrame->stVFrame.u64pts);

	return stFrame->stVFrame.u64pts;
}
*/
import "C"

import (
	"strconv"
	"unsafe"
	"application/pkg/common"
)

type delayed_start struct {
	Name string
	Id int
	Timestamp uint64
}

func (p delayed_start) GetName() string {
	return p.Name
}

func (p delayed_start) GetId() int {
	return p.Id
}

func (p delayed_start) Create(id int, params map[string][]string) (common.Processing,int,string) {
	timestampString, ok := params["timestamp"]
	if (!ok || len(timestampString) <= 0) {
		return nil, 0, "timestamp not specified"
	}	


	timestamp, err := strconv.ParseUint(timestampString[0], 10, 64)
	if err != nil {
		return nil, 0, "timestamp must be int"
	}

	var v delayed_start
	v.Name = "delayed_start"
	v.Id = id
	v.Timestamp = timestamp
	return v,id,""
}

func (p delayed_start) Destroy() {
}

func (p delayed_start) Callback(data unsafe.Pointer) {
	frameTime := uint64(C.getTimestamp(data))
	if frameTime >= p.Timestamp{
		sendToEncoders(p.Id, data)
	}
}

func init() {
	var v delayed_start
	v.Name = "delayed_start"
	v.Id = -1
	register(v)
}

//http://213.141.129.12:8080/cam1/api/mpp/channel/start?channelId=1&width=1920&height=1080&fps=30
//http://213.141.129.12:8080/cam1/api/processing/create?processingName=delayed_start&timestamp=300000000
//http://213.141.129.12:8080/cam1/api/pipeline/create?encoderName=H264_1920_1080_1M
//http://213.141.129.12:8080/cam1/api/processing/subscribeChannel?processingId=1&channelId=1
//http://213.141.129.12:8080/cam1/api/encoder/subscribeProcessing?processingId=1&encoderId=1
