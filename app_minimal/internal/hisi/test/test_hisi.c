#include <stdio.h>
#include "../hisi_external.h"

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

    return 0;
}
