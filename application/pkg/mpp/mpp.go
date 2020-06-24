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

    "application/pkg/mpp/utils"
    "application/pkg/utils/chip"

    "application/pkg/buildinfo"
    "application/pkg/logger"

    "os/exec"

    //"application/pkg/mpp/vo"

    "application/pkg/mpp/ai"

    "time"
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

    ai.Init()

    //Set initial PTS
    utils.InitPTS( uint64(time.Now().UnixNano() / 1000) )

   
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
