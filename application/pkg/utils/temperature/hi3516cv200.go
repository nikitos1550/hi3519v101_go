// +build hi3516cv200

package temperature

import "application/pkg/utils"

func init() {
	utils.WriteDevMem32(0x20270110, 0x60FA0000)
}

func Get() (float32, error) {
	var tempCode uint32 = utils.ReadDevMem32(0x20270114)
	var temp float32 = ((float32(tempCode&0xFF) * 180) / 256) - 40

	return temp, nil
}
