//+build arm
//+build hi3516ev200
//+build debug

package errmpp

import (
	"strconv"
)

func Resolve(code int64) string {
	switch code {

	default:
		out := "unknown error " + strconv.FormatInt(code, 16)
		return out
	}
}


