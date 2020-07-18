//+build nobuild

package schedule

/*
#include "processing.h"

//unsigned long long getTimestamp(void* data) {
//	VIDEO_FRAME_INFO_S* stFrame = data;
//	return stFrame->stVFrame.u64pts;
//}

unsigned long long getTimestamp(void* data) {
    VIDEO_FRAME_INFO_S* stFrame = data;

    //printf("getTimestamp %llu \n", stFrame->stVFrame.u64pts);

    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        return stFrame->stVFrame.u64pts;
    #elif HI_MPP == 4
        return stFrame->stVFrame.u64PTS;
    #endif
}

*/
import "C"

import (
	"strconv"
	"unsafe"
	"application/pkg/common"
	"application/pkg/logger"
)

type State int
var (
	PENDING State = 0
	STARTED State = 1
	FINISHED State = 2
)

type schedule struct {
	Name string
	Id int
	StartTimestamp uint64
	StopTimestamp uint64
	CurrentState State
}

func (p schedule) GetName() string {
	return p.Name
}

func (p schedule) GetId() int {
	return p.Id
}

func (p schedule) Create(id int, params map[string][]string) (common.Processing,int,string) {
	timestampString, ok := params["StartTimestamp"]
	if (!ok || len(timestampString) <= 0) {
		return nil, 0, "StartTimestamp not specified"
	}

	startTimestamp, err := strconv.ParseUint(timestampString[0], 10, 64)
	if err != nil {
		return nil, 0, "StartTimestamp must be int"
	}

	timestampString, ok = params["StopTimestamp"]
	if (!ok || len(timestampString) <= 0) {
		return nil, 0, "StopTimestamp not specified"
	}

	stopTimestamp, err := strconv.ParseUint(timestampString[0], 10, 64)
	if err != nil {
		return nil, 0, "StopTimestamp must be int"
	}

	v := &schedule{
		Name: "schedule",
		Id: id,
		StartTimestamp: startTimestamp,
		StopTimestamp: stopTimestamp,
		CurrentState: PENDING,
	}

	return v,id,""
}

func (p schedule) Destroy() {
}

func (p schedule) updateState(frameTime uint64) {
	if (p.CurrentState == PENDING) {
		if frameTime >= p.StartTimestamp{
			logger.Log.Debug().
				Uint64("p.StartTimestamp", p.StartTimestamp).
				Uint64("frameTime", frameTime).
				Int("p.CurrentState", int(p.CurrentState)).
				Int("newState", int(STARTED)).
				Msg("State was changed")

			p.CurrentState = STARTED
		}
	} else if (p.CurrentState == STARTED) {
		if frameTime >= p.StopTimestamp{
			logger.Log.Debug().
				Uint64("p.StopTimestamp", p.StopTimestamp).
				Uint64("frameTime", frameTime).
				Int("p.CurrentState", int(p.CurrentState)).
				Int("newState", int(FINISHED)).
				Msg("State was changed")

			p.CurrentState = FINISHED
			//
		}
	}
}

func (p schedule) Callback(data unsafe.Pointer) {
	frameTime := uint64(C.getTimestamp(data))
	//p.updateState(frameTime)
	//
	//if (p.CurrentState == STARTED) {
	//	sendToEncoders(p.Id, data)
	//}

	if frameTime < p.StartTimestamp {
		return
	}

	if frameTime > p.StopTimestamp {
		return
	}

	sendToEncoders(p.Id, data)
}

func init() {
	var v schedule
	v.Name = "schedule"
	v.Id = -1
	register(v)
}

//http://213.141.129.12:8080/cam1/api/mpp/channel/start?channelId=1&width=1920&height=1080&fps=30
//http://213.141.129.12:8080/cam1/api/processing/create?processingName=schedule&StartTimestamp=1593081840000000&StopTimestamp=1593081900000000
//http://213.141.129.12:8080/cam1/api/encoder/create?encoderName=H264_1920_1080_1M
//http://213.141.129.12:8080/cam1/api/processing/subscribeChannel?processingId=1&channelId=1
//http://213.141.129.12:8080/cam1/api/encoder/subscribeProcessing?processingId=1&encoderId=1
//http://213.141.129.12:8080/cam1/api/files/record/start?encoderId=1
