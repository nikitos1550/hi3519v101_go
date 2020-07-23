package recorder

import (
    "github.com/pkg/errors"
    "github.com/valyala/bytebufferpool"

    "application/archive/record"
    "application/core/mpp/frames"
    "application/core/logger"
)

func (r *Recorder) Start(name string) error {
    r.Lock()
    defer r.Unlock()

    if r.record != nil {
        return errors.New("TODO")
    }

    rec, err := record.New(r.path, name, r.codec) //uuid.New().String())
    if err != nil {
        return errors.Wrap(err, "Start record failed")
    }

    r.record = rec

    return nil
}

func (r *Recorder) Stop() (*record.Record, error) {
    r.Lock()
    defer r.Unlock()

    if r.record == nil {
        return nil, errors.New("Not inited")
    }

    rec := r.record
    rec.Close()
    r.record = nil

    return rec, nil
}

func (r *Recorder) SetPreview(jpeg []byte) error {
    r.RLock()
    defer r.RUnlock()

    return r.record.SetPreview(jpeg)
}

func (r *Recorder) processFrame(f frames.FrameItem) {
    r.RLock()
    defer r.RUnlock()

    //logger.Log.Trace().
    //    Uint64("delta", f.Info.Pts - r.lastPts).
    //    Msg("recorder new frame")
    //r.lastPts = f.Info.Pts

    if r.record != nil {
        s, err := r.source.GetStorage()
        if err != nil {
            return
        }

        buf := bytebufferpool.Get()
        defer bytebufferpool.Put(buf)

        _, err = s.WriteItemTo(f, buf)
        if err != nil {
            logger.Log.Warn().
                Str("reason", err.Error()).
                Msg("Can get frame")
            return
        }

        n, err := r.record.Write(f.Info.Pts, buf.B)
        if err != nil {
            logger.Log.Warn().
                Str("reson", err.Error()).
                Int("n", n).
                Msg("Can write frame to file")
        }
    }
}

func (r *Recorder) Status() bool {
    r.RLock()
    defer r.RUnlock()

    if r.record != nil {
        return true
    }

    return false
}
