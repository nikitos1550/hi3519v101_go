//+build streamerWs

package ws

import (
    _"log"
)


type Hub struct {
    vencHub     bool
    clients     map[*Client]bool
    register    chan *Client
    unregister  chan *Client
}

func newHub() *Hub {
    return &Hub{
        clients:    make(map[*Client]bool),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (h *Hub) runHub() {
    //register/subscribe to some allowed video source
    for {
        select {
            case client := <-h.register:
                h.clients[client] = true
            case client := <-h.unregsiter:
                if _, ok := h.clients[client]; ok {
                    delete(h.clients, client)
                }
        }
    }
}

