//+build 386 amd64
//+build host
//+build debug

package error

import (
        "strconv"
)

func Resolve(code int64) string {
	panic("This can`t be invoked in host build!")
	return nil
}

