package mpp

import (
	"application/pkg/mpp/sys"

	//"application/pkg/mpp/cmos"
	"application/pkg/mpp/isp"
	"application/pkg/mpp/mipi"
	"application/pkg/mpp/venc"
	"application/pkg/mpp/vi"
	"application/pkg/mpp/vpss"
)

func Init() {
	//cmos.Test() //TEST

	//panic("TEST")

	systemInit()
	//
	//
	sys.Init()
	mipi.Init()
	isp.Init()
	vi.Init()
	vpss.Init()

	venc.Init()

	//init sample videopipeline
	vpss.SampleChannel0()

}
