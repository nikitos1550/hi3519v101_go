//+build streamerWs

package ws

import (
    "log"
    "application/pkg/openapi"
)

func init() {
    openapi.AddWsRoute("wsVideo",      "/video",     "GET",      wsVideo)
}

func Init() {}

func wsVideo(w http.ResponseWriter, r *http.Request) {
    log.Println("wsEcho")

    c, err := openapi.Upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()

    //check video id, check corresponding hub

    //create new client
    //register new client to corresponding hub
    //start client goroutines
}

