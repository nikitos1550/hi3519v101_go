package vpss

import (
    "sync"

    "application/core/mpp/connection"
)

type Channel struct {
    Id                  int                                                     `json:"id"`
    name                string                                                  //TODO

    Params              Parameters                                              `json:"parameters"`
    stat                statistics                                              `json:"-"`

    //Created             bool                                                    `json:"created"`
    //Locked              bool                                                    `json:"locked"`

    mutex               sync.RWMutex                                            `json:"-"`

    depth               int                                                     `json:"-"`

    rawClients          map[connection.ClientRawFrame] *chan connection.Frame   `json:"-"`//TODO
    rawClientsMutex     sync.RWMutex                                            `json:"-"`

    bindClients         map[connection.ClientBind] connection.BindInformation   `json:"-"`//TODO
    bindClientsMutex    sync.RWMutex                                            `json:"-"`

    //sync rutine
    rutineRun           bool                                                    `json:"-"`
    rutineCtrl          chan bool                                               `json:"-"`

    //async rutine
    sendFrame           chan connection.Frame
    //newFrame            chan connection.Frame
    //releaseFrame        chan connection.Frame
    //rutineDone          chan bool
    lastPts             uint64
}

type Parameters struct {
    Width       int     `json:"width"`
    Height      int     `json:"height"`
    Fps         int     `json:"fps"`
    CropX       int     `json:"cropx,omitempty"`
    CropY       int     `json:"cropy,omitempty"`
    CropWidth   int     `json:"cropwidth,omitempty"`
    CropHeight  int     `json:"cropheight,omitempty"`
}

type statistics struct {
    Count       uint64  `json:"count"`
    Drops       uint64  `json:"drops"`

    PeriodAvg   float64 `json:"averageperiod"`

    TsPrev      uint64  `json:"-"`
    TsLast      uint64  `json:"-"`
}
