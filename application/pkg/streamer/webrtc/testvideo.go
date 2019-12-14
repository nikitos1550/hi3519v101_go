package webrtc

import (
    "log"
    "os"
)

var (
    buffer          []byte
    frameOffset     []int
    frameSize       []int
    frames          int
)

func loadTestVideo() {
    file, err := os.Open("/opt/testvideo.h264")
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


func parseTestVideo() {
    var i int
    var counter int = 0
    for i=0;i<(len(buffer)-4);i++ {
        //log.Println(buffer[i], " ", buffer[i+1], " ", buffer[i+2], " ", buffer[i+3])
        found := 0
        //if  (buffer[i]  == 0 && buffer[i+1] == 0 && buffer[i+2] == 1) {
        //    found = 3
        //} else 
        if  (buffer[i]  == 0 && buffer[i+1] == 0 && buffer[i+2] == 0 && buffer[i+3] == 1) {
            found = 4
        }

        if (found > 0) {
            nalType := buffer[i+found] & 0x1F
            //log.Println("Found ", i, " NAL ", nalType)

            counter++
            frameOffset = append(frameOffset, i)
            frameSize   = append(frameSize, 0)
            if (len(frameSize)>1) {
                frameSize[counter-2] = frameOffset[counter-1]-frameOffset[counter-2]
            }
            i=i+found
        }
    }
    frameSize[counter-1] = len(buffer)-frameOffset[counter-1] //last frame szie

    frames=counter
    //log.Println("Total found ", frames)
    //log.Println("offsets: ",frameOffset)
    //log.Println("sizes: ", frameSize)
    //frames[i]
}

func getFrameTestVideo(i int) []byte {
    if (i==1) {
        tmp := make([]byte, frameSize[0]+frameSize[1])
        //log.Println("tmp size ", frameSize[0]+frameSize[1]-2)
        copy(tmp[0:frameSize[0]], buffer[frameOffset[0]:frameOffset[0]+frameSize[0]])
        copy(tmp[frameSize[0]:frameSize[0]+frameSize[1]], buffer[frameOffset[1]:frameOffset[1]+frameSize[1]])
        //log.Println(tmp)
        return tmp
    }

    //return buffer[frameOffset[i]+1:frameOffset[i]+frameSize[i]]
    return buffer[frameOffset[i]:frameOffset[i]+frameSize[i]]
}

