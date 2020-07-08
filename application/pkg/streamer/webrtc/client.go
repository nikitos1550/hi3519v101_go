package webrtc

import (
    "errors"

    "github.com/pion/webrtc/v2/pkg/media"

    "application/pkg/mpp/connection"
    "application/pkg/mpp/frames"

    "application/pkg/logger"
)

//connection.ClientEncodedData interface implementation

func (w *webrtcServer) RegisterEncodedDataSource(connection.SourceEncodedData, connection.EncodedDataParams) error {
    if w == nil {
        return errors.New("Null pointer")
    }

    w.Lock()
    defer w.Unlock()

    return nil
}

func (w *webrtcServer) DisconnectEncodedDataSource(connection.SourceEncodedData) error {
    if w == nil {
        return errors.New("Null pointer")
    }

    w.Lock()
    defer w.Unlock()

    return nil
}

func (w *webrtcServer) GetNotificator() *chan frames.FrameItem {
    if w == nil {
        return nil
    }

    w.Lock()
    defer w.Unlock()

    return nil
}

////////////////////////////////////////////////////////////////////////////////

func (w *webrtcServer) rutine() {

    storage, err := w.source.GetStorage()
    if err != nil {
      logger.Log.Fatal().
            Str("info", err.Error()).
            Msg("WebRTC pushData GetStorage")
    }

    for {
        new := <- w.Notify

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
    //TODO
}

