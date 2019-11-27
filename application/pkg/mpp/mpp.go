package mpp

import (
    "application/pkg/mpp/sys"

    _"application/pkg/mpp/cmos"
    "application/pkg/mpp/mipi"
    "application/pkg/mpp/isp"
    "application/pkg/mpp/vi"
    "application/pkg/mpp/vpss"
    _"application/pkg/mpp/venc"

)

func Init() {
    systemInit()
    //
    //
    sys.Init()
    mipi.Init()
    isp.Init()
    vi.Init()
    vpss.Init()
}
