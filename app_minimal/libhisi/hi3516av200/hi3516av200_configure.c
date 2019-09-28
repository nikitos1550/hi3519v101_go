#include "../hisi_external.h"
#include "hi3516av200_mpp.h"

int hi3516av200_vpps_configure();
int hi3516av200_venc_configure();

int hisi_configure() {

    hi3516av200_vpps_configure();
    hi3516av200_venc_configure();

    return ERR_NONE;
}

int hi3516av200_vpps_configure() {

}

int hi3516av200_venc_configure() {

}
