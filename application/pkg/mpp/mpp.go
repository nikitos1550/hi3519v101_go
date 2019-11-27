package mpp

import (
    "application/pkg/koloader"

    "application/pkg/mpp/sys"

    _"application/pkg/mpp/cmos"
    _"application/pkg/mpp/mipi"
    _"application/pkg/mpp/isp"
    _"application/pkg/mpp/vi"
    _"application/pkg/mpp/vpss"
    _"application/pkg/mpp/venc"

)

func Init() {
    koloader.LoadAll()
    //
    //
    sys.Init()
    
}
