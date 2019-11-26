// +build hi3516cv100

package temperature

import "errors"

func Init() {
    //hi3516cv100 family doesn`t have known internal temperature sensor
}

func Get() (float32, error) {
    return 0, errors.New("not supported by hardware")
}

