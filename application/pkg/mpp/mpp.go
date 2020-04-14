package mpp

import (
	"application/pkg/mpp/sys"

	//"application/pkg/mpp/cmos"
	"application/pkg/mpp/isp"
	"application/pkg/mpp/mipi"
	"application/pkg/mpp/venc"
	"application/pkg/mpp/vi"
	"application/pkg/mpp/vpss"

	"application/pkg/mpp/ai"

    "application/pkg/buildinfo"
    "application/pkg/logger"

    //"time"
)

func Init(devInfo DeviceInfo) {

	systemInit(devInfo)
    //time.Sleep(1 * time.Second)
	sys.Init()
    //time.Sleep(1 * time.Second)
	mipi.Init()
    //time.Sleep(1 * time.Second)
    if (buildinfo.Family == "hi3516cv500") {
        logger.Log.Debug().
            Msg("hi3516cv500 family vi init before isp!")
	    vi.Init()
        isp.Init()
    } else {
        isp.Init()
        //time.Sleep(1 * time.Second)
        vi.Init()
    }
    //time.Sleep(1 * time.Second)
    vpss.Init()
    //time.Sleep(1 * time.Second)
	ai.Init()
    //time.Sleep(1 * time.Second)
	venc.Init()

	//init sample videopipeline
	vpss.SampleChannel0()
}
