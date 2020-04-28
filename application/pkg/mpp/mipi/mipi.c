#include "mipi.h"

#if defined(HI3516CV100)
    //there is no mipi subsystem
#endif // defined(HI3516CV100)

#if defined(HI3516CV200) || defined(HI3516AV100)
int mpp_mipi_init(error_in *err, mpp_mipi_init_in *in) {
    int general_error_code = 0;

    int fd;

    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "open /dev/hi_mipi")
         err->general = fd;
        return ERR_GENERAL;
    }

    combo_dev_attr_t stcomboDevAttr;

    memcpy(&stcomboDevAttr, in->mipi, sizeof(combo_dev_attr_t));

    *error_code = ioctl(fd, HI_MIPI_SET_DEV_ATTR, &stcomboDevAttr);
    if (*error_code != 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_SET_DEV_ATTR")        
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    close(fd);

    return ERR_NONE;
}
#endif // defined(HI3516CV200) || defined(HI3516AV100)

#if defined(HI3516CV300) || defined(HI3516AV200)
int mpp_mipi_init(error_in *err, mpp_mipi_init_in * in) {
    int general_error_code = 0;
    
    int fd;

    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "open /dev/hi_mipi");
        err->general = fd;    
        return ERR_GENERAL;   
    }                         

    combo_dev_attr_t stcomboDevAttr;

    memcpy(&stcomboDevAttr, in->mipi, sizeof(combo_dev_attr_t));
    stcomboDevAttr.devno = 0;

    general_error_code = ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno);
    if (general_error_code != 0){
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_RESET_MIPI");
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    general_error_code = ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno); 
    if (general_error_code != 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_RESET_SENSOR");
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    general_error_code = ioctl(fd, HI_MIPI_SET_DEV_ATTR, &stcomboDevAttr);
    if (general_error_code != 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_SET_DEV_ATTR");
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }
    
    #if defined(HI3516CV300)
    usleep(10000);
    #endif

    general_error_code = ioctl(fd, HI_MIPI_UNRESET_MIPI, &stcomboDevAttr.devno);
    if (general_error_code != 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_UNRESET_MIPI");
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    general_error_code = ioctl(fd, HI_MIPI_UNRESET_SENSOR, &stcomboDevAttr.devno); 
    if (general_error_code != 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_UNRESET_SENSOR");
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    close(fd);

    return ERR_NONE;
}
#endif // defined(HI3516CV300) || defined(HI3516AV200)

#if defined(HI3516CV500) || defined(HI3516EV200) || defined(HI3519AV100) || defined(HI3559AV100)
int mpp_mipi_init(int *error_code, mpp_mipi_init_in *in) {
    *error_code = 0;

    int fd;
    
    fd = open( "/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
         GO_LOG_MIPI(LOGGER_ERROR, "open /dev/hi_mipi")     
        *error_code = fd;
        return ERR_GENERAL;
    }

    lane_divide_mode_t lane_divide_mode = LANE_DIVIDE_MODE_0;

    *error_code = ioctl(fd, HI_MIPI_SET_HS_MODE, &lane_divide_mode);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_SET_HS_MODE")
        close(fd);  
        return ERR_GENERAL; 
    }

    combo_dev_t devno = 0;

    *error_code = ioctl(fd, HI_MIPI_ENABLE_MIPI_CLOCK, &devno);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_ENABLE_MIPI_CLOCK")
        close(fd);
        return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_RESET_MIPI, &devno);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_RESET_MIPI")
        close(fd);
        return ERR_GENERAL; 
    }

    sns_clk_source_t SnsDev = 0;

    *error_code = ioctl(fd, HI_MIPI_ENABLE_SENSOR_CLOCK, &SnsDev);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_ENABLE_SENSOR_CLOCK")
        close(fd);
        return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_RESET_SENSOR, &SnsDev);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_RESET_SENSOR")
        close(fd);
        return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_SET_DEV_ATTR, in->mipi);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_SET_DEV_ATTR")
        close(fd);
        return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_UNRESET_MIPI, &devno);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_UNRESET_MIPI")
        close(fd);
        return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_UNRESET_SENSOR, &SnsDev);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_UNRESET_SENSOR")
        close(fd);
        return ERR_GENERAL; 
    }

    close(fd);

    return ERR_NONE;
}
#endif // defined(HI3516CV500) || defined(HI3516EV200) || defined(HI3519AV100) || defined(HI3559AV100)
