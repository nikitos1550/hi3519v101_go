// +build hi3516av100

//go:generate rm -f hi3516av100_kobin.go
//go:generate go run -tags "generate hi3516av100" ./generate.go --output hi3516av100_kobin.go --tag hi3516av100 --dir ../../sdk/hi3516av100/ko --pkg koloader --source ./hi3516av100.go

package koloader

var (
    ModulesList = [...][2]string {
        [2]string{"mmz.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
        [2]string{"hi_media.ko", ""},
        [2]string{"hi3516a_base.ko", ""},
        [2]string{"hi3516a_sys.ko", "vi_vpss_online=0 sensor=NULL"},
    }

    minimalModulesList = [...]string {
        "mmz.ko",
        "hi_media.ko",
        "hi3516a_base.ko",
        "hi3516a_sys.ko",
    }
)

