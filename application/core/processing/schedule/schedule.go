package schedule

import (
	"sync"
    //"time"

    "github.com/pkg/errors"

    "application/core/mpp/connection"
)

type Schedule struct {
    sync.RWMutex

	name            string

    startTimestamp  uint64
    stopTimestamp   uint64

    params          connection.FrameCompatibility

    sourceRaw       connection.SourceRawFrame
    rawFramesCh     chan connection.Frame
    rutineStop      chan bool
    rutineDone      chan bool

    clientRaw       connection.ClientRawFrame
    clientCh        *chan connection.Frame

    //finishCb        func ()
}

func New(name string, forward bool) (*Schedule, error) {
    if name == "" {
        return nil, errors.New("Name can`t be empty")
    }

    var start uint64
    var stop uint64

    if forward == true {
        stop = ^uint64(0)
    }

    return &Schedule{
        name: name,
        startTimestamp: start,
        stopTimestamp: stop,
    }, nil
}

func (s *Schedule) Delete() error  {
    s.Lock()
    defer s.Unlock()

    if s.sourceRaw != nil {
        return errors.New("Forward can`t be destroyed because sourced")
    }

    if s.clientRaw != nil {
        return errors.New("Forward can`t be destroyed because of client")
    }

    return nil
}

////////////////////////////////////////////////////////////////////////////////

func (s *Schedule) Name() string {
    s.RLock()
    defer s.RUnlock()

    return s.name
}

func (s *Schedule) FullName() string {
    s.RLock()
    defer s.RUnlock()

    return "schedule:"+s.name
}

////////////////////////////////////////////////////////////////////////////////

func (s *Schedule) SetForward() {
    s.Lock()
    defer s.Unlock()

    s.startTimestamp    = 0
    s.stopTimestamp     = ^uint64(0)
}

/*
func (s *Schedule) SetTimeNano(start uint64, stop uint64) {
    s.Lock()
    defer s.Unlock()

    s.startTimestamp    = start
    s.stopTimestamp     = stop
}

func (s *Schedule) SetTime(start time.Time, stop time.Time) {
    s.Lock()
    defer s.Unlock()

    s.startTimestamp    = uint64(start.UnixNano())
    s.stopTimestamp     = uint64(stop.UnixNano())
}

func (s *Schedule) SetTimeWithDuration(start time.Time, duration time.Duration) {
    s.Lock()
    defer s.Unlock()

    s.startTimestamp    = uint64(start.UnixNano())
    s.stopTimestamp     = uint64(start.Add(duration).UnixNano())
}

func (s *Schedule) GetTimeNano() (uint64, uint64) {
    s.RLock()
    defer s.RUnlock()

    return s.startTimestamp, s.stopTimestamp
}

func (s *Schedule) GetTime() (time.Time, time.Time) {
    s.RLock()
    defer s.RUnlock()

    return time.Unix(0, int64(s.startTimestamp)), time.Unix(0, int64(s.stopTimestamp))
}
*/
