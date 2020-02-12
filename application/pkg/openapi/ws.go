//+build openapi
//+build debug

package openapi

import (
	"log"
	"net/http"
    //"github.com/gorilla/websocket"
)

func init() {
    AddWsRoute("wsEcho",      "/debug/echo",     "GET",      wsEcho)
}

func wsEcho(w http.ResponseWriter, r *http.Request) {
    log.Println("wsEcho")

    c, err := Upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()

    log.Println("WS connection established")

    for {
        mt, message, err := c.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        log.Printf("recv: %s", message)
        err = c.WriteMessage(mt, message)
        if err != nil {
            log.Println("write:", err)
            break
        }
    }
}


