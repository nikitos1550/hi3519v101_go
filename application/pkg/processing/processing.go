//+build processing

package processing

/*
#include "processing.h"

//TODO move to venc package!!!
int sendToEncoder(error_in *err, unsigned int vencId, void* frame) {
    #if HI_MPP == 1
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_SendFrame, vencId, frame);
    #elif HI_MPP >= 2
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_SendFrame, vencId, frame, -1);
    #endif

    return ERR_NONE;
}

*/
import "C"

import (
	"log"

	"application/pkg/common"
	"application/pkg/mpp/vpss"
)

func Init() {}

type ActiveProcessing struct {
	Proc common.Processing
	InputChannel int
	InputProcessing int
	Encoders map[int] common.Encoder
	Processings map[int] bool
}

var (
	Processings map[string] common.Processing
	ActiveProcessings map[int] ActiveProcessing
	lastProcessingId int
)

func init() {
	ActiveProcessings = make(map[int] ActiveProcessing)
	lastProcessingId = 0
}

func register(processing common.Processing){
	if (Processings == nil){
		Processings = make(map[string] common.Processing)
	}

	_, exists := Processings[processing.GetName()]
	if (exists) {
		log.Fatal("processing already exists", processing.GetName())
	}

	Processings[processing.GetName()] = processing
}

func CreateProcessing(processingName string, params map[string][]string)  (int, string)  {
	processing, exists := Processings[processingName]
	if (!exists) {
		return -1, "Processing not found"
	}

	lastProcessingId++

	p,err,errString := processing.Create(lastProcessingId, params)
	if (err < 0){
		return err,errString
	}

	activeProcessing := ActiveProcessing{
		Proc: p,
		InputChannel: -1,
		InputProcessing: -1,
		Encoders: make(map[int] common.Encoder),
		Processings: make(map[int] bool),
	}

	ActiveProcessings[lastProcessingId] = activeProcessing

	return lastProcessingId, ""
}

func DeleteProcessing(processingId int)  (int, string)  {
	activeProcessing, processingExists := ActiveProcessings[processingId]
	if (!processingExists) {
		return -1, "Processing not created"
	}

	if (activeProcessing.InputChannel != -1){
		err, errorString := vpss.UnsubscribeChannel(activeProcessing.InputChannel, activeProcessing.Proc)
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

	activeProcessing.Proc.Destroy()
	delete(ActiveProcessings, processingId)

	return 0, ""
}

func SubscribeEncoderToProcessing(processingId int, encoder common.Encoder)  (int, string)  {
	processing, exists := ActiveProcessings[processingId]
	if (!exists) {
		return -1, "Processing not created"
	}

	_, exists = processing.Encoders[encoder.GetId()]
	if (exists) {
		return -1, "Already subscribed"
	}
	
	processing.Encoders[encoder.GetId()] = encoder
	ActiveProcessings[processingId] = processing

	return 0, ""
}

func UnsubscribeEncoderToProcessing(processingId int, encoder common.Encoder)  (int, string)  {
	processing, exists := ActiveProcessings[processingId]
	if (!exists) {
		return -1, "Processing not created"
	}

	_, exists = processing.Encoders[encoder.GetId()]
	if (!exists) {
		return -1, "Encoder not subscribed"
	}

	delete(processing.Encoders, encoder.GetId())	

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

	activeProcessing.Processings[subscribeProcessingId] = true
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

