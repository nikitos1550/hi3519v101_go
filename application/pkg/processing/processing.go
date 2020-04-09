//+build processing

package processing

/*

#include "../mpp/include/mpp_v3.h"

void sendToClients(unsigned int processingId, VIDEO_FRAME_INFO_S* frame) {
	sendToEncoders(processingId, frame);
}

void sendToEncoder(unsigned int vencId, VIDEO_FRAME_INFO_S* frame) {
}

*/
import "C"

import (
	"log"
	"unsafe"
)

type ActiveProcessing struct {
	Name string
	Callback unsafe.Pointer
	Encoders map[int] bool
	Processings map[unsafe.Pointer] bool
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
		Callback: processing.Callback,
		Encoders: make(map[int] bool),
		Processings: make(map[unsafe.Pointer] bool),
	}

	lastProcessingId++
	ActiveProcessings[lastProcessingId] = activeProcessing

	return lastProcessingId, ""
}