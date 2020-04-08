//+build openapi

package vpss

import (
    "log"
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
	Clients map[unsafe.Pointer] bool
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
		return 1, "Channel already exists"
	}

	CreateChannel(channel)

	Channels[channel.ChannelId] = channel
	return 0, ""
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

func SubscribeChannel(channelId int, callback unsafe.Pointer)  (int, string)  {
	channel, channelExists := Channels[channelId]
	if (!channelExists) {
		return 1, "Channel does not exist"
	}

	_, callbackExists := channel.Clients[callback]
	if (callbackExists) {
		return 1, "Already subscribed"
	}

	channel.Clients[callback] = true
	Channels[channelId] = channel
	log.Println("SubscribeChannel", len(Channels[channelId].Clients))
	
	return 0, ""
}