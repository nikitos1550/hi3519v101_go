//+build debug
//+build hi3516cv100

package error

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
