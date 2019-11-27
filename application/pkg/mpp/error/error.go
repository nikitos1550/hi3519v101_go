//+build !debug

package error

import (
	"strconv"
)

func Resolve(code int) string {
	switch code {
	default:
		out := "error " + strconv.FormatInt(int64(code), 10)
		return out
	}
}