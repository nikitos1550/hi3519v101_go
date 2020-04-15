//+build arm
//+build hi3516cv500

package temperature

import "application/pkg/utils"

func init() {
    utils.WriteDevMem32(0x120300B4, 0xCBA00000)
}

func Get() (float32, error) {
    var tempCode uint32 = utils.ReadDevMem32(0x120300BC)
    var temp float32 = ((( float32(tempCode & 0x3FF)-136)/793)*165)-40
    return temp, nil
}

