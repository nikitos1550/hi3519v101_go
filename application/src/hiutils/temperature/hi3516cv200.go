// +build hi3516cv200

package temperature

import "application/hiutils"

func initTemperature() {
    hiutils.WriteDevMem32(0x20270110, 0x60FA0000)
}

func getTemperature() float32 {
    var tempCode uint32 = hiutils.ReadDevMem32(0x20270114)
    var temp float32 = (( float32(tempCode & 0xFF)*180)/256)-40
    return temp
}

