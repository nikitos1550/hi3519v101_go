package himpp3

// #include "himpp3.h"
import "C"


func himpp3sysinit() {
	C.himpp3_sys_init()
}

func himpp3_vi_init() {
	C.himpp3_vi_init()
}

func himpp3_mipi_init() {
	C.himpp3_mipi_init()
}

func himpp3_isp_init() {
	C.himpp3_isp_init
}

func himpp3_vpss_init() {
	C.himpp3_vpss_init()
}

func himpp3_venc_init() {
	C.himpp3_venc_init()
}


