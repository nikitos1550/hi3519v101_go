package connection

import (
    "application/core/mpp/frames"
)

type FrameCompatibility struct {
    Width   int
    Height  int
    Fps     int
}

type SourceRawFrame interface {
    AddRawFrameClient(ClientRawFrame) error                                     //will be called by 3rd party, can`t be called from client object
    RemoveRawFrameClient(ClientRawFrame) error                                  //will be called by 3rd party, can`t be called from client object
    FullName() string
}

type ClientRawFrame interface {
    RegisterRawFrameSource(SourceRawFrame, FrameCompatibility) (*chan Frame, error) //will be called by source object
    UnregisterRawFrameSource(SourceRawFrame) error                                  //will be called by source object
    FullName() string
}

type BindClientType  int
const (
    Encoder BindClientType  = 1
)

type BindInformation struct {
    ClientType  BindClientType
    Id          int
}

type SourceBind interface {
    AddBindClient(ClientBind) error                                             //will be called by 3rd party, can`t be called from client object
    RemoveBindClient(ClientBind) error                                          //will be called by 3rd party, can`t be called from client object
    FullName() string
}

type ClientBind interface {
    RegisterBindSource(SourceBind, FrameCompatibility) (BindInformation, error) //will be called by source object
	UnregisterBindSource(SourceBind) error                                      //will be called by source object
    FullName() string
}


type SourceEncodedData interface {
    AddEncodedDataClient(ClientEncodedData) error                               //will be called by 3rd party, can`t be called from client object
    RemoveEncodedDataClient(ClientEncodedData) error                            //will be called by 3rd party, can`t be called from client object
    GetStorage() (*frames.Frames, error)                                        //will be called by client object
    FullName() string
}

type CodecType int
const (
    MJPEG   CodecType = 1
    H264    CodecType = 2
    H265    CodecType = 3
)

type EncodedDataParams struct {
    Codec   CodecType
}

type ClientEncodedData interface {
    RegisterEncodedDataSource(SourceEncodedData, EncodedDataParams) (*chan frames.FrameItem, error) //will be called by source object //TODO add params
    UnregisterEncodedDataSource(SourceEncodedData) error                                            //will be called by source object
    FullName() string
}

func ConnectRawFrame(s SourceRawFrame, c ClientRawFrame) error {
    err := s.AddRawFrameClient(c)
    if err != nil {
        return err
    }

    return nil
}

func DisconnectRawFrame(s SourceRawFrame, c ClientRawFrame) error {
    err := s.RemoveRawFrameClient(c)
    if err != nil {
        return err
    }

    return nil
}

func ConnectBind(s SourceBind, c ClientBind) error {
    err := s.AddBindClient(c)
    if err != nil {
        return err
    }

    return nil
}

func DisconnectBind(s SourceBind, c ClientBind) error {
    err := s.RemoveBindClient(c)
    if err != nil {
        return err
    }

    return nil
}

func ConnectEncodedData(s SourceEncodedData, c ClientEncodedData) error {
    err := s.AddEncodedDataClient(c)
    if err != nil {
        return err
    }

    return nil
}

func DisconnectEncodedData(s SourceEncodedData, c ClientEncodedData) error {
    err := s.RemoveEncodedDataClient(c)
    if err != nil {
        return err
    }

    return nil
}
