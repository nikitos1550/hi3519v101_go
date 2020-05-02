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

    "os/exec"
)

func Init(devInfo DeviceInfo) {
    cmos.Init()

    vi.CheckFlags()

	systemInit(devInfo)
    logger.Log.Debug().
        Msg("OS and chip inited")

    //echo "all=4" > /proc/umap/logmpp

    cmd := exec.Command("sh", "-c", "echo \"all=9\" > /proc/umap/logmpp")
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.Log.Error().
            Msg("Can`t increase logmpp level")
	}
    logger.Log.Debug().
        Msg("logmpp level increased")

    sys.Init(devInfo.Chip)

    if (buildinfo.Family != "hi3516cv100") {
        mipi.Init()
    }

    cmos.Register()

    if (buildinfo.Family == "hi3516cv500" || buildinfo.Family == "hi3516ev200") {
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
