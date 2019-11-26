package main

import (
    "log"
    _"application/pkg/koloader"
    _"application/pkg/mpp"
    "application/pkg/openapi"
    _"application/pkg/utils/chip"
    "application/pkg/utils/temperature"
)

func main() {

    log.Println("daemon")

    //koloader.LoadMinimal()
    //mpp.Init()
    openapi.Init()
    temperature.Init()

    log.Println("loaded")

    select {}
}
