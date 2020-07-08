package webrtc

import (
    "errors"
    "sync"

    "application/pkg/mpp/connection"
    "application/pkg/mpp/frames"
)

type webrtcServer struct {
    sync.RWMutex

    clients         map[string] *WebrtcSession
    clientsMutex    sync.RWMutex

    source          connection.SourceEncodedData
    Notify          chan frames.FrameItem

    rutineCtrl      chan bool
    rutineStop      chan bool
}

const (
    maxId   = 1024
)

var (
    servers       map[int] *webrtcServer
    serversMutex  sync.RWMutex
)

func init() {
    servers = make(map[int] *webrtcServer)
}

func Init() {}

func Create() (*webrtcServer, error) {
    serversMutex.Lock()
    defer serversMutex.Unlock()

    var server webrtcServer
    server.clients = make(map[string] *WebrtcSession)
    server.rutineCtrl = make(chan bool)
    server.rutineStop = make(chan bool)

    go server.rutine()

    return nil, nil
}

func GetById(id int) (*webrtcServer, error) {

    return nil, nil
}

//func DeleteById(id int) error {
//    serversMutex.Lock()
//    defer serversMutex.Unlock()
//
//    item, exist := servers[id]
//    if !exist {
//        return errors.New("No such instance")
//    }
//
//    if item.source != nil {
//        return errors.New("Can`t delete, because sourced")
//    }
//
//    //TODO force clients disconnect
//
//    delete(servers, id)
//
//    return nil
//}

func Delete(w *webrtcServer) error {
    serversMutex.Lock()
    defer serversMutex.Unlock()

    /*
    for i:=0; i < maxId; i++ {
        item := jpegs[i]
        if j == item {
            if item.source != nil {
                return errors.New("Can`t delete, because sourced")
            }

            delete(jpegs, i)

            return nil
        }
    }
    */

    return errors.New("No such instance")
}
