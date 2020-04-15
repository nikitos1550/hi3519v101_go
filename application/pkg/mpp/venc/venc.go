package venc

import (
	//"log"
	"application/pkg/logger"
)

func Init() {
    loopInit()
	readEncoders()

}

func init() {
}

func SubsribeEncoder(encoderId int, ch chan []byte) {
    encoder, encoderExists := ActiveEncoders[encoderId]
    if !encoderExists {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Failed to find encoder")
        return
    }
		
    _, exists := encoder.Channels[ch]
    if (exists) {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Already subscribed")
        return
    }

    encoder.Channels[ch] = true
    ActiveEncoders[encoderId] = encoder
}

func RemoveSubscription(encoderId int, ch chan []byte) {
    encoder, encoderExists := ActiveEncoders[encoderId]
    if !encoderExists {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Failed to find encoder")
        return
    }
		
    _, exists := encoder.Channels[ch]
    if (exists) {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Not subscribed")
        return
    }

    delete(encoder.Channels, ch)
    ActiveEncoders[encoderId] = encoder
}

