package mpp

import (
	"application/pkg/mpp/sys"

	"application/pkg/mpp/cmos"
	"application/pkg/mpp/isp"
	"application/pkg/mpp/mipi"
	"application/pkg/mpp/venc"
	"application/pkg/mpp/vi"
	"application/pkg/mpp/vpss"

	//"application/pkg/mpp/ai"

    "application/pkg/buildinfo"
    "application/pkg/logger"

)

func Init(devInfo DeviceInfo) {
    cmos.Init()

	systemInit(devInfo)
    logger.Log.Debug().
        Msg("OS and chip inited")

    sys.Init(devInfo.Chip)

    if (buildinfo.Family != "hi3516cv100") {
        mipi.Init()
    }

    if (buildinfo.Family == "hi3516cv500") {
	    vi.Init()
        isp.Init()
    } else {
        isp.Init()
        vi.Init()
    }

    vpss.Init()

    logger.Log.Debug().
            Msg("VENC init")
	venc.Init()


}
