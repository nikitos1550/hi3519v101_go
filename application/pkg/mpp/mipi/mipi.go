package mipi

//#include "mipi.h"
import "C"

import (
    //"unsafe"
    //"errors"
    "application/pkg/mpp/cmos"
    "application/pkg/logger"
    "application/pkg/buildinfo"
)

var (
    //mipi unsafe.Pointer
)

func Init() {

    if buildinfo.Family != "hi3516cv100" {


        var inErr C.error_in
        var in C.mpp_mipi_init_in

        switch cmos.S.Data() {
            case cmos.LVDS, cmos.SubLVDS, cmos.HISPI:
                in.data_type = C.uint(C.INPUT_MODE_LVDS)

                if (cmos.S.MipiLVDSAttr() == nil) {
                    logger.Log.Fatal().
                        Msg("MIPI LVDS attrs missed")
                }

                in.mipi_lvds_attr = cmos.S.MipiLVDSAttr()
            //case cmos.SubLVDS:
            //    in.data_type = C.uint(C.INPUT_MODE_SUBLVDS)
            //
            //    if (cmos.S.MipiLVDSAttr() == nil) {
            //        logger.Log.Fatal().
            //            Msg("MIPI LVDS attrs missed")   
            //    }
            //
            //    in.mipi_lvds_attr = cmos.S.MipiLVDSAttr()
            //case cmos.HISPI:
            //    in.data_type = C.uint(C.INPUT_MODE_HISPI)
            //
            //    if (cmos.S.MipiLVDSAttr() == nil) {
            //        logger.Log.Fatal().
            //            Msg("MIPI LVDS attrs missed")   
            //    }
            //
            //    in.mipi_lvds_attr = cmos.S.MipiLVDSAttr()
            case cmos.MIPI:
                in.data_type = C.uint(C.INPUT_MODE_MIPI)

                if (cmos.S.MipiMIPIAttr() == nil) {
                    logger.Log.Fatal().
                        Msg("MIPI MIPI attrs missed")   
                }

                in.mipi_mipi_attr = cmos.S.MipiMIPIAttr()
            case cmos.DC:
                if  buildinfo.Family == "hi3516cv200" ||
                    buildinfo.Family == "hi3516av100" {
                        in.data_type = C.uint(C.INPUT_MODE_CMOS_33V)
                } else if buildinfo.Family == "hi3516cv100" {
                    logger.Log.Fatal().
                        Msg("hi3516cv100 has no mipi!!!")
                } else {
                    in.data_type = C.uint(C.INPUT_MODE_CMOS)
                }
            default:
            logger.Log.Fatal().
                Msg("MIPI unsupported data type")
        }

        err := C.mpp_mipi_init(&inErr, &in)

        if err != C.ERR_NONE {
            logger.Log.Fatal().
                //Str("error", errors.New("MIPI error TODO").Error()).
                Str("error", C.GoString(inErr.name)).
                Int("code", int(inErr.code)).
                Msg("MIPI")
        }

    }

    logger.Log.Debug().
        Msg("MIPI inited")
}

//export go_logger_mipi
func go_logger_mipi(level C.int, msgC *C.char) {
        logger.CLogger("MIPI", int(level), C.GoString(msgC))
}
