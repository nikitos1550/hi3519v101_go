#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <pthread.h>
#include <signal.h>

#define     FPS 25
#define ENC_WIDTH 3840
#define ENC_HEIGHT 2160

//#define SECOND_CHANNEL

#include "hi_common.h"
#include "hi_comm_sys.h"
#include "hi_comm_vb.h"
#include "hi_comm_isp.h"
#include "hi_comm_vi.h"
#include "hi_comm_vo.h"
#include "hi_comm_venc.h"
#include "hi_comm_vpss.h"
//#include "hi_comm_vdec.h"
#include "hi_comm_region.h"
#include "hi_comm_adec.h"
#include "hi_comm_aenc.h"
#include "hi_comm_ai.h"
#include "hi_comm_ao.h"
#include "hi_comm_aio.h"
#include "hi_defines.h"

#include "mpi_sys.h"
#include "mpi_vb.h"
#include "mpi_vi.h"
#include "mpi_vo.h"
#include "mpi_venc.h"
#include "mpi_vpss.h"
//#include "mpi_vdec.h"
#include "mpi_region.h"
#include "mpi_adec.h"
#include "mpi_aenc.h"
#include "mpi_ai.h"
#include "mpi_ao.h"
#include "mpi_isp.h"
#include "mpi_ae.h"
#include "mpi_awb.h"
#include "mpi_af.h"
#include "hi_vreg.h"
#include "hi_sns_ctrl.h"

#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <sys/poll.h>
#include <sys/time.h>
#include <fcntl.h>

#include "hi_mipi.h"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/uio.h>
#include <stdint.h>
#include "irisd.h"

#include <sys/ioctl.h>
#include <fcntl.h>

combo_dev_attr_t LVDS_4lane_SENSOR_IMX226_12BIT_8M_NOWDR_ATTR =
{
    .devno = 0,
    /* input mode */
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {0, 0, 4248, 2182},

    {
        .lvds_attr = {
            .raw_data_type    = RAW_DATA_12BIT,
            .wdr_mode         = HI_WDR_MODE_NONE,
            .sync_mode        = LVDS_SYNC_MODE_SAV,
            .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
            .fid_type         = {LVDS_FID_NONE, HI_TRUE},
            .data_endian      = LVDS_ENDIAN_BIG,
            .sync_code_endian = LVDS_ENDIAN_BIG,
            .lane_id = { -1, -1, 0, -1, 1, 2, -1, 3, -1, -1, -1, -1},
            .sync_code = {
                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane 0
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}
                },

                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane 1
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}
                },

                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane2
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}
                },

                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane3
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}
                },
            }
        }
    }
};

// 8lane 30fps
combo_dev_attr_t LVDS_8lane_SENSOR_IMX226_12BIT_8M_NOWDR_ATTR =
{
    .devno = 0,
    /* input mode */
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {0, 0, 4248, 2182},

    .lvds_attr =
    {
        .raw_data_type    = RAW_DATA_12BIT,
	.wdr_mode         = HI_WDR_MODE_NONE,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_NONE, HI_TRUE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {0, 1, 2, -1, 3, 4, -1, 5, 6, 7, -1, -1},
        .sync_code =
        {
            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane 0
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane 1
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane2
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane3
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane4
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane5
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane6
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane7
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },
        }
    }
};

combo_dev_attr_t LVDS_10lane_SENSOR_IMX226_10BIT_8M_NOWDR_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {0, 0, 4248, 2182},
    .lvds_attr = 
    {
        .raw_data_type    = RAW_DATA_10BIT,
        .wdr_mode         = HI_WDR_MODE_NONE,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_NONE, HI_TRUE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1},
        .sync_code = 
        {
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 0
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 1
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 2
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 3
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 4
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 5
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 6
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 7
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 8
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 9
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
        }
    }
};
combo_dev_attr_t LVDS_10lane_SENSOR_IMX226_12BIT_12M_NOWDR_ATTR = 
{
    .devno = 0,
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {0, 0, 4100, 3100},
    .lvds_attr = 
    {
        .raw_data_type    = RAW_DATA_12BIT,
        .wdr_mode         = HI_WDR_MODE_NONE,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_NONE, HI_TRUE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1},
        .sync_code = 
        {
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane 0
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane 1
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane2
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane3
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane4
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane5
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane6
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane7
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
			  {{0xab0, 0xb60, 0x800, 0x9d0},      // lane8
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
			  {{0xab0, 0xb60, 0x800, 0x9d0},      // lane9
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
        }
    }
};

// 4Kx3K crop to 3kx3k
combo_dev_attr_t LVDS_10lane_SENSOR_IMX226_12BIT_9M_NOWDR_ATTR = 
{
    .devno = 0,
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {570, 22, 3000, 3000},
    .lvds_attr = 
    {
        .raw_data_type    = RAW_DATA_12BIT,
        .wdr_mode         = HI_WDR_MODE_NONE,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_NONE, HI_TRUE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1},
        .sync_code = 
        {
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane 0
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane 1
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane2
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane3
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane4
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane5
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane6
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane7
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
			  {{0xab0, 0xb60, 0x800, 0x9d0},      // lane8
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
			  {{0xab0, 0xb60, 0x800, 0x9d0},      // lane9
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
        }
    }
};


/* 6lane 12bit 30fps*/
combo_dev_attr_t LVDS_6lane_SENSOR_IMX274_12BIT_8M_NOWDR_ATTR =
{
    .devno = 0,
    /* input mode */
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {12, 40, 3840, 2160},

    .lvds_attr = 
    {
        .raw_data_type    = RAW_DATA_12BIT,
        .wdr_mode         = HI_WDR_MODE_NONE,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_NONE, HI_FALSE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {-1, 0, 1, -1, 2, 3, -1, 4, 5, -1, -1, -1},
        .sync_code = 
        {
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane 0
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},

            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane 1
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},

            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane2
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},

            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane3
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},

            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane4
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},

            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane5
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}}
        }
    }
};

// 10lane 30fps
combo_dev_attr_t LVDS_10lane_SENSOR_IMX274_10BIT_8M_2WDR1_ATTR =
{
    .devno = 0,
    /* input mode */
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {12, 40, 3840, 2160},

    .lvds_attr = 
    {
        .raw_data_type    = RAW_DATA_10BIT,
        .wdr_mode         = HI_WDR_MODE_DOL_2F,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_IN_SAV, HI_TRUE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1},
        .sync_code = 
        {
            {{0x2ac,0x2d8,0x201,0x275},      // lane 0
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

            {{0x2ac,0x2d8,0x201,0x275},      // lane 1
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

            {{0x2ac,0x2d8,0x201,0x275},      // lane 2
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

            {{0x2ac,0x2d8,0x201,0x275},      // lane 3
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

            {{0x2ac,0x2d8,0x201,0x275},      // lane 4
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

            {{0x2ac,0x2d8,0x201,0x275},      // lane 5
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

             {{0x2ac,0x2d8,0x201,0x275},      // lane 6
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

             {{0x2ac,0x2d8,0x201,0x275},      // lane 7
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

             {{0x2ac,0x2d8,0x201,0x275},      // lane 8
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

             {{0x2ac,0x2d8,0x201,0x275},      // lane 9
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},
        }
    }
};

VI_DEV_ATTR_S DEV_ATTR_LVDS_BASE =
{
    /* interface mode */
    VI_MODE_LVDS,
    /* multiplex mode */
    VI_WORK_MODE_1Multiplex,
    /* r_mask    g_mask    b_mask*/
    {0xFFF00000,    0x0},
    /* progessive or interleaving */
    VI_SCAN_PROGRESSIVE,
    /*AdChnId*/
    { -1, -1, -1, -1},
    /*enDataSeq, only support yuv*/
    VI_INPUT_DATA_YUYV,

    /* synchronization information */
    {
        /*port_vsync   port_vsync_neg     port_hsync        port_hsync_neg        */
        VI_VSYNC_PULSE, VI_VSYNC_NEG_LOW, VI_HSYNC_VALID_SINGNAL, VI_HSYNC_NEG_HIGH, VI_VSYNC_VALID_SINGAL, VI_VSYNC_VALID_NEG_HIGH,

        /*hsync_hfb    hsync_act    hsync_hhb*/
        {
            0,            1280,        0,
            /*vsync0_vhb vsync0_act vsync0_hhb*/
            0,            720,        0,
            /*vsync1_vhb vsync1_act vsync1_hhb*/
            0,            0,            0
        }
    },
    /* use interior ISP */
    VI_PATH_ISP,
    /* input data type */
    VI_DATA_TYPE_RGB,
    /* bRever */
    HI_FALSE,
    /* DEV CROP */
    {0, 0, 1920, 1080},
    {
        {
            {1920, 1080},
            HI_FALSE,

        },
        {
            VI_REPHASE_MODE_NONE,
            VI_REPHASE_MODE_NONE
        }
    }
};


typedef struct {
	VENC_CHN channel;
	int fd;
	int fd_id;
	int stream_id;
	Irisd *irisd;
	unsigned char ts_type; //0 - rtsp 90kHz base, 1 - c50 ms
} Capture;


void *init_hisi();
Frame *capture_hisi(Capture *);

int vb_init(void);
int isp_init(void);
int vi_init(void);
int vpss_init(void);
int venc_init(void);

HI_S32 SAMPLE_COMM_VI_SetMipiAttr(void);
HI_VOID* Test_ISP_Run(HI_VOID *param);
int VencGetH264Stream(VENC_CHN *stream);
void resolve_mppv2_errors(HI_S32 error);


void capture_cb(int fd, short flag, void *arg) {
	Frame *frame = capture_hisi(arg);
	Capture *capture = (Capture *)arg;
	//printf("capture_cb capture->channel = %d\n", capture->channel);
	//if (capture->channel == 0) {
		deliver_frame(capture->irisd, frame);
	//} else {
	//	free(frame);
	//}
}


Capture capture1; 
Capture capture2; 

int configure_hisi(Irisd *irisd) {
	init_hisi(irisd);

	capture1.channel = 0;
	capture1.irisd = irisd;
	capture1.fd = HI_MPI_VENC_GetFd(0);
	capture1.fd_id = 0;
	capture1.ts_type = 1;
	//printf("stream 0 configured as c50 ms timestamps [%d]\n", capture1.ts_type);

	irisd->streams[irisd->stream_count].id = irisd->stream_count;
	irisd->stream_count++;

	struct event *v_evt_0 = event_new(irisd->base, capture1.fd, EV_READ|EV_PERSIST, capture_cb, &capture1);
	event_add(v_evt_0,NULL);  

	#ifdef SECOND_CHANNEL
	capture2.channel = 1;
	capture2.irisd = irisd;
	capture2.fd = HI_MPI_VENC_GetFd(1);
	capture2.fd_id = 1;
	capture2.ts_type = 1;
	//printf("stream 1 configured as rtsp 90kHz base timestamps [%d]\n", capture2.ts_type);


	irisd->streams[irisd->stream_count].id = irisd->stream_count;
	irisd->stream_count++;

	struct event *v_evt_1 = event_new(irisd->base, capture2.fd, EV_READ|EV_PERSIST, capture_cb, &capture2);
	event_add(v_evt_1,NULL);  
	#endif

}

void *init_hisi(Irisd *irisd) {
	if (vb_init() != 0) exit(1);
	fprintf(stderr, "vb_init ok\n");

  	if (isp_init() != 0) exit(1);
	fprintf(stderr, "isp_init ok\n");

	if (vi_init() != 0) exit(1);
	fprintf(stderr, "vi_init ok\n");

	if (vpss_init() != 0) exit(1);
	fprintf(stderr, "vpss_init ok\n");

	if (venc_init_0() != 0) exit(1);
	fprintf(stderr, "venc_init_0 ok\n");

	#ifdef SECOND_CHANNEL
	if (venc_init_1() != 0) exit(1);
	fprintf(stderr, "venc_init_1 ok\n");
	#endif

	fprintf(stderr, "Encoding started\n");

	/*
	irisd->streams[irisd->stream_count].id = irisd->stream_count;
	irisd->stream_count++;

	Capture *c = (Capture *)calloc(1, sizeof(Capture));
	c->channel = 0;
	c->irisd = irisd;
	c->fd = HI_MPI_VENC_GetFd(0);
	*/
	//return c;
}

long timevaldiff(struct timeval *starttime, struct timeval *finishtime) {
	long msec = 0;

	if (finishtime->tv_usec > starttime->tv_usec) {
		 msec+=(finishtime->tv_usec-starttime->tv_usec);
	} else {
		msec+=finishtime->tv_usec;
		msec+=(1000000-starttime->tv_usec);
	}

  	return msec;
}

Frame *capture_hisi(Capture *r) {
	VENC_CHN VencChn = r->fd_id;//r->channel;

	HI_S32 s32Ret;
	VENC_CHN_STAT_S stStat;
	VENC_STREAM_S stStream;

	memset(&stStream, 0, sizeof(stStream));
    s32Ret = HI_MPI_VENC_Query(VencChn, &stStat);
    if (HI_SUCCESS != s32Ret) {
    	printf("HI_MPI_VENC_Query chn[%d] failed with %#x!\n", VencChn, s32Ret);                    
    }

	if (0 == stStat.u32CurPacks) {
		printf("stStat.u32CurPacks == 0\n");
        //SAMPLE_PRT("NOTE: Current  frame is NULL!\n");
                       //continue;
    }  

	stStream.pstPack = (VENC_PACK_S*)malloc(sizeof(VENC_PACK_S) * stStat.u32CurPacks);
    if (NULL == stStream.pstPack) {
		printf("malloc stream pack failed!\n");
                        
    }

    stStream.u32PackCount = stStat.u32CurPacks;
    s32Ret = HI_MPI_VENC_GetStream(VencChn, &stStream, HI_TRUE);
    if (HI_SUCCESS != s32Ret) {
    	free(stStream.pstPack);
        stStream.pstPack = NULL;
        printf("HI_MPI_VENC_GetStream failed with %#x!\n", s32Ret);
    }

  	int j;
  	int len = 0;
  	for(j = 0; j < stStream.u32PackCount;j++) {
    	len += stStream.pstPack[j].u32Len;// + stStream.pstPack[j].u32Len[1];
  	}
  	len += stStream.u32PackCount*sizeof(struct iovec) + sizeof(Frame);

  	Frame *frame = (Frame *)calloc(1,len);

 	frame->nal_count = stStream.u32PackCount;
  	frame->nals = (void *)frame + sizeof(Frame);
  	char *ptr = (void *)frame + sizeof(Frame) + stStream.u32PackCount*sizeof(struct iovec);

	if (r->ts_type == 0) {
		frame->dts = stStream.pstPack[0].u64PTS/1000*90;
	} else {	
  		struct timeval tp;
  		gettimeofday(&tp, NULL);
  		frame->dts = tp.tv_sec*1000 + (tp.tv_usec + 500)/ 1000;
	}

  	// Here will be other tracks and streams
  	frame->stream_id = r->channel;//0;

  	for(j = 0; j < stStream.u32PackCount;j++) {
    	frame->nals[j].iov_len = stStream.pstPack[j].u32Len;//[0] + stStream.pstPack[j].u32Len[1];
    	frame->nals[j].iov_base = ptr;
    	ptr += frame->nals[j].iov_len;

    	memcpy(frame->nals[j].iov_base, stStream.pstPack[j].pu8Addr, stStream.pstPack[j].u32Len);
    	//if(stStream.pstPack[j].u32Len[1]) {
    		//memcpy(frame->nals[j].iov_base + stStream.pstPack[j].u32Len[0], stStream.pstPack[j].pu8Addr[1], stStream.pstPack[j].u32Len[1]);      
    	//}
    	int l = frame->nals[j].iov_len - 4;
    	char *p = frame->nals[j].iov_base;
    	p[0] = l >> 24;
    	p[1] = l >> 16;
    	p[2] = l >> 8;
    	p[3] = l >> 0;
    	if(5 == (p[4]&31)) {
      		frame->keyframe = 1;
    	}
  	}

  	if((s32Ret = HI_MPI_VENC_ReleaseStream(VencChn, &stStream))) {
    	fprintf(stderr, "failed to release stream: %#x\n", s32Ret);
  	}
  	free(stStream.pstPack);

  	return frame;
}

////////////////////////////////////////////////////////////////////////////////

int vb_init(void){
	HI_S32 			ret;
	VB_CONF_S 		stVbConf;
    MPP_SYS_CONF_S	stSysConf;

	memset(&stVbConf,0,sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt = 128;

    /*video buffer*/   
    stVbConf.astCommPool[0].u32BlkSize =(CEILING_2_POWER(3840, 64) * CEILING_2_POWER(2160, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt = 10;

	#ifdef SECOND_CHANNEL
    stVbConf.astCommPool[1].u32BlkSize =(CEILING_2_POWER(1920, 64) * CEILING_2_POWER(1080, 64) * 1.5);
    stVbConf.astCommPool[1].u32BlkCnt = 10;
	#endif

	ret = HI_MPI_SYS_Exit();
	if (ret != HI_SUCCESS) {
		fprintf(stderr, "HI_MPI_SYS_Exit failed: ");
		resolve_mppv2_errors(ret);
		return 1;
	}

	ret = HI_MPI_VB_Exit();
	if (ret != HI_SUCCESS) {
		fprintf(stderr, "HI_MPI_VB_Exit failed: ");
		resolve_mppv2_errors(ret);
		return 1;
	}

	MPP_VERSION_S stVersion;

	ret = HI_MPI_SYS_GetVersion(&stVersion);
	if(ret != HI_SUCCESS) {
    	fprintf(stderr, "HI_MPI_SYS_GetVersion failed: ");
		resolve_mppv2_errors(ret);
		return 1;
	}
	fprintf(stderr, "MPI version is %s\n", stVersion.aVersion);

	ret = HI_MPI_VB_SetConf(&stVbConf);
	if(ret != HI_SUCCESS) {
		fprintf(stderr, "HI_MPI_VB_SetConf failed: ");
		resolve_mppv2_errors(ret);
		return 1;
	}

	ret = HI_MPI_VB_Init();
	if (ret != HI_SUCCESS) {
		fprintf(stderr, "HI_MPI_VB_Init failed: ");
		resolve_mppv2_errors(ret);
		return 1;
	}

	stSysConf.u32AlignWidth = 64;

	ret = HI_MPI_SYS_SetConf(&stSysConf);
	if (ret != HI_SUCCESS) {
		fprintf(stderr, "HI_MPI_SYS_SetConf failed: ");
		resolve_mppv2_errors(ret);
		return 1;
	}

	ret = HI_MPI_SYS_Init();
	if(ret != HI_SUCCESS) {
		fprintf(stderr, "HI_MPI_SYS_Init failed: ");
		resolve_mppv2_errors(ret);
		return 1;
	}

	return 0;
}

HI_S32 SAMPLE_COMM_VI_SetMipiAttr(void) {

    HI_S32 fd;
    combo_dev_attr_t *pstcomboDevAttr, stcomboDevAttr;

    /* mipi reset unrest */
    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0)
    {
        printf("warning: open hi_mipi dev failed\n");
        return -1;
    }
    
	#ifdef CMOS_IMX274
		#ifdef CMOS_WDR
		pstcomboDevAttr = &LVDS_10lane_SENSOR_IMX274_10BIT_8M_2WDR1_ATTR;
		#else
 		pstcomboDevAttr = &LVDS_6lane_SENSOR_IMX274_12BIT_8M_NOWDR_ATTR;
		#endif
	#endif
	#ifdef CMOS_IMX226
		pstcomboDevAttr = &LVDS_8lane_SENSOR_IMX226_12BIT_8M_NOWDR_ATTR;
	#endif

    memcpy(&stcomboDevAttr, pstcomboDevAttr, sizeof(combo_dev_attr_t));
    stcomboDevAttr.devno = 0;

  /* 1.reset mipi */
    if(ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno)) {
        printf("HI_MIPI_RESET_MIPI failed\n");
        close(fd);
        return -1;
   	}

    /* 2.reset sensor */
    if(ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno)) {
    	printf("HI_MIPI_RESET_SENSOR failed\n");
        close(fd);
        return -1;
    }

    if (ioctl(fd, HI_MIPI_SET_DEV_ATTR, pstcomboDevAttr)) {
        printf("set mipi attr failed\n");
        close(fd);
        return -1;
    }

    /* 4.unreset mipi */
    if(ioctl(fd, HI_MIPI_UNRESET_MIPI, &stcomboDevAttr.devno)) {
        printf("HI_MIPI_UNRESET_MIPI failed\n");
        close(fd);
        return -1;
    }

    /* 5.unreset sensor */
    if(ioctl(fd, HI_MIPI_UNRESET_SENSOR, &stcomboDevAttr.devno)) {
        printf("HI_MIPI_UNRESET_SENSOR failed\n");
        close(fd);
        return -1;
    }

    close(fd);

    return HI_SUCCESS;
}

HI_VOID* Test_ISP_Run(HI_VOID *param){
	ISP_DEV IspDev = 0;
    HI_MPI_ISP_Run(IspDev);

    return HI_NULL;
}

static pthread_t gs_IspPid;

int isp_init(void){
	HI_S32 			ret;

    ret = HI_MPI_ISP_Exit(0);
    if (ret != HI_SUCCESS) {
        printf("%s: HI_MPI_ISP_Exit failed with %#x!\n", \
               __FUNCTION__, ret);
        return ret;
    }

	ret = SAMPLE_COMM_VI_SetMipiAttr();
    if (HI_SUCCESS != ret) {
      printf("%s: mipi init failed!\n", __FUNCTION__);
      return HI_FAILURE;
    }

////////////////////////////
    ISP_DEV IspDev = 0;
    //HI_S32 ret;
    ISP_PUB_ATTR_S stPubAttr;
    ALG_LIB_S stLib;

	#ifdef CMOS_IMX226
	const ISP_SNS_OBJ_S *g_pstSnsObj[2] =  {&stSnsImx226Obj, HI_NULL};
	#endif
	#ifdef CMOS_IMX274
	const ISP_SNS_OBJ_S *g_pstSnsObj[2] =  {&stSnsImx274Obj, HI_NULL};
	#endif

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;
    ALG_LIB_S stAfLib;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    stAfLib.s32Id = 0;
    strncpy(stAeLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
    strncpy(stAfLib.acLibName, HI_AF_LIB_NAME, sizeof(HI_AF_LIB_NAME)); 

	if (g_pstSnsObj[0]->pfnRegisterCallback != HI_NULL) {
        ret = g_pstSnsObj[0]->pfnRegisterCallback(IspDev, &stAeLib, &stAwbLib);
        if (ret != HI_SUCCESS)
        {
            printf("%s: sensor_register_callback failed with %#x!\n", __FUNCTION__, ret);
            return ret;
        }
    } else {
        printf("%s: sensor_register_callback failed with HI_NULL!\n",  __FUNCTION__);
    }

    /* 2. register hisi ae lib */
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AE_LIB_NAME);
    ret = HI_MPI_AE_Register(IspDev, &stLib);
    if (ret != HI_SUCCESS) {
        printf("%s: HI_MPI_AE_Register failed!\n", __FUNCTION__);
        return ret;
    }

    /* 3. register hisi awb lib */
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AWB_LIB_NAME);
    ret = HI_MPI_AWB_Register(IspDev, &stLib);
    if (ret != HI_SUCCESS) {
        printf("%s: HI_MPI_AWB_Register failed!\n", __FUNCTION__);
        return ret;
    }

    /* 4. register hisi af lib */
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AF_LIB_NAME);
    ret = HI_MPI_AF_Register(IspDev, &stLib);
    if (ret != HI_SUCCESS) {
        printf("%s: HI_MPI_AF_Register failed!\n", __FUNCTION__);
        return ret;
    }

    /* 5. isp mem init */
    ret = HI_MPI_ISP_MemInit(IspDev);
    if (ret != HI_SUCCESS) {
        printf("%s: HI_MPI_ISP_Init failed!\n", __FUNCTION__);
        return ret;
    }

    /* 6. isp set WDR mode */
    ISP_WDR_MODE_S stWdrMode;
	
	#ifdef CMOS_WDR
		//TODO
	 	stWdrMode.enWDRMode  = WDR_MODE_2To1_LINE;//WDR_MODE_NONE;//enWDRMode;
	#else
    	stWdrMode.enWDRMode  = WDR_MODE_NONE;
	#endif

    ret = HI_MPI_ISP_SetWDRMode(0, &stWdrMode);    
    if (HI_SUCCESS != ret) {
        printf("start ISP WDR failed!\n");
        return ret;
    }

	#ifdef CMOS_IMX226
	stPubAttr.enBayer				= BAYER_RGGB;
    stPubAttr.f32FrameRate          = FPS;//30;
    stPubAttr.stWndRect.s32X        = 0;
    stPubAttr.stWndRect.s32Y        = 0;
    stPubAttr.stWndRect.u32Width    = 3840;
    stPubAttr.stWndRect.u32Height   = 2160;
    stPubAttr.stSnsSize.u32Width    = 3840;
    stPubAttr.stSnsSize.u32Height   = 2160;    
	#endif
	#ifdef CMOS_IMX274
	stPubAttr.enBayer				= BAYER_RGGB;
    stPubAttr.f32FrameRate          = FPS;//30;
    stPubAttr.stWndRect.s32X        = 0;
    stPubAttr.stWndRect.s32Y        = 0;
    stPubAttr.stWndRect.u32Width    = 3840;
    stPubAttr.stWndRect.u32Height   = 2160;
    stPubAttr.stSnsSize.u32Width    = 3840;
    stPubAttr.stSnsSize.u32Height   = 2160;    
	#endif


    ret = HI_MPI_ISP_SetPubAttr(IspDev, &stPubAttr);
    if (ret != HI_SUCCESS) {
        printf("%s: HI_MPI_ISP_SetPubAttr failed with %#x!\n", __FUNCTION__, ret);
        return ret;
    }

    /* 8. isp init */
    ret = HI_MPI_ISP_Init(IspDev);
    if (ret != HI_SUCCESS) {
        printf("%s: HI_MPI_ISP_Init failed!\n", __FUNCTION__);
        return ret;
    }

#if 1
    if (0 != pthread_create(&gs_IspPid, 0, (void* (*)(void*))Test_ISP_Run, NULL)) {
        printf("%s: create isp running thread failed!\n", __FUNCTION__);
        return 1;
    }
#else
	/* configure thread priority */
	if (1) {
		#include <sched.h>

		pthread_attr_t attr;
		struct sched_param param;
		int newprio = 50;

		pthread_attr_init(&attr);

		if (1) {
			int policy = 0;
			int min, max;

			pthread_attr_getschedpolicy(&attr, &policy);
			printf("-->default thread use policy is %d --<\n", policy);

			pthread_attr_setschedpolicy(&attr, SCHED_RR);
			pthread_attr_getschedpolicy(&attr, &policy);
			printf("-->current thread use policy is %d --<\n", policy);

			switch (policy) {
				case SCHED_FIFO:
					printf("-->current thread use policy is SCHED_FIFO --<\n");
					break;

				case SCHED_RR:
					printf("-->current thread use policy is SCHED_RR --<\n");
					break;

				case SCHED_OTHER:
					printf("-->current thread use policy is SCHED_OTHER --<\n");
					break;

				default:
					printf("-->current thread use policy is UNKNOW --<\n");
					break;
			}

			min = sched_get_priority_min(policy);
			max = sched_get_priority_max(policy);

			printf("-->current thread policy priority range (%d ~ %d) --<\n", min, max);
		}

		pthread_attr_getschedparam(&attr, &param);

		printf("-->default isp thread priority is %d , next be %d --<\n", param.sched_priority, newprio);
		param.sched_priority = newprio;
		pthread_attr_setschedparam(&attr, &param);

		if (0 != pthread_create(&gs_IspPid, &attr, (void* (*)(void*))HI_MPI_ISP_Run, NULL)) {
			printf("%s: create isp running thread failed!\n", __FUNCTION__);
			return HI_FAILURE;
		}

		pthread_attr_destroy(&attr);
	}
#endif


	ISP_DRC_ATTR_S stDRC;
	ret = HI_MPI_ISP_GetDRCAttr(IspDev, &stDRC);
    if (ret != HI_SUCCESS) {
        printf("%s: HI_MPI_ISP_GetDRCAttr!\n", __FUNCTION__);
        return 1;
    }
	printf("DRC is %d\n", stDRC.bEnable);
	stDRC.bEnable = 1;
	stDRC.enOpType = 0;

	ret = HI_MPI_ISP_SetDRCAttr(IspDev, &stDRC);
    if (ret != HI_SUCCESS) {
        printf("%s: HI_MPI_ISP_SetDRCAttr!\n", __FUNCTION__);
        return 1;
    }


	return 0;
}

int vi_init(void){
	HI_S32 			ret;

   	HI_S32 s32Ret;
    HI_S32 s32IspDev = 0;
    ISP_WDR_MODE_S stWdrMode;
    VI_DEV_ATTR_S  stViDevAttr;
    
    memset(&stViDevAttr,0,sizeof(stViDevAttr));

	memcpy(&stViDevAttr, &DEV_ATTR_LVDS_BASE, sizeof(stViDevAttr));
	#ifdef CMOS_IMX226
    stViDevAttr.stDevRect.s32X = 100;
	#endif
	#ifdef CMOS_IMX274
    stViDevAttr.stDevRect.s32X = 0;//100;
	#endif
    stViDevAttr.stDevRect.s32Y = 0;
    stViDevAttr.stDevRect.u32Width  = 3840;
    stViDevAttr.stDevRect.u32Height = 2160;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Width  = 3840;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Height = 2160;
    stViDevAttr.stBasAttr.stSacleAttr.bCompress = HI_FALSE;

    s32Ret = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (s32Ret != HI_SUCCESS) {
        printf("HI_MPI_VI_SetDevAttr failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }

	//TODO
    #ifdef CMOS_WDR
    s32Ret = HI_MPI_ISP_GetWDRMode(s32IspDev, &stWdrMode);
    if (s32Ret != HI_SUCCESS)
    {
        printf("HI_MPI_ISP_GetWDRMode failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }


    if (stWdrMode.enWDRMode)  //wdr mode
    {
        VI_WDR_ATTR_S stWdrAttr;

        stWdrAttr.enWDRMode = stWdrMode.enWDRMode;
        stWdrAttr.bCompress = HI_TRUE;

        s32Ret = HI_MPI_VI_SetWDRAttr(0, &stWdrAttr);
        if (s32Ret)
        {
            printf("HI_MPI_VI_SetWDRAttr failed with %#x!\n", s32Ret);
            return HI_FAILURE;
        }
    }
    #endif
    
    s32Ret = HI_MPI_VI_EnableDev(0);
    if (s32Ret != HI_SUCCESS) {
        printf("HI_MPI_VI_EnableDev failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }

    RECT_S stCapRect;
    SIZE_S stTargetSize;

    stCapRect.s32X = 0;
    stCapRect.s32Y = 0;
    stCapRect.u32Width  = 3840;//4000;
    stCapRect.u32Height = 2160;//3000;
    stTargetSize.u32Width = stCapRect.u32Width;
    stTargetSize.u32Height = stCapRect.u32Height;

    VI_CHN_ATTR_S stChnAttr;

	memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));
    stChnAttr.enCapSel = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width = stTargetSize.u32Width ;
    stChnAttr.stDestSize.u32Height =  stTargetSize.u32Height ;
    stChnAttr.enPixFormat = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   /* sp420 or sp422 */

    stChnAttr.bMirror = HI_FALSE;
    stChnAttr.bFlip = HI_FALSE;

    stChnAttr.s32SrcFrameRate = FPS;//30;
    stChnAttr.s32DstFrameRate = FPS;//30;
    stChnAttr.enCompressMode = COMPRESS_MODE_NONE;

    s32Ret = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (s32Ret != HI_SUCCESS) {
        printf("HI_MPI_VI_SetChnAttr failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }
////////////LDC////////////////////


#ifdef CMOS_LDC

VI_LDC_ATTR_S stLDCAttr;
//First enable VI devices and VI channel.
//Initialize LDC attributes.
stLDCAttr.bEnable = HI_TRUE;
stLDCAttr.stAttr.enViewType = LDC_VIEW_TYPE_ALL;
//LDC_VIEW_TYPE_CROP;
stLDCAttr.stAttr.s32CenterXOffset = 0;
stLDCAttr.stAttr.s32CenterYOffset = 0;
stLDCAttr.stAttr.s32Ratio = 58;
stLDCAttr.stAttr.s32MinRatio = 0;
//Set LDC attributes.
s32Ret = HI_MPI_VI_SetLDCAttr(0, &stLDCAttr);
if (HI_SUCCESS != s32Ret) {
    printf("Set vi LDC attr err:0x%x\n", s32Ret);
    return HI_FAILURE;
}
//Obtain LDC attributes.
s32Ret = HI_MPI_VI_GetLDCAttr (0, &stLDCAttr);
if (HI_SUCCESS != s32Ret) {
    printf("Get vi LDC attr err:0x%x\n", s32Ret);
    return HI_FAILURE;
}
#endif

////////////

    s32Ret = HI_MPI_VI_EnableChn(0);
    if (s32Ret != HI_SUCCESS) {
        printf(" HI_MPI_VI_EnableChn failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }

	return 0;
}

int vpss_init(void){
	HI_S32 			ret, s32Ret;

  	VPSS_GRP VpssGrp = 0;
    VPSS_CHN VpssChn = 0;
    VPSS_GRP_ATTR_S stVpssGrpAttr;
    VPSS_CHN_ATTR_S stVpssChnAttr;
    VPSS_CHN_MODE_S stVpssChnMode;

    VpssGrp = 0;

	stVpssGrpAttr.u32MaxW = 3840;//4000;
	stVpssGrpAttr.u32MaxH = 2160;//3000;
	stVpssGrpAttr.bIeEn = HI_FALSE;
	stVpssGrpAttr.bNrEn = HI_TRUE;//HI_TRUE;//HI_FALSE;//HI_TRUE;
	stVpssGrpAttr.stNrAttr.enNrType = VPSS_NR_TYPE_VIDEO;
	stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrRefSource = VPSS_NR_REF_FROM_RFR;//VPSS_NR_REF_FROM_CHN0, VPSS_NR_REF_FROM_SRC
    stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrOutputMode = VPSS_NR_OUTPUT_NORMAL;//VPSS_NR_OUTPUT_DELAY NORMAL
	stVpssGrpAttr.stNrAttr.u32RefFrameNum = 2;
	stVpssGrpAttr.bHistEn = HI_FALSE;
	stVpssGrpAttr.bDciEn = HI_FALSE;
	stVpssGrpAttr.enDieMode = VPSS_DIE_MODE_NODIE;
	stVpssGrpAttr.enPixFmt = PIXEL_FORMAT_YUV_SEMIPLANAR_420;//SAMPLE_PIXEL_FORMAT;
    stVpssGrpAttr.bStitchBlendEn   = HI_FALSE;

    s32Ret = HI_MPI_VPSS_CreateGrp(VpssGrp, &stVpssGrpAttr);
    if (s32Ret != HI_SUCCESS) {
        printf("HI_MPI_VPSS_CreateGrp failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }

    s32Ret = HI_MPI_VPSS_StartGrp(VpssGrp);
    if (s32Ret != HI_SUCCESS) {
        printf("HI_MPI_VPSS_StartGrp failed with %#x\n", s32Ret);
        return HI_FAILURE;
    }

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

	stSrcChn.enModId  = HI_ID_VIU;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 0;
    
    stDestChn.enModId  = HI_ID_VPSS;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;
    
    s32Ret = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (s32Ret != HI_SUCCESS) {
            printf("failed with %#x!\n", s32Ret);
            return HI_FAILURE;
    }



	VpssChn = 0;
    stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
    stVpssChnMode.bDouble        = HI_FALSE;
    stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    stVpssChnMode.u32Width       = 3840;// 3840;//3840;//4000;
    stVpssChnMode.u32Height      = 2160;//2160;//2160;//3000;
    stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE;//COMPRESS_MODE_SEG;
    memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));
    stVpssChnAttr.s32SrcFrameRate = FPS;
    stVpssChnAttr.s32DstFrameRate = FPS;

	s32Ret = HI_MPI_VPSS_SetChnAttr(VpssGrp, VpssChn, &stVpssChnAttr);
	if (s32Ret != HI_SUCCESS) {
    	printf("HI_MPI_VPSS_SetChnAttr failed with %#x\n", s32Ret);
        return HI_FAILURE;
    }

	s32Ret = HI_MPI_VPSS_SetChnMode(VpssGrp, VpssChn, &stVpssChnMode);
    if (s32Ret != HI_SUCCESS) {
    	printf("%s failed with %#x\n", __FUNCTION__, s32Ret);
        return HI_FAILURE;
    }     

    /*    VPSS_CROP_INFO_S CropInfo;
        CropInfo.bEnable = 1;
        CropInfo.enCropCoordinate = VPSS_CROP_ABS_COOR;
        CropInfo.stCropRect.s32X = 0;
        CropInfo.stCropRect.s32Y = 0;
        CropInfo.stCropRect.u32Width = 3840;
        CropInfo.stCropRect.u32Height = 1080;

        s32Ret = HI_MPI_VPSS_SetChnCrop(VpssGrp, VpssChn, &CropInfo);
    if (s32Ret != HI_SUCCESS) {
        printf("HI_MPI_VPSS_SetChnCrop failed with %#x\n", s32Ret);
        return HI_FAILURE;
    }*/

	s32Ret = HI_MPI_VPSS_EnableChn(VpssGrp, VpssChn);
    if (s32Ret != HI_SUCCESS) {
        printf("HI_MPI_VPSS_EnableChn failed with %#x\n", s32Ret);
        return HI_FAILURE;
    }

	/////////////////////////
	#ifdef SECOND_CHANNEL
	VpssChn = 1;
    stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
    stVpssChnMode.bDouble        = HI_FALSE;
    stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    stVpssChnMode.u32Width       = 2560;// 3840;//3840;//4000;
    stVpssChnMode.u32Height      = 1440;//2160;//2160;//3000;
    stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE;//COMPRESS_MODE_SEG;
    memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));
    stVpssChnAttr.s32SrcFrameRate = FPS;
    stVpssChnAttr.s32DstFrameRate = FPS;

	s32Ret = HI_MPI_VPSS_SetChnAttr(VpssGrp, VpssChn, &stVpssChnAttr);
	if (s32Ret != HI_SUCCESS) {
    	printf("HI_MPI_VPSS_SetChnAttr failed with %#x\n", s32Ret);
        return HI_FAILURE;
    }

	s32Ret = HI_MPI_VPSS_SetChnMode(VpssGrp, VpssChn, &stVpssChnMode);
    if (s32Ret != HI_SUCCESS) {
    	printf("%s failed with %#x\n", __FUNCTION__, s32Ret);
        return HI_FAILURE;
    }     

	VPSS_CROP_INFO_S CropInfo;
	CropInfo.bEnable = 1;
	CropInfo.enCropCoordinate = VPSS_CROP_ABS_COOR;
	CropInfo.stCropRect.s32X = 0;
	CropInfo.stCropRect.s32Y = 0;
	CropInfo.stCropRect.u32Width = 2560;
	CropInfo.stCropRect.u32Height = 720;

	s32Ret = HI_MPI_VPSS_SetChnCrop(VpssGrp, VpssChn, &CropInfo);
    if (s32Ret != HI_SUCCESS) {
        printf("HI_MPI_VPSS_SetChnCrop failed with %#x\n", s32Ret);
        return HI_FAILURE;
    }

	s32Ret = HI_MPI_VPSS_EnableChn(VpssGrp, VpssChn);
    if (s32Ret != HI_SUCCESS) {
        printf("HI_MPI_VPSS_EnableChn failed with %#x\n", s32Ret);
        return HI_FAILURE;
    }
	#endif


	return 0;
}

int venc_init_0(void){
	HI_S32 			ret;

    HI_S32 s32Ret;
    VENC_CHN_ATTR_S stVencChnAttr;
    VENC_ATTR_H264_S stH264Attr;
    VENC_ATTR_H264_CBR_S    stH264Cbr;
	VENC_ATTR_H264_VBR_S    stH264Vbr;
	VENC_ATTR_H264_AVBR_S    stH264AVbr;
	VENC_ATTR_H264_FIXQP_S  stH264FixQp;

  	VENC_ATTR_H265_S        stH265Attr;
    VENC_ATTR_H265_CBR_S    stH265Cbr;
    VENC_ATTR_H265_VBR_S    stH265Vbr;
    VENC_ATTR_H265_FIXQP_S  stH265FixQp;

    stVencChnAttr.stVeAttr.enType = PT_H264;
    stH264Attr.u32MaxPicWidth = 3840;//4000;
    stH264Attr.u32MaxPicHeight = 2160;//3000;
    stH264Attr.u32PicWidth = 3840;//3840;//4000;/*the picture width*/
    stH264Attr.u32PicHeight = 2160;//2160;//3000;/*the picture height*/
    stH264Attr.u32BufSize  = 3840*2160*2;//4000 * 3000 * 2;/*stream buffer size*/
    stH264Attr.u32Profile  = 2;/*0: baseline; 1:MP; 2:HP;  3:svc_t */
    stH264Attr.bByFrame = HI_TRUE;//HI_FALSE;//HI_TRUE;//HI_TRUE;/*get stream mode is slice mode or frame mode?*/
	//stH264Attr.u32BFrameNum = 0;/* 0: not support B frame; >=1: number of B frames */
	//stH264Attr.u32RefNum = 1;/* 0: default; number of refrence frame*/

	memcpy(&stVencChnAttr.stVeAttr.stAttrH264e, &stH264Attr, sizeof(VENC_ATTR_H264_S));

 	stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CBR;
    stH264Cbr.u32Gop            = FPS;//30*3;
    stH264Cbr.u32StatTime       = 1; 
    stH264Cbr.u32SrcFrmRate      = FPS;//30;// input (vi) frame rate 
    stH264Cbr.fr32DstFrmRate = FPS;//30;// target frame rate 
	stH264Cbr.u32BitRate = 1024*16;//1024*1;//1024*1;//30;
	stH264Cbr.u32FluctuateLevel = 1; // average bit rate 
	memcpy(&stVencChnAttr.stRcAttr.stAttrH264Cbr, &stH264Cbr, sizeof(VENC_ATTR_H264_CBR_S));

	stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
	stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 2;

	s32Ret = HI_MPI_VENC_CreateChn(0, &stVencChnAttr);
    if (HI_SUCCESS != s32Ret) {
        printf("HI_MPI_VENC_CreateChn [%d] faild with %#x!\n", 2, s32Ret);
        return s32Ret;
    }

    s32Ret = HI_MPI_VENC_StartRecvPic(0);
    if (HI_SUCCESS != s32Ret) {
        printf("HI_MPI_VENC_StartRecvPic faild with%#x!\n", s32Ret);
        return HI_FAILURE;
    }

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId = HI_ID_VPSS;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 0;

    stDestChn.enModId = HI_ID_VENC;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;

    s32Ret = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (s32Ret != HI_SUCCESS) {
        printf("failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }

	return 0;
}

int venc_init_1(void){
	HI_S32 			ret;

    HI_S32 s32Ret;
    VENC_CHN_ATTR_S stVencChnAttr;
    VENC_ATTR_H264_S stH264Attr;
    VENC_ATTR_H264_CBR_S    stH264Cbr;
	VENC_ATTR_H264_VBR_S    stH264Vbr;
	VENC_ATTR_H264_AVBR_S    stH264AVbr;
	VENC_ATTR_H264_FIXQP_S  stH264FixQp;

  	VENC_ATTR_H265_S        stH265Attr;
    VENC_ATTR_H265_CBR_S    stH265Cbr;
    VENC_ATTR_H265_VBR_S    stH265Vbr;
    VENC_ATTR_H265_FIXQP_S  stH265FixQp;

    stVencChnAttr.stVeAttr.enType = PT_H264;
    stH264Attr.u32MaxPicWidth = 2560;//4000;
    stH264Attr.u32MaxPicHeight = 720;//3000;
    stH264Attr.u32PicWidth = 2560;//3840;//4000;/*the picture width*/
    stH264Attr.u32PicHeight = 720;//2160;//3000;/*the picture height*/
    stH264Attr.u32BufSize  = 2560*720*1;//4000 * 3000 * 2;/*stream buffer size*/
    stH264Attr.u32Profile  = 2;/*0: baseline; 1:MP; 2:HP;  3:svc_t */
    stH264Attr.bByFrame = HI_TRUE;//HI_FALSE;//HI_TRUE;//HI_TRUE;/*get stream mode is slice mode or frame mode?*/
	//stH264Attr.u32BFrameNum = 0;/* 0: not support B frame; >=1: number of B frames */
	//stH264Attr.u32RefNum = 1;/* 0: default; number of refrence frame*/

	memcpy(&stVencChnAttr.stVeAttr.stAttrH264e, &stH264Attr, sizeof(VENC_ATTR_H264_S));

 	stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CBR;
    stH264Cbr.u32Gop            = FPS;//30*3;
    stH264Cbr.u32StatTime       = 1; 
    stH264Cbr.u32SrcFrmRate      = FPS;//30;// input (vi) frame rate 
    stH264Cbr.fr32DstFrmRate = FPS;//30;// target frame rate 
	stH264Cbr.u32BitRate = 1024*8;//1024*1;//1024*1;//30;
	stH264Cbr.u32FluctuateLevel = 1; // average bit rate 
	memcpy(&stVencChnAttr.stRcAttr.stAttrH264Cbr, &stH264Cbr, sizeof(VENC_ATTR_H264_CBR_S));

	stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
	stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 2;

	s32Ret = HI_MPI_VENC_CreateChn(1, &stVencChnAttr);
    if (HI_SUCCESS != s32Ret) {
        printf("HI_MPI_VENC_CreateChn [%d] faild with %#x!\n", 2, s32Ret);
        return s32Ret;
    }

    s32Ret = HI_MPI_VENC_StartRecvPic(1);
    if (HI_SUCCESS != s32Ret) {
        printf("HI_MPI_VENC_StartRecvPic faild with%#x!\n", s32Ret);
        return HI_FAILURE;
    }

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId = HI_ID_VPSS;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 1;

    stDestChn.enModId = HI_ID_VENC;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 1;

    s32Ret = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (s32Ret != HI_SUCCESS) {
        printf("failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }

	return 0;
}


/*
	Errors resolvers
*/
void resolve_mppv2_errors(HI_S32 error) { switch(error) {
	case 0xA0028003://HI_ERR_SYS_ILLEGAL_PARAM 
	fprintf(stderr, "\tThe parameter configuration is invalid.\n");
	break;
	case 0xA0028006://HI_ERR_SYS_NULL_PTR 
	fprintf(stderr, "\tThe pointer is null.\n");
	break;
	case 0xA0028009://HI_ERR_SYS_NOT_PERM 
	fprintf(stderr, "\tThe operation is forbidden.\n");
	break;
	case 0xA0028010://HI_ERR_SYS_NOTREADY 
	fprintf(stderr, "\tThe system control attributes are not configured.\n");
	break;
	case 0xA0028012://HI_ERR_SYS_BUSY 
	fprintf(stderr, "\tThe system is busy.\n");
	break;
	case 0xA002800C://HI_ERR_SYS_NOMEM 
	fprintf(stderr, "\tThe memory fails to be allocated due to some causes such as insufficient system memory.\n");
	break;
	case 0xA0018003://HI_ERR_VB_ILLEGAL_PARAM 
	fprintf(stderr, "\tThe parameter configuration is invalid.\n");
	break;
	case 0xA0018005://HI_ERR_VB_UNEXIST 
	fprintf(stderr, "\tThe VB pool does not exist.\n");
	break;
	case 0xA0018006://HI_ERR_VB_NULL_PTR 
	fprintf(stderr, "\tThe pointer is null.\n");
	break;
	case 0xA0018009://HI_ERR_VB_NOT_PERM 
	fprintf(stderr, "\tThe operation is forbidden.\n");
	break;
	case 0xA001800C://HI_ERR_VB_NOMEM 
	fprintf(stderr, "\tThe memory fails to be allocated.\n");
	break;
	case 0xA001800D://HI_ERR_VB_NOBUF 
	fprintf(stderr, "\tThe buffer fails to be allocated.\n");
	break;
	case 0xA0018010://HI_ERR_VB_NOTREADY 
	fprintf(stderr, "\tThe system control attributes are not configured.\n");
	break;
	case 0xA0018012://HI_ERR_VB_BUSY 
	fprintf(stderr, "\tThe system is busy.\n");
	break;
	case 0xA0018040://HI_ERR_VB_2MPOOLS 
	fprintf(stderr, "\tToo many VB pools are created.\n");
	break;
	case 0xA0108001://HI_ERR_VI_INVALID_DEVID 
	fprintf(stderr, "\tThe VI device ID is invalid.\n");
	break;
	case 0xA0108002://HI_ERR_VI_INVALID_CHNID 
	fprintf(stderr, "\tThe VI channel ID is invalid.\n");
	break;
	case 0xA0108003://HI_ERR_VI_INVALID_PARA 
	fprintf(stderr, "\tThe VI parameter is invalid.\n");
	break;
	case 0xA0108006://HI_ERR_VI_INVALID_NULL_PTR 
	fprintf(stderr, "\tThe pointer of the input parameter is null.\n");
	break;
	case 0xA0108007://HI_ERR_VI_FAILED_NOTCONFIG 
	fprintf(stderr, "\tThe attributes of the video device are not set.\n");
	break;
	case 0xA0108008://HI_ERR_VI_NOT_SUPPORT 
	fprintf(stderr, "\tThe operation is not supported.\n");
	break;
	case 0xA0108009://HI_ERR_VI_NOT_PERM 
	fprintf(stderr, "\tThe operation is forbidden.\n");
	break;
	case 0xA010800C://HI_ERR_VI_NOMEM 
	fprintf(stderr, "\tThe memory fails to be allocated.\n");
	break;
	case 0xA010800E://HI_ERR_VI_BUF_EMPTY 
	fprintf(stderr, "\tThe VI buffer is empty.\n");
	break;
	case 0xA010800F://HI_ERR_VI_BUF_FULL 
	fprintf(stderr, "\tThe VI buffer is full.\n");
	break;
	case 0xA0108010://HI_ERR_VI_SYS_NOTREADY 
	fprintf(stderr, "\tThe VI system is not initialized.\n");
	break;
	case 0xA0108012://HI_ERR_VI_BUSY 
	fprintf(stderr, "\tThe VI system is busy.\n");
	break;
	case 0xA0108040://HI_ERR_VI_FAILED_NOTENABLE 
	fprintf(stderr, "\tThe VI device or VI channel is not enabled.\n");
	break;
	case 0xA0108041://HI_ERR_VI_FAILED_NOTDISABLE 
	fprintf(stderr, "\tThe VI device or VI channel is not disabled.\n");
	break;
	case 0xA0108042://HI_ERR_VI_FAILED_CHNOTDISABLE 
	fprintf(stderr, "\tThe VI channel is not disabled.\n");
	break;
	case 0xA0108043://HI_ERR_VI_CFG_TIMEOUT 
	fprintf(stderr, "\tThe video attribute configuration times out.\n");
	break;
	case 0xA0108045://HI_ERR_VI_INVALID_WAYID 
	fprintf(stderr, "\tThe video channel ID is invalid.\n");
	break;
	case 0xA0108046://HI_ERR_VI_INVALID_PHYCHNID 
	fprintf(stderr, "\tThe physical video channel ID is invalid.\n");
	break;
	case 0xA0108047://HI_ERR_VI_FAILED_NOTBIND 
	fprintf(stderr, "\tThe video channel is not bound.\n");
	break;
	case 0xA0108048://HI_ERR_VI_FAILED_BINDED 
	fprintf(stderr, "\tThe video channel is bound.\n");
	break;
	case 0xA0078001://HI_ERR_VPSS_INVALID_DEVID 
	fprintf(stderr, "\tThe VPSS group ID is invalid.\n");
	break;
	case 0xA0078002://HI_ERR_VPSS_INVALID_CHNID 
	fprintf(stderr, "\tThe VPSS channel ID is invalid.\n");
	break;
	case 0xA0078003://HI_ERR_VPSS_ILLEGAL_PARAM 
	fprintf(stderr, "\tThe VPSS parameter is invalid.\n");
	break;
	case 0xA0078004://HI_ERR_VPSS_EXIST 
	fprintf(stderr, "\tA VPSS group is created.\n");
	break;
	case 0xA0078005://HI_ERR_VPSS_UNEXIST 
	fprintf(stderr, "\tNo VPSS group is created.\n");
	break;
	case 0xA0078006://HI_ERR_VPSS_NULL_PTR 
	fprintf(stderr, "\tThe pointer of the input parameter is null.\n");
	break;
	case 0xA0078008://HI_ERR_VPSS_NOT_SUPPORT 
	fprintf(stderr, "\tThe operation is not supported.\n");
	break;
	case 0xA0078009://HI_ERR_VPSS_NOT_PERM 
	fprintf(stderr, "\tThe operation is forbidden.\n");
	break;
	case 0xA007800C://HI_ERR_VPSS_NOMEM 
	fprintf(stderr, "\tThe memory fails to be allocated.\n");
	break;
	case 0xA007800D://HI_ERR_VPSS_NOBUF 
	fprintf(stderr, "\tThe buffer pool fails to be allocated.\n");
	break;
	case 0xA007800E://HI_ERR_VPSS_BUF_EMPTY 
	fprintf(stderr, "\tThe picture queue is empty.\n");
	break;
	case 0xA0078010://HI_ERR_VPSS_NOTREADY 
	fprintf(stderr, "\tThe VPSS is not initialized.\n");
	break;
	case 0xA0078012://HI_ERR_VPSS_BUSY 
	fprintf(stderr, "\tThe VPSS is busy.\n");
	break;
	case 0xA0088002://HI_ERR_VENC_INVALID_CHNID 
	fprintf(stderr, "\tThe channel ID is invalid.\n");
	break;
	case 0xA0088003://HI_ERR_VENC_ILLEGAL_PARAM 
	fprintf(stderr, "\tThe parameter is invalid.\n");
	break;
	case 0xA0088004://HI_ERR_VENC_EXIST 
	fprintf(stderr, "\tThe device, channel or resource to be created or applied for exists.\n");
	break;
	case 0xA0088005://HI_ERR_VENC_UNEXIST 
	fprintf(stderr, "\tThe device, channel or resource to be used or destroyed does not exist.\n");
	break;
	case 0xA0088006://HI_ERR_VENC_NULL_PTR 
	fprintf(stderr, "\tThe parameter pointer is null.\n");
	break;
	case 0xA0088007://HI_ERR_VENC_NOT_CONFIG 
	fprintf(stderr, "\tNo parameter is set before use.\n");
	break;
	case 0xA0088008://HI_ERR_VENC_NOT_SUPPORT 
	fprintf(stderr, "\tThe parameter or function is not supported.\n");
	break;
	case 0xA0088009://HI_ERR_VENC_NOT_PERM 
	fprintf(stderr, "\tThe operation, for example, modifying static parameters, is forbidden.\n");
	break;
	case 0xA008800C://HI_ERR_VENC_NOMEM 
	fprintf(stderr, "\tThe memory fails to be allocated due to some causes such as insufficient system memory.\n");
	break;
	case 0xA008800D://HI_ERR_VENC_NOBUF 
	fprintf(stderr, "\tThe buffer fails to be allocated due to some causes such as oversize of the data buffer applied for.\n");
	break;
	case 0xA008800E://HI_ERR_VENC_BUF_EMPTY 
	fprintf(stderr, "\tThe buffer is empty.\n");
	break;
	case 0xA008800F://HI_ERR_VENC_BUF_FULL 
	fprintf(stderr, "\tThe buffer is full.\n");
	break;
	case 0xA0088010://HI_ERR_VENC_SYS_NOTREADY 
	fprintf(stderr, "\tThe system is not initialized or the corresponding module is not loaded.\n");
	break;
	case 0xA0088012://HI_ERR_VENC_BUSY 
	fprintf(stderr, "\tThe VENC system is busy.\n");
	break;
	case 0xA01C8006://HI_ERR_ISP_NULL_PTR 
	fprintf(stderr, "\tThe input pointer is null.\n");
	break;
	case 0xA01C8003://HI_ERR_ISP_ILLEGAL_PARAM 
	fprintf(stderr, "\tThe input parameter is invalid.\n");
	break;
	case 0xA01C8043://HI_ERR_ISP_SNS_UNREGISTER 
	fprintf(stderr, "\tThe sensor is not registered.\n");
	break;
	case 0xA01C0041://HI_ERR_ISP_MEM_NOT_INIT 
	fprintf(stderr, "\tThe external registers are not initialized.\n");
	break;
	case 0xA01C0040://HI_ERR_ISP_NOT_INIT 
	fprintf(stderr, "\tThe ISP is not initialized.\n");
	break;
	case 0xA01C0044://HI_ERR_ISP_INVALID_ADDR 
	fprintf(stderr, "\tThe address is invalid.\n");
	break;
	case 0xA01C0042://HI_ERR_ISP_ATTR_NOT_CFG 
	fprintf(stderr, "\tThe attribute is not configured.\n");
	break;
	default:
	fprintf(stderr, "\tUnrecognized error!\n");
}}

//////////
/*
combo_dev_attr_t LVDS_4lane_SENSOR_IMX178_12BIT_5M_NOWDR_ATTR =
{
    .input_mode = INPUT_MODE_LVDS,
    {
        .lvds_attr = {
            .img_size = {2592, 1944},
            HI_WDR_MODE_NONE,
            LVDS_SYNC_MODE_SAV,
            RAW_DATA_12BIT,
            LVDS_ENDIAN_BIG,
            LVDS_ENDIAN_BIG,
            .lane_id = {0, 1, 2, 3, -1, -1, -1, -1},
            .sync_code = { 
                    {{0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},
                    
                    {{0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},
                    
                    {{0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},
                    
                    {{0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},
                        
                    {{0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},
                    
                    {{0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}} 
                }
        }
    }
};
*/


