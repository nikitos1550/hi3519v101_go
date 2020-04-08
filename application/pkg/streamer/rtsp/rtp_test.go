//////+ b u i l d  streamerRtsp

package rtsp

import "testing"
import "log"

func TestSum(t *testing.T) {
	var testp packetizer
	var nal [100]byte

	log.Println(testp.NalH264ToRtp(nal))
	t.Errorf("Test error")
}
