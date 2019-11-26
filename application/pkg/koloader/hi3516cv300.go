// +build hi3516cv300

//go:generate rm -f hi3516cv300_kobin.go
//go:generate go run -tags "generate hi3516cv300" ./generate.go --output hi3516cv300_kobin.go --tag hi3516cv300 --dir ../../sdk/hi3516cv300/ko --pkg koloader --source ./hi3516cv300.go

package koloader

var (

    ModulesList = [...][2]string {
        [2]string{"sys_config.ko", "vi_vpss_online=0"},
        [2]string{"hi_osal.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
        [2]string{"hi3516cv300_base.ko", ""},
        [2]string{"hi3516cv300_sys.ko", "vi_vpss_online=0 sensor=NULL mem_total={memTotalSize}"},
    }
)

