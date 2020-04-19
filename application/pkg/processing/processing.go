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

	"application/pkg/mpp/vpss"
)

func Init() {}

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

func DeleteProcessing(processingId int)  (int, string)  {
	activeProcessing, processingExists := ActiveProcessings[processingId]
	if (!processingExists) {
		return -1, "Processing not created"
	}

	if (activeProcessing.InputChannel != -1){
		err, errorString := vpss.UnsubscribeChannel(activeProcessing.InputChannel, processingId)
		if err < 0 {
			return err, errorString
		}
	}

	if (activeProcessing.InputProcessing != -1){
		err, errorString := UnsubscribeProcessingToProcessing(activeProcessing.InputProcessing, processingId)
		if err < 0 {
			return err, errorString
		}
	}

	delete(ActiveProcessings, processingId)

	return 0, ""
}

func SubscribeEncoderToProcessing(processingId int, encoderId int)  (int, string)  {
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

func UnsubscribeEncoderToProcessing(processingId int, encoderId int)  (int, string)  {
	processing, exists := ActiveProcessings[processingId]
	if (!exists) {
		return -1, "Processing not created"
	}

	_, exists = processing.Encoders[encoderId]
	if (!exists) {
		return -1, "Encoder not subscribed"
	}

	delete(processing.Encoders, encoderId)	

	return 0, ""
}

func SubscribeProcessingToProcessing(processingId int, subscribeProcessingId int)  (int, string)  {
	activeProcessing, processingExists := ActiveProcessings[processingId]
	if (!processingExists) {
		return -1, "Main processing not created"
	}

	subscribeProcessing, subscribeProcessingExists := ActiveProcessings[subscribeProcessingId]
	if (!subscribeProcessingExists) {
		return -1, "Subscribe processing not created"
	}

	_, exists := activeProcessing.Processings[subscribeProcessingId]
	if (exists) {
		return -1, "Already subscribed"
	}

	activeProcessing.Processings[subscribeProcessingId] = subscribeProcessing.Callback
	ActiveProcessings[processingId] = activeProcessing

	subscribeProcessing.InputProcessing = processingId
	ActiveProcessings[subscribeProcessingId] = subscribeProcessing

	return 0, ""
}

func UnsubscribeProcessingToProcessing(processingId int, subscribeProcessingId int)  (int, string)  {
	activeProcessing, processingExists := ActiveProcessings[processingId]
	if (!processingExists) {
		return -1, "Main processing not created"
	}

	_, subscribeProcessingExists := ActiveProcessings[subscribeProcessingId]
	if (!subscribeProcessingExists) {
		return -1, "Subscribe processing not created"
	}

	_, exists := activeProcessing.Processings[subscribeProcessingId]
	if (!exists) {
		return -1, "Processing not subscribed"
	}

	delete(activeProcessing.Processings, subscribeProcessingId)	

	return 0, ""
}

