//+build streamerRtsp

package rtsp

import (
    "bytes"
)

var keyData = []byte{0x00, 0x00, 0x00, 0x01}

func GetNal(encoder string, data []byte) byte {
    if (encoder == "h265"){
        return (data[0]&0x7E)>>1;
    }

    return data[0]&0x1F;
}

func IsSps(encoder string, data []byte) bool {
    nal := GetNal(encoder, data)

    if (encoder == "h265"){
        return nal == 33
    }

    return nal == 7
}

func IsPps(encoder string, data []byte) bool {
    nal := GetNal(encoder, data)

    if (encoder == "h265"){
        return nal == 34
    }

    return nal == 8
}

func ExtractSps(encoder string, data []byte) []byte{
    payloads := bytes.Split(data, keyData)
    for _, payload := range payloads {
        if (len(payload) <= 0){
            continue
        }

        if (IsSps(encoder, payload)) {
            return payload
        }
    }

    return []byte{}
}

func ExtractPps(encoder string, data []byte) []byte{
    payloads := bytes.Split(data, keyData)
    for _, payload := range payloads {
        if (len(payload) <= 0){
            continue
        }

        if (IsPps(encoder, payload)) {
            return payload
        }
    }

    return []byte{}
}
