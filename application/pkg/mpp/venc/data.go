package venc

import (
    "application/pkg/logger"
)

func SubsribeEncoderData(encoderId int, ch chan []byte) {
    encoder, encoderExists := ActiveEncoders[encoderId]
    if !encoderExists {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Failed to find encoder")
        return
    }
		
    _, exists := encoder.DataChannels[ch]
    if (exists) {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Already subscribed")
        return
    }

    encoder.DataChannels[ch] = true
    ActiveEncoders[encoderId] = encoder
}

func RemoveDataSubscription(encoderId int, ch chan []byte) {
    encoder, encoderExists := ActiveEncoders[encoderId]
    if !encoderExists {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Failed to find encoder")
        return
    }
		
    _, exists := encoder.DataChannels[ch]
    if (!exists) {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Not subscribed")
        return
    }


    delete(encoder.DataChannels, ch)
    ActiveEncoders[encoderId] = encoder
}