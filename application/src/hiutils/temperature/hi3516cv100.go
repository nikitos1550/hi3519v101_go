// +build hi3516cv100

package temperature

func initTemperature() {
    //hi3516cv100 family doesn`t have known internal temperature sensor
}

func getTemperature() float32 {
    var temp float32 = -999
    return temp
}

