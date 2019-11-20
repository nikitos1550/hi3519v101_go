// +build hi3516av100

package koloader

var (

    modules = [...][2]string {
        [2]string{"mmz.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
        [2]string{"hi_media.ko", ""},
        [2]string{"hi3516a_base.ko", ""},
        [2]string{"hi3516a_sys.ko", "vi_vpss_online=0 sensor=NULL"},
    }
)

