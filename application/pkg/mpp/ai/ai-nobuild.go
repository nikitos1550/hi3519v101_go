//+build !hi3516cv200,!hi3516av100,!hi3516av200,!hi3516cv300

package ai

//#include "ai.h"
import "C"

func Init() {}

func IsAudioExistTmp() bool { //temporary function to check if there any audio avalible in the system
    return false
}

func RemoveOpus(ch chan []byte) {
//    delete(Clients, ch)
}

func SubsribeOpus(ch chan []byte) {
//    Clients[ch] = true
}

