package main

import (
    "fmt"
    "net/http"

    "github.com/google/uuid"

    "application/archive/recorder"
    "application/core/mpp/connection"

    "application/core/logger"
)

var(
    recorderObj     *recorder.Recorder

    //rutineStop      chan bool
    //rutineDone      chan bool
)

func initRecorder() {
    recorderObj, _ = recorder.New("testrecorder", *flagArchiveRawPath)

    err := connection.ConnectEncodedData(encoderH264Main, recorderObj)
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Connect main encoder to recorder failed")
    }
}

func recorderStatus(w http.ResponseWriter, r *http.Request) {
    if recorderObj.Status() == true {
        fmt.Fprintf(w, "Recording")
    } else {
        fmt.Fprintf(w, "Idle")
    }
}

func recorderStart(w http.ResponseWriter, r *http.Request) {
    if recorderObj.Status() == true {
        fmt.Fprintf(w, "Already recording")
        return
    }

    encoderH264Main.Stop()
    logger.Log.Trace().Msg("encoderH264Main.Stop()")
    encoderH264Main.Reset()
    logger.Log.Trace().Msg("encoderH264Main.Reset()")
    scheduleObj.SetForward()
    logger.Log.Trace().Msg("scheduleObj.SetForward()")
    recorderObj.Start(uuid.New().String())
    logger.Log.Trace().Msg("recorderObj.Start(uuid.New().String())")
    encoderH264Main.Start()
    logger.Log.Trace().Msg("encoderH264Main.Start()")

    preview, err := jpegSmall.GetJpeg()
    if err == nil {
        recorderObj.SetPreview(preview)
    }
    logger.Log.Trace().Msg("Started")
}

func recorderStop(w http.ResponseWriter, r *http.Request) {
    if recorderObj.Status() == false {
        fmt.Fprintf(w, "Idle, nothing to stop")
        return
    }

    rec, err := recorderObj.Stop()
    if err != nil {
        fmt.Fprintf(w, "Can`t finilize recording: %s", err.Error())
        return
    }

    encoderH264Main.Stop()
    encoderH264Main.Reset()

    archiveMutex.Lock()
    defer archiveMutex.Unlock()

    archive[rec.Name] = archiveItem{
        record: rec,
    }
}

//func recorderSchedule(w http.ResponseWriter, r *http.Request) {
//    if recorderObj.Status() == true {
//        fmt.Fprintf(w, "Already recording")
//        return
//    }
//}