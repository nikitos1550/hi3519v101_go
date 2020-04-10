package venc

import (
	"encoding/json"
	"io/ioutil"
	"log"
//	"net/http"

//	"application/pkg/openapi"
)

type Encoder struct {
    VencId int
    Format string 
    Width int 
    Height int 
    Bitrate int 
}

type encoderInfo struct {
    Name string
    Format string 
    Width int 
    Height int 
    Bitrate int 
}

var Encoders map[string] Encoder

func init() {
    //openapi.AddApiRoute("listEncoders", "/encoders", "GET", listEncoders)
	Encoders = make(map[string] Encoder)
//	readEncoders() //moved to venc.go Init()
}

func readEncoders() {
	path := "/opt/configs/encoders.json"
    data, err := ioutil.ReadFile(path)
    if err != nil {
		log.Fatal("Failed to read records from file " + path)
		return
    }
    
	err = json.Unmarshal(data, &Encoders)
    if err != nil {
        log.Fatal("Failed to parse records from file " + path)
    }
}

/*
func listEncoders(w http.ResponseWriter, r *http.Request)  {
	var encodersInfo []encoderInfo
	for name, encoder := range Encoders {
		info := encoderInfo{
			Name: name,
			Format: encoder.Format,
			Width: encoder.Width,
			Height: encoder.Height,
			Bitrate: encoder.Bitrate,
		}
	
		encodersInfo = append(encodersInfo, info)
	}
	openapi.ResponseSuccessWithDetails(w, encodersInfo)
}
*/
