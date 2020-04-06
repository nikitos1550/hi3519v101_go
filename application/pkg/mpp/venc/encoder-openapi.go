//+build openapi

package venc

import (
        "encoding/json"
        "io/ioutil"
        "log"
        "net/http"

        "application/pkg/openapi"
)

func init() {
    openapi.AddApiRoute("listEncoders", "/encoders", "GET", listEncoders)
        Encoders = make(map[string] Encoder)
        readEncoders()
}

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

