#include "mipi.h"

#if HI_MPP > 1
static void mpp_mipi_set_attrs(combo_dev_attr_t *stcomboDevAttr, mpp_mipi_init_in *in) {
    memset(stcomboDevAttr, 0, sizeof(combo_dev_attr_t));

    #if HI_MPP == 1
        //mpp1 doesn`t have mipi subsystem
    #elif HI_MPP == 2
    //mpp2
    //typedef struct {
    //    input_mode_t          input_mode;               // input mode: MIPI/LVDS/SUBLVDS/HISPI/DC 
    //    union
    //    {
    //        mipi_dev_attr_t     mipi_attr;              // for MIPI configuration 
    //        lvds_dev_attr_t     lvds_attr;              // for LVDS/SUBLVDS/HISPI configuration 
    //    };
    //} combo_dev_attr_t;
    //
    //typedef struct {
    //    raw_data_type_e       raw_data_type;            // raw data type: 8/10/12/14 bit
    //    short                 lane_id[MIPI_LANE_NUM];   // lane_id: -1 - disable 
    //} mipi_dev_attr_t;
    //
    //typedef struct {
    //    img_size_t          img_size;                   // oringnal sensor input image size 
    //    wdr_mode_e          wdr_mode;                   // WDR mode 
    //    lvds_sync_mode_e    sync_mode;                  // sync mode: SOL, SAV 
    //    raw_data_type_e     raw_data_type;              // raw data type: 8/10/12/14 bit 
    //    lvds_bit_endian     data_endian;                // data endian: little/big 
    //    lvds_bit_endian     sync_code_endian;           // sync code endian: little/big 
    //    short               lane_id[LVDS_LANE_NUM];     // lane_id: -1 - disable
    //
    //    // each vc has 4 params, sync_code[i]:
    //    //sync_mode is SYNC_MODE_SOL: SOF, EOF, SOL, EOL
    //    //sync_mode is SYNC_MODE_SAV: invalid sav, invalid eav, valid sav, valid eav
    //    unsigned short      sync_code[LVDS_LANE_NUM][WDR_VC_NUM][SYNC_CODE_NUM];
    //} lvds_dev_attr_t;
        stcomboDevAttr->input_mode                  = in->data_type;

        if (in->data_type == INPUT_MODE_MIPI) {
            combo_dev_attr_t *tmp = (combo_dev_attr_t*)in->mipi_mipi_attr;
            //memcpy(&stcomboDevAttr->mipi_attr, &tmp->mipi_attr, sizeof(mipi_dev_attr_t));
            memcpy(stcomboDevAttr, tmp, sizeof(combo_dev_attr_t));
        }
        else if (in->data_type == INPUT_MODE_LVDS || 
            in->data_type == INPUT_MODE_SUBLVDS ||
            in->data_type == INPUT_MODE_HISPI) {
            combo_dev_attr_t *tmp = (combo_dev_attr_t*)in->mipi_lvds_attr;
            //memcpy(&stcomboDevAttr->mipi_attr, &tmp->lvds_attr, sizeof(lvds_dev_attr_t));
            memcpy(stcomboDevAttr, tmp, sizeof(combo_dev_attr_t));
        }


    #elif HI_MPP == 3
    //mpp3
    //typedef struct {
    //    COMBO_DEV               devno;                  // device number, select sensor0 and sensor 1
    //    input_mode_t            input_mode;             // input mode: MIPI/LVDS/SUBLVDS/HISPI/DC 
    //
    //    union
    //    {
    //        mipi_dev_attr_t     mipi_attr;
    //        lvds_dev_attr_t     lvds_attr;
    //    };
    //} combo_dev_attr_t;
    //
    //typedef struct {
    //    raw_data_type_e       raw_data_type;            // raw data type: 8/10/12/14 bit
    //    mipi_wdr_mode_e       wdr_mode;                 // MIPI WDR mode
    //    short                 lane_id[MIPI_LANE_NUM];   // lane_id: -1 - disable
    //
    //    union
    //    {
    //        short data_type[WDR_VC_NUM];                // used by the HI_MIPI_WDR_MODE_DT
    //    };
    //} mipi_dev_attr_t;
    //
    //typedef struct {
    //    img_size_t          img_size;                   // oringnal sensor input image size
    //    raw_data_type_e     raw_data_type;              // raw data type: 8/10/12/14 bit
    //    wdr_mode_e          wdr_mode;                   // WDR mode
    //
    //    lvds_sync_mode_e    sync_mode;                  // sync mode: SOF, SAV 
    //    lvds_vsync_type_t   vsync_type;                 // normal, share, hconnect 
    //    lvds_fid_type_t     fid_type;                   // frame identification code 
    //    
    //    lvds_bit_endian     data_endian;                // data endian: little/big 
    //    lvds_bit_endian     sync_code_endian;           // sync code endian: little/big 
    //    short               lane_id[LVDS_LANE_NUM];     // lane_id: -1 - disable
    //
    //    // each vc has 4 params, sync_code[i]:
    //    //sync_mode is SYNC_MODE_SOF: SOL, EOL, SOF, EOF
    //    //sync_mode is SYNC_MODE_SAV: valid sav, valid eav, invalid sav, invalid eav
    //    unsigned short      sync_code[LVDS_LANE_NUM][WDR_VC_NUM][SYNC_CODE_NUM];
    //} lvds_dev_attr_t;
        stcomboDevAttr->devno                       = 0;
        stcomboDevAttr->input_mode                  = in->data_type;

        if (in->data_type == INPUT_MODE_MIPI) {
			combo_dev_attr_t *tmp = (combo_dev_attr_t*)in->mipi_mipi_attr;
            //memcpy(&stcomboDevAttr->mipi_attr, &tmp->mipi_attr, sizeof(mipi_dev_attr_t));
            memcpy(stcomboDevAttr, tmp, sizeof(combo_dev_attr_t));

            /*
            stcomboDevAttr.mipi_attr.raw_data_type      = mpp_get_raw_data_type(in->pixel_bitness);
            stcomboDevAttr.mipi_attr.wdr_mode
            stcomboDevAttr.mipi_attr.lane_id
            stcomboDevAttr.mipi_attr.lane_id
            */
        }
        else if (in->data_type == INPUT_MODE_LVDS || 
            in->data_type == INPUT_MODE_SUBLVDS ||
            in->data_type == INPUT_MODE_HISPI) {
            combo_dev_attr_t *tmp = (combo_dev_attr_t*)in->mipi_lvds_attr;
            //memcpy(&stcomboDevAttr->mipi_attr, &tmp->lvds_attr, sizeof(lvds_dev_attr_t));
            memcpy(stcomboDevAttr, tmp, sizeof(combo_dev_attr_t));

            /*
            stcomboDevAttr.lvds_attr.img_size.width     = in->mipi_crop_width;
            stcomboDevAttr.lvds_attr.img_size.height    = in->mipi_crop_height;

            stcomboDevAttr.lvds_attr.raw_data_type      = mpp_get_raw_data_type(in->pixel_bitness);

            //LVDS_SYNC_MODE_SOF = 0,
            //LVDS_SYNC_MODE_SAV,  
            stcomboDevAttr.lvds_attr.sync_mode

            stcomboDevAttr.lvds_attr.vsync_type.sync_type
            stcomboDevAttr.lvds_attr.vsync_type.hblank1
            stcomboDevAttr.lvds_attr.vsync_type.hblank1

            stcomboDevAttr.lvds_attr.fid_type.fid
            stcomboDevAttr.lvds_attr.fid_type.output_fil

            //LVDS_ENDIAN_LITTLE  = 0x0,
            //LVDS_ENDIAN_BIG     = 0x1,
            stcomboDevAttr.lvds_attr.data_endian
            stcomboDevAttr.lvds_attr.sync_code_endian

            stcomboDevAttr.lvds_attr.
            stcomboDevAttr.lvds_attr.
            */
        }
        stcomboDevAttr->devno                       = 0; //!!!
    #elif HI_MPP == 4
    //mpp4
    //typedef struct {
    //    combo_dev_t         devno;              // device number
    //    input_mode_t        input_mode;         // input mode: MIPI/LVDS/SUBLVDS/HISPI/DC
    //    mipi_data_rate_t    data_rate;
    //    // MIPI Rx device crop area (corresponding to the oringnal sensor input image size)
    //    img_rect_t          img_rect;
    //
    //    union {
    //        mipi_dev_attr_t     mipi_attr;
    //        lvds_dev_attr_t     lvds_attr;
    //    };
    //} combo_dev_attr_t;
    //
    //typedef struct {
    //    data_type_t input_data_type;        // data type: 8/10/12/14/16 bit
    //    mipi_wdr_mode_t wdr_mode;           // MIPI WDR mode
    //    short lane_id[MIPI_LANE_NUM];       // lane_id: -1 - disable
    //
    //    union {
    //        short data_type[WDR_VC_NUM];    // used by the HI_MIPI_WDR_MODE_DT
    //    };
    //} mipi_dev_attr_t;
    //
    //typedef struct {
    //    data_type_t input_data_type;        // data type: 8/10/12/14 bit
    //    wdr_mode_t wdr_mode;                // WDR mode
    //
    //    lvds_sync_mode_t sync_mode;         // sync mode: SOF, SAV
    //    lvds_vsync_attr_t vsync_attr;       // normal, share, hconnect
    //    lvds_fid_attr_t fid_attr;           // frame identification code
    //
    //    lvds_bit_endian_t data_endian;      // data endian: little/big
    //    lvds_bit_endian_t sync_code_endian; // sync code endian: little/big
    //    short lane_id[LVDS_LANE_NUM];       // lane_id: -1 - disable
    //
    //    // each vc has 4 params, sync_code[i]:
    //    //sync_mode is SYNC_MODE_SOF: SOF, EOF, SOL, EOL
    //    //sync_mode is SYNC_MODE_SAV: invalid sav, invalid eav, valid sav, valid eav
    //    unsigned short sync_code[LVDS_LANE_NUM][WDR_VC_NUM][SYNC_CODE_NUM];
    //} lvds_dev_attr_t;
        stcomboDevAttr->devno                       = 0;
        stcomboDevAttr->input_mode                  = in->data_type;

        if (in->data_type == INPUT_MODE_MIPI) {
            combo_dev_attr_t *tmp = (combo_dev_attr_t*)in->mipi_mipi_attr;
            //memcpy(&stcomboDevAttr->mipi_attr, &tmp->mipi_attr, sizeof(mipi_dev_attr_t));
            memcpy(stcomboDevAttr, tmp, sizeof(combo_dev_attr_t));
        }
        else if (in->data_type == INPUT_MODE_LVDS || 
            in->data_type == INPUT_MODE_SUBLVDS ||
            in->data_type == INPUT_MODE_HISPI) {
            combo_dev_attr_t *tmp = (combo_dev_attr_t*)in->mipi_lvds_attr;
            //memcpy(&stcomboDevAttr->mipi_attr, &tmp->lvds_attr, sizeof(lvds_dev_attr_t));
            memcpy(stcomboDevAttr, tmp, sizeof(combo_dev_attr_t));
        }

        stcomboDevAttr->devno                       = 0; //!!!

    #endif
}

#endif

/*
combo_dev_attr_t MIPI_4lane_CHN0_SENSOR_IMX327_12BIT_2M_NOWDR_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI,
    .data_rate = MIPI_DATA_RATE_X1,
    .img_rect = {0, 0, 1920, 1080},
     
    {
        .mipi_attr =    
        {
            DATA_TYPE_RAW_12BIT,
            HI_MIPI_WDR_MODE_NONE,
            {0, 1, 2, 3}
        }    
    }
};
*/

int mpp_mipi_init(error_in *err, mpp_mipi_init_in *in) {
    #if HI_MPP == 1
        //there is no mipi subsystem in hi3516cv100 family
    #elif HI_MPP >= 2

        int general_error_code = 0;

        int fd;

        fd = open("/dev/hi_mipi", O_RDWR);
        if (fd < 0) {
            RETURN_ERR_GENERAL(err, "open /dev/hi_mipi", fd);
        }

        combo_dev_attr_t stcomboDevAttr;

        #if 1
            mpp_mipi_set_attrs(&stcomboDevAttr, in);
        #else
            memcpy(&stcomboDevAttr, &MIPI_4lane_CHN0_SENSOR_IMX327_12BIT_2M_NOWDR_ATTR, sizeof(combo_dev_attr_t));
            #if HI_MPP >= 3
                stcomboDevAttr.devno = 0; //TODO
            #endif
        #endif

        #if HI_MPP == 4
    	    lane_divide_mode_t lane_divide_mode = LANE_DIVIDE_MODE_0;

            general_error_code = ioctl(fd, HI_MIPI_SET_HS_MODE, &lane_divide_mode);
            if (general_error_code != HI_SUCCESS) {
                close(fd);  
                RETURN_ERR_GENERAL(err, "HI_MIPI_SET_HS_MODE", general_error_code);
            }

            general_error_code = ioctl(fd, HI_MIPI_ENABLE_MIPI_CLOCK, &stcomboDevAttr.devno);   //&devno);
            if (general_error_code != HI_SUCCESS) {
                close(fd);
                RETURN_ERR_GENERAL(err, "HI_MIPI_ENABLE_MIPI_CLOCK", general_error_code);
            }

            general_error_code = ioctl(fd, HI_MIPI_ENABLE_SENSOR_CLOCK, &stcomboDevAttr.devno); //&SnsDev);
            if (general_error_code != HI_SUCCESS) {
                close(fd);
                RETURN_ERR_GENERAL(err, "HI_MIPI_ENABLE_SENSOR_CLOCK", general_error_code); 
            }
        #endif

        #if HI_MPP >= 3
            general_error_code = ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno);  //&devno);
            if (general_error_code != HI_SUCCESS) {
                close(fd);
                RETURN_ERR_GENERAL(err, "HI_MIPI_RESET_MIPI", general_error_code); 
            }

            general_error_code = ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno);    //&SnsDev);
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

        #if HI_MPP >= 3
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

        close(fd);
    #endif

    return ERR_NONE;
}
