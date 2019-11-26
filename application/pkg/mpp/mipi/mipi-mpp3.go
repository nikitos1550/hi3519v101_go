//+build hi3516cv300 hi3516av200

package mipi

/*
int hi3516av200_mipi_init(struct hi3516av200_cmos * c) {
    int error_code = 0;
    int fd;

    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
        printf("C DEBUG: TODO\n");
        return ERR_GENERAL;
    }

    combo_dev_attr_t stcomboDevAttr;

    memcpy(&stcomboDevAttr, c->mipidev, sizeof(combo_dev_attr_t));
    stcomboDevAttr.devno = 0;

    printf("stcomboDevAttr memcpy ok\n");

    if(ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno)) {
        printf("C DEBUG: TODO\n");
        close(fd);
        return ERR_GENERAL;
    }

    if(ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno)) {
        printf("C DEBUG: TODO HI_MIPI_RESET_SENSOR failed\n");
        close(fd);
        return ERR_GENERAL;
    }

    if (ioctl(fd, HI_MIPI_SET_DEV_ATTR, &stcomboDevAttr)) {
        printf("set mipi attr failed\n");
        close(fd);
        return ERR_GENERAL;
    }

    if(ioctl(fd, HI_MIPI_UNRESET_MIPI, &stcomboDevAttr.devno)) {
        printf("C DEBUG: TODO HI_MIPI_UNRESET_MIPI failed\n");
        close(fd);
        return ERR_GENERAL;
    }

    if(ioctl(fd, HI_MIPI_UNRESET_SENSOR, &stcomboDevAttr.devno)) {
        printf("C DEBUG: TODO HI_MIPI_UNRESET_SENSOR failed\n");
        close(fd);
        return ERR_GENERAL;
    }

    close(fd);

    return ERR_NONE;
}
*/
import "C"