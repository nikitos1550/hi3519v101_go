package sample

import (
    "application/pkg/mpp/connection"
)

type sample struct {
    //TODO
}

func (s *sample) Init() error {
    return nil
}

func (s *sample) DeInit() error {
    return nil
}

func (s *sample) GetConnection() connection.Connection {
    return connection.Connection{
        Name: "processing", // + strconv.Itoa(s.Id),
        CType: connection.Push,
        EncoderId: 0,
        MaxWidth: 0,
        MaxHeight: 0,
        MaxFps: 0,
    }
}

func (s *sample) CanAcceptFrame() bool {
    return false
}

func (s *sample) PushFrame(connection.Frame) error {
    return nil
}

//type Client interface {
//    GetConnection() Connection
//    PushData(Frame) error
//}

//func GetConnection() Connection
//func PushData(Frame) error {}
