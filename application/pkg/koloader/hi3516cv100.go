// +build hi3516cv100

package koloader

var (

    modules = [...][2]string {
        [2]string{"mmz.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
        [2]string{"hi3518_base.ko", ""},
        [2]string{"hi3518_sys.ko", ""},
    }
)

