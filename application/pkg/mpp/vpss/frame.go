package vpss

//#include "vpss.h"
import "C"

type VpssFrame C.VIDEO_FRAME_INFO_S

type Frame struct {
    frame   *VpssFrame
    pts     uint64
}
