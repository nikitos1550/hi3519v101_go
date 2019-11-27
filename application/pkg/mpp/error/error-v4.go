//+build debug
//+build hi3516cv500 hi3516ev200 hi3519av100 hi3559av100

package error

import (
	"strconv"
)

func Resolve(code int) string {
	switch code {

	default:
		out := "unknown error " + strconv.FormatInt(int64(code), 10)
		return out
	}
}


