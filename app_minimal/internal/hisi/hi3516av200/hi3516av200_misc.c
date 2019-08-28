#include "../hisi_external.h"

int hisi_get_cmos_info(unsigned int * width, unsigned int * height, unsigned int * fps) {
    return ERR_NOT_IMPLEMENTED;
}

/*
int inittemperature() {
        devmem(0x120A0110, 0x60FA0000, NULL);
        return 0;
}
*/

int hisi_get_temperature(float * t) {
    /*
        uint32_t read;
        devmem(0x120A0118, -1, &read);
        printf("C DEBUG: temperature code 0x%lx\n", read & 0x3FF);
        printf("C DEBUG: temperature C %.1f\n", (float)((( (read & 0x3FF)-125) / 806.0 ) * 165) - 40 );
        return (float)(((( (read & 0x3FF) - 125) / 806.0 ) *165) - 40);
    */
    return ERR_NOT_IMPLEMENTED;
}
