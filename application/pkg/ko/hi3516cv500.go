// +build hi3516cv500

//go:generate rm -f kobin_hi3516cv500.go
//go:generate go run -tags "generate hi3516cv500" ./generate.go --output kobin_hi3516cv500.go --tag hi3516cv500 --dir ../../sdk/hi3516cv500/ko/ --pkg ko --source ./hi3516cv500.go

package ko

var (
	ModulesList = [...][2]string{
		[2]string{"sys_config.ko", "chip={chipName} sensors=sns0=NULL,sns1=NULL, g_cmos_yuv_flag=0"}, //TODO!!!!!!!!!!!!!!!!
		[2]string{"hi_osal.ko", "anony=1 mmz_allocator=hisi mmz=anonymous,0,0x{memStartAddr},{memMppSize}M"},
		[2]string{"hi3516cv500_base.ko", ""},
		[2]string{"hi3516cv500_sys.ko", ""},
	}

	MinimalModulesList = [...]string{
		"sys_config.ko",
		"hi_osal.ko",
		"hi3516cv500_base.ko",
		"hi3516cv500_sys.ko",
	}
)
