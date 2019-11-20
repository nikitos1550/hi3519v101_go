// +build hi3516cv200

package koloader

var (

    modules = [...][2]string {
        [2]string{"mmz.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
        [2]string{"hi_media.ko", ""},
        [2]string{"hi3518e_base.ko", ""},
        [2]string{"hi3518e_sys.ko", "vi_vpss_online=0 sensor=NULL"},
    }
)

