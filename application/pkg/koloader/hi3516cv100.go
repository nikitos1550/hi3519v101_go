// +build hi3516cv100

//go:generate rm -f hi3516cv100_kobin.go
//go:generate go run -tags "generate hi3516cv100" ./generate.go --output hi3516cv100_kobin.go --tag hi3516cv100 --dir ../../sdk/hi3516cv100/ko --pkg koloader --source ./hi3516cv100.go

package koloader

var (

    Modules = [...][2]string {
        [2]string{"mmz.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
        [2]string{"hi3518_base.ko", ""},
        [2]string{"hi3518_sys.ko", ""},
    }
)

