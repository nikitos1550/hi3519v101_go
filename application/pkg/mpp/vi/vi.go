package vi

//#include "../include/mpp.h"
import "C"

import (
    "flag"
    "application/pkg/logger"
    "application/pkg/mpp/cmos"
)

var (
    flipX bool
    flipY bool
    x0 int
    y0 int
    width int
    height int
    fps int
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
}

func Init() {
    /*
    logger.Log.Debug().
        Uint("x0", x0).
        Uint("y0", y0).
        Int("width", width).
        Int("height", height).
        Int("fps", fps).
        Msg("VI cmd params")
    */

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

    //mirror =  true

    err := initFamily()
    if err != nil {
        logger.Log.Fatal().
            Str("error", err.Error()).
            Msg("VI")
    }
    logger.Log.Debug().
        Msg("VI inited")
}
