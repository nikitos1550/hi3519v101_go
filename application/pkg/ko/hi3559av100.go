// +build hi3559av100

package ko

var (
        ModulesList = [...][2]string{
		[2]string{"sys_config.ko", "g_online_flag=0 sensors=sns0=imx334,sns1=imx334,sns2=imx334,sns3=imx334,sns4=imx334,sns5=imx334,sns6=imx334,sns7=imx334"},
		[2]string{"hi_osal.ko", "anony=1 mmz_allocator=hisi mmz=anonymous,0,{mem_start_addr},{mem_mpp_size}"},
        	[2]string{"hi3559av100_base.ko", ""},
        	[2]string{"hi3559av100_sys.ko", ""},
	}

        MinimalModulesList = [...]string{
		"sys_config.ko",
        	"hi_osal.ko",
        	"hi3559av100_base.ko",
        	"hi3559av100_sys.ko",
	}

)

