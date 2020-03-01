// +build hi3516cv300

//go:generate rm -f hi3516cv300_kobin.go
//go:generate go run -tags "generate hi3516cv300" ./generate.go --output hi3516cv300_kobin.go --tag hi3516cv300 --dir ../../sdk/hi3516cv300/ko/ --pkg ko --source ./hi3516cv300.go

package ko

var (
	ModulesList = [...][2]string{
		[2]string{"sys_config.ko", "vi_vpss_online=0"}, //TODO!!!!!!!!!
		[2]string{"hi_osal.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
		[2]string{"hi3516cv300_base.ko", ""},
		[2]string{"hi3516cv300_sys.ko", "vi_vpss_online=0 sensor=NULL mem_total={memTotalSize}"},
		[2]string{"hi3516cv300_region.ko", ""},
		[2]string{"hi3516cv300_vgs.ko", "vgs_clk_frequency=$vgs_frequency"},
		[2]string{"hi3516cv300_viu.ko", "detect_err_frame=10 viu_clk_frequency=$viu_frequency isp_div=$isp_div input_mode=$intf_mode"},
		[2]string{"hi3516cv300_isp.ko", "update_pos=0  proc_param=30 port_init_delay=0"},
		[2]string{"hi3516cv300_vpss.ko", "vpss_clk_frequency=$vpss_frequency"},
		[2]string{"hi3516cv300_vou.ko", "vou_mode=$vou_intf_mode"},
		//    #insmod hi3516cv300_vou.ko detectCycle=0 vou_mode=$vou_intf_mode #close dac detect
		//    #insmod hi3516cv300_vou.ko transparentTransmit=1 vou_mode=$vou_intf_mode #enable transparentTransmit
		[2]string{"hi3516cv300_rc.ko", ""},
		[2]string{"hi3516cv300_venc.ko", ""},
		[2]string{"hi3516cv300_chnl.ko", ""},
		[2]string{"hi3516cv300_vedu.ko", "vedu_clk_frequency=$vedu_frequency"},
		[2]string{"hi3516cv300_h264e.ko", ""},
		[2]string{"hi3516cv300_h265e.ko", ""},
		[2]string{"hi3516cv300_jpege.ko", ""},
		[2]string{"hi3516cv300_ive.ko", "save_power=1 ive_clk_frequency=$ive_frequency"},
		[2]string{"hi3516cv300_sensor.ko", "sensor_bus_type=$bus_type sensor_clk_frequency=$sensor_clk_freq sensor_pinmux_mode=$pinmux_mode"},
		[2]string{"hi3516cv300_pwm.ko", ""},
		[2]string{"hi_piris.ko", ""},
		[2]string{"hi3516cv300_aio.ko", ""},
		[2]string{"hi3516cv300_ai.ko", ""},
		[2]string{"hi3516cv300_ao.ko", ""},
		[2]string{"hi3516cv300_aenc.ko", ""},
		[2]string{"hi3516cv300_adec.ko", ""},
		[2]string{"hi_acodec.ko", ""},
		[2]string{"hi_mipi.ko", ""},
	}

	minimalModulesList = [...]string{
		"sys_config.ko",
		"hi_osal.ko",
		"hi3516cv300_base.ko",
		"hi3516cv300_sys.ko",
	}
)

/*
insert_ko()
{
    //cv300
    insert_sns
    insmod sys_config.ko vi_vpss_online=$b_arg_online

    # driver load
    insmod hi_osal.ko mmz=anonymous,0,$mmz_start,$mmz_size anony=1 || report_error
    insmod hi3516cv300_base.ko

    insmod hi3516cv300_sys.ko vi_vpss_online=$b_arg_online sensor=$SNS_TYPE mem_total=$mem_total

    insmod hi3516cv300_region.ko
    insmod hi3516cv300_vgs.ko vgs_clk_frequency=$vgs_frequency

    insmod hi3516cv300_viu.ko detect_err_frame=10 viu_clk_frequency=$viu_frequency isp_div=$isp_div input_mode=$intf_mode
    insert_isp;
    insmod hi3516cv300_vpss.ko vpss_clk_frequency=$vpss_frequency
    insmod hi3516cv300_vou.ko vou_mode=$vou_intf_mode
    #insmod hi3516cv300_vou.ko detectCycle=0 vou_mode=$vou_intf_mode #close dac detect
    #insmod hi3516cv300_vou.ko transparentTransmit=1 vou_mode=$vou_intf_mode #enable transparentTransmit

    insmod hi3516cv300_rc.ko
    insmod hi3516cv300_venc.ko
    insmod hi3516cv300_chnl.ko
    insmod hi3516cv300_vedu.ko vedu_clk_frequency=$vedu_frequency
    insmod hi3516cv300_h264e.ko
    insmod hi3516cv300_h265e.ko
    insmod hi3516cv300_jpege.ko
    insmod hi3516cv300_ive.ko save_power=1 ive_clk_frequency=$ive_frequency
    insmod hi3516cv300_sensor.ko sensor_bus_type=$bus_type sensor_clk_frequency=$sensor_clk_freq sensor_pinmux_mode=$pinmux_mode
    insmod hi3516cv300_pwm.ko

    insmod extdrv/hi_piris.ko
    insert_audio

    insmod hi_mipi.ko
*/
