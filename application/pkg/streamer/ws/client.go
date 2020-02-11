package ws

import (
    "log"

    "github.com/gorilla/websocket"
)

type Client struct {
    hub *Hub
    // The websocket connection.
    conn *websocket.Conn
}
var Clients []Client

func (c *Client) readPump() {
    defer func() {
        c.hub.unregister <- c
        c.conn.Close()
    }()
    //some initial setup for client
    for {
        _, message, err := c.conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("error: %v", err)
            }
            break
        }
        log.Println(message)
    }
}

func (c *Client) writePump() {
}

