package recorder

import (
    "github.com/pkg/errors"
    //"github.com/valyala/bytebufferpool"

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

        //buf := bytebufferpool.Get()
        //defer bytebufferpool.Put(buf)

        var buf []byte

        sps, err := s.SPS()
        if err != nil {
            logger.Log.Warn().Str("reason", err.Error()).Msg("SPS get")
        }
        //logger.Log.Trace().Int("len", len(sps)).Msg("sps")
        if len(sps) == 0 {
            logger.Log.Warn().Msg("SPS len 0")
        }

        pps, err := s.PPS()
        if err != nil {
            logger.Log.Warn().Str("reason", err.Error()).Msg("PPS get")
        }
        if len(pps) == 0 {
            logger.Log.Warn().Msg("PPS len 0")
        }
        //logger.Log.Trace().Int("len", len(pps)).Msg("pps")

        r.record.SetSPSPPS(sps, pps)//r.record.ConfigureTs(sps, pps)


        //_, err = s.WriteItemTo(f, buf)
        _, err = s.ReadItemAlloc(f, &buf)
        if err != nil {
            logger.Log.Warn().
                Str("reason", err.Error()).
                Msg("Can get frame")
            return
        }

        //n, err := r.record.Write(f.Info.Pts, buf.B)
        n, err := r.record.Write(f.Info.Pts, buf)
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
