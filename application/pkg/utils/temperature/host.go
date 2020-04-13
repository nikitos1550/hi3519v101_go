//+build 386 amd64
//+build host

package temperature

import "errors"

func Get() (float32, error) {
    return 0, errors.New("not supported by host")
}

