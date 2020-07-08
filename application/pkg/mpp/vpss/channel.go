package vpss

import (
    "sync"

    "application/pkg/mpp/connection"
)

type Channel struct {
    Id                  int                                                     `json:"id"`

    Params              Parameters                                              `json:"parameters"`
    stat                statistics                                              `json:"-"`

    Created             bool                                                    `json:"created"`
    Locked              bool                                                    `json:"locked"`

    mutex               sync.RWMutex                                            `json:"-"`

    depth               int                                                     `json:"-"`

    rawClients          map[connection.ClientRawFrame] bool                     //TODO
    rawClientsMutex     sync.RWMutex                                            `json:"-"`

    bindClients         map[connection.ClientBind] connection.BindInformation   //TODO
    bindClientsMutex    sync.RWMutex                                            `json:"-"`

    rutineRun           bool                                                    `json:"-"`
    rutineCtrl          chan bool                                               `json:"-"`
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
