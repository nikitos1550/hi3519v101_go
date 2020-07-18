package webrtc

import (
    "errors"

    "github.com/pion/webrtc/v2/pkg/media"
    "github.com/valyala/bytebufferpool"

    "application/core/mpp/connection"
    "application/core/mpp/frames"
    "application/core/logger"
)

const (
    defaultRecvChannelSize = 2
)

//connection.ClientEncodedData interface implementation

func (w *webrtcServer) RegisterEncodedDataSource(source connection.SourceEncodedData, params connection.EncodedDataParams) (*chan frames.FrameItem, error) {
    if w == nil {
        return nil, errors.New("Null pointer")
    }

    w.Lock()
    defer w.Unlock()

    if w.deleted == true {
        logger.Log.Error().
            Msg("Jpeg invoked deleted instance")
        return nil, errors.New("Invoked deleted instance")
    }

    if w.source != nil {
        return nil, errors.New("Already sourced")
    }

    if params.Codec != connection.H264 {
        return nil, errors.New("Only codec h264 supported")
    }

    w.source = source

    if defaultRecvChannelSize == 0 {
        w.notify = make(chan frames.FrameItem)
    } else {
        w.notify = make(chan frames.FrameItem, defaultRecvChannelSize)
    }

    w.rutineStop = make(chan bool)
    w.rutineDone = make(chan bool)

    go w.rutine()

    return &w.notify, nil
}

func (w *webrtcServer) UnregisterEncodedDataSource(source connection.SourceEncodedData) error {
    if w == nil {
        return errors.New("Null pointer")
    }

    w.Lock()
    defer w.Unlock()

    if w.deleted == true {
        logger.Log.Error().
            Msg("webrtc invoked deleted instance")
        return errors.New("Invoked deleted instance")
    }

    if w.source == nil {
        return errors.New("Not sourced")
    }

    w.rutineStop <- true
    <- w.rutineDone

    w.source = nil


    return nil
}

//func (w *webrtcServer) GetNotificator() *chan frames.FrameItem {
//    if w == nil {
//        return nil
//    }
//
//    if w.deleted == true {
//        logger.Log.Error().
//            Msg("webrtc invoked deleted instance")
//        return nil
//    }
//
//    return &w.notify
//}

////////////////////////////////////////////////////////////////////////////////

func (w *webrtcServer) rutine() {

    //storage, err := w.source.GetStorage()
    //if err != nil {
    //  logger.Log.Fatal().
    //        Str("info", err.Error()).
    //        Msg("WebRTC pushData GetStorage")
    //}

    //logger.Log.Trace().
    //    Msg("webrtc rutine started")

    for {
        //select {
            //case frame := <- w.notify:
            frame := <- w.notify
            {
                //logger.Log.Trace().
                //    Int("id", m.id).
                //    Str("name", m.name).
                //    Msg("Mjpeg rutine new frame")

                //var time0 int64
                //var time1 int64
                //var time2 int64

                //time0 = time.Now().UnixNano()

                w.clientsMutex.RLock()

                if len(w.clients) > 0 {

                    //var buf []byte
                    //buf = make([]byte, frame.Size)

                    buf := bytebufferpool.Get()

                    //err := w.getFrame(frame, buf)
                    //if err != nil {
                    //    logger.Log.Error().
                    //        Str("reason", err.Error()).
                    //        Msg("WEBRTC TODO")
                    //}

                    w.writeFrameTo(frame, buf)

                    //time1 = time.Now().UnixNano()

                    for _, client := range(w.clients) {
                        //if client.keyframe == false && frame.Info.Type != 1 {
                        //
                        //} else {
                            var h264Err error
                            if h264Err = client.videoTrack.WriteSample(media.Sample{Data: buf.B, Samples: 90000}); h264Err != nil {
                                logger.Log.Error().
                                    Str("error", h264Err.Error()).
                                    Msg("WebRTC")
                            }
                        //    client.keyframe = true
                        //}
                    }

                    bytebufferpool.Put(buf)

                    //time2 = time.Now().UnixNano()
                }
                w.clientsMutex.RUnlock()

                /*
                logger.Log.Trace().
                    Int64("delta0", (time1-time0)/(1000*1000)).
                    Int64("delta1", (time2-time1)/(1000*1000)).
                    Int64("delta2", (time2-time0)/(1000*1000)).
                    Int("size", frame.Size).
                    Msg("webrtc timing")
                */
            }
//            case <- w.rutineStop:
//                w.rutineDone <- true
//
//                logger.Log.Debug().
//                    Msg("webrtc rutine done")
//
//                return
//        }
    }

    /*
    for {
        new := <- w.notify

        if len(w.clients) != 0 {

            var buf []byte = make([]byte, new.Size)
            n, err := storage.ReadItem(new, buf)

            if err != nil {
                logger.Log.Error().
                    Str("info", err.Error()).
                    Int("n", n).
                    Int("buf", len(buf)).
                    Int("length", new.Size).
                    Msg("sendVideoData storage.ReadItem")
            }

            w.clientsMutex.RLock()
            for _, client := range(w.clients) {
                if ok, _ := client.getConnected(); ok == true {
                    var h264Err error
                    if h264Err = client.videoTrack.WriteSample(media.Sample{Data: buf, Samples: 90000}); h264Err != nil {
                        logger.Log.Error().
                            Str("error", h264Err.Error()).
                            Msg("WebRTC")
                    }
                    //if err = session.AudioTrack.WriteSample(media.Sample{Data: buf, Samples: 960}); err != nil {
                    //    log.Println("Webrtc audio: ", err)
                    //}
                }
            }
            w.clientsMutex.RUnlock()
        }
    }
    */
}

