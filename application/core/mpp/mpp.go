package mpp

import (
    "time"
    "os/exec"

	"application/core/mpp/sys"
	"application/core/mpp/cmos"
	"application/core/mpp/isp"
	"application/core/mpp/mipi"
	"application/core/mpp/venc"
	"application/core/mpp/vi"
	"application/core/mpp/vpss"
    "application/core/mpp/utils"
    "application/core/utils/chip"
    "application/core/compiletime"
    "application/core/logger"
    //"application/pkg/mpp/vo"
    //"application/pkg/mpp/ai"
)

func Init(devInfo DeviceInfo) {
    cmos.Init()

    vi.CheckFlags()

    //TODO perform system cleanup as in hi3516av200 for all families
    closePrev()

	systemInit(devInfo)
    logger.Log.Debug().
        Msg("OS and chip inited")

    logger.Log.Trace().
        Str("chip", chip.Detect(utils.MppId())).
        Msg("MPP")

    //echo "all=4" > /proc/umap/logmpp


    if false {
        cmd := exec.Command("sh", "-c", "echo \"all=9\" > /proc/umap/logmpp")
	    _, err := cmd.CombinedOutput()
        if err != nil {
		    logger.Log.Error().
                Msg("Can`t increase logmpp level")
	    }
        logger.Log.Debug().
            Msg("logmpp level increased")
    }


    sys.Init(devInfo.Chip)

    if (compiletime.Family != "hi3516cv100") {
        mipi.Init()
    }

    cmos.Register()

    if (compiletime.Family == "hi3516cv500" || compiletime.Family == "hi3516ev200") {
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

    //ai.Init()

    //Set initial PTS

    utils.InitPTS( uint64(time.Now().UnixNano() / 1000) )

    //update pts each minute
    ticker := time.NewTicker(1 * time.Minute)
    quit := make(chan struct{})
    go func() {
        for {
        select {
            case <- ticker.C:
                utils.SyncPTS( uint64(time.Now().UnixNano() / 1000) )
            case <- quit:
                ticker.Stop()
                return
            }
        }
    }()


    //vo.Init() //FOR TEST, onlu hi3516cv500
}
