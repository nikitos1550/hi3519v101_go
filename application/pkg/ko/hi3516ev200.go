// +build hi3516ev200

//g-o:generate rm -f kobin_hi3516ev200.go
//g-o:generate go run -tags "generate hi3516ev200" ./generate.go --output kobin_hi3516ev200.go --tag hi3516ev200 --dir ../../sdk/hi3516ev200/ko/ --pkg ko --source ./hi3516ev200.go

package ko

/*
YUV_TYPE0=0;                # 0 -- raw, 1 --DC, 3 --bt656
CHIP_TYPE=hi3516ev200;      # chip type
*/

var (
        ModulesList = [...][2]string{
                [2]string{"sys_config.ko", "chip={chip} sensors={cmos} g_cmos_yuv_flag={g_cmos_yuv_flag} board={board}"}, //TODO!!!!!!!!!!!!!!!!
                [2]string{"hi_osal.ko", "anony=1 mmz_allocator=hisi mmz=anonymous,0,{mem_start_addr},{mem_mpp_size}"},
                [2]string{"hi3516ev200_base.ko", ""},
                [2]string{"hi3516ev200_sys.ko", ""},
        	[2]string{"hi3516ev200_tde.ko", ""},
        	[2]string{"hi3516ev200_rgn.ko", ""},
        	[2]string{"hi3516ev200_vgs.ko", ""},
		[2]string{"hi3516ev200_vi.ko", ""},
		[2]string{"hi3516ev200_isp.ko", ""},
		[2]string{"hi3516ev200_vpss.ko", ""},
        	[2]string{"hi3516ev200_vo.ko", ""},
		//[2]string{"hifb.ko video="hifb:vram0_size:1620"     # default fb0:D1
        	[2]string{"hi3516ev200_chnl.ko", ""},
        	[2]string{"hi3516ev200_vedu.ko", ""},
        	[2]string{"hi3516ev200_rc.ko", ""},
        	[2]string{"hi3516ev200_venc.ko", ""},
        	[2]string{"hi3516ev200_h264e.ko", ""},
        	[2]string{"hi3516ev200_h265e.ko", ""},
        	[2]string{"hi3516ev200_jpege.ko", ""},
        	[2]string{"hi3516ev200_ive.ko", "save_power={save_power}"},
        	[2]string{"hi_pwm.ko", ""},
        	[2]string{"hi_sensor_i2c.ko", ""},
        	[2]string{"hi_sensor_spi.ko", ""},
        	[2]string{"hi3516ev200_aio.ko", ""},
        	[2]string{"hi3516ev200_ai.ko", ""},
        	[2]string{"hi3516ev200_ao.ko", ""},
        	[2]string{"hi3516ev200_aenc.ko", ""},
        	[2]string{"hi3516ev200_adec.ko", ""},
        	[2]string{"hi3516ev200_acodec.ko", ""},
		//[2]string{"hi_tlv320aic31.ko", ""},
	        [2]string{"hi_mipi_rx.ko", ""},
		//[2]string{"hi_user.ko", ""},
	}

        MinimalModulesList = [...]string{
                "sys_config.ko",
                "hi_osal.ko",
                "hi3516ev200_base.ko",
                "hi3516ev200_sys.ko",
        }

)

/*
 insmod sys_config.ko chip=$CHIP_TYPE sensors=$SNS_TYPE0 g_cmos_yuv_flag=$YUV_TYPE0 board=$BOARD
        insmod hi_osal.ko anony=1 mmz_allocator=hisi mmz=anonymous,0,$mmz_start,$mmz_size || report_error
        insmod hi3516ev200_base.ko
        insmod hi3516ev200_sys.ko
        insmod hi3516ev200_tde.ko
        insmod hi3516ev200_rgn.ko
        insmod hi3516ev200_vgs.ko
        insmod hi3516ev200_vi.ko
        insert_isp;


insert_isp()
{
        insmod hi3516ev200_isp.ko
}



        insmod hi3516ev200_vpss.ko
        insmod hi3516ev200_vo.ko
        insmod hifb.ko video="hifb:vram0_size:1620"     # default fb0:D1
        insmod hi3516ev200_chnl.ko
        insmod hi3516ev200_vedu.ko
        insmod hi3516ev200_rc.ko
        insmod hi3516ev200_venc.ko
        insmod hi3516ev200_h264e.ko
        insmod hi3516ev200_h265e.ko
        insmod hi3516ev200_jpege.ko
        insmod hi3516ev200_ive.ko save_power=0
        insmod extdrv/hi_pwm.ko
        insmod extdrv/hi_sensor_i2c.ko
        insmod extdrv/hi_sensor_spi.ko
        insert_sil9024; # BT1120
#       insert_adv7179; # BT656
        insert_audio

insert_audio()
{
        insmod hi3516ev200_aio.ko
        insmod hi3516ev200_ai.ko 
        insmod hi3516ev200_ao.ko
        insmod hi3516ev200_aenc.ko
        insmod hi3516ev200_adec.ko
        insmod hi3516ev200_acodec.ko
#       insmod extdrv/hi_tlv320aic31.ko
        echo "insert audio"
}


        insmod hi_mipi_rx.ko
#       insmod hi_user.ko
*/
