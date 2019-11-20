// +build hi3516av200

package temperature

import "application/hiutils"

func initTemperature() {
    hiutils.WriteDevMem32(0x120A0110, 0x60FA0000)
}

func getTemperature() float32 {
    var tempCode uint32 = hiutils.ReadDevMem32(0x120A0118)
    var temp float32 = ((( float32(tempCode & 0x3FF)-125)/806)*165)-40
    return temp
}
