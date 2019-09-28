#include <stdio.h>
#include "hisi_external.h"

int main(void) {
    /*
    struct cmos_params {
    unsigned int fps;
    unsigned int x0;
    unsigned int y0;
    unsigned int width;
    unsigned int height;
    //unsigned int wdr;
    unsigned int undistortion;
    unsigned int undistortion_x;
    unsigned int undistortion_y;
    unsigned int undistortion_ratio;
    unsigned int undistortion_minratio;
    };
    */
    int error_code = 0;
    struct cmos_params cp;
    error_code = hisi_init(&cp);

    struct channel_params cparams;
    cparams.width=3840;
    cparams.height=2160;
    cparams.fps=30;

    error_code = hisi_channel_enable(0, &cparams);
    printf("hisi_channel_enable(0, &cparams) returns %d\n", error_code);

    struct encoder_static_params esparams;
    esparams.codec = CODEC_MJPEG;
    esparams.rc = RC_CBR;
    esparams.profile = PROFILE_MJPEG_BASELINE;
    esparams.width = 3840;
    esparams.height = 2160;
    esparams.fps = 30;
    esparams.channel_id = 0;

    struct encoder_dynamic_params edparams;
    edparams.cbr.bitrate = 1024;
    edparams.cbr.gop = 30;
    edparams.cbr.stattime = 1;
    edparams.cbr.fluctuate = 1;

    error_code = hisi_encoder_create(0, &esparams, &edparams);
    printf("hisi_encoder_create(0, &esparams, &edparams) returns %d\n", error_code);

    error_code = hisi_encoder_create(0, &esparams, &edparams);
    printf("hisi_encoder_create(0, &esparams, &edparams) returns %d\n", error_code);

    return 0;
}
