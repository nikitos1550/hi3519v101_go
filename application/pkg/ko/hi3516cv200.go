// +build hi3516cv200

//go:generate rm -f kobin_hi3516cv200.go
//go:generate go run -tags "generate hi3516cv200" ./generate.go --output kobin_hi3516cv200.go --tag hi3516cv200 --dir ../../sdk/hi3516cv200/ko/ --pkg ko --source ./hi3516cv200.go

package ko

var (
	ModulesList = [...][2]string{
		[2]string{"sys_config.ko", "vi_vpss_online=$b_arg_online sensor=$SNS_TYPE pin_mux_select=0"}, //TODO!!!!!!!!!!!!!!!!!!!!!

		[2]string{"mmz.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
		[2]string{"hi_media.ko", ""},
		[2]string{"hi3518e_base.ko", ""},
		[2]string{"hi3518e_sys.ko", "vi_vpss_online=0 sensor=NULL"},

		[2]string{"hi3518e_tde.ko", ""},
		[2]string{"hi3518e_region.ko", ""},
		[2]string{"hi3518e_vgs.ko", ""},
		//[2]string{"hi3518e_isp.ko", "update_pos=1"}, //ov9750
		[2]string{"hi3518e_isp.ko", "update_pos=0 proc_param=1"}, //rest cmoses
		[2]string{"hi3518e_viu.ko", "detect_err_frame=10"},
		[2]string{"hi3518e_vpss.ko", "rfr_frame_comp=1"},
		[2]string{"hi3518e_vou.ko", ""},
		//[2]string{"hi3518e_vou.ko", "transparentTransmit=1"}, //enable transparentTransmit
		//[2]string{"hifb.ko", 'video="hifb:vram0_size:1620"'}, //default pal
		[2]string{"hi3518e_rc.ko", ""},
		[2]string{"hi3518e_venc.ko", ""},
		[2]string{"hi3518e_chnl.ko", "ChnlLowPower=1"},
		[2]string{"hi3518e_h264e.ko", "H264eMiniBufMode=1"},
		[2]string{"hi3518e_jpege.ko", ""},
		[2]string{"hi3518e_ive.ko", "save_power=0"},
		//[2]string{"hi3518e_ive.ko", ""},
		[2]string{"sensor_i2c.ko", ""},
		[2]string{"pwm.ko", ""},
		[2]string{"piris.ko", ""},

		[2]string{"acodec.ko", ""},
		[2]string{"hi3518e_aio.ko", ""},
		[2]string{"hi3518e_ai.ko", ""},
		[2]string{"hi3518e_ao.ko", ""},
		[2]string{"hi3518e_aenc.ko", ""},
		[2]string{"hi3518e_adec.ko", ""},
		[2]string{"hi_mipi.ko", ""},
	}

	minimalModulesList = [...]string{
		"mmz.ko",
		"hi_media.ko",
		"hi3518e_base.ko",
		"hi3518e_sys.ko",
	}
)
