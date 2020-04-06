//+build arm
//+build hi3516av100
//+build debug

package error

import (
	"strconv"
)

//TODO add rest codes
func Resolve(code int64) string {
	switch code {
	case 0xA0028003:
		return "HI_ERR_SYS_ILLEGAL_PARAM (The parameter configuration is invalid)"
	case 0xA0028006:
		return "HI_ERR_SYS_NULL_PTR (The pointer is null)"
	case 0xA0028009:
		return "HI_ERR_SYS_NOT_PERM (The operation is forbidden)"
	case 0xA0028010: 
		return "HI_ERR_SYS_NOTREADY (The system control attributes are not configured)"
	case 0xA0028012:
		return "HI_ERR_SYS_BUSY (The system is busy)"
	case 0xA002800C:
		return "HI_ERR_SYS_NOMEM (The memory fails to be allocated due to some causes such as insufficient system memory)"
	case 0xA0018003: 
		return "HI_ERR_VB_ILLEGAL_PARAM (The parameter configuration is invalid)"
	case 0xA0018005: 
		return "HI_ERR_VB_UNEXIST (The VB pool does not exist)"
	case 0xA0018006: 
		return "HI_ERR_VB_NULL_PTR (The pointer is null)"
	case 0xA0018009:
		return "HI_ERR_VB_NOT_PERM  (The operation is forbidden)"
	case 0xA001800C: 
		return "HI_ERR_VB_NOMEM (The memory fails to be allocated)"
	case 0xA001800D: 
		return "HI_ERR_VB_NOBUF (The buffer fails to be allocated)"
	case 0xA0018010: 
		return "HI_ERR_VB_NOTREADY (The system control attributes are not configured)"
	case 0xA0018012: 
		return "HI_ERR_VB_BUSY (The system is busy)"
	case 0xA0018040: 
		return "HI_ERR_VB_2MPOOLS (Too many VB pools are created)"
	case 0xA0108001: 
		return "HI_ERR_VI_INVALID_DEVID (The VI device ID is invalid)"
	case 0xA0108002: 
		return "HI_ERR_VI_INVALID_CHNID (The VI channel ID is invalid)"
	case 0xA0108003: 
		return "HI_ERR_VI_INVALID_PARA (The VI parameter is invalid)"
	case 0xA0108006: 
		return "HI_ERR_VI_INVALID_NULL_PTR (The pointer of the input parameter is null)"
	case 0xA0108007: 
		return "HI_ERR_VI_FAILED_NOTCONFIG (The attributes of the video device are not set)"
	case 0xA0108008: 
		return "HI_ERR_VI_NOT_SUPPORT (The operation is not supported)"
	case 0xA0108009: 
		return "HI_ERR_VI_NOT_PERM (The operation is forbidden)"
	case 0xA010800C: 
		return "HI_ERR_VI_NOMEM (The memory fails to be allocated)"
	case 0xA010800E: 
		return "HI_ERR_VI_BUF_EMPTY (The VI buffer is empty)"
	case 0xA010800F: 
		return "HI_ERR_VI_BUF_FULL (The VI buffer is full)"
	case 0xA0108010: 
		return "HI_ERR_VI_SYS_NOTREADY (The VI system is not initialized)"
	case 0xA0108012: 
		return "HI_ERR_VI_BUSY (The VI system is busy)"
	case 0xA0108040: 
		return "HI_ERR_VI_FAILED_NOTENABLE (The VI device or VI channel is not enabled)"
	case 0xA0108041: 
		return "HI_ERR_VI_FAILED_NOTDISABLE (The VI device or VI channel is not disabled)"
	case 0xA0108042: 
		return "HI_ERR_VI_FAILED_CHNOTDISABLE (The VI channel is not disabled)"
	case 0xA0108043: 
		return "HI_ERR_VI_CFG_TIMEOUT (The video attribute configuration times out)"
	case 0xA0108045: 
		return "HI_ERR_VI_INVALID_WAYID(The video channel ID is invalid)"
	case 0xA0108046: 
		return "HI_ERR_VI_INVALID_PHYCHNID (The physical video channel ID is invalid)"
	case 0xA0108047: 
		return "HI_ERR_VI_FAILED_NOTBIND (The video channel is not bound)"
	case 0xA0108048: 
		return "HI_ERR_VI_FAILED_BINDED (The video channel is bound)"
	case 0xA0078001: 
		return "HI_ERR_VPSS_INVALID_DEVID () VPSS group ID is invalid)"
	case 0xA0078002: 
		return "HI_ERR_VPSS_INVALID_CHNID (The VPSS channel ID is invalid)"
	case 0xA0078003: 
		return "HI_ERR_VPSS_ILLEGAL_PARAM (The VPSS parameter is invalid)"
	case 0xA0078004: 
		return "HI_ERR_VPSS_EXIST (A VPSS group is created)"
	case 0xA0078005: 
		return "HI_ERR_VPSS_UNEXIST (No VPSS group is created)"
	case 0xA0078006: 
		return "HI_ERR_VPSS_NULL_PTR (The pointer of the input parameter is null)"
	case 0xA0078008: 
		return "HI_ERR_VPSS_NOT_SUPPORT (The operation is not supported)"
	case 0xA0078009: 
		return "HI_ERR_VPSS_NOT_PERM (The operation is forbidden)"
	case 0xA007800C: 
		return "HI_ERR_VPSS_NOMEM (The memory fails to be allocated)"
	case 0xA007800D: 
		return "HI_ERR_VPSS_NOBUF (The buffer pool fails to be allocated)"
	case 0xA007800E: 
		return "HI_ERR_VPSS_BUF_EMPTY (The picture queue is empty)"
	case 0xA0078010: 
		return "HI_ERR_VPSS_NOTREADY (The VPSS is not initialized)"
	case 0xA0078012: 
		return "HI_ERR_VPSS_BUSY (The VPSS is busy)"
	case 0xA0088002:
		return "HI_ERR_VENC_INVALID_CHNID  (The channel ID is invalid)"
	case 0xA0088003: 
		return "HI_ERR_VENC_ILLEGAL_PARAM (The parameter is invalid)"
	case 0xA0088004:
		return "HI_ERR_VENC_EXIST (The device, channel or resource to be created or applied for exists)"
	case 0xA0088005: 
		return "HI_ERR_VENC_UNEXIST (The device, channel or resource to be used or destroyed does not exist)"
	case 0xA0088006: 
		return "HI_ERR_VENC_NULL_PTR (The parameter pointer is null)"
	case 0xA0088007:// 
		return "HI_ERR_VENC_NOT_CONFIG (No parameter is set before use)"
	case 0xA0088008:
		return "HI_ERR_VENC_NOT_SUPPORT (The parameter or function is not supported)"
	case 0xA0088009: 
		return "HI_ERR_VENC_NOT_PERM (The operation, for example, modifying static parameters, is forbidden)"
	case 0xA008800C: 
		return "HI_ERR_VENC_NOMEM (The memory fails to be allocated due to some causes such as insufficient system memory)"
	case 0xA008800D: 
		return "HI_ERR_VENC_NOBUF (The buffer fails to be allocated due to some causes such as oversize of the data buffer applied for)"
	case 0xA008800E: 
		return "HI_ERR_VENC_BUF_EMPTY (The buffer is empty)"
	case 0xA008800F: 
		return "HI_ERR_VENC_BUF_FULL (The buffer is full)"
	case 0xA0088010: 
		return "HI_ERR_VENC_SYS_NOTREADY (The system is not initialized or the corresponding module is not loaded)"
	case 0xA0088012: 
		return "HI_ERR_VENC_BUSY (The VENC system is busy)"
	case 0xA01C8006: 
		return "HI_ERR_ISP_NULL_PTR (The input pointer is null)"
	case 0xA01C8003: 
		return "HI_ERR_ISP_ILLEGAL_PARAM (The input parameter is invalid)"
	case 0xA01C8043: 
		return "HI_ERR_ISP_SNS_UNREGISTER (sensor is not registered)"
	case 0xA01C0041:
		return "HI_ERR_ISP_MEM_NOT_INIT (The external registers are not initialized)"
	case 0xA01C0040: 
		return "HI_ERR_ISP_NOT_INIT (The ISP is not initialized)"
	case 0xA01C0044: 
		return "HI_ERR_ISP_INVALID_ADDR (The address is invalid)"
	case 0xA01C0042: 
		return "HI_ERR_ISP_ATTR_NOT_CFG (The attribute is not configured)"
	default:
		out := "unknown error " + strconv.FormatInt(code, 16)
		return out
	}
}




