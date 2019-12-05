//+build hi3516cv300 hi3516av200

package getter

import (
    "log"
)

/*
#include "../include/"

//TODO poll data get lopp
//int getter_data_loop() {
//  watch for new data
//  if new data, call go callback
//  //
//}

// int getter_venc_register(int venc_channel_id) {}
// init getter_venc_unregister(int venc_channel_id) {}

*/
import "C"

func registerVencEncodedChannel() { }
func unRegisterVencEncodedChannel() { }

//func RegisgerVpssRawChannel() { }
//func UnRegisgerVpssRawChannel() { }

//func RegisgerViRawChannel() { }
//func UnRegisgerViRawChannel() { }

func registerConsumer() { }
func unRegisterConsumer() { }
