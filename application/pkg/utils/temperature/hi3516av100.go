// +build hi3516av100

package temperature

func Init() {
    //hi3516av100 family doesn`t have known internal temperature sensor
}

func Get() float32 {
    var temp float32 = -999

    return temp
}


