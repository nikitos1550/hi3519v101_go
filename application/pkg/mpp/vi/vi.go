package vi

//#include "vi.h"
import "C"

import (
    "flag"
    "application/pkg/buildinfo"
    "application/pkg/logger"
    "application/pkg/mpp/cmos"
    "application/pkg/mpp/errmpp"
)

var (
    flipX bool
    flipY bool
    x0 int
    y0 int
    width int
    height int
    fps int

    ldc bool
    ldcOffsetX int
    ldcOffsetY int
    ldcK  int
)

func Width() int {
    return width
}
func Height() int {
    return height
}
func Fps() int {
    return fps
}

func init() {
    flag.BoolVar(&flipY, "vi-flip-y", false, "Flip image relative to y axis")
    flag.BoolVar(&flipX, "vi-flip-x", false, "flip image relative to x axis")

    flag.IntVar(&x0, "vi-x0", 0, "top left x point to capture from")
    flag.IntVar(&y0, "vi-y0", 0, "top left y point to capture from")
    flag.IntVar(&width, "vi-width", -1, "width of capture image")
    flag.IntVar(&height, "vi-height", -1, "height of capture image")
    flag.IntVar(&fps, "vi-fps", -1, "base framerate, should be less or equal cmos")

    if buildinfo.Family == "hi3516av100" {
        /*
        When the resolution of the captured VI picture is not greater than D1, the value range of s32Ratio is [0, 480].
        When the resolution of the captured VI picture is greater than D1 but not greater than 720p, the value range of s32Ratio is [0, 433].
        When the resolution of the captured VI picture is greater than 720p but not greater than 1080p, the value range of s32Ratio is [0, 400].
        When the resolution of the captured VI picture is greater than 1080p but not greater than 2304 x 1536, the value range of s32Ratio is [0, 300].
        When the resolution of the captured VI picture is greater than 2304 x 1536 but not greater than 5 megapixels, the value range of s32Ratio is [0, 168].
        */
        flag.BoolVar(&ldc, "vi-ldc", false, "LDC enable")
        flag.IntVar(&ldcOffsetX, "vi-ldc-offset-x", 0, "LDC x offset from center [-75;75]")
        flag.IntVar(&ldcOffsetY, "vi-ldc-offset-y", 0, "LDC y offset from center [-75;75]")
        flag.IntVar(&ldcK, "vi-ldc-k", 0, "LDC coefficient [0;168]")
    }

    if buildinfo.Family == "hi3516av200" {
        flag.BoolVar(&ldc, "vi-ldc", false, "LDC enable")
        flag.IntVar(&ldcOffsetX, "vi-ldc-offset-x", 0, "LDC x offset from center [-127;127]")
        flag.IntVar(&ldcOffsetY, "vi-ldc-offset-y", 0, "LDC y offset from center [-127;127]")
        flag.IntVar(&ldcK, "vi-ldc-k", 0, "LDC coefficient [-300;500]")
    }
}

func Params() {
    if width == -1 {
        width = cmos.S.Width()
    }
    if height == -1 {
        height = cmos.S.Height()
    }
    if fps == -1 {
        fps = cmos.S.Fps()
    }

    if x0<0 || x0 > cmos.S.Width() {
        logger.Log.Fatal().
            Int("vi-x0", x0).
            Msg("vi-x0 should be positive")
    }
    if y0<0 || y0 > cmos.S.Height() {
        logger.Log.Fatal().
            Int("vi-y0", y0).
            Msg("vi-y0 should be positive")
    }
    if width < x0 || width > cmos.S.Width() {
        logger.Log.Fatal().
            Int("vi-width", width).
            Int("vi-x0", x0).
            Int("cmos-width", int(cmos.S.Width())).
            Msg("vi-width should be greater than x0 and less or equal than cmos width")
    }
    if height < y0 || height > cmos.S.Height() {
        logger.Log.Fatal().
            Int("vi-height", height).
            Int("vi-y0", x0).
            Int("cmos-height", int(cmos.S.Height())).
            Msg("vi-height should be greater than y0 and less or equal than cmos height")
    }
    if (width - x0) % 2 != 0 {
        logger.Log.Fatal().
            Int("vi-captured-width", (width - x0)).
            Msg("captured width (vi-width - vi-x0) should be aligned by 2 pixels")
    }
    if (height - y0) % 2 != 0 {
        logger.Log.Fatal().
            Int("vi-captured-height", (height - x0)).
            Msg("captured height (vi-height - vi-y0) should be aligned by 2 pixels")
    }
    if (width - x0) < C.VPSS_MIN_IMAGE_WIDTH {
        logger.Log.Fatal().
            Int("vi-captured-width", (width - x0)).
            Int("vi-minimal-width", int(C.VPSS_MIN_IMAGE_WIDTH)).
            Msg("captured width (vi-width - vi-x0) should be greater than minimal width")
    }
    if (height - y0) < C.VPSS_MIN_IMAGE_HEIGHT {
        logger.Log.Fatal().
            Int("vi-captured-width", (height - x0)).
            Int("vi-minimal-width", int(C.VPSS_MIN_IMAGE_HEIGHT)).
            Msg("captured height (vi-height - vi-y0) should be greater than minimal height")
    }
    if fps < 0 || fps == 0 || fps > cmos.S.Fps() {
        logger.Log.Fatal().
            Int("vi-fps", fps).
            Int("cmos-fps", cmos.S.Fps()).
            Msg("vi-fps should be greater than 0 and less or equal cmos fps")
    }

}

func Init() {
    var inErr C.error_in
    var in C.mpp_vi_init_in

    //TODO move LDS to Params
    if buildinfo.Family == "hi3516av100" {
        if ldc == true {
            if ldcOffsetX < -75 || ldcOffsetX > 75 {
                logger.Log.Fatal().
                    Int("ldc-offset-x", ldcOffsetX).
                    Msg("vi-ldc-offset-x should be [-75;75]")
            }
            if ldcOffsetY < -75 || ldcOffsetY > 75 {
                logger.Log.Fatal().
                    Int("ldc-offset-y", ldcOffsetY).
                    Msg("vi-ldc-offset-y should be [-75;75]")
            }
            if ldcK < 0 || ldcK > 168 {
                logger.Log.Fatal().
                    Int("ldc-k", ldcK).
                    Msg("vi-ldc-k should be [0;168]")
            }

            in.ldc = 1
            in.ldc_offset_x = C.int(ldcOffsetX)
            in.ldc_offset_y = C.int(ldcOffsetY)
            in.ldc_k = C.int(ldcK)
        }
    }

    if buildinfo.Family == "hi3516av200" {
        if ldc == true {
            if ldcOffsetX < -127 || ldcOffsetX > 127 {
                logger.Log.Fatal().
                    Int("ldc-offset-x", ldcOffsetX).
                    Msg("vi-ldc-offset-x should be [-127;127]")
            }
            if ldcOffsetY < -127 || ldcOffsetY > 127 {
                logger.Log.Fatal().
                    Int("ldc-offset-y", ldcOffsetY).
                    Msg("vi-ldc-offset-y should be [-127;127]")
            }
            if ldcK < -300 || ldcK > 500 {
                logger.Log.Fatal().
                    Int("ldc-k", ldcK).
                    Msg("vi-ldc-k should be [-300;500]")
            }

            in.ldc = 1
            in.ldc_offset_x = C.int(ldcOffsetX)
            in.ldc_offset_y = C.int(ldcOffsetY)
            in.ldc_k = C.int(ldcK)
        }
    }

    if flipY == true {
        in.mirror = 1
    }
    if flipX == true {
        in.flip = 1
    }

    in.videv = cmos.S.ViDev()
    in.cmos_width = C.uint(cmos.S.Width())
    in.cmos_height = C.uint(cmos.S.Height())
    in.x0 = C.uint(x0)
    in.y0 = C.uint(y0)
    in.width = C.uint(width)
    in.height = C.uint(height)
    in.cmos_fps = C.uint(cmos.S.Fps())
    in.fps = C.uint(fps)

    logger.Log.Trace().
        Uint("mirror", uint(in.mirror)).
        Uint("flip", uint(in.flip)).
        Uint("cmos_width", uint(in.cmos_width)).
        Uint("cmos_height", uint(in.cmos_height)).
        Uint("x0", uint(in.x0)).
        Uint("y0", uint(in.y0)).
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("cmos_fps", uint(in.cmos_fps)).
        Uint("fps", uint(in.fps)).
        Uint("ldc", uint(in.ldc)).
        Int("ldc-offset-x", int(in.ldc_offset_x)).
        Int("ldc-offset-y", int(in.ldc_offset_y)).
        Int("ldc-k", int(in.ldc_k)).
        Msg("VI params")

    err := C.mpp_vi_init(&inErr, &in)

    if err != C.ERR_NONE {
        logger.Log.Fatal().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VI")
    }
    logger.Log.Debug().
        Msg("VI inited")
}

//export go_logger_vi
func go_logger_vi(level C.int, msgC *C.char) {
        logger.CLogger("VI", int(level), C.GoString(msgC))
}
