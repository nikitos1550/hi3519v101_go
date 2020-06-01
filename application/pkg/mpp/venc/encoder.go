package venc

import (
	"encoding/json"
	"io/ioutil"
	"log"

    "application/pkg/processing"
)

type PredefinedEncoder struct {
    Format string 
    Width int 
    Height int 
    Bitrate int 
}

type ActiveEncoder struct {
	VencId int 
	ProcessingId int 
	Format string 
	Width int 
	Height int 
	Bitrate int 
	Channels map[chan []byte]bool
	DataChannels map[chan []byte]bool
}

var PredefinedEncoders map[string] PredefinedEncoder
var ActiveEncoders map[int] ActiveEncoder
var lastEncoderId int

func init() {
	PredefinedEncoders = make(map[string] PredefinedEncoder)
	ActiveEncoders = make(map[int] ActiveEncoder)
	lastEncoderId = 0
}

func readEncoders() {
	path := "/opt/configs/encoders.json"
    data, err := ioutil.ReadFile(path)
    if err != nil {
		log.Fatal("Failed to read records from file " + path)
		return
    }
    
	err = json.Unmarshal(data, &PredefinedEncoders)
    if err != nil {
        log.Fatal("Failed to parse records from file " + path)
    }
}

func CreatePredefinedEncoder(encoderName string) (int, string)  {
	encoder, encoderExists := PredefinedEncoders[encoderName]
	if (!encoderExists) {
		return -1, "Failed to find encoder  " + encoderName
	}

	lastEncoderId++
	activeEncoder := ActiveEncoder{
		VencId: lastEncoderId, 
		ProcessingId: -1,
		Format: encoder.Format,
		Width: encoder.Width,
		Height: encoder.Height,
		Bitrate: encoder.Bitrate,
		Channels: make(map[chan []byte]bool),
		DataChannels: make(map[chan []byte]bool),
	}

	createVencEncoder(activeEncoder)

	ActiveEncoders[lastEncoderId] = activeEncoder

	return lastEncoderId, ""
}

func CreateDummyEncoder() (int, string)  {
	lastEncoderId++
	activeEncoder := ActiveEncoder{
		VencId: lastEncoderId, 
		ProcessingId: -1,
		Format: "",
		Width: 0,
		Height: 0,
		Bitrate: 0,
		Channels: make(map[chan []byte]bool),
		DataChannels: make(map[chan []byte]bool),
	}

	ActiveEncoders[lastEncoderId] = activeEncoder

	return lastEncoderId, ""
}

func DeleteEncoder(encoderId int) (int, string) {
	encoder, encoderExists := ActiveEncoders[encoderId]
	if (!encoderExists) {
		return -1, "Failed to find encoder"
	}

	if (encoder.ProcessingId != -1){
		err, errorString := processing.UnsubscribeEncoderToProcessing(encoder.ProcessingId, encoder)
		if err < 0 {
			return err, errorString
		}
	}

	deleteVencEncoder(encoder)
	delete(ActiveEncoders, encoderId)

	return 0, ""
}

func (encoder ActiveEncoder) GetId() int {
	return encoder.VencId
}

func (encoder ActiveEncoder) DataCallback(data []byte) {
	for ch,_ := range encoder.DataChannels {
		if (cap(ch) <= len(ch)) {
			<-ch
		}

		ch <- data
	}
}
