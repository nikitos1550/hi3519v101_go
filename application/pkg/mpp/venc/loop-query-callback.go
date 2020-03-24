package venc

/*
#include "loop.h"
//Should be here to export go callback
*/
import "C"

import (
	//"log"
	"unsafe"
)

//export go_callback_receive_data
func go_callback_receive_data(venc_channel C.int, seq C.uint, data_pointer *C.data_from_c, data_num C.int) {
	vencChannel := int(venc_channel)
	num := int(data_num)

	channels, exists := EncoderSubscriptions[vencChannel]
	if (!exists) {
		return
	}

	dataFromC := (*[1 << 10]C.data_from_c)(unsafe.Pointer(data_pointer))[:num:num]
	length := 0
	for i := 0; i < num; i++ {
		length = length + int(dataFromC[i].length)
	}

	data := make([]byte, length)

	offset := 0
	for i := 0; i < num; i++ {
		p := (*[1 << 28]byte)(unsafe.Pointer(dataFromC[i].data))[:dataFromC[i].length:dataFromC[i].length]
		n := copy(data[offset:], p)
		offset = offset + n
	}

	for ch,enabled := range channels {
		if (enabled){
			if (cap(ch) <= len(ch)) {
				//log.Println("Channel is full. Capacity ", cap(ch), " Length ", len(ch), "Skip element")
				<-ch
			}

			ch <- data
		}
	}
}
