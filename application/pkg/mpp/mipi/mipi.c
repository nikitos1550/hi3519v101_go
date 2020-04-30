#include "mipi.h"

#if HI_MPP == 1
    //there is no mipi subsystem in hi3516cv100 family
#elif HI_MPP >= 2
int mpp_mipi_init(error_in *err, mpp_mipi_init_in *in) {
    int general_error_code = 0;

    int fd;

    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
        RETURN_ERR_GENERAL(err, "open /dev/hi_mipi", fd);
    }

    combo_dev_attr_t stcomboDevAttr;
    memcpy(&stcomboDevAttr, in->mipi, sizeof(combo_dev_attr_t));

    #if HI_MPP == 3
        stcomboDevAttr.devno = 0; //TODO

        general_error_code = ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno);
        if (general_error_code != 0){
            close(fd);
            RETURN_ERR_GENERAL(err, "HI_MIPI_RESET_MIPI", general_error_code);
        }

        general_error_code = ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno); 
        if (general_error_code != 0) {
            close(fd);
            RETURN_ERR_GENERAL(err, "HI_MIPI_RESET_SENSOR", general_error_code);
        }
    #endif

    #if HI_MPP == 4
    	lane_divide_mode_t lane_divide_mode = LANE_DIVIDE_MODE_0;

        general_error_code = ioctl(fd, HI_MIPI_SET_HS_MODE, &lane_divide_mode);
        if (general_error_code != HI_SUCCESS) {
            close(fd);  
            RETURN_ERR_GENERAL(err, "HI_MIPI_SET_HS_MODE", general_error_code);
        }

        combo_dev_t devno = 0;

        general_error_code = ioctl(fd, HI_MIPI_ENABLE_MIPI_CLOCK, &devno);
        if (general_error_code != HI_SUCCESS) {
            close(fd);
            RETURN_ERR_GENERAL(err, "HI_MIPI_ENABLE_MIPI_CLOCK", general_error_code);
        }

        general_error_code = ioctl(fd, HI_MIPI_RESET_MIPI, &devno);
        if (general_error_code != HI_SUCCESS) {
            close(fd);
            RETURN_ERR_GENERAL(err, "HI_MIPI_RESET_MIPI", general_error_code); 
        }

        sns_clk_source_t SnsDev = 0;

        general_error_code = ioctl(fd, HI_MIPI_ENABLE_SENSOR_CLOCK, &SnsDev);
        if (general_error_code != HI_SUCCESS) {
            close(fd);
            RETURN_ERR_GENERAL(err, "HI_MIPI_ENABLE_SENSOR_CLOCK", general_error_code); 
       	}

        general_error_code = ioctl(fd, HI_MIPI_RESET_SENSOR, &SnsDev);
        if (general_error_code != HI_SUCCESS) {
            close(fd);
            RETURN_ERR_GENERAL(err, "HI_MIPI_RESET_SENSOR", general_error_code); 
		}
	#endif

    general_error_code = ioctl(fd, HI_MIPI_SET_DEV_ATTR, &stcomboDevAttr);
    if (general_error_code != 0) {
        close(fd);
        RETURN_ERR_GENERAL(err, "HI_MIPI_SET_DEV_ATTR", general_error_code); 
    }

    #if HI_MPP == 3
   		#if defined(HI3516CV300)
       		usleep(10000);
        #endif

        general_error_code = ioctl(fd, HI_MIPI_UNRESET_MIPI, &stcomboDevAttr.devno);
        if (general_error_code != 0) {
            close(fd);
            RETURN_ERR_GENERAL(err, "HI_MIPI_UNRESET_MIPI", general_error_code); 
        }

        general_error_code = ioctl(fd, HI_MIPI_UNRESET_SENSOR, &stcomboDevAttr.devno); 
        if (general_error_code != 0) {
            close(fd);
            RETURN_ERR_GENERAL(err, "HI_MIPI_UNRESET_SENSOR", general_error_code); 
        }
	#endif

    #if HI_MPP == 4
    	general_error_code = ioctl(fd, HI_MIPI_UNRESET_MIPI, &devno);
        if (general_error_code != HI_SUCCESS) {
            close(fd);
            RETURN_ERR_GENERAL(err, "HI_MIPI_UNRESET_MIPI", general_error_code); 
        }

        general_error_code = ioctl(fd, HI_MIPI_UNRESET_SENSOR, &SnsDev);
        if (general_error_code != HI_SUCCESS) {
            close(fd);
            RETURN_ERR_GENERAL(err, "HI_MIPI_UNRESET_SENSOR", general_error_code); 
     	}
	#endif

    close(fd);

    return ERR_NONE;
}
#endif
