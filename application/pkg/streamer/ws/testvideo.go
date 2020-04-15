//+build streamerWs

package ws

import (
    "log"
    "os"
)

var (
    buffer          []byte
    //frameOffset     []int
    //frameSize       []int
    //frames          int
)

func loadTestVideo() {
    file, err := os.Open("/opt/testvideo.mp4")
    if err != nil {
      panic(err)
    }

    fileinfo, err := file.Stat()
    if err != nil {
      panic(err)
    }

    filesize := fileinfo.Size()
    buffer = make([]byte, filesize)

    bytesread, err := file.Read(buffer)
    if err != nil {
        panic(err)
    }

    log.Println("Test video bytes read: ", bytesread)
}


