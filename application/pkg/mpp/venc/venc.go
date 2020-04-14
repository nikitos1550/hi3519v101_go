package venc

import (
	//"log"
	"application/pkg/logger"
)

func Init() {
    loopInit()
	readEncoders()

}

var (
        EncoderSubscriptions map[int]map[chan []byte]bool
)

func init() {
        EncoderSubscriptions = make(map[int]map[chan []byte]bool)
}

func SubsribeEncoder(encoderId int, ch chan []byte) {
        encoder, encoderExists := ActiveEncoders[encoderId]
        if !encoderExists {
			logger.Log.Error().
				Int("encoderId", encoderId).
				Msg("Failed to find encoder")
            return
        }
		
        channels, exists := EncoderSubscriptions[encoder.VencId]
        if !exists || !hasSubscription(encoder.VencId) {
                CreateEncoder(encoder)
                channels = make(map[chan []byte]bool)
        } else {
                addVenc(encoder.VencId)
        }

        channels[ch] = true
        EncoderSubscriptions[encoder.VencId] = channels
}

func hasSubscription(vencId int) bool {
        for _, value := range EncoderSubscriptions[vencId] {
                if value {
                        return true
                }
        }
        return false
}

func RemoveSubscription(encoderId int, ch chan []byte) {
        encoder, encoderExists := ActiveEncoders[encoderId]
        if !encoderExists {
                //log.Println("Failed to find encoder ", encoderId)
		logger.Log.Error().
                        Int("encoderId", encoderId).
                        Msg("Failed to find encoder")
                return
        }

        EncoderSubscriptions[encoder.VencId][ch] = false

        if !hasSubscription(encoder.VencId) {
                //log.Println("No subscriptions for ", encoder.VencId, " remove venc")
		logger.Log.Debug().
                        Int("vencId", encoder.VencId).
                        Msg("remove venc as No subscriptions")

                DeleteEncoder(encoder) //delVenc(encoder.VencId)
        }
}

