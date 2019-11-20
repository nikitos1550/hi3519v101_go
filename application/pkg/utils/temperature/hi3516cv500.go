// +build hi3516cv500

package temperature

import "application/pkg/utils"

func Init() {
    utils.WriteDevMem32(0x120300B4, 0xCBA00000)
}

func Get() float32 {
    var tempCode uint32 = hiutils.ReadDevMem32(0x120300BC)
    var temp float32 = ((( float32(tempCode & 0x3FF)-136)/793)*165)-40
    return temp
}

