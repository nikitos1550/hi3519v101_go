//+build hi3516av200 hi3516cv300

package mpp

/*
    //Should be here to export go callback
*/
import "C"

import (
    "log"
)

//export go_callback_receive_data
func go_callback_receive_data() {
    log.Println("go_callback_receive_data()")
}
