package vpss

//#include "vpss.h"
import "C"

import (
    "time"
    "unsafe"

    "application/pkg/mpp/errmpp"
    "application/pkg/logger"
    "application/pkg/buildinfo"
)

func sendDataToClients(c *channel) {
    periodTreshhold := uint64(1500000 / c.params.Fps)
    periodWait      := uint64(5000 / c.params.Fps)

    logger.Log.Trace().
        Int("channel", c.id).
        Str("name", "sendDataToClients").
        Uint64("treshhold", periodTreshhold).
        Uint64("wait", periodWait).
        Msg("VPSS rutine started")

    for {
        if (!c.started){
            break
        }

        var err C.int
        var inErr C.error_in
        var frame unsafe.Pointer
        var pts C.ulonglong

        //hi3516cv100 family doesn`t provide blocking getFrame call
        if buildinfo.Family == "hi3516cv100" {  //TODO
            time.Sleep(1 * time.Second)         //now we will just sleep here
        }

        err = C.mpp_receive_frame(&inErr, C.uint(c.id), &frame, &pts, C.uint(periodWait));

        if err != C.ERR_NONE {
            logger.Log.Warn().
                Int("channel", c.id).
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

        c.mutex.RLock()         //clients list read lock
        {
            for processing, _ := range c.clients {
                processing.Callback(frame)
            }

            //for _, _ := range c.clents2 {
            //    //TODO
            //}
        }
        c.mutex.RUnlock()

        err = C.mpp_release_frame(&inErr, C.uint(c.id));

        if err != C.ERR_NONE {
            logger.Log.Error().
                Int("channel", c.id).
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("VPSS failed release frame")
        }
    }

    c.rutineStop <- true //TODO improve channel stop

    logger.Log.Trace().
        Int("channel", c.id).
        Str("name", "sendDataToClients").
        Msg("VPSS rutine stopped")
}

