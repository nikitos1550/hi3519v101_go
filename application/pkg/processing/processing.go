//+build processing

package processing

/*

#include "../mpp/include/mpp_v3.h"
#include "processing.h"

void sendToClients(unsigned int processingId, VIDEO_FRAME_INFO_S* frame) {
	sendToEncoders(processingId, frame);
}

int sendToEncoder(unsigned int vencId, void* frame) {
 	return HI_MPI_VENC_SendFrame(vencId, frame, -1);
}

*/
import "C"

import (
	"log"
	"unsafe"
)

type ActiveProcessing struct {
	Name string
	InputChannel int
	InputProcessing int
	Callback unsafe.Pointer
	Encoders map[int] bool
	Processings map[int] unsafe.Pointer
}

type Processing struct {
	Name string
	Callback unsafe.Pointer
}

var (
	Processings map[string] Processing
	ActiveProcessings map[int] ActiveProcessing
	lastProcessingId int
)

func Init() {
}

func init() {
	Processings = make(map[string] Processing)
	ActiveProcessings = make(map[int] ActiveProcessing)
	lastProcessingId = 0
}

func register(name string, callback unsafe.Pointer){
	_, exists := Processings[name]
	if (exists) {
		log.Fatal("processing already exists", name)
	}

	processing := Processing{
		Name: name,
		Callback: callback,
	}
	
	Processings[name] = processing
}

func CreateProcessing(processingName string)  (int, string)  {
	processing, exists := Processings[processingName]
	if (!exists) {
		return -1, "Processing not found"
	}

	activeProcessing := ActiveProcessing{
		Name: processing.Name,
		InputChannel: -1,
		InputProcessing: -1,
		Callback: processing.Callback,
		Encoders: make(map[int] bool),
		Processings: make(map[int] unsafe.Pointer),
	}

	lastProcessingId++
	ActiveProcessings[lastProcessingId] = activeProcessing

	return lastProcessingId, ""
}

func SubscribeProcessing(processingId int, encoderId int)  (int, string)  {
	processing, exists := ActiveProcessings[processingId]
	if (!exists) {
		return -1, "Processing not created"
	}

	_, exists = processing.Encoders[encoderId]
	if (exists) {
		return -1, "Already subscribed"
	}
	
	processing.Encoders[encoderId] = true
	ActiveProcessings[processingId] = processing

	return 0, ""
}
