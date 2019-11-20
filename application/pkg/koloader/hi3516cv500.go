// +build hi3516cv500

package koloader

var (

    modules = [...][2]string {
        [2]string{"sys_config.ko", "chip={chipName} sensors=sns0=NULL,sns1=NULL, g_cmos_yuv_flag=0"},
        [2]string{"hi_osal.ko", "anony=1 mmz_allocator=hisi mmz=anonymous,0,0x{memStartAddr},{memMppSize}M"},
        [2]string{"hi3516cv500_base.ko", ""},
        [2]string{"hi3516cv500_sys.ko", ""},
    }
)

