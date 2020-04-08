package temperature

import "errors"

func Get() (float32, error) {
    return 0, errors.New("not supported by host")
}

