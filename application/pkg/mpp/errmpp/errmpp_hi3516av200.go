//+build arm
//+build hi3516av200
//+build debug

package errmpp

func resolveCode(code uint) (string, string) {
	switch code {
	case 0xA0028003:
		return "HI_ERR_SYS_ILLEGAL_PARAM", "The parameter configuration is invalid"
	case 0xA0028006:
		return "HI_ERR_SYS_NULL_PTR", "The pointer is null"
	case 0xA0028009:
		return "HI_ERR_SYS_NOT_PERM", "The operation is forbidden"
	case 0xA0028010:
		return "HI_ERR_SYS_NOTREADY", "The system control attributes are not configured"
	case 0xA0028012:
		return "HI_ERR_SYS_BUSY", "The system is busy"
	case 0xA002800C:
		return "HI_ERR_SYS_NOMEM", "The memory fails to be allocated due to some causes such as insufficient system memory"
	case 0xA0018003:
		return "HI_ERR_VB_ILLEGAL_PARAM", "The parameter configuration is invalid"
	case 0xA0018005:
		return "HI_ERR_VB_UNEXIST", "The VB pool does not exist"
	case 0xA0018006:
		return "HI_ERR_VB_NULL_PTR", "The pointer is null"
	case 0xA0018009:
		return "HI_ERR_VB_NOT_PERM" ,"The operation is forbidden"
	case 0xA001800C:
		return "HI_ERR_VB_NOMEM" ,"The memory fails to be allocated)"
	case 0xA001800D:
		return "HI_ERR_VB_NOBUF" ,"The buffer fails to be allocated"
	case 0xA0018010:
		return "HI_ERR_VB_NOTREADY", "The system control attributes are not configured"
	case 0xA0018012:
		return "HI_ERR_VB_BUSY", "The system is busy"
	case 0xA0018040:
		return "HI_ERR_VB_2MPOOLS", "Too many VB pools are created"
	case 0xA0108001:
		return "HI_ERR_VI_INVALID_DEVID" ,"The VI device ID is invalid"
	case 0xA0108002:
		return "HI_ERR_VI_INVALID_CHNID", "The VI channel ID is invalid"
	case 0xA0108003:
		return "HI_ERR_VI_INVALID_PARA" ,"The VI parameter is invalid"
	case 0xA0108006:
		return "HI_ERR_VI_INVALID_NULL_PTR", "The pointer of the input parameter is null"
	case 0xA0108007:
		return "HI_ERR_VI_FAILED_NOTCONFIG" ,"The attributes of the video device are not"
	case 0xA0108008:
		return "HI_ERR_VI_NOT_SUPPORT", "The operation is not supported"
	case 0xA0108009:
		return "HI_ERR_VI_NOT_PERM" ,"The operation is forbidden"
	case 0xA010800C:
		return "HI_ERR_VI_NOMEM" ,"The memory fails to be allocated"
	case 0xA010800E:
		return "HI_ERR_VI_BUF_EMPTY" ,"The VI buffer is empty"
	case 0xA010800F:
		return  "HI_ERR_VI_BUF_FULL", "The VI buffer is full"
	case 0xA0108010:
		return "HI_ERR_VI_SYS_NOTREADY", "The VI system is not initialized"
	case 0xA0108012:
		return "HI_ERR_VI_BUSY", "The VI system is busy"
	case 0xA0108040:
		return "HI_ERR_VI_FAILED_NOTENABLE", "The VI device or VI channel is not enabled"
	case 0xA0108041:
		return "HI_ERR_VI_FAILED_NOTDISABLE", "The VI device or VI channel is not disabled"
	case 0xA0108042:
		return "HI_ERR_VI_FAILED_CHNOTDISABLE", "The VI channel is not disabled"
	case 0xA0108043:
		return  "HI_ERR_VI_CFG_TIMEOUT", "The video attribute configuration times out"
	case 0xA0108044:
		return  "HI_ERR_VI_NORM_UNMATCH", "Mismatch occurs"
	case 0xA0108045:
		return  "HI_ERR_VI_INVALID_WAYID", "The video channel ID is invalid"
	case 0xA0108046:
		return  "HI_ERR_VI_INVALID_PHYCHNID", "The physical video channel ID is invalid"
	case 0xA0108047:
		return  "HI_ERR_VI_FAILED_NOTBIND", "The video channel is not bound"
	case 0xA0108048:
		return  "HI_ERR_VI_FAILED_BINDED", "The video channel is bound"
	case 0xA0108049:
		return  "HI_ERR_VI_DIS_PROCESS_FAIL", "The DIS fails to run"
	case 0xA00F8001:
		return  "HI_ERR_VO_INVALID_DEVID", "The device ID does not fall within the value range"
	case 0xA00F8002:
		return  "HI_ERR_VO_INVALID_CHNID", "The channel ID does not fall within the value range"
	case 0xA00F8003:
		return  "HI_ERR_VO_ILLEGAL_PARAM", "The parameter value does not fall within the value range"
	case 0xA00F8006:
		return  "HI_ERR_VO_NULL_PTR", "The pointer is null"
	case 0xA00F8008:
		return  "HI_ERR_VO_NOT_SUPPORT", "The operation is not supported"
	case 0xA00F8009:
		return  "HI_ERR_VO_NOT_PERMIT", "The operation is forbidden"
	case 0xA00F800C:
		return  "HI_ERR_VO_NO_MEM", "The memory is insufficient"
	case 0xA00F8010:
		return  "HI_ERR_VO_SYS_NOTREADY", "The system is not initialized"
	case 0xA00F8012:
		return  "HI_ERR_VO_BUSY", "The resources are unavailable"
	case 0xA00F8040:
		return  "HI_ERR_VO_DEV_NOT_CONFIG", "The VO device is not configured"
	case 0xA00F8041:
		return  "HI_ERR_VO_DEV_NOT_ENABLE", "The VO device is not enabled"
	case 0xA00F8042:
		return  "HI_ERR_VO_DEV_HAS_ENABLED", "The VO device has been enabled"
	case 0xA00F8043:
		return  "HI_ERR_VO_DEV_HAS_BINDED", "The device has been bound"
	case 0xA00F8044:
		return  "HI_ERR_VO_DEV_NOT_BINDED", "The device is not bound"
	case 0xA00F8045:
		return  "HI_ERR_VO_VIDEO_NOT_ENABLE", "The video layer is not enabled"
	case 0xA00F8046:
		return  "HI_ERR_VO_VIDEO_NOT_DISABLE", "The video layer is not disabled"
	case 0xA00F8047:
		return  "HI_ERR_VO_VIDEO_NOT_CONFIG", "The video layer is not configured"
	case 0xA00F8048:
		return  "HI_ERR_VO_CHN_NOT_DISABLE", "The VO channel is not disabled"
	case 0xA00F8049:
		return  "HI_ERR_VO_CHN_NOT_ENABLE", "No VO channel is enabled"
	case 0xA00F804A:
		return  "HI_ERR_VO_CHN_NOT_CONFIG", "The VO channel is not configured"
	case 0xA00F804B:
		return  "HI_ERR_VO_CHN_NOT_ALLOC", "No VO channel is allocated"
	case 0xA00F804C:
		return  "HI_ERR_VO_INVALID_PATTERN", "The pattern is invalid"
	case 0xA00F804D:
		return  "HI_ERR_VO_INVALID_POSITION", "The cascade position is invalid"
	case 0xA00F804E:
		return  "HI_ERR_VO_WAIT_TIMEOUT", "Waiting times out"
	case 0xA00F804F:
		return  "HI_ERR_VO_INVALID_VFRAME", "The video frame is invalid"
	case 0xA00F8050:
		return  "HI_ERR_VO_INVALID_RECT_PARA", "The rectangle parameter is invalid"
	case 0xA00F8051:
		return  "HI_ERR_VO_SETBEGIN_ALREADY", "The SETBEGIN MPI has been configured"
	case 0xA00F8052:
		return  "HI_ERR_VO_SETBEGIN_NOTYET", "The SETBEGIN MPI is not configured"
	case 0xA00F8053:
		return  "HI_ERR_VO_SETEND_ALREADY", "The SETEND MPI has been configured"
	case 0xA00F8054:
		return  "HI_ERR_VO_SETEND_NOTYET", "The SETEND MPI is not configured"
	case 0xA00F8065:
		return  "HI_ERR_VO_GFX_NOT_DISABLE", "The graphics layer is not disabled"
	case 0xA00F8066:
		return  "HI_ERR_VO_GFX_NOT_BIND", "The graphics layer is not bound"
	case 0xA00F8067:
		return  "HI_ERR_VO_GFX_NOT_UNBIND", "The graphics layer is not unbound"
	case 0xA00F8068:
		return  "HI_ERR_VO_GFX_INVALID_ID", "The graphics layer ID does not fall within the value range"
	case 0xA00F806b:
		return  "HI_ERR_VO_CHN_AREA_OVERLAP", "The VO channel areas overlap"
	case 0xA00F806d:
		return  "HI_ERR_VO_INVALID_LAYERID", "The video layer ID does not fall within the value range"
	case 0xA00F806e:
		return  "HI_ERR_VO_VIDEO_HAS_BINDED", "The video layer has been bound"
	case 0xA00F806f:
		return  "HI_ERR_VO_VIDEO_NOT_BINDED", "The video layer is not bound"
	case 0xA0078001:
		return  "HI_ERR_VPSS_INVALID_DEVID", "The VPSS group ID is invalid"
	case 0xA0078002:
		return  "HI_ERR_VPSS_INVALID_CHNID", "The VPSS channel ID is invalid"
	case 0xA0078003:
		return  "HI_ERR_VPSS_ILLEGAL_PARAM", "The VPSS parameter is invalid"
	case 0xA0078004:
		return  "HI_ERR_VPSS_EXIST", "A VPSS group is created"
	case 0xA0078005:
		return  "HI_ERR_VPSS_UNEXIST", "No VPSS group is created"
	case 0xA0078006:
		return  "HI_ERR_VPSS_NULL_PTR", "The pointer of the input parameter is null"
	case 0xA0078008:
		return  "HI_ERR_VPSS_NOT_SUPPORT", "The operation is not supported"
	case 0xA0078009:
		return  "HI_ERR_VPSS_NOT_PERM", "The operation is forbidden"
	case 0xA007800C:
		return  "HI_ERR_VPSS_NOMEM", "The memory fails to be allocated"
	case 0xA007800D:
		return  "HI_ERR_VPSS_NOBUF", "The buffer pool fails to be allocated"
	case 0xA007800E:
		return  "HI_ERR_VPSS_BUF_EMPTY", "The picture queue is empty"
	case 0xA0078010:
		return  "HI_ERR_VPSS_NOTREADY", "The VPSS is not initialized"
	case 0xA0078012:
		return  "HI_ERR_VPSS_BUSY", "The VPSS is busy"
	case 0xA0088002:
		return  "HI_ERR_VENC_INVALID_CHNID", "The channel ID is invalid"
	case 0xA0088003:
		return  "HI_ERR_VENC_ILLEGAL_PARAM", "The parameter is invalid"
	case 0xA0088004:
		return  "HI_ERR_VENC_EXIST", "The device, channel or resource to be created or applied for exists"
	case 0xA0088005:
		return  "HI_ERR_VENC_UNEXIST", "The device, channel or resource to be used or destroyed does not exist"
	case 0xA0088006:
		return  "HI_ERR_VENC_NULL_PTR", "The parameter pointer is null"
	case 0xA0088007:
		return  "HI_ERR_VENC_NOT_CONFIG", "No parameter is set before use"
	case 0xA0088008:
		return  "HI_ERR_VENC_NOT_SUPPORT", "The parameter or function is not supported"
	case 0xA0088009:
		return  "HI_ERR_VENC_NOT_PERM", "The operation, for example, modifying static parameters, is forbidden"
	case 0xA008800C:
		return  "HI_ERR_VENC_NOMEM", "The memory fails to be allocated due to some causes such as insufficient system memory"
	case 0xA008800D:
		return  "HI_ERR_VENC_NOBUF", "The buffer fails to be allocated due to some causes such as oversize of the data buffer applied for"
	case 0xA008800E:
		return  "HI_ERR_VENC_BUF_EMPTY", "The buffer is empty"
	case 0xA008800F:
		return  "HI_ERR_VENC_BUF_FULL", "The buffer is full"
	case 0xA0088010:
		return  "HI_ERR_VENC_SYS_NOTREADY", "The system is not initialized or the corresponding module is not loaded"
	case 0xA0088012:
		return  "HI_ERR_VENC_BUSY", "The VENC system is busy"
	case 0xA0098001:
		return  "HI_ERR_VDA_INVALID_DEVID", "The device ID exceeds the valid range"
	case 0xA0098002:
		return  "HI_ERR_VDA_INVALID_CHNID", "The channel ID exceeds the valid range"
	case 0xA0098003:
		return  "HI_ERR_VDA_ILLEGAL_PARAM", "The parameter value exceeds its valid range"
	case 0xA0098004:
		return  "HI_ERR_VDA_EXIST", "The device, channel, or resource to be created or applied for already exists"
	case 0xA0098005:
		return  "HI_ERR_VDA_UNEXIST", "The device, channel, or resource to be used or destroyed does not exist"
	case 0xA0098006:
		return  "HI_ERR_VDA_NULL_PTR", "The pointer is null"
	case 0xA0098007:
		return  "HI_ERR_VDA_NOT_CONFIG", "The system or VDA channel is not configured"
	case 0xA0098008:
		return  "HI_ERR_VDA_NOT_SUPPORT", "The parameter or function is not supported"
	case 0xA0098009:
		return  "HI_ERR_VDA_NOT_PERM", "The operation, for example, attempting to modify the value of a static parameter, is forbidden"
	case 0xA009800C:
		return  "HI_ERR_VDA_NOMEM", "The memory fails to be allocated due to some causes such as insufficient system memory"
	case 0xA009800D:
		return  "HI_ERR_VDA_NOBUF", "The buffer fails to be allocated due to some causes such as oversize of the data buffer applied for"
	case 0xA009800E:
		return  "HI_ERR_VDA_BUF_EMPTY", "The buffer is empty"
	case 0xA009800F:
		return  "HI_ERR_VDA_BUF_FULL", "The buffer is full"
	case 0xA0098010:
		return  "HI_ERR_VDA_SYS_NOTREADY", "The system is not initialized or the corresponding module is not loaded"
	case 0xA0098012:
		return  "HI_ERR_VDA_BUSY", "The system is busy"
	case 0xA0038001:
		return  "HI_ERR_RGN_INVALID_DEVID", "The device ID exceeds the valid range"
	case 0xA0038002:
		return  "HI_ERR_RGN_INVALID_CHNID", "The channel ID is incorrect or the region handle is invalid"
	case 0xA0038003:
		return  "HI_ERR_RGN_ILLEGAL_PARAM", "The parameter value exceeds its valid range"
	case 0xA0038004:
		return  "HI_ERR_RGN_EXIST", "The device, channel, or resource to be created already exists"
	case 0xA0038005:
		return  "HI_ERR_RGN_UNEXIST", "The device, channel, or resource to be used or destroyed does not exist"
	case 0xA0038006:
		return  "HI_ERR_RGN_NULL_PTR", "The pointer is null"
	case 0xA0038007:
		return  "HI_ERR_RGN_NOT_CONFIG", "The module is not configured"
	case 0xA0038008:
		return  "HI_ERR_RGN_NOT_SUPPORT", "The parameter or function is not supported"
	case 0xA0038009:
		return  "HI_ERR_RGN_NOT_PERM", "The operation, for example, attempting to modify the value of a static parameter, is forbidden"
	case 0xA003800C:
		return  "HI_ERR_RGN_NOMEM", "The memory fails to be allocated due to some causes such as insufficient system memory"
	case 0xA003800D:
		return  "HI_ERR_RGN_NOBUF", "The buffer fails to be allocated due to some causes such as oversize of the data buffer applied for"
	case 0xA003800E:
		return  "HI_ERR_RGN_BUF_EMPTY", "The buffer is empty"
	case 0xA003800F:
		return  "HI_ERR_RGN_BUF_FULL", "The buffer is full"
	case 0xA0038010:
		return  "HI_ERR_RGN_NOTREADY", "The system is not initialized or the corresponding module is not loaded"
	case 0xA0038011:
		return  "HI_ERR_RGN_BADADDR", "The address is invalid"
	case 0xA0038012:
		return  "HI_ERR_RGN_BUSY", "The system is busy"
	case 0xA0158001:
		return  "HI_ERR_AI_INVALID_DEVID", "The AI device ID is invalid"
	case 0xA0158002:
		return  "HI_ERR_AI_INVALID_CHNID", "The AI channel ID is invalid"
	case 0xA0158003:
		return  "HI_ERR_AI_ILLEGAL_PARAM", "The settings of the AI parameters are invalid"
	case 0xA0158005:
		return  "HI_ERR_AI_NOT_ENABLED", "The AI device or AI channel is not enabled"
	case 0xA0158006:
		return  "HI_ERR_AI_NULL_PTR", "The input parameter pointer is null"
	case 0xA0158007:
		return  "HI_ERR_AI_NOT_CONFIG", "The attributes of an AI device are not set"
	case 0xA0158008:
		return  "HI_ERR_AI_NOT_SUPPORT", "The operation is not supported"
	case 0xA0158009:
		return  "HI_ERR_AI_NOT_PERM", "The operation is forbidden"
	case 0xA015800C:
		return  "HI_ERR_AI_NOMEM", "The memory fails to be allocated"
	case 0xA015800D:
		return  "HI_ERR_AI_NOBUF", "The AI buffer is insufficient"
	case 0xA015800E:
		return  "HI_ERR_AI_BUF_EMPTY", "The AI buffer is empty"
	case 0xA015800F:
		return  "HI_ERR_AI_BUF_FULL", "The AI buffer is full"
	case 0xA0158010:
		return  "HI_ERR_AI_SYS_NOTREADY", "The AI system is not initialized"
	case 0xA0158012:
		return  "HI_ERR_AI_BUSY", "The AI system is busy"
	case 0xA0158041:
		return  "HI_ERR_AI_VQE_ERR", "A VQE processing error occurs in the AI channel"
	case 0xA0168001:
		return  "HI_ERR_AO_INVALID_DEVID", "The AO device ID is invalid"
	case 0xA0168002:
		return  "HI_ERR_AO_INVALID_CHNID", "The AO channel ID is invalid"
	case 0xA0168003:
		return  "HI_ERR_AO_ILLEGAL_PARAM", "The settings of the AO parameters are invalid"
	case 0xA0168005:
		return "HI_ERR_AO_NOT_ENABLED", "The AO device or AO channel is not enabled"
	case 0xA0168006:
		return  "HI_ERR_AO_NULL_PTR", "The output parameter pointer is null"
	case 0xA0168007:
		return  "HI_ERR_AO_NOT_CONFIG", "The attributes of an AO device are not set"
	case 0xA0168008:
		return  "HI_ERR_AO_NOT_SUPPORT", "The operation is not supported"
	case 0xA0168009:
		return  "HI_ERR_AO_NOT_PERM", "The operation is forbidden"
	case 0xA016800C:
		return  "HI_ERR_AO_NOMEM", "The system memory is insufficient"
	case 0xA016800D:
		return  "HI_ERR_AO_NOBUF", "The AO buffer is insufficient"
	case 0xA016800E:
		return  "HI_ERR_AO_BUF_EMPTY", "The AO buffer is empty"
	case 0xA016800F:
		return  "HI_ERR_AO_BUF_FULL", "The AO buffer is full"
	case 0xA0168010:
		return  "HI_ERR_AO_SYS_NOTREADY", "The AO system is not initialized"
	case 0xA0168012:
		return  "HI_ERR_AO_BUSY", "The AO system is busy"
	case 0xA0168041:
		return  "HI_ERR_AO_VQE_ERR", "A VQE processing error occurs in the AO channel"
	case 0xA0178001:
		return  "HI_ERR_AENC_INVALID_DEVID", "The AENC device ID is invalid"
	case 0xA0178002:
		return  "HI_ERR_AENC_INVALID_CHNID", "The AENC channel ID is invalid"
	case 0xA0178003:
		return  "HI_ERR_AENC_ILLEGAL_PARAM", "The settings of the AENC parameters are invalid"
	case 0xA0178004:
		return  "HI_ERR_AENC_EXIST", "An AENC channel is created"
	case 0xA0178005:
		return  "HI_ERR_AENC_UNEXIST", "An AENC channel is not created"
	case 0xA0178006:
		return  "HI_ERR_AENC_NULL_PTR", "The input parameter pointer is null"
	case 0xA0178007:
		return  "HI_ERR_AENC_NOT_CONFIG", "The AENC channel is not configured"
	case 0xA0178008:
		return  "HI_ERR_AENC_NOT_SUPPORT", "The operation is not supported"
	case 0xA0178009:
		return  "HI_ERR_AENC_NOT_PERM", "The operation is forbidden"
	case 0xA017800C:
		return  "HI_ERR_AENC_NOMEM", "The system memory is insufficient"
	case 0xA017800D:
		return  "HI_ERR_AENC_NOBUF", "The buffer for AENC channels fails to be allocated"
	case 0xA017800E:
		return  "HI_ERR_AENC_BUF_EMPTY", "The AENC channel buffer is empty"
	case 0xA017800F:
		return  "HI_ERR_AENC_BUF_FULL" ,"The AENC channel buffer is full"
	case 0xA0178010:
		return  "HI_ERR_AENC_SYS_NOTREADY" ,"The system is not initialized"
	case 0xA0178040:
		return  "HI_ERR_AENC_ENCODER_ERR", "An AENC data error occurs"
	case 0xA0178041:
		return  "HI_ERR_AENC_VQE_ERR A VQE", "processing error occurs in the AENC channel"    //TODO
	case 0xA0188001:
		return  "HI_ERR_ADEC_INVALID_DEVID", "The ADEC device is invalid"
	case 0xA0188002:
		return  "HI_ERR_ADEC_INVALID_CHNID", "The ADEC channel ID is invalid"
	case 0xA0188003:
		return  "HI_ERR_ADEC_ILLEGAL_PARAM", "The settings of the ADEC parameters are invalid"
	case 0xA0188004:
		return  "HI_ERR_ADEC_EXIST", "An ADEC channel is created"
	case 0xA0188005:
		return  "HI_ERR_ADEC_UNEXIST", "An ADEC channel is not created"
	case 0xA0188006:
		return  "HI_ERR_ADEC_NULL_PTR", "The input parameter pointer is null"
	case 0xA0188007:
		return  "HI_ERR_ADEC_NOT_CONFIG", "The attributes of an ADEC channel are not set"
	case 0xA0188008:
		return  "HI_ERR_ADEC_NOT_SUPPORT", "The operation is not supported"
	case 0xA0188009:
		return  "HI_ERR_ADEC_NOT_PERM", "The operation is forbidden"
	case 0xA018800C:
		return  "HI_ERR_ADEC_NOMEM", "The system memory is insufficient"
	case 0xA018800D:
		return  "HI_ERR_ADEC_NOBUF", "The buffer for ADEC channels fails to be allocated"
	case 0xA018800E:
		return  "HI_ERR_ADEC_BUF_EMPTY", "The ADEC channel buffer is empty"
	case 0xA018800F:
		return  "HI_ERR_ADEC_BUF_FULL", "The ADEC channel buffer is full"
	case 0xA0188010:
		return  "HI_ERR_ADEC_SYS_NOTREADY", "The system is not initialized"
	case 0xA0188040:
		return  "HI_ERR_ADEC_DECODER_ERR", "An ADEC data error occurs"
	case 0xA0188041:
		return  "HI_ERR_ADEC_BUF_LACK", "An insufficient buffer occurs in the ADEC channel"
	case 0xA02D800E:
		return  "HI_ERR_VGS_BUF_EMPTY", "The VGS jobs, tasks, or nodes are used up"
	case 0xA02D8003:
		return  "HI_ERR_VGS_ILLEGAL_PARAM", "The VGS parameter value is invalid"
	case 0xA02D8006:
		return  "HI_ERR_VGS_NULL_PTR", "The input parameter pointer is null"
	case 0xA02D8008:
		return  "HI_ERR_VGS_NOT_SUPPORT", "The operation is not supported"
	case 0xA02D8009:
		return "HI_ERR_VGS_NOT_PERMITTED", "The operation is forbidden"
	case 0xA02D800D:
		return  "HI_ERR_VGS_NOBUF", "The memory fails to be allocated"
	case 0xA02D8010:
		return  "HI_ERR_VGS_SYS_NOTREADY", "The system is not initialized"
	case 0xA033800D:
		return  "HI_ERR_FISHEYE_NOBUF", "The memory fails to be allocated"
	case 0xA033800E:
		return  "HI_ERR_FISHEYE_BUF_EMPTY", "The jobs, tasks, or nodes of the fisheye subsystem are used up"
	case 0xA0338006:
		return  "HI_ERR_FISHEYE_NULL_PTR", "The pointer of the input parameter is null"
	case 0xA0338003:
		return "HI_ERR_FISHEYE_ILLEGAL_PARAM", "The configuration of fisheye parameters is invalid"
	case 0xA0338010:
		return  "HI_ERR_FISHEYE_SYS_NOTREADY", "The system is not initialized"
	case 0xA0338008:
		return  "HI_ERR_FISHEYE_NOT_SUPPORT", "This operation is not supported"
	case 0xA0338009:
		return  "HI_ERR_FISHEYE_NOT_PERMITTED", "This operation is not allowed"
	case 0xA01C8006:
		return "HI_ERR_ISP_NULL_PTR", "The input pointer is null"
	case 0xA01C8003:
		return "HI_ERR_ISP_ILLEGAL_PARAM", "The input parameter is invalid"
	case 0xA01C8008:
		return "HI_ERR_ISP_NOT_SUPPORT", "This function is not supported by the ISP"
	case 0xA01C8043:
		return "HI_ERR_ISP_SNS_UNREGISTER", "The sensor is not registered"
	case 0xA01C8041:
		return "HI_ERR_ISP_MEM_NOT_INIT", "The external registers are not initialized"
	case 0xA01C8040:
		return "HI_ERR_ISP_NOT_INIT", "The ISP is not initialized"
	case 0xA01C8044:
		return "HI_ERR_ISP_INVALID_ADDR", "The address is invalid"
	case 0xA01C8042:
		return "HI_ERR_ISP_ATTR_NOT_CFG", "The attribute is not configured"
	case 0xA01C8045:
		return "HI_ERR_ISP_NOMEM", "The memory is insufficient"
	case 0xA01C8046:
		return "HI_ERR_ISP_NO_INT", "The ISP module has no interrupt"
	default:
		return "unknown", "unknown"
	}
}
