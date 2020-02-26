// +build hi3516cv200

//go:generate rm -f hi3516cv200_kobin.go
//go:generate go run -tags "generate hi3516cv200" ./generate.go --output hi3516cv200_kobin.go --tag hi3516cv200 --dir ../../sdk/hi3516cv200/ko --pkg ko --source ./hi3516cv200.go

package ko

var (
	ModulesList = [...][2]string{
		[2]string{"mmz.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
		[2]string{"hi_media.ko", ""},
		[2]string{"hi3518e_base.ko", ""},
		[2]string{"hi3518e_sys.ko", "vi_vpss_online=0 sensor=NULL"},
	}

	minimalModulesList = [...]string{
		"mmz.ko",
		"hi_media.ko",
		"hi3518e_base.ko",
		"hi3518e_sys.ko",
	}
)
