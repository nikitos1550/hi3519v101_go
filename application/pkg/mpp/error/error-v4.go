//+build debug
//+build hi3516cv500 hi3516ev200 hi3519av100 hi3559av100

package error

import (
	"strconv"
)

func Resolve(code int64) string {
	switch code {

	default:
		out := "unknown error " + strconv.FormatInt(code, 16)
		return out
	}
}


/*

0xA0028003 HI_ERR_SYS_ILLEGAL_PARAM The parameter configuration is invalid.
0xA0028006 HI_ERR_SYS_NULL_PTR The pointer is null.
0xA0028008 HI_ERR_SYS_NOT_SUPPORT The function is not supported.
0xA0028009 HI_ERR_SYS_NOT_PERM The operation is forbidden.
0xA0028010 HI_ERR_SYS_NOTREADY The system control attributes are not configured.
0xA0028012 HI_ERR_SYS_BUSY The system is busy.
0xA002800C HI_ERR_SYS_NOMEM The memory fails to be allocated due to some causes such as insufficient system memory.

0xA0018003 HI_ERR_VB_ILLEGAL_PARAM The parameter configuration is invalid.
0xA0018005 HI_ERR_VB_UNEXIST The VB pool does not exist.
0xA0018006 HI_ERR_VB_NULL_PTR The pointer is null.
0xA0018009 HI_ERR_VB_NOT_PERM The operation is forbidden.
0xA001800C HI_ERR_VB_NOMEM The memory fails to be allocated.
0xA001800D HI_ERR_VB_NOBUF The buffer fails to be allocated.
0xA0018010 HI_ERR_VB_NOTREADY The system control attributes are not configured.
0xA0018012 HI_ERR_VB_BUSY The system is busy.
0xA0018013 HI_ERR_VB_SIZE_NOT_ENOUGH The VB block size is too small.

0xA0018040 HI_ERR_VB_2MPOOLS Too many VB pools are created.

*/
