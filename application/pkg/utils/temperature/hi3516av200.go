// +build hi3516av200

package temperature

import "application/pkg/utils"

func Init() {
    utils.WriteDevMem32(0x120A0110, 0x60FA0000)
}

func Get() (float32, error) {
    var tempCode uint32 = utils.ReadDevMem32(0x120A0118)
    var temp float32 = ((( float32(tempCode & 0x3FF)-125)/806)*165)-40

    return temp, nil
}
