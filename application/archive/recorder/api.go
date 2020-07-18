package recorder

import (
    "github.com/pkg/errors"

    "application/archive/record"
    "application/core/mpp/frames"
)

func (r *Recorder) Init() error {
    r.Lock()
    defer r.Unlock()

    if r.record != nil {
        return errors.New("Already recording")
    }

    return nil
}

func (r *Recorder) InitWithSchedule() error {
    r.Lock()
    defer r.Unlock()

    if r.record != nil {
        return errors.New("Already recording")
    }

    return nil
}

func (r *Recorder) Start() error {
    r.Lock()
    defer r.Unlock()

    return nil
}

func (r *Recorder) Stop() (*record.Record, error) {
    r.Lock()
    defer r.Unlock()

    return nil, nil
}

func (r *Recorder) Record() (*record.Record, error) {
    return nil, nil
}

func (r *Recorder) processFrame(f frames.FrameItem) {
    r.RLock()
    defer r.RUnlock()

}

func (r *Recorder) Info() {}
