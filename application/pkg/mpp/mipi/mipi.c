#include "mipi.h"

#if defined(HI_MPP_V1)
    //there is no mipi subsystem in hi3516cv100 family
#elif defined(HI_MPP_V2) \
    || defined(HI_MPP_V3) \
    || defined(HI_MPP_V4)
int mpp_mipi_init(error_in *err, mpp_mipi_init_in *in) {
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

    #if defined(HI_MPP_V3)
        stcomboDevAttr.devno = 0; //TODO

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
    #endif

    #if defined(HI_MPP_V4)
    	lane_divide_mode_t lane_divide_mode = LANE_DIVIDE_MODE_0;

        general_error_code = ioctl(fd, HI_MIPI_SET_HS_MODE, &lane_divide_mode);
        if (general_error_code != HI_SUCCESS) {
        	GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_SET_HS_MODE");
            close(fd);  
            err->general = general_error_code;
            return ERR_GENERAL; 
        }

        combo_dev_t devno = 0;

        general_error_code = ioctl(fd, HI_MIPI_ENABLE_MIPI_CLOCK, &devno);
        if (general_error_code != HI_SUCCESS) {
        	GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_ENABLE_MIPI_CLOCK");
            close(fd);
            err->general = general_error_code;
            return ERR_GENERAL; 
        }

        general_error_code = ioctl(fd, HI_MIPI_RESET_MIPI, &devno);
        if (general_error_code != HI_SUCCESS) {
        	GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_RESET_MIPI");
            close(fd);
            err->general = general_error_code;
            return ERR_GENERAL; 
        }

        sns_clk_source_t SnsDev = 0;

        general_error_code = ioctl(fd, HI_MIPI_ENABLE_SENSOR_CLOCK, &SnsDev);
        if (general_error_code != HI_SUCCESS) {
        	GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_ENABLE_SENSOR_CLOCK");
            close(fd);
            err->general = general_error_code;
            return ERR_GENERAL; 
       	}

        general_error_code = ioctl(fd, HI_MIPI_RESET_SENSOR, &SnsDev);
        if (general_error_code != HI_SUCCESS) {
        	GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_RESET_SENSOR");
            close(fd);
            err->general = general_error_code;
            return ERR_GENERAL; 
		}
	#endif

    general_error_code = ioctl(fd, HI_MIPI_SET_DEV_ATTR, &stcomboDevAttr);
    if (general_error_code != 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_SET_DEV_ATTR"); 
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    #if defined(HI_MPP_V3)
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
	#endif

    #if defined(HI_MPP_V4)
    	general_error_code = ioctl(fd, HI_MIPI_UNRESET_MIPI, &devno);
        if (general_error_code != HI_SUCCESS) {
        	GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_UNRESET_MIPI");
            close(fd);
            err->general = general_error_code;
            return ERR_GENERAL; 
        }

        general_error_code = ioctl(fd, HI_MIPI_UNRESET_SENSOR, &SnsDev);
        if (general_error_code != HI_SUCCESS) {
        	GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_UNRESET_SENSOR");
            close(fd);
            err->general = general_error_code;
            return ERR_GENERAL; 
     	}
	#endif

    close(fd);

    return ERR_NONE;
}
#endif
