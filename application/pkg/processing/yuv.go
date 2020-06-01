//+build processing

package processing

/*
#include "processing.h"

int getData(void* frame, void** out) {
	VIDEO_FRAME_INFO_S* stFrame = frame;

	char* data = (char*)malloc(stFrame->stVFrame.u32Width * stFrame->stVFrame.u32Height * 3 / 2);

	int ySize = stFrame->stVFrame.u32Width * stFrame->stVFrame.u32Height;
	HI_U8* pUserPageAddr = (HI_U8*)HI_MPI_SYS_Mmap(stFrame->stVFrame.u32PhyAddr[0], ySize);
	if (NULL == pUserPageAddr) {
		return 0;
	}
	memcpy(data, pUserPageAddr, ySize);
	HI_MPI_SYS_Munmap(pUserPageAddr, ySize);

	int uvSize = stFrame->stVFrame.u32Width * stFrame->stVFrame.u32Height / 4;
	pUserPageAddr = (HI_U8*)HI_MPI_SYS_Mmap(stFrame->stVFrame.u32PhyAddr[1], uvSize);
	if (NULL == pUserPageAddr) {
		return 0;
	}
	memcpy(data + ySize, pUserPageAddr, uvSize);
	HI_MPI_SYS_Munmap(pUserPageAddr, uvSize);

	pUserPageAddr = (HI_U8*)HI_MPI_SYS_Mmap(stFrame->stVFrame.u32PhyAddr[2], uvSize);
	if (NULL == pUserPageAddr) {
		return 0;
	}
	memcpy((data + ySize + uvSize), pUserPageAddr, uvSize);
	HI_MPI_SYS_Munmap(pUserPageAddr, uvSize);

	*out = data;
	return ySize + 2 * uvSize;
}

void releaseData(void* data) {
	free(data);
}
*/
import "C"

import (
	"unsafe"
	"application/pkg/common"
)

type yuv struct {
	Name string
	Id int
}

func (p yuv) GetName() string {
	return p.Name
}

func (p yuv) GetId() int {
	return p.Id
}

func (p yuv) Create(id int, params map[string][]string) (common.Processing,int,string) {
	var v yuv
	v.Name = "yuv"
	v.Id = id
	return v,id,""
}

func (p yuv) Destroy() {
}

func (p yuv) Callback(data unsafe.Pointer) {
	var out unsafe.Pointer
	size := int(C.getData(data, &out))
	
	dest := make([]byte, size)
    copy(dest, (*(*[1024]byte)(unsafe.Pointer(out)))[:size:size])
	
	C.releaseData(out)

	sendDataToEncoders(p.Id, dest)
}

func init() {
	var v yuv
	v.Name = "yuv"
	v.Id = -1
	register(v)
}


//http://213.141.129.12:8080/cam1/api/mpp/channel/start?channelId=1&width=1920&height=1080&fps=30
//http://213.141.129.12:8080/cam1/api/processing/create?processingName=yuv
//http://213.141.129.12:8080/cam1/api/encoder/create_dummy
//http://213.141.129.12:8080/cam1/api/processing/subscribeChannel?processingId=1&channelId=1
//http://213.141.129.12:8080/cam1/api/encoder/subscribeProcessing?processingId=1&encoderId=1