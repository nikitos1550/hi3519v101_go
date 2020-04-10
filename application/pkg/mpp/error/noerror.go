//+build !debug

package error

import (
	"strconv"
)

func Resolve(code int64) string {
	switch code {
	default:
		out := "error " + strconv.FormatInt(code, 16)
		return out
	}
}
