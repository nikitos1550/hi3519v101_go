//+build arm
//+build hi3516ev200

package mpp

import (
    //"log"
    //"application/pkg/logger"
    //"os"

	"application/pkg/ko"
    //"application/pkg/utils"
    //"application/pkg/mpp/error"
)

//TODO rework this mess
func systemInit() {
	//This family originally pack all reg init to sy_conf ko module
	ko.UnloadAll()
	ko.LoadAll()
}
