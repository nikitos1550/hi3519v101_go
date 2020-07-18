//+build arm
//+build hi3516ev200

package temperature

//import "errors"
import "application/core/utils"

func init() {
    utils.WriteDevMem32(0x120280B4, 0xCBA00000)
}

func Get() (float32, error) {
    //var tempCode uint32 = utils.ReadDevMem32(0x120300BC)
    //var temp float32 = ((( float32(tempCode & 0x3FF)-136)/793)*165)-40
    //return temp, nil

    var tempCode uint32 = utils.ReadDevMem32(0x120280BC)
    var temp float32 = ((( float32(tempCode & 0x3FF)-117)/798)*165)-40
    return temp, nil

    //return 0, errors.New("TODO, not implemented yet")

}

