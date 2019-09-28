#include "../hisi_external.h"
#include "../hisi_utils.h"

#include "hi3516av200_mpp.h"

#include <stddef.h>
#include <stdio.h>

int hisi_get_chipid(unsigned int * chip) {
    HI_MPI_SYS_GetChipId(chip);
    return ERR_NONE;
}

int hisi_get_temperature(float * t) {
    uint32_t read;
    devmem(0x120A0118, -1, &read);
    printf("C DEBUG: temperature 0x%lx code, %.1f C\n", read & 0x3FF, (float)((( (read & 0x3FF)-125) / 806.0 ) * 165) - 40 );
    *t = (float)(((( (read & 0x3FF) - 125) / 806.0 ) *165) - 40);

    return ERR_NONE;
}
