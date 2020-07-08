package venc

import (
    "errors"

    "application/pkg/mpp/connection"
	"application/pkg/mpp/frames"
)

//connection.SourceEncodedData interface implementation

func (e *Encoder) AddEncodedDataClient(client connection.ClientEncodedData) error {
    if e == nil {
        return errors.New("Null pointer")
    }

    e.mutex.RLock()
    defer e.mutex.RUnlock()

    if e.Created == false {
        return errors.New("Encoder is not created")
    }

    e.clientsMutex.Lock()
    defer e.clientsMutex.Unlock()

    _, exist := e.clients[client]
    if exist {
        return errors.New("Client already added")
    }

    var codec connection.CodecType

    switch e.Params.Codec {
        case MJPEG:
            codec = connection.MJPEG
        case H264:
            codec = connection.H264
        case H265:
            codec = connection.H265
    }

    var encodedDataParams connection.EncodedDataParams = connection.EncodedDataParams{
        Codec: codec,
    }

    err := client.RegisterEncodedDataSource(e, encodedDataParams)
    if err != nil {
        return err
    }

    e.clients[client] = client.GetNotificator()

	return nil
}

func (e *Encoder) RemoveEncodedDataClient(client connection.ClientEncodedData) error {
    if e == nil {
        return errors.New("Null pointer")
    }

    e.mutex.RLock()
    defer e.mutex.RUnlock()

    if e.Created == false {
        return errors.New("Encoder is not created")
    }

    e.clientsMutex.Lock()
    defer e.clientsMutex.Unlock()

    _, exist := e.clients[client]
    if !exist {
        return errors.New("Client is not in list")
    }

    err := client.UnregisterEncodedDataSource(e)
    if err != nil {
        return err
    }

    delete(e.clients, client)

	return nil
}

//func (e *Encoder) UnregisterEncodedDataClient(connection.ClientEncodedData) error {
//    if e == nil {
//        return errors.New("Null pointer")
//    }
//
//	return nil
//}

func (e *Encoder) GetStorage() (*frames.Frames, error) {
    if e == nil {
        return nil, errors.New("Null pointer")
    }

    e.mutex.RLock()
    defer e.mutex.RUnlock()

    if e.Created == false {
        return nil, errors.New("Encoder is not created")
    }

    if e.Started == false { //TODO
        return nil, errors.New("Encoder is not started")
    }

    return &e.storage, nil
}
