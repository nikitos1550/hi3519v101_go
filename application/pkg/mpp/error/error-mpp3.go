//+build hi3516av200 hi3516cv300

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