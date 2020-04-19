//+build openapi

package vpss

import (
    "sync"
	"unsafe"
    //"errors"

    "application/pkg/logger"
)

type Channel struct {
	ChannelId  int
	Width int
	Height int
	Fps int
	CropX int                           //Not used
	CropY int                           //Not used
	CropWidth int                       //Not used
	CropHeight int                      //Not used
    Mutex sync.RWMutex                  //Not used
	Started bool                        //is channel active
	Clients map[int] unsafe.Pointer     //int - id processing, callback processing
    Clients2 map[int] chan unsafe.Pointer
}

var (
	Channels map[int] Channel
)

func init() {
	Channels = make(map[int] Channel)
}

func Init() {
    err := initFamily()
    if err != nil {
        logger.Log.Fatal().
            Str("error", err.Error()).
            Msg("VPSS")
    }
    logger.Log.Debug().
        Msg("VPSS inited")

}

func StartChannel(channel Channel)  (int, string)  {
	_, channelExists := Channels[channel.ChannelId]
	if (channelExists) {
        //return -1, errors.New("Channel already exists")
		return -1, "Channel already exists"
	}

	createChannel(channel)

	Channels[channel.ChannelId] = channel
	return channel.ChannelId, ""
}

func StopChannel(channelId int)  (int, string)  {
	channel, channelExists := Channels[channelId]
	if (!channelExists) {
		return -1, "Channel does not exist"
	}

	destroyChannel(channel)

	delete(Channels, channelId)
	return 0, ""
}

func SubscribeChannel(channelId int, processingId int, callback unsafe.Pointer)  (int, string)  {
	channel, channelExists := Channels[channelId]
	if (!channelExists) {
		return -1, "Channel does not exist"
	}

	_, callbackExists := channel.Clients[processingId]
	if (callbackExists) {
		return -1, "Already subscribed"
	}

	channel.Clients[processingId] = callback
	Channels[channelId] = channel
	
	return 0, ""
}

func UnsubscribeChannel(channelId int, processingId int)  (int, string)  {
	channel, channelExists := Channels[channelId]
	if (!channelExists) {
		return -1, "Channel does not exist"
	}

	_, callbackExists := channel.Clients[processingId]
	if (!callbackExists) {
		return -1, "Not subscribed"
	}

	delete(channel.Clients, processingId)

	return 0, ""
}
