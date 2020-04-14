package venc

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type PredefinedEncoder struct {
    Format string 
    Width int 
    Height int 
    Bitrate int 
}

type ActiveEncoder struct {
    VencId int 
    Format string 
    Width int 
    Height int 
    Bitrate int 
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
		Format: encoder.Format,
		Width: encoder.Width,
		Height: encoder.Height,
		Bitrate: encoder.Bitrate,
	}

	CreateEncoder(activeEncoder)

	ActiveEncoders[lastEncoderId] = activeEncoder

	return lastEncoderId, ""
}
