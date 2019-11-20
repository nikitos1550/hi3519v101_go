// +build hi3516cv300

package temperature

import "application/hiutils"

func initTemperature() {
    hiutils.WriteDevMem32(0x1203009C, 0xCFA00000)
}

func getTemperature() float32 {
    var tempCode uint32 = hiutils.ReadDevMem32(0x120300A4)
    var temp float32 = ((( float32(tempCode & 0x3FF)-125)/806)*165)-40
    return temp
}

