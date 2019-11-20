// +build hi3516cv100

package temperature

func Init() {
    //hi3516cv100 family doesn`t have known internal temperature sensor
}

func Get() float32 {
    var temp float32 = -999

    return temp
}

