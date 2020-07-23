package vpss

//#include "vpss.h"
import "C"

import (
    "errors"
    "time"
    "unsafe"
    "sync"

    "application/core/mpp/errmpp"
    "application/core/logger"
    "application/core/compiletime"
    "application/core/mpp/connection"
)

const (
    defaultDepth = 1
)

func (c *Channel) rutineStart() error {
    if c == nil {
        return errors.New("Null pointer")
    }

    c.rutineRun = true

    c.sendFrame = make(chan connection.Frame, defaultDepth)
    //c.sendFrame = make(chan connection.Frame)
    go c.sendToClients()
    go c.rutine()

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


    //var inErr C.error_in
    //C.mpp_frame_test(&inErr, C.uint(c.Id))

    for {
        if (!c.rutineRun) {
            break
        }

        var err C.int
        var inErr C.error_in
        var mppFrame unsafe.Pointer
        //var rawFrame C.VIDEO_FRAME_INFO_S
        var pts C.ulonglong

        //hi3516cv100 family doesn`t provide blocking getFrame call
        if compiletime.Family == "hi3516cv100" {
            time.Sleep(time.Duration(periodWait) * time.Millisecond)         //TODO set proper value
        }

        var frame connection.Frame

        err = C.mpp_receive_frame(&inErr, C.uint(c.Id), &mppFrame, &pts, C.uint(periodWait), unsafe.Pointer(&frame.FrameMPP));

        if err != C.ERR_NONE {
            logger.Log.Warn().
                Int("channel", c.Id).
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("VPSS failed receive frame")
            continue
        }
        
        //var delta uint64 = uint64(pts) - c.lastPts
        //if delta > 40000 {
        //    logger.Log.Trace().
        //        Uint64("delta", delta).
        //        Int("id", c.Id).
        //        Msg("vpss new frame")
        //    c.lastPts = uint64(pts)
        //}

        //as stat is not so important, we doesn`t lock on stat write
		//{
		//	c.stat.TsPrev = c.stat.TsLast
		//	c.stat.TsLast = uint64(pts)
		//	c.stat.Count++
		//
		//	var period uint64 = c.stat.TsLast - c.stat.TsPrev
		//
		//	if period > periodTreshhold {
		//		if (c.stat.Count > 1) {
		//			//TODO estimate amount of dropped frames
		//			c.stat.Drops++
		//		}
		//	} else {
		//		c.stat.PeriodAvg += ((float64(period) - c.stat.PeriodAvg) / float64(c.stat.Count))
		//	}
		//}

        //var frame connection.Frame

        //frame.Frame     = mppFrame
        //frame.Frame     = unsafe.Pointer(&frame.FrameMPP)

        frame.Pts       = uint64(pts)
        frame.Wg        = new(sync.WaitGroup)
        //frame.FrameMPP  = rawFrame

        //logger.Log.Trace().Uint64("pts", frame.Pts).Msg("New frame")

        //c.sendFrame<-frame
        //logger.Log.Trace().Uint64("pts", frame.Pts).Msg("Frame sent")

        ///////////c.releaseFrame(frame)

		select {
			case c.sendFrame<-frame:
                //logger.Log.Trace().
                //    Uint64("pts", frame.Pts).
                //    Msg("VPSS send client frame")
				break
            default:
                c.releaseFrame(frame)
                logger.Log.Warn().
                    Uint64("pts", frame.Pts).
                    Msg("VPSS dropping raw frame")
				break
		}


    }

    close(c.sendFrame)
    //c.rutineCtrl <- true //TODO improve channel stop

    //logger.Log.Trace().
    //    Int("channel", c.Id).
    //    Str("name", "sendDataToClients").
    //    Msg("VPSS rutine stopped")
}

/*
func (c *Channel) waitForNewFrame() {
	for {
		var err C.int
		var inErr C.error_in
		var mppFrame unsafe.Pointer
		var pts C.ulonglong

		//hi3516cv100 family doesn`t provide blocking getFrame call
		if compiletime.Family == "hi3516cv100" {
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

		var frame connection.Frame

		frame.Frame = mppFrame
		frame.Pts   = uint64(pts)
		frame.Wg    = new(sync.WaitGroup)

		c.newFrame<-frame
	}
}
*/

func (c *Channel) sendToClients() {
    var emptyFrame connection.Frame

    for {
        //select {
        //case frame := <-c.sendFrame:
        frame := <-c.sendFrame
            //logger.Log.Trace().Uint64("pts", frame.Pts).Msg("Sending to clients")

            if frame == emptyFrame {
                logger.Log.Trace().
                    Msg("sendToClients finished")
                c.rutineCtrl <- true
                return
            }

	        c.rawClientsMutex.RLock()
	        {
		        frame.Wg.Add(len(c.rawClients))
		        for client, notify := range c.rawClients {
			        select {
			        case *notify <- frame:
				        //logger.Log.Trace().Uint64("pts", frame.Pts).Msg("Adding 1 receiver wg")
				        break
			        default:
				        frame.Wg.Add(-1)
                        logger.Log.Warn().
                            Str("client", client.FullName()).
                            Msg("VPSS client droppped frame")

				        break
			        }
		        }
		        //logger.Log.Trace().
		        //    Uint64("pts", frame.Pts).
		        //    Msg("VPSS rutine start waiting")
		        frame.Wg.Wait()
		        //logger.Log.Trace().
		        //    Uint64("pts", frame.Pts).
		        //    Msg("VPSS rutine waiting done")
	        }
	        c.rawClientsMutex.RUnlock()

            c.releaseFrame(frame)
            /*
            var err C.int
            var inErr C.error_in

            err = C.mpp_release_frame(&inErr, C.uint(c.Id));

            logger.Log.Trace().
                Uint64("pts", frame.Pts).
                Msg("VPSS release frame")

            if err != C.ERR_NONE {
                logger.Log.Error().
                    Int("channel", c.Id).
                    Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                    Msg("VPSS failed release frame")
            }
            */
        //}
    }
}

func (c *Channel) releaseFrame(frame connection.Frame) {
    var err C.int
    var inErr C.error_in

    //logger.Log.Trace().Uint64("pts", frame.Pts).Msg("Release")

    //for {
    err = C.mpp_release_frame(&inErr, C.uint(c.Id), unsafe.Pointer(&frame.FrameMPP));

    //logger.Log.Trace().
    //    Uint64("pts", frame.Pts).
    //    Msg("VPSS release frame")

    if err != C.ERR_NONE {
        logger.Log.Error().
            Int("channel", c.Id).
            Uint64("pts", frame.Pts).
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VPSS failed release frame")
    }

    //if err == C.ERR_NONE {break}

    //}
}
