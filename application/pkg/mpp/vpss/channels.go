//+build openapi

package vpss

import (
    "sync"
	"unsafe"
)

type Channel struct {
	ChannelId  int
	Width int
	Height int
	Fps int
	CropX int
	CropY int
	CropWidth int
	CropHeight int
    Mutex sync.RWMutex
	Started bool
	Clients map[int] unsafe.Pointer
}

var (
	Channels map[int] Channel
)

func init() {
	Channels = make(map[int] Channel)
}

func StartChannel(channel Channel)  (int, string)  {
	_, channelExists := Channels[channel.ChannelId]
	if (channelExists) {
		return -1, "Channel already exists"
	}

	CreateChannel(channel)

	Channels[channel.ChannelId] = channel
	return channel.ChannelId, ""
}

func StopChannel(channelId int)  (int, string)  {
	channel, channelExists := Channels[channelId]
	if (!channelExists) {
		return 1, "Channel does not exist"
	}

	DestroyChannel(channel)

	delete(Channels, channelId)
	return 0, ""
}

func SubscribeChannel(channelId int, processingId int, callback unsafe.Pointer)  (int, string)  {
	channel, channelExists := Channels[channelId]
	if (!channelExists) {
		return 1, "Channel does not exist"
	}

	_, callbackExists := channel.Clients[processingId]
	if (callbackExists) {
		return 1, "Already subscribed"
	}

	channel.Clients[processingId] = callback
	Channels[channelId] = channel
	
	return 0, ""
}

func UnsubscribeChannel(channelId int, processingId int)  (int, string)  {
	channel, channelExists := Channels[channelId]
	if (!channelExists) {
		return 1, "Channel does not exist"
	}

	_, callbackExists := channel.Clients[processingId]
	if (!callbackExists) {
		return 1, "Not subscribed"
	}

	delete(channel.Clients, processingId)

	return 0, ""
}