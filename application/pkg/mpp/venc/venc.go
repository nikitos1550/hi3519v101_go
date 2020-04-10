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

func SubsribeEncoder(encoderId string, ch chan []byte) {
        encoder, encoderExists := Encoders[encoderId]
        if !encoderExists {
                //log.Println("Failed to find encoder ", encoderId)
		logger.Log.Error().
			Str("encoderId", encoderId).
			Msg("Failed to find encoder")
                return
        }

        channels, exists := EncoderSubscriptions[encoder.VencId]
        if !exists {
                createEncoder(encoder)
                channels = make(map[chan []byte]bool)
        } else if !hasSubscription(encoder.VencId) {
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

func RemoveSubscription(encoderId string, ch chan []byte) {
        encoder, encoderExists := Encoders[encoderId]
        if !encoderExists {
                //log.Println("Failed to find encoder ", encoderId)
		logger.Log.Error().
                        Str("encoderId", encoderId).
                        Msg("Failed to find encoder")
                return
        }

        EncoderSubscriptions[encoder.VencId][ch] = false

        if !hasSubscription(encoder.VencId) {
                //log.Println("No subscriptions for ", encoder.VencId, " remove venc")
		logger.Log.Error().
                        Int("vencId", encoder.VencId).
                        Msg("remove venc as No subscriptions")

                deleteEncoder(encoder) //delVenc(encoder.VencId)
        }
}

