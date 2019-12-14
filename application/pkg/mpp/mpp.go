package mpp

import (
    "application/pkg/mpp/sys"

    //_"application/pkg/mpp/cmos"
    "application/pkg/mpp/mipi"
    "application/pkg/mpp/isp"
    "application/pkg/mpp/vi"
    "application/pkg/mpp/vpss"
    "application/pkg/mpp/venc"

    "application/pkg/mpp/getloop"
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

    getloop.Init()

    //init sample videopipeline
    vpss.SampleChannel0()

    //venc.SampleH264()
    venc.SampleMjpeg()

}
