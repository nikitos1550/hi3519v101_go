//+build arm
//+build hi3516av100 hi3516av200 hi3516cv100 hi3516cv200 hi3516cv300 hi3516cv500 hi3516ev200 hi3519av100 hi3559av100

package venc

//#include "venc.h"
import "C"

import (
    "unsafe"
    "sync"
    "errors"

    "application/core/logger"
    "application/core/mpp/frames"
)

//var mutex = &sync.Mutex{}

var (
    mutex       sync.Mutex
    encoders    map[int] *Encoder
)

//export go_callback_receive_data
func go_callback_receive_data(id C.int, info_pointer *C.info_from_c, data_pointer *C.data_from_c, data_num C.int) { 

    mutex.Lock()
    defer mutex.Unlock()

    encoder, exist := encoders[int(id)]

    if !exist {
        logger.Log.Warn().
            Int("id", int(id)).
            Msg("VENC loop callback invoked for unknown id")
            return
    }

    //if int(id) >= channelsAmount {
    //    logger.Log.Error().
    //        Int("id", int(id)).
    //        Msg("go_callback_receive_data id not valid")
    //        return
    //}

    //moved below
    //channels[id].mutex.RLock()          //read lock
    //defer channels[id].mutex.RUnlock()

    //if channels[id].created == false {
    //    logger.Log.Error().
    //        Int("id", int(id)).
    //        Msg("go_callback_receive_data encoder is not created")
    //    return
    //}

    //if channels[id].started == false {
    //    logger.Log.Error().
    //        Int("id", int(id)).
    //        Msg("go_callback_receive_data encoder is not started")
    //    return
    //}


    num := int(data_num)

    var infoFromC *C.info_from_c = info_pointer

    /*
    var refType string

    switch infoFromC.ref_type {
        case C.BASE_IDRSLICE:
            refType = "BASE_IDRSLICE"
        case C.BASE_PSLICE_REFTOIDR:
            refType = "BASE_PSLICE_REFTOIDR"
        case C.BASE_PSLICE_REFBYBASE:
            refType = "BASE_PSLICE_REFBYBASE"
        case C.BASE_PSLICE_REFBYENHANCE:
            refType = "BASE_PSLICE_REFBYENHANCE"
        case C.ENHANCE_PSLICE_REFBYENHANCE:
            refType = "ENHANCE_PSLICE_REFBYENHANCE"
        case C.ENHANCE_PSLICE_NOTFORREF:
            refType = "ENHANCE_PSLICE_NOTFORREF"
        default:
            refType = "unknown"
    }
    */

    //logger.Log.Debug().
    //    Uint64("pts", uint64(infoFromC.pts)).
    //    Str("refType", refType).
    //    Msg("VENC")
    

    //infoFromC.ref_type:
    //BASE_IDRSLICE = 0           //IDR frame at the base layer
    //BASE_PSLICE_REFTOIDR        //P-frame at the base layer, referenced by other frames at the base layer and references only IDR frames
    //BASE_PSLICE_REFBYBASE       //P-frame at the base layer, referenced by other frames at the base layer
    //BASE_PSLICE_REFBYENHANCE    //P-frame at the base layer, referenced by frames at the enhance layer
    //ENHANCE_PSLICE_REFBYENHANCE
    //ENHANCE_PSLICE_NOTFORREF

    dataFromC := (*[1 << 10]C.data_from_c)(unsafe.Pointer(data_pointer))[:num:num]
    length := 0
    for i := 0; i < num; i++ {
        length = length + int(dataFromC[i].length)
    }

    //data := make([]byte, length)

    offset := 0
    var p [][]byte = make([][]byte, num)

    for i := 0; i < num; i++ {
        p[i] = (*[1 << 28]byte)(unsafe.Pointer(dataFromC[i].data))[:dataFromC[i].length:dataFromC[i].length]
        //n := copy(data[offset:], p)
        offset = offset + int(dataFromC[i].length)
    }

    //if num > 1 {
    //logger.Log.Trace().
    //    Int("num", num).
    //    Msg("VENC loop")
    //}

    //logger.Log.Trace().
    //    Int("length", length).
    //    Uint64("pts", uint64(infoFromC.pts)).
    //    Str("refType", refType).
    //     Msg("VENC")

    var info frames.FrameInfo
    info.Seq = uint32(infoFromC.seq)
    info.Pts = uint64(infoFromC.pts)
    if  infoFromC.ref_type == C.BASE_IDRSLICE {
        info.Type = 1
    }

    encoder.clientsMutex.RLock() //encoders[id].clientsMutex.RLock()
    doCopy := false
    if len(encoder.clients) > 0 { //if len(encoders[id].clients) > 0 {
        doCopy = true
    }
    encoder.clientsMutex.RUnlock() //encoders[id].clientsMutex.RUnlock()

    if doCopy {
        //slot, err := encoders[id].storage.WritevNext(p, info)
        slot, err := encoder.storage.WritevNext(p, info)
        if err != nil {
            logger.Log.Error().
                Str("reason", err.Error()).
                Msg("VENC loop")
        }

        var item frames.FrameItem
        item.Slot = slot
        item.Size = length
        item.Info = info

        //encoders[id].clientsMutex.RLock()          //read lock
        encoder.clientsMutex.RLock()
        {
            //for client, notify := range encoders[id].clients {
            for client, notify := range encoder.clients {
                if notify != nil {
                    select {
                    case *notify <- item:
                        break
                    default:
                        logger.Log.Warn().
                            Str("name", client.FullName()).
                            Msg("VENC LOOP client dropped frame")
                    }
                    //if notify != nil && len(*notify) < cap(*notify) {
                    //*notify <- item
                    //logger.Log.Trace().
                    //    Int("slot", slot).
                    //    Int("len ch", len(*notify)).
                    //    Int("cap ch", cap(*notify)).
                    //    Msg("VENC LOOP sent to client")
                }
            }
        }

        //encoders[id].clientsMutex.RUnlock()
        encoder.clientsMutex.RUnlock()
    }
    /*
    //for ch, enabled := range encoders[id].clients {
    //    if (enabled){
    //        if (cap(ch) <= len(ch)) {
    //            <-ch
    //        }
    //
    //        var tmpData ChannelEncoder
    //        tmpData.Data = data
    //        tmpData.Pts = uint64(infoFromC.pts)
    //
    //        ch <- tmpData
    //    }
    //}
    */
}


//Rules:
//addVenc/delVenc functions should operate one in a time.
//There should be some sync for them run:
//1) exlusive mutex
//2) query (based obviosly on go channels)

func addToLoop(id int, e *Encoder, codec Codec) error {
    mutex.Lock()
	defer mutex.Unlock()

    var errorCode C.uint
    var codecIn C.uint

    switch codec {
        case MJPEG:
            codecIn = C.CODEC_MJPEG
        case H264:
            codecIn= C.CODEC_H264
        case H265:
            codecIn = C.CODEC_H265
        default:
            return errors.New("mpp_data_loop_add unknown codec")
    }

    err := C.mpp_data_loop_add(&errorCode, C.uint(id), codecIn)
    if err != C.ERR_NONE {
        return errors.New("mpp_data_loop_add TODO error")
    }

    encoders[id] = e

    return nil
}

func removeFromLoop(id int) error {
    mutex.Lock()
	defer mutex.Unlock()

    var errorCode C.uint

    err := C.mpp_data_loop_del(&errorCode, C.uint(id))
    if err != C.ERR_NONE {
        return errors.New("mpp_data_loop_del TODO error")
    }

    delete(encoders, id)

    return nil
}


func loopInit() error {
    var errorCode C.uint

    err := C.mpp_data_loop_init(&errorCode)
    if err != C.ERR_NONE {
        return errors.New("mpp_data_loop_init TODO error")
    }

    encoders = make(map[int] *Encoder)

    return nil
}
