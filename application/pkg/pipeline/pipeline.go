//+build nobuild

package pipeline

import (
	"application/pkg/mpp/venc"
	"application/pkg/mpp/vpss"
	"application/pkg/processing"

	"application/pkg/common"
	"application/pkg/logger"
)

func Init() {
}

func CreatePipeline(encoderName string)  (int, string)  {
	encoder, encoderExists := venc.PredefinedEncoders[encoderName]
	if (!encoderExists) {
		return -1, "Failed to find encoder  " + encoderName
	}

	var freeChannel int = 0
	for{
		_, channelExists := vpss.Channels[freeChannel]
		if (!channelExists) {
			break
		}
		freeChannel++
	}
	logger.Log.Info().Int("channelId", freeChannel).Msg("Found free channel ")

	channel := createChannelFromEncoder(encoder)
	channel.ChannelId = freeChannel
	logger.Log.Info().Int("channelId", channel.ChannelId).Msg("Starting Channel")
	channelId, err := vpss.StartChannel(channel)
	if channelId < 0 {
		return channelId, err
	}
	logger.Log.Info().Int("channelId", channelId).Msg("Channel was created")

	params := make(map[string][]string)
	processingId, err := processing.CreateProcessing("proxy", params)
	if processingId < 0 {
		return processingId, err
	}
	logger.Log.Info().Int("processingId", processingId).Msg("Processing was created")

	encoderId, err := venc.CreatePredefinedEncoder(encoderName)
	if encoderId < 0 {
		return encoderId, err
	}
	logger.Log.Info().Int("encoderId", encoderId).Msg("Encoder was created")

	activeEncoder, _ := venc.ActiveEncoders[encoderId]

	activeProcessing, exists := processing.ActiveProcessings[processingId]
	if (!exists) {
		return -1, "Processing not created"
	}

	errId, err := vpss.SubscribeChannel(channelId, activeProcessing.Proc)
	if errId < 0 {
		return errId, err
	}
	activeProcessing.InputChannel = channelId
	processing.ActiveProcessings[processingId] = activeProcessing
	logger.Log.Info().Int("channelId", channelId).Int("processingId", processingId).Msg("Subscribed to channel")

	errId, err = processing.SubscribeEncoderToProcessing(processingId, activeEncoder)
	if errId < 0 {
		return errId, err
	}
	activeEncoder.ProcessingId = processingId
	venc.ActiveEncoders[encoderId] = activeEncoder
	logger.Log.Info().Int("processingId", processingId).Int("encoderId", encoderId).Msg("Subscribed to processing")

	return encoderId, ""
}

func createChannelFromEncoder(encoder venc.PredefinedEncoder)  (vpss.Channel)  {
	ch := vpss.Channel{
		ChannelId: -1,
		Width: encoder.Width,
		Height: encoder.Height,
		Fps: 30,
		CropX: 0,
		CropY: 0,
		CropWidth: 0,
		CropHeight: 0,
		Started: true,
		Clients: make(map[common.Processing] bool),
	}

	return ch
}
