// +build hi3516cv300

package temperature

import "application/pkg/utils"

func init() {
    utils.WriteDevMem32(0x1203009C, 0xCFA00000)
}

func getTemperature() (float32, error) {
    var tempCode uint32 = hiutils.ReadDevMem32(0x120300A4)
    var temp float32 = ((( float32(tempCode & 0x3FF)-125)/806)*165)-40

    return temp, nil
}

