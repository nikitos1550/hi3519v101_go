//+build hi3516cv100

//g-o:generate rm -f kobin_hi3516cv100.go
//g-o:generate go run -tags "generate hi3516cv100" ./generate.go --output kobin_hi3516cv100.go --tag hi3516cv100 --dir ../../sdk/hi3516cv100/ko/ --pkg ko --source ./hi3516cv100.go

package ko

var (
	ModulesList = [...][2]string{
		[2]string{"mmz.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
		[2]string{"hi3518_base.ko", ""},
		[2]string{"hi3518_sys.ko", ""},
		[2]string{"hiuser.ko", ""},
		[2]string{"hi3518_tde.ko", ""},
		[2]string{"hi3518_dsu.ko", ""},
		[2]string{"hi3518_viu.ko", ""},
		[2]string{"hi3518_isp.ko", ""},
		[2]string{"hi3518_vpss.ko", ""},
		[2]string{"hi3518_vou.ko", ""},
		//[2]string{"hi3518_vou.ko", "detectCycle=0"}, //close dac detect
		//[2]string{"hifb.ko", 'video="hifb:vram0_size:1620"'},
		[2]string{"hi3518_venc.ko", ""},
		[2]string{"hi3518_group.ko", ""},
		[2]string{"hi3518_chnl.ko", ""},
		[2]string{"hi3518_h264e.ko", ""},
		[2]string{"hi3518_jpege.ko", ""},
		[2]string{"hi3518_rc.ko", ""},
		[2]string{"hi3518_region.ko", ""},
		[2]string{"hi3518_vda.ko", ""},
		[2]string{"hi3518_ive.ko", ""},
		[2]string{"hi_i2c.ko", ""},
		//[2]string{"gpioi2c.ko", ""},
		//[2]string{"gpioi2c_ex.ko", ""},
		[2]string{"pwm.ko", ""},
		//[2]string{"sil9024.ko", "norm=5"}, //720P@60fps
		//[2]string{"hi3518_isp.ko", ""},
		[2]string{"acodec.ko", ""},
		[2]string{"hidmac.ko", ""},
		[2]string{"hi3518_sio.ko", ""},
		[2]string{"hi3518_ai.ko", ""},
		[2]string{"hi3518_ao.ko", ""},
		[2]string{"hi3518_aenc.ko", ""},
		[2]string{"hi3518_adec.ko", ""},
		[2]string{"ssp_sony.ko", ""},
	}

	MinimalModulesList = [...]string{
		"mmz.ko",
		"hi3518_base.ko",
		"hi3518_sys.ko",
	}
)
