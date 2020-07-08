package vpss

//#include "vpss.h"
import "C"

import (
    "errors"
    "time"
    "unsafe"

    "application/pkg/mpp/errmpp"
    "application/pkg/logger"
    "application/pkg/buildinfo"
    "application/pkg/mpp/connection"
)

const (
    defaultDepth = 1
)

func (c *Channel) rutineStart() error {
    if c == nil {
        return errors.New("Null pointer")
    }

    c.rutineRun = true

    go func() {
        c.rutine()
    }()

    err := mppChangeDepth(c.Id, defaultDepth)
    if err != nil {
        return err
    }

    c.depth = defaultDepth

    return nil
}

func (c *Channel) rutineStop() error {
    if c == nil {
		return errors.New("Null pointer")
	}

    c.rutineRun = false

    _ = <- c.rutineCtrl

    err := mppChangeDepth(c.Id, 0)
    if err != nil {
        return err
    }

    c.depth = 0

    return nil
}

func (c *Channel) rutine() {
	if c == nil {
		logger.Log.Fatal().
			Msg("VPSS rutine null pointer")
		return
	}

    periodTreshhold := uint64(1500000 / c.Params.Fps)   //TODO why this value
    periodWait      := uint64(5000 / c.Params.Fps)      //TODO why this value

    logger.Log.Trace().
        Int("channel", c.Id).
        Str("name", "sendDataToClients").
        Uint64("treshhold,us", periodTreshhold).
        Uint64("wait,ms", periodWait).
        Msg("VPSS rutine started")

    for {
        if (!c.rutineRun) {
            break
        }

        var err C.int
        var inErr C.error_in
        var mppFrame unsafe.Pointer
        var pts C.ulonglong

        //hi3516cv100 family doesn`t provide blocking getFrame call
        if buildinfo.Family == "hi3516cv100" {
            time.Sleep(time.Duration(periodWait) * time.Millisecond)         //TODO set proper value
        }

        err = C.mpp_receive_frame(&inErr, C.uint(c.Id), &mppFrame, &pts, C.uint(periodWait));

        if err != C.ERR_NONE {
            logger.Log.Warn().
                Int("channel", c.Id).
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("VPSS failed receive frame")
            continue
        }

        //as stat is not so important, we doesn`t lock on stat write
        c.stat.TsPrev = c.stat.TsLast
        c.stat.TsLast = uint64(pts)
        c.stat.Count++

        var period uint64 = c.stat.TsLast - c.stat.TsPrev

        if period > periodTreshhold {
            if (c.stat.Count > 1) {
                //TODO estimate amount of dropped frames
                c.stat.Drops++
            }
        } else {
            c.stat.PeriodAvg += ((float64(period) - c.stat.PeriodAvg) / float64(c.stat.Count))
        }

        var frame connection.Frame
        frame.Frame = mppFrame
        frame.Pts = uint64(pts)

        c.rawClientsMutex.RLock()
        {
            for client, _ := range c.rawClients {
				client.PushRawFrame(frame)
            }
        }
        c.rawClientsMutex.RUnlock()

        err = C.mpp_release_frame(&inErr, C.uint(c.Id));

        if err != C.ERR_NONE {
            logger.Log.Error().
                Int("channel", c.Id).
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("VPSS failed release frame")
        }
    }

    c.rutineCtrl <- true //TODO improve channel stop

    logger.Log.Trace().
        Int("channel", c.Id).
        Str("name", "sendDataToClients").
        Msg("VPSS rutine stopped")
}
