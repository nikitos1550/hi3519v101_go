//+build openapi

package vpss

import (
    "sync"

    "application/pkg/common"
)

type channel struct {
    id          int

    params      Parameters
    stat        statistics

    started     bool

    mutex       sync.RWMutex
    clients     map[common.Processing] bool //int - id processing, callback processing
    rutineStop  chan bool
}

type Parameters struct {
    Width       int
    Height      int
    Fps         int
    CropX       int                           //Not used
    CropY       int                           //Not used
    CropWidth   int                       //Not used
    CropHeight  int                      //Not used
}

type statistics struct {
    Count       uint64  `json:"count"`
    Drops       uint64  `json:"drops"`

    PeriodAvg   float64 `json:"period"`

    TsPrev      uint64
    TsLast      uint64
}

var (
    channels []channel
)

func init() {
    channels = make([]channel, channelsAmount)

    for i := 0; i < channelsAmount; i++ {
		channels[i].id = i
	}
}
