// +build openapi

package openapi

import (
	"log"
	"net/http"
    "github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{} // use default options

func init() {
    AddWsRoute("wsEcho",      "/echo",     "GET",      wsEcho)
}

func wsEcho(w http.ResponseWriter, r *http.Request) {
    log.Println("wsEcho")

    c, err := Upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()

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


