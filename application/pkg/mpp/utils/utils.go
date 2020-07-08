//+build arm arm64
//+build hi3516av100 hi3516av200 hi3516cv100 hi3516cv200 hi3516cv300 hi3516cv500 hi3516ev200 hi3519av100 hi3559av100

package utils

// #include "utils.h"
import "C"

import (
    "errors"
    "application/pkg/logger"
)

var (
    lastTS uint64
)

func Version() string {
	var ver C.MPP_VERSION_S
	C.HI_MPI_SYS_GetVersion(&ver)
	mppVersion := C.GoString(&ver.aVersion[0])

	return mppVersion
}

func MppId() uint32 {
	var id C.HI_U32
	C.HI_MPI_SYS_GetChipId(&id)

	//log.Println("ChipID=", id)

	return uint32(id)
}

//After the current system PTS (unit:μs) is fine-tuned, the PTS does not roll back. When
//multiple chips are synchronized, the difference between the clock sources of the boards may
//be significant. Therefore, you are recommended to tune the PTS once a second.
func SyncPTS(pts uint64) error {
    //err := C.HI_MPI_SYS_SyncPts(C.HI_U64(pts))
    err := C.SyncPts(C.HI_U64(pts))

    if err != 0 {
        logger.Log.Warn().
            Uint64("pts", pts).
            Uint64("delta", pts-lastTS).
           Msg("SyncPTS")
        return errors.New("Some SyncPTS error")
    } else {
        logger.Log.Debug().
            Uint64("pts", pts).
            Uint64("delta", pts-lastTS).
           Msg("SyncPTS")
    }
    lastTS = pts
    return nil
}

//Regardless of the original system PTS, initializing the PTS base forces the current system
//PTS to u64PtsBase. Therefore, you are recommended to call this MPI before a media service
//is enabled. For example, you can call this MPI immediately when the OS starts. If a media
//service is enabled, you can call HI_MPI_SYS_SyncPts to tune the PTS.
func InitPTS(pts uint64) error {
    lastTS = pts

    //err := C.HI_MPI_SYS_InitPtsBase(C.HI_U64(pts))
    err := C.InitPtsBase(C.HI_U64(pts))

    if err != 0 {
        logger.Log.Warn().
            Uint64("pts", pts).
            Msg("InitPTS")
        return errors.New("Some InitPTS error")
    } else {
        logger.Log.Debug().
            Uint64("pts", pts).
            Msg("InitPTS")
    }
    return nil
}

func BindVpssVenc(vpssId int, vencId int) error {
    var errIn C.error_in
    var in C.mpp_bind_vpss_venc_in

    in.vpss_id = C.int(vpssId)
    in.venc_id = C.int(vencId)

    err := C.mpp_bind_vpss_venc(&errIn, &in)
    if err != C.ERR_NONE {
        return errors.New("BindVpssVenc error TODO")
    }

    return nil
}

func UnBindVpssVenc(vpssId int, vencId int) error {
    var errIn C.error_in
    var in C.mpp_bind_vpss_venc_in

    in.vpss_id = C.int(vpssId)
    in.venc_id = C.int(vencId)

    err := C.mpp_unbind_vpss_venc(&errIn, &in)
    if err != C.ERR_NONE {
        return errors.New("BindVpssVenc error TODO")
    }

    return nil
}

